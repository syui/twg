package postp

import (
	"fmt"
	"net/url"
	"github.com/urfave/cli/v2"
	"github.com/syui/twg/oauth"
)

func Postp(c *cli.Context) error {
	if c.NArg() > 0 {
		mes := c.Args().First()
		api := oauth.GetOAuthApi()
		v := url.Values{}
		v.Set("tweet_mode", "extended")
		tweet, err := api.PostTweet(mes, v)
		if err != nil {
		  panic(err)
		}
		fmt.Println(tweet.Id)
	} else {
		fmt.Println("twg p 'message'")
	}
	return nil
}
