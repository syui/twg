package notify

import (
	"fmt"
	"log"
	"net/url"
	"github.com/urfave/cli/v2"
	"github.com/syui/twg/oauth"
	"github.com/syui/twg/color"
)

func GetNotify(c *cli.Context) error {
	api := oauth.GetOAuthApi()
	v := url.Values{}
	s := c.Args().First()
	if c.NArg() > 0 {
		v.Set("count",s)
	} else {
		v.Set("count","10")
	}
	v.Set("tweet_mode", "extended")
	mentions, err := api.GetMentionsTimeline(v)
	if err != nil {
		log.Fatal(err)
	}
	for _, mention := range mentions {
		rname := "\"@" + mention.User.ScreenName
		fmt.Println(color.Cyan(rname), mention.FullText)
		fmt.Println("Re:twg mm",color.Red(mention.Id), color.Cyan(rname), "$message\"")
		if mention.InReplyToStatusID != 0 {
			fmt.Println("src:twg mm",color.Blue(mention.InReplyToStatusID), color.Cyan(rname), "$message\"")
		}
		fmt.Println("-----------------------------------------")
	}
	return nil
}
