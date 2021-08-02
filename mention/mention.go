package mention

import (
	"fmt"
	"log"
	"net/url"
	"github.com/urfave/cli/v2"
	"github.com/syui/twg/oauth"
	"github.com/syui/twg/color"
)

func GetMention(c *cli.Context) error {
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
		fmt.Println(color.Cyan(mention.User.ScreenName), mention.FullText, mention.Id, mention.InReplyToStatusID)
	}
	return nil
}

func GetTimeLineId(c *cli.Context) error {
	api := oauth.GetOAuthApi()
	v := url.Values{}
	s := c.Args().First()
	if c.NArg() > 0 {
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
			fmt.Println(color.Cyan(tweet.User.ScreenName), "RT", color.Red(rname), retweet.FullText, tweet.Id, retweet.InReplyToStatusID)
		} else {
				fmt.Println(color.Cyan(tweet.User.ScreenName), tweet.FullText, tweet.Id, tweet.InReplyToStatusID)
		}
		tweeturl := tweet.Entities.Urls
		if  len(tweeturl) != 0 {
			fmt.Println(color.Blue(tweeturl[0].Expanded_url))
		}
	}
	return nil
}

func GetUserTimeLineId(c *cli.Context) error {
	name := c.Args().First()
	api := oauth.GetOAuthApi()
	v := url.Values{}
	v.Set("screen_name",name)
	if c.NArg() == 0 {
		s := c.Args().First()
		v.Set("count",s)
	} else if c.NArg() == 2 {
		s := c.Args().Get(1)
		v.Set("count",s)
	} else {
		v.Set("count","10")
	}
	tweets, err := api.GetUserTimeline(v)
	if err != nil {
	  panic(err)
	}
	for _, tweet := range tweets {
		tweeturl := tweet.Entities.Urls
		retweet := tweet.RetweetedStatus


		if  len(tweeturl) != 0 {
			fmt.Println(color.Blue(tweeturl[0].Expanded_url))
		}
		if retweet != nil {
		      rname := "@" + tweet.Entities.User_mentions[0].Screen_name
		      fmt.Println(color.Cyan(tweet.User.ScreenName), "RT", color.Red(rname), retweet.FullText, retweet.Id)
		} else {
		      						fmt.Println(color.Cyan(tweet.User.ScreenName), tweet.FullText, tweet.Id, tweet.InReplyToStatusID)
		}
		if  len(tweeturl) != 0 {
			fmt.Println(color.Blue(tweeturl[0].Expanded_url))
		}
	}
	return nil
}
