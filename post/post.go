package post

import (
	"fmt"
	"net/url"
	"github.com/urfave/cli/v2"
	"github.com/syui/twg/oauth"
	"github.com/syui/twg/color"
)

func Post(c *cli.Context) error {
	if c.NArg() > 0 {
		mes := c.Args().First()
		api := oauth.GetOAuthApi()
		v := url.Values{}
		v.Set("tweet_mode", "extended")
		tweet, err := api.PostTweet(mes, v)
		if err != nil {
		  panic(err)
		}
		fmt.Println(color.Cyan(tweet.User.ScreenName), tweet.FullText)
	} else {
		fmt.Println("twg p 'message'")
	}
	return nil
}
