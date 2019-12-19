package timeline

import (
	"fmt"
	"net/url"
	"github.com/urfave/cli/v2"
	"github.com/syui/twg/oauth"
	"github.com/syui/twg/color"
)

func GetTimeLine(c *cli.Context) error {
	api := oauth.GetOAuthApi()
	v := url.Values{}
	s := c.Args().First()
	if c.NArg() > 0 {
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
