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
	"gitlab.com/syui/twg/oauth"
	"github.com/martinlindhe/imgcat/lib"
	"github.com/fatih/color"
	"github.com/ChimeraCoder/anaconda"
	"github.com/urfave/cli"
	"github.com/nfnt/resize"
)

func ViewImageUser(filename string) {
	dir := filepath.Join(os.Getenv("HOME"), ".config", "twg", "img")
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
	} else if pos == ".jpg" || pos == ".jpeg" {
		jpeg.Encode(out, m, nil)
	} else if pos == ".gif" {
		gif.Encode(out, m, nil)
	}
}

func GetImage(url string, file string){
	dir := filepath.Join(os.Getenv("HOME"), ".config", "twg", "img")
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		panic(err)
	}
	dirIcon := filepath.Join(dir, file)
	if b := Exists(dirIcon); b {
		return
	}
	img, _ := os.Create(dirIcon)
	defer img.Close()
	fmt.Println(url)
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
	blue := color.New(color.FgBlue).SprintFunc()
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
		tweeturl := tweet.Entities.Urls
		if  len(tweeturl) != 0 {
			fmt.Println(blue(tweeturl[0].Expanded_url))
		}
	}
	return
}

func ItermGetTimeLineOption(c *cli.Context) error {
	blue := color.New(color.FgBlue).SprintFunc()
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
		tweeturl := tweet.Entities.Urls
		if  len(tweeturl) != 0 {
			fmt.Println(blue(tweeturl[0].Expanded_url))
		}
	}
	return nil
}

func ItermRunStream() {
	blue := color.New(color.FgBlue).SprintFunc()
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
		tweeturl := v.Entities.Urls
		if  len(tweeturl) != 0 {
			fmt.Println(blue(tweeturl[0].Expanded_url))
		}
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
