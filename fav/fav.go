package fav

import (
	"fmt"
	"strconv"
	//"net/url"
	"github.com/urfave/cli/v2"
	"github.com/syui/twg/oauth"
	"github.com/syui/twg/color"
)

func Fav(c *cli.Context) error {
	if c.NArg() > 0 {
		id,_ := strconv.ParseInt(c.Args().First(), 10, 64)
		api := oauth.GetOAuthApi()
		//v := url.Values{}
		//v.Set("tweet_mode", "extended")
		tweet, err := api.Favorite(id)
		if err != nil {
		  panic(err)
		}
		fmt.Println(color.Cyan(tweet.User.ScreenName), tweet.FullText)
	} else {
		fmt.Println("twg f $id")
	}
	return nil
}

