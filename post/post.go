package post

import (
	"fmt"
	"net/url"
	"github.com/urfave/cli"
	"gitlab.com/syui/twg/oauth"
)

func Post(c *cli.Context) error {
	if len(c.Args()) > 0 {
		mes := c.Args().First()
		api := oauth.GetOAuthApi()
		v := url.Values{}
		v.Set("tweet_mode", "extended")
		tweet, err := api.PostTweet(mes, v)
		if err != nil {
		  panic(err)
		}
		fmt.Println(tweet.User.ScreenName, tweet.FullText)
	} else {
		fmt.Println("twg p 'message'")
	}
	return nil
}
