package oauth

import (
	"flag"
	"fmt"
	"log"
	"os"
	"io/ioutil"
	"net/url"
	"encoding/json"
	"path/filepath"
	"github.com/ChimeraCoder/anaconda"
	"github.com/skratchdot/open-golang/open"
	"github.com/hokaccha/go-prettyjson"
	"github.com/mrjones/oauth"
)

var ckey string
var cskey string

type Oauth struct {
	AdditionalData struct {
		ScreenName string `json:"screen_name"`
		UserID     string `json:"user_id"`
	} `json:"AdditionalData"`
	Secret string `json:"Secret"`
	Token  string `json:"Token"`
}

func GetOAuthApi() *anaconda.TwitterApi {
	var o Oauth
	anaconda.SetConsumerKey(ckey)
	anaconda.SetConsumerSecret(cskey)
	dir := filepath.Join(os.Getenv("HOME"), ".config", "twg")
	dirConf := filepath.Join(dir, "user.json")
	_, err := os.Stat(dirConf)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		panic(err)
	}
	file,err := ioutil.ReadFile(dirConf)
	if err != nil {
		fmt.Printf("$ twg --oauth", err)
		os.Exit(1)
	}
	json.Unmarshal(file, &o)
	api := anaconda.NewTwitterApi(o.Token, o.Secret)
	return api
}

func RunOAuth() {
	anaconda.SetConsumerKey(ckey)
	anaconda.SetConsumerSecret(cskey)
	dir := filepath.Join(os.Getenv("HOME"), ".config", "twg")
	dirConf := filepath.Join(dir, "user.json")
	//dirTweet := filepath.Join(dir, "tweet.json")
	_, err := os.Stat(dirConf)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		panic(err)
	}
	flag.Parse()
	c := oauth.NewConsumer(
		string(ckey),
		string(cskey),
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		})

	requestToken, u, err := c.GetRequestTokenAndUrl("oob")
	if err != nil {
	    log.Fatal(err)
	}

	fmt.Print("\nopen url: ", string(u))
	fmt.Print("\ninput pin: ")
	open.Run(u)

	verificationCode := ""
	fmt.Scanln(&verificationCode)
	accessToken, err := c.AuthorizeToken(requestToken, verificationCode)
	if err != nil {
		log.Fatal(err)
	}
	outputJSON, err := json.Marshal(&accessToken)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nYour token is %s\n", outputJSON)
	jat, _ := prettyjson.Marshal(accessToken)
	fmt.Printf("\nYour token is %s\n", jat)
	ioutil.WriteFile(dirConf, outputJSON, os.ModePerm)
	return
}

func GetOAuthTimeLine() {
	api := GetOAuthApi()
	v := url.Values{}
	v.Set("count","10")
	v.Set("tweet_mode", "extended")
	tweets, err := api.GetHomeTimeline(v)
	if err != nil {
	  panic(err)
	}
	for _, tweet := range tweets {
		fmt.Println(tweet.User.ScreenName, tweet.FullText)
	}

	return
}

func RunStream() {
	api := GetOAuthApi()
	v := url.Values{}
	v.Set("count","1")
	v.Set("tweet_mode", "extended")
	tweets, err := api.GetHomeTimeline(v)
	if err != nil {
	  panic(err)
	}
	for _, tweet := range tweets {
		fmt.Println(tweet.User.ScreenName, tweet.FullText)
	}
	s := api.UserStream(v)
	for t := range s.C {
	  switch v := t.(type) {
	  case anaconda.Tweet:
	    fmt.Printf("%-15s: %s\n", v.User.ScreenName, v.FullText)
	  case anaconda.EventTweet:
	    switch v.Event.Event {
	    case "favorite":
	      sn := v.Source.ScreenName
	      tw := v.TargetObject.Text
	      fmt.Printf("Favorited by %-15s: %s\n", sn, tw)
	    case "unfavorite":
	      sn := v.Source.ScreenName
	      tw := v.TargetObject.Text
	      fmt.Printf("UnFavorited by %-15s: %s\n", sn, tw)
	    }
	  }
	}
	return
}

func FirstRunOAuth() {
	var o Oauth
	dir := filepath.Join(os.Getenv("HOME"), ".config", "twg")
	dirConf := filepath.Join(dir, "user.json")
	_, err := os.Stat(dirConf)
	if err := os.MkdirAll(dir, os.ModePerm); err!= nil {
		panic(err)
	}
	file,err := ioutil.ReadFile(dirConf)
	if err != nil {
		RunOAuth()
	} else {
		json.Unmarshal(file, &o)
		fmt.Printf("login: %s\n", o.AdditionalData.ScreenName)
		GetOAuthTimeLine()
	}
	return
}
