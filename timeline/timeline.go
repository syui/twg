package timeline

import (
	"fmt"
	"net/url"
	"gitlab.com/syui/twg/oauth"
)

func GetTimeLine() {
	api := oauth.GetOAuthApi()
	v := url.Values{}
	v.Set("count","10")
	tweets, err := api.GetHomeTimeline(v)
	if err != nil {
	  panic(err)
	}
	for _, tweet := range tweets {
	  fmt.Println(tweet.User.ScreenName, tweet.FullText)
	}
	return
}

