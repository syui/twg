package icon

import (
	"bytes"
	"io"
	"os"
	"log"
	"fmt"
	"net/url"
	"net/http"
	"io/ioutil"
	"path/filepath"
	"encoding/json"
	"image"
	"image/png"
	"image/jpeg"
	"gitlab.com/syui/twg/oauth"
	"github.com/martinlindhe/imgcat/lib"
	"github.com/fatih/color"
	"github.com/ChimeraCoder/anaconda"
	"github.com/urfave/cli"
	"github.com/nfnt/resize"
)


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

func ImageResize(dir string) {
	pos := filepath.Ext(dir)
	file, err := os.Open(dir)
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(file)
	if err != nil {
		return
	}
	file.Close()
	m := resize.Resize(20, 20, img, resize.Lanczos3)
	out, err := os.Create(dir)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	if pos == ".png" {
		png.Encode(out, m)
	} else if pos == ".jpg" {
		jpeg.Encode(out, m, nil)
	}
}

func GetImage(url string, file string){
	dir := filepath.Join(os.Getenv("HOME"), ".config", "twg")
	dirIcon := filepath.Join(dir, file)
	if b := Exists(dirIcon); b {
		return
	}
	img, _ := os.Create(dirIcon)
	defer img.Close()
	//fmt.Println(url)
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	io.Copy(img, resp.Body)
	f, _ := os.Open(dirIcon)
	buf := new(bytes.Buffer)
	io.Copy(buf, f)
	ImageResize(dirIcon)
	return
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
		url := tweet.User.ProfileImageURL
		pos := filepath.Ext(url)
		file := name + pos
		GetImage(url, file)
		ViewImageUser(file)
		fmt.Println(cyan(tweet.User.ScreenName), tweet.FullText)
	}
	return
}

func ItermGetTimeLineOption(c *cli.Context) error {
	cyan := color.New(color.FgCyan).SprintFunc()
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
		fmt.Println(cyan(tweet.User.ScreenName), tweet.FullText)
	}
	return nil
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
		url := v.User.ProfileImageURL
		pos := filepath.Ext(url)
		file := name + pos
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
		url := tweet.User.ProfileImageURL
		pos := filepath.Ext(url)
		file := name + pos
		GetImage(url, file)
		ViewImageUser(file)
		fmt.Println(cyan(tweet.User.ScreenName), tweet.FullText)
	}
	return nil
}

func CheckOAuth() (check bool){
	//var o oauth.Oauth
	dir := filepath.Join(os.Getenv("HOME"), ".config", "twg")
	dirConf := filepath.Join(dir, "user.json")
	if b := Exists(dirConf); b {
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
