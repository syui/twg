package icon

import (
	"bytes"
	"io"
	"os"
	"log"
	"fmt"
	"net/url"
	"net/http"
	"path/filepath"
	"image"
	"image/gif"
	"image/png"
	"image/jpeg"
	"github.com/martinlindhe/imgcat/lib"
	"github.com/ChimeraCoder/anaconda"
	"github.com/urfave/cli"
	"github.com/nfnt/resize"
	"gitlab.com/syui/twg/path"
	"gitlab.com/syui/twg/color"
	"gitlab.com/syui/twg/oauth"
)

var dir = path.Dir
var dirVerify = path.DirVerify
var dirUser = path.DirUser
var dirImg = path.DirImg
//var api = oauth.oauth.GetOAuthApi()
//var v = url.Values{}

func ViewImageUser(filename string) {
	f := filepath.Join(dirImg, filename)
	file, _ := os.Open(f)
	imgcat.Cat(file, os.Stdout)
	return
}

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func ImageResize(dirIcon string) {
	pos := filepath.Ext(dirIcon)
	file, err := os.Open(dirIcon)
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(file)
	if err != nil {
		return
	}
	file.Close()
	m := resize.Resize(20, 20, img, resize.Lanczos3)
	out, err := os.Create(dirIcon)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	if pos == ".png" {
		png.Encode(out, m)
	} else if pos == ".jpg" || pos == ".jpeg" {
		jpeg.Encode(out, m, nil)
	} else if pos == ".gif" {
		gif.Encode(out, m, nil)
	}
}

func GetImage(url string, file string){
	if err := os.MkdirAll(dirImg, os.ModePerm); err != nil {
		panic(err)
	}
	dirIcon := filepath.Join(dirImg, file)
	if b := Exists(dirIcon); b {
		return
	}
	img, _ := os.Create(dirIcon)
	defer img.Close()
	//fmt.Println(url, "-> ", dirIcon)
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	io.Copy(img, resp.Body)
	f, _ := os.Open(dirIcon)
	buf := new(bytes.Buffer)
	io.Copy(buf, f)
	ImageResize(dirIcon)
	return
}

func ItermGetTimeLine() {
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
		url := tweet.User.ProfileImageURL
		pos := filepath.Ext(url)
		file := name + pos
		GetImage(url, file)
		ViewImageUser(file)
		tweeturl := tweet.Entities.Urls
		retweet := tweet.RetweetedStatus
		if retweet != nil {
		      rname := "@" + tweet.Entities.User_mentions[0].Screen_name
		      fmt.Println(color.Cyan(tweet.User.ScreenName), "RT", color.Red(rname), retweet.FullText)
		} else {
		      fmt.Println(color.Cyan(tweet.User.ScreenName), tweet.FullText)
		}
		if  len(tweeturl) != 0 {
			fmt.Println(color.Blue(tweeturl[0].Expanded_url))
		}
	}
	return
}

func ItermGetTimeLineOption(c *cli.Context) error {
	api := oauth.GetOAuthApi()
	v := url.Values{}
	s := c.Args().First()
	if len(c.Args()) > 0 {
		v.Set("count",s)
	} else {
		v.Set("count","10")
	}
	v.Set("tweet_mode", "extended")
	tweets, err := api.GetHomeTimeline(v)
	if err != nil {
	  panic(err)
	}
	for _, tweet := range tweets {
		name := tweet.User.ScreenName
		url := tweet.User.ProfileImageURL
		pos := filepath.Ext(url)
		file := name + pos
		GetImage(url, file)
		ViewImageUser(file)
		retweet := tweet.RetweetedStatus
		if retweet != nil {
		      rname := "@" + tweet.Entities.User_mentions[0].Screen_name
		      fmt.Println(color.Cyan(tweet.User.ScreenName), "RT", color.Red(rname), retweet.FullText)
		} else {
		      fmt.Println(color.Cyan(tweet.User.ScreenName), tweet.FullText)
		}
		tweeturl := tweet.Entities.Urls
		if  len(tweeturl) != 0 {
			fmt.Println(color.Blue(tweeturl[0].Expanded_url))
		}
	}
	return nil
}

func ItermRunStream() {
	api := oauth.GetOAuthApi()
	v := url.Values{}
	v.Set("tweet_mode", "extended")
	s := api.UserStream(v)
	for t := range s.C {
	  switch v := t.(type) {
	  case anaconda.Tweet:
		name := v.User.ScreenName
		url := v.User.ProfileImageURL
		pos := filepath.Ext(url)
		file := name + pos
		GetImage(url, file)
		ViewImageUser(file)
		tweeturl := v.Entities.Urls
		retweet := v.RetweetedStatus
		if retweet != nil {
		      rname := "@" + v.Entities.User_mentions[0].Screen_name
		      fmt.Println(color.Cyan(v.User.ScreenName), "RT", color.Red(rname), retweet.FullText)
		} else {
		      fmt.Println(color.Cyan(v.User.ScreenName), v.FullText)
		}

		if  len(tweeturl) != 0 {
			fmt.Println(color.Blue(tweeturl[0].Expanded_url))
		}
	  case anaconda.EventTweet:
	    switch v.Event.Event {
	    case "favorite":
	      sn := v.Source.ScreenName
	      tw := v.TargetObject.FullText
	      fmt.Printf("Favorited by %-15s: %s\n", color.Yellow(sn), tw)
	    case "unfavorite":
	      sn := v.Source.ScreenName
	      tw := v.TargetObject.FullText
	      fmt.Printf("UnFavorited by %-15s: %s\n", color.Yellow(sn), tw)
	    }
	  }
	}
	return
}

func ItermUser(c *cli.Context) error {
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
		url := tweet.User.ProfileImageURL
		pos := filepath.Ext(url)
		file := name + pos
		GetImage(url, file)
		ViewImageUser(file)
		fmt.Println(color.Cyan(tweet.User.ScreenName), tweet.FullText)
	}
	return nil
}

func CheckOAuth() (check bool){
	if b := Exists(dirVerify); b {
		check = true
	} else {
		check = false
	}
	return check
}

func FirstItermCommand() {
	check := CheckOAuth()
	if check == true {
		ItermGetTimeLine()
	} else {
		oauth.RunOAuth()
	}
}
