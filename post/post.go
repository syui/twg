package post

import (
	"fmt"
	"github.com/urfave/cli"
	"gitlab.com/syui/twg/oauth"
)

func Post(c *cli.Context) error {
	if len(c.Args()) > 0 {
		mes := c.Args().First()
		api := oauth.GetOAuthApi()
		tweet, err := api.PostTweet(mes, nil)
		if err != nil {
		  panic(err)
		}
		fmt.Println(tweet.User.ScreenName, tweet.Text)
	} else {
		fmt.Println("twg p 'message'")
	}
	return nil
}
