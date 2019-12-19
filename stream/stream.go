package stream

import (
	"fmt"
	"net/url"
	"github.com/urfave/cli/v2"
	"github.com/ChimeraCoder/anaconda"
	"github.com/syui/twg/oauth"
	"github.com/syui/twg/color"
)

func RunStream(c *cli.Context, o string) error {
	api := oauth.GetOAuthApi()
	v := url.Values{}
	v.Set("tweet_mode", "extended")
	s := api.UserStream(v)
	if o == "user" || o == "u" || o == "normal" {
		fmt.Println("UserStream")
		s = api.UserStream(v)
	} else if o == "site" || o == "s" {
		fmt.Println("SiteStream")
		s = api.SiteStream(v)
	} else if o == "public" {
		fmt.Println("PublicStreamSample")
		s = api.PublicStreamSample(v)
	} else {
		return nil
	}
	for t := range s.C {
		switch v := t.(type) {
		case anaconda.Tweet:
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
	return nil
}
