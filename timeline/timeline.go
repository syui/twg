package timeline

import (
	"fmt"
	"net/url"
	"github.com/urfave/cli"
	"gitlab.com/syui/twg/oauth"
)

func GetTimeLine(c *cli.Context) error {
	api := oauth.GetOAuthApi()
	v := url.Values{}
	s := c.Args().First()
	if len(c.Args()) > 0 {
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
		fmt.Println(tweet.User.ScreenName, "RT", rname, retweet.FullText)
	  } else {
		fmt.Println(tweet.User.ScreenName, tweet.FullText)
	  }

	}
	return nil
}
