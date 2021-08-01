package mm

import (
	"fmt"
	"net/url"
	"github.com/urfave/cli/v2"
	"github.com/syui/twg/oauth"
	"github.com/syui/twg/color"
)

func Mm(c *cli.Context) error {
	if c.NArg() > 1 {
		id := c.Args().First()
		mes := c.Args().Get(1)
		api := oauth.GetOAuthApi()
		v := url.Values{}
		v.Set("in_reply_to_status_id", id)
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
