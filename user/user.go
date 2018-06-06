package user

import (
	"fmt"
	"net/url"
	"github.com/urfave/cli"
	"gitlab.com/syui/twg/oauth"
)

func User(c *cli.Context) error {
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
	  fmt.Println(tweet.User.ScreenName, tweet.FullText)
	}
	return nil
}
