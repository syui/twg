package mention

import (
	"fmt"
	"log"
	"net/url"
	"github.com/urfave/cli/v2"
	"github.com/syui/twg/oauth"
	"github.com/syui/twg/color"
)

func GetMentionId(c *cli.Context) error {
	api := oauth.GetOAuthApi()
	v := url.Values{}
	s := c.Args().First()
	if c.NArg() > 0 {
		v.Set("count",s)
	} else {
		v.Set("count","30")
	}
	v.Set("tweet_mode", "extended")
	mentions, err := api.GetMentionsTimeline(v)
	if err != nil {
		log.Fatal(err)
	}
	for _, mention := range mentions {
		rname := "@" + mention.User.ScreenName
		fmt.Println(color.Cyan(mention.User.ScreenName), mention.FullText)
		fmt.Println("re:twg mm",color.Red(mention.Id), "\"", color.Cyan(rname), "$message\"")
		if mention.InReplyToStatusID != 0 {
			fmt.Println("src:twg mm",color.Blue(mention.InReplyToStatusID), "\"", color.Cyan(rname), "$message\"")
		}
		fmt.Println("---------------------------------")
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
		v.Set("count","30")
	}
	v.Set("tweet_mode", "extended")
	tweets, err := api.GetHomeTimeline(v)
	if err != nil {
		panic(err)
	}
	for _, tweet := range tweets {
		retweet := tweet.RetweetedStatus
		rname := "@" + tweet.User.ScreenName
		if retweet != nil {
			fmt.Println(color.Cyan(tweet.User.ScreenName), "RT", color.Red(rname), retweet.FullText)
			fmt.Println("Re:twg mm",color.Red(tweet.Id), "\"", color.Cyan(rname), "$message\"")
			if tweet.InReplyToStatusID != 0 {
				fmt.Println("src:twg mm",color.Blue(tweet.InReplyToStatusID), "\"", color.Cyan(rname), "$message\"")
			}
		} else {
			fmt.Println(color.Cyan(tweet.User.ScreenName), tweet.FullText)
			fmt.Println("Re:twg mm",color.Red(tweet.Id), "\"", color.Cyan(rname), "$message\"")
			if tweet.InReplyToStatusID != 0 {
				fmt.Println("src:twg mm",color.Blue(tweet.InReplyToStatusID), "\"", color.Cyan(rname), "$message\"")
			}
		}
		tweeturl := tweet.Entities.Urls
		if  len(tweeturl) != 0 {
			fmt.Println(color.Blue(tweeturl[0].Expanded_url))
		}
		fmt.Println("---------------------------------")
	}
	return nil
}

func GetUserTimeLineId(c *cli.Context) error {
	name := c.Args().First()
	api := oauth.GetOAuthApi()
	v := url.Values{}
	v.Set("screen_name",name)
	v.Set("count","30")
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
			fmt.Println(color.Cyan(tweet.User.ScreenName), "RT", color.Red(rname), retweet.FullText)
			fmt.Println("Re:twg mm",color.Red(retweet.Id), "\"", color.Cyan(rname), "$message\"")
			if retweet.InReplyToStatusID != 0 {
				fmt.Println("Re(mention):twg mm",color.Blue(retweet.InReplyToStatusID), "\"", color.Cyan(rname), "$message\"")
			}
		} else {
			fmt.Println(color.Cyan(tweet.User.ScreenName), tweet.FullText, tweet.Id, tweet.InReplyToStatusID)
		}
		if  len(tweeturl) != 0 {
			fmt.Println(color.Blue(tweeturl[0].Expanded_url))
		}
	}
	return nil
}

