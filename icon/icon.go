package icon

import (
	"os"
	"fmt"
	"net/url"
	"path/filepath"
	"gitlab.com/syui/twg/oauth"
	"github.com/martinlindhe/imgcat/lib"
	"github.com/fatih/color"
	"github.com/ChimeraCoder/anaconda"
	//"log"
	//"image/jpeg"
	//"bytes"
	//"image"
	//"image/color"
	//"github.com/mattn/go-sixel"
)

func ItermGetTimeLine() {
	dir := filepath.Join(os.Getenv("HOME"), ".config", "twg")
	f := filepath.Join(dir, "user.jpg")
	file, _ := os.Open(f)

	//lib : go-sixel
	//img, _ := jpeg.Decode(file)
	//buf := new(bytes.Buffer)
	//jpeg.Encode(buf, img, nil)
	//if err := sixel.NewEncoder(os.Stdout).Encode(img); err != nil {
	//    log.Fatal(err)
	//}

	// lib : imgcat
	imgcat.Cat(file, os.Stdout)
	api := oauth.GetOAuthApi()
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


func ItermRunStream() {
	dir := filepath.Join(os.Getenv("HOME"), ".config", "twg")
	f := filepath.Join(dir, "user.jpg")
	file, _ := os.Open(f)
	imgcat.Cat(file, os.Stdout)
	api := oauth.GetOAuthApi()
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
