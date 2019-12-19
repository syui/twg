package search

import (
	"fmt"
	"net/url"
	"github.com/urfave/cli/v2"
	"github.com/syui/twg/oauth"
	"github.com/syui/twg/color"
)

func Search(c *cli.Context) error {
	mes := c.Args().First()
	api := oauth.GetOAuthApi()
	v := url.Values{}
	v.Set("tweet_mode", "extended")
	if c.NArg() == 0 {
		s := c.Args().First()
		v.Set("count",s)
	} else if c.NArg() == 2 {
		s := c.Args().Get(1)
		v.Set("count",s)
	} else {
		v.Set("count","10")
	}
	searchResult, _ := api.GetSearch(mes, v)
	for _ , tweet := range searchResult.Statuses {
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
	return nil
}
