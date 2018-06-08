package icon

import (
	"bytes"
	"io"
	"os"
	"fmt"
	"net/url"
	"path/filepath"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"gitlab.com/syui/twg/oauth"
	"github.com/martinlindhe/imgcat/lib"
	"github.com/fatih/color"
	"github.com/ChimeraCoder/anaconda"
	"github.com/urfave/cli"
)

func ViewImage() {
	dir := filepath.Join(os.Getenv("HOME"), ".config", "twg")
	f := filepath.Join(dir, "user.jpg")
	file, _ := os.Open(f)
	imgcat.Cat(file, os.Stdout)
	return
}

func ViewImageUser(filename string) {
	dir := filepath.Join(os.Getenv("HOME"), ".config", "twg")
	f := filepath.Join(dir, filename)
	file, _ := os.Open(f)
	imgcat.Cat(file, os.Stdout)
	return
}

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func GetImage(url string, file string) {
	dir := filepath.Join(os.Getenv("HOME"), ".config", "twg")
	dirIcon := filepath.Join(dir, file)
	if b := Exists(dirIcon); b {
		return
	}
	img, _ := os.Create(dirIcon)
	defer img.Close()
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	io.Copy(img, resp.Body)
	f, _ := os.Open(dirIcon)
	buf := new(bytes.Buffer)
	io.Copy(buf, f)
}

func GetUserName() (name string){
	var o oauth.UserVerifyCredentials
	dir := filepath.Join(os.Getenv("HOME"), ".config", "twg")
	dirUser := filepath.Join(dir, "verify.json")
	file,err := ioutil.ReadFile(dirUser)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(file, &o)
	name = o.ScreenName
	return
}

func ItermGetTimeLine() {
	cyan := color.New(color.FgCyan).SprintFunc()
	api := oauth.GetOAuthApi()
	v := url.Values{}
	v.Set("count","10")
	v.Set("tweet_mode", "extended")
	tweets, err := api.GetHomeTimeline(v)
	if err != nil {
	  panic(err)
	}
	for _, tweet := range tweets {
		name := tweet.User.ScreenName
		file := name + ".jpg"
		url := tweet.User.ProfileImageURL
		GetImage(url, file)
		ViewImageUser(file)
		fmt.Println(cyan(tweet.User.ScreenName), tweet.FullText)
	}
	return
}

func ItermRunStream() {
	api := oauth.GetOAuthApi()
	v := url.Values{}
	v.Set("tweet_mode", "extended")
	s := api.UserStream(v)
	cyan := color.New(color.FgCyan).SprintFunc()
        yellow := color.New(color.FgYellow).SprintFunc()
	for t := range s.C {
	  switch v := t.(type) {
	  case anaconda.Tweet:
		name := v.User.ScreenName
		file := name + ".jpg"
		url := v.User.ProfileImageURL
		GetImage(url, file)
		ViewImageUser(file)
		fmt.Println(cyan(v.User.ScreenName), v.FullText)
	  case anaconda.EventTweet:
	    switch v.Event.Event {
	    case "favorite":
	      sn := v.Source.ScreenName
	      tw := v.TargetObject.FullText
	      fmt.Printf("Favorited by %-15s: %s\n", yellow(sn), tw)
	    case "unfavorite":
	      sn := v.Source.ScreenName
	      tw := v.TargetObject.FullText
	      fmt.Printf("UnFavorited by %-15s: %s\n", yellow(sn), tw)
	    }
	  }
	}
	return
}

func ItermUser(c *cli.Context) error {
	cyan := color.New(color.FgCyan).SprintFunc()
	name := c.Args().First()
	api := oauth.GetOAuthApi()
	v := url.Values{}
	v.Set("screen_name",name)
	v.Set("count", "10")
	tweets, err := api.GetUserTimeline(v)
	if err != nil {
	  panic(err)
	}
	for _, tweet := range tweets {
		name := tweet.User.ScreenName
		file := name + ".jpg"
		url := tweet.User.ProfileImageURL
		GetImage(url, file)
		ViewImageUser(file)
		fmt.Println(cyan(tweet.User.ScreenName), tweet.FullText)
	}
	return nil
}
