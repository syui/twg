package oauth

import (
	"flag"
	"fmt"
	"log"
	"os"
	"io"
	"bytes"
	"net/url"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"path/filepath"
	"github.com/ChimeraCoder/anaconda"
	"github.com/skratchdot/open-golang/open"
	"github.com/mrjones/oauth"
	"github.com/fatih/color"
	//"github.com/hokaccha/go-prettyjson"
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

type UserVerifyCredentials struct {
	ContributorsEnabled bool   `json:"contributors_enabled"`
	CreatedAt           string `json:"created_at"`
	DefaultProfile      bool   `json:"default_profile"`
	DefaultProfileImage bool   `json:"default_profile_image"`
	Description         string `json:"description"`
	Entities            struct {
		Description struct {
			Urls []interface{} `json:"urls"`
		} `json:"description"`
		URL struct {
			Urls []struct {
				DisplayURL  string  `json:"display_url"`
				ExpandedURL string  `json:"expanded_url"`
				Indices     []int64 `json:"indices"`
				URL         string  `json:"url"`
			} `json:"urls"`
		} `json:"url"`
	} `json:"entities"`
	FavouritesCount                int64  `json:"favourites_count"`
	FollowRequestSent              bool   `json:"follow_request_sent"`
	FollowersCount                 int64  `json:"followers_count"`
	Following                      bool   `json:"following"`
	FriendsCount                   int64  `json:"friends_count"`
	GeoEnabled                     bool   `json:"geo_enabled"`
	HasExtendedProfile             bool   `json:"has_extended_profile"`
	ID                             int64  `json:"id"`
	IDStr                          string `json:"id_str"`
	IsTranslationEnabled           bool   `json:"is_translation_enabled"`
	IsTranslator                   bool   `json:"is_translator"`
	Lang                           string `json:"lang"`
	ListedCount                    int64  `json:"listed_count"`
	Location                       string `json:"location"`
	Name                           string `json:"name"`
	NeedsPhoneVerification         bool   `json:"needs_phone_verification"`
	Notifications                  bool   `json:"notifications"`
	ProfileBackgroundColor         string `json:"profile_background_color"`
	ProfileBackgroundImageURL      string `json:"profile_background_image_url"`
	ProfileBackgroundImageURLHTTPS string `json:"profile_background_image_url_https"`
	ProfileBackgroundTile          bool   `json:"profile_background_tile"`
	ProfileBannerURL               string `json:"profile_banner_url"`
	ProfileImageURL                string `json:"profile_image_url"`
	ProfileImageURLHTTPS           string `json:"profile_image_url_https"`
	ProfileLinkColor               string `json:"profile_link_color"`
	ProfileSidebarBorderColor      string `json:"profile_sidebar_border_color"`
	ProfileSidebarFillColor        string `json:"profile_sidebar_fill_color"`
	ProfileTextColor               string `json:"profile_text_color"`
	ProfileUseBackgroundImage      bool   `json:"profile_use_background_image"`
	Protected                      bool   `json:"protected"`
	ScreenName                     string `json:"screen_name"`
	Status                         struct {
		Contributors interface{} `json:"contributors"`
		Coordinates  interface{} `json:"coordinates"`
		CreatedAt    string      `json:"created_at"`
		Entities     struct {
			Hashtags     []interface{} `json:"hashtags"`
			Symbols      []interface{} `json:"symbols"`
			Urls         []interface{} `json:"urls"`
			UserMentions []interface{} `json:"user_mentions"`
		} `json:"entities"`
		FavoriteCount        int64       `json:"favorite_count"`
		Favorited            bool        `json:"favorited"`
		Geo                  interface{} `json:"geo"`
		ID                   int64       `json:"id"`
		IDStr                string      `json:"id_str"`
		InReplyToScreenName  interface{} `json:"in_reply_to_screen_name"`
		InReplyToStatusID    interface{} `json:"in_reply_to_status_id"`
		InReplyToStatusIDStr interface{} `json:"in_reply_to_status_id_str"`
		InReplyToUserID      interface{} `json:"in_reply_to_user_id"`
		InReplyToUserIDStr   interface{} `json:"in_reply_to_user_id_str"`
		IsQuoteStatus        bool        `json:"is_quote_status"`
		Lang                 string      `json:"lang"`
		Place                interface{} `json:"place"`
		RetweetCount         int64       `json:"retweet_count"`
		Retweeted            bool        `json:"retweeted"`
		Source               string      `json:"source"`
		Text                 string      `json:"text"`
		Truncated            bool        `json:"truncated"`
	} `json:"status"`
	StatusesCount  int64       `json:"statuses_count"`
	Suspended      bool        `json:"suspended"`
	TimeZone       interface{} `json:"time_zone"`
	TranslatorType string      `json:"translator_type"`
	URL            string      `json:"url"`
	UtcOffset      interface{} `json:"utc_offset"`
	Verified       bool        `json:"verified"`
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

func GetUserIcon() {
	var o UserVerifyCredentials
	dir := filepath.Join(os.Getenv("HOME"), ".config", "twg")
	dirUser := filepath.Join(dir, "verify.json")
	dirIcon := filepath.Join(dir, "user.jpg")
	file,err := ioutil.ReadFile(dirUser)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(file, &o)

	//fmt.Printf("download user icon : %s", o.ProfileImageURL)
	img, _ := os.Create(dirIcon)
	defer img.Close()
	url := o.ProfileImageURL
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	io.Copy(img, resp.Body)
	//b, _ := io.Copy(img, resp.Body)
	//fmt.Println("/ size: ", b)
	f, _ := os.Open(dirIcon)
	buf := new(bytes.Buffer)
	io.Copy(buf, f)
}

func RunOAuth() {
	anaconda.SetConsumerKey(ckey)
	anaconda.SetConsumerSecret(cskey)
	dir := filepath.Join(os.Getenv("HOME"), ".config", "twg")
	dirConf := filepath.Join(dir, "user.json")
	dirUser := filepath.Join(dir, "verify.json")
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
	//fmt.Printf("\nYour token is %s\n", outputJSON)
	//jat, _ := prettyjson.Marshal(accessToken)
	//fmt.Printf("\nYour token is %s\n", jat)
	ioutil.WriteFile(dirConf, outputJSON, os.ModePerm)

	// write : ~/.config/twg/profile.json
	client, err := c.MakeHttpClient(accessToken)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Get(
		"https://api.twitter.com/1.1/account/verify_credentials.json")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bit, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bit))
	ioutil.WriteFile(dirUser, bit, os.ModePerm)
	GetUserIcon()
	return
}

func GetOAuthTimeLine() {
	api := GetOAuthApi()
	v := url.Values{}
	v.Set("count","10")
	v.Set("tweet_mode", "extended")
	tweets, err := api.GetHomeTimeline(v)
	cyan := color.New(color.FgCyan).SprintFunc()
	if err != nil {
	  panic(err)
	}
	for _, tweet := range tweets {
		fmt.Println(cyan(tweet.User.ScreenName), tweet.FullText)
	}

	return
}

func RunStream() {
	api := GetOAuthApi()
	v := url.Values{}
	v.Set("tweet_mode", "extended")
	s := api.UserStream(v)
	cyan := color.New(color.FgCyan).SprintFunc()
        yellow := color.New(color.FgYellow).SprintFunc()
	for t := range s.C {
	  switch v := t.(type) {
	  case anaconda.Tweet:
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
		red := color.New(color.FgRed).SprintFunc()
		fmt.Printf("login: %s\n", red(o.AdditionalData.ScreenName))
		GetOAuthTimeLine()
	}
	return
}

