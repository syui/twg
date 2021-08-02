package del

import (
	"fmt"
	"strconv"
	"github.com/urfave/cli/v2"
	"github.com/syui/twg/oauth"
	"github.com/syui/twg/color"
)

func Del(c *cli.Context) error {
	if c.NArg() > 0 {
		id,_ := strconv.ParseInt(c.Args().First(), 10, 64)
		api := oauth.GetOAuthApi()
		v,_ := strconv.ParseBool("t")
		tweet, err := api.DeleteTweet(id,v)
		if err != nil {
		  panic(err)
		}
		fmt.Println(color.Cyan(tweet.User.ScreenName), tweet.FullText)
	} else {
		fmt.Println("twg d $id")
	}
	return nil
}
