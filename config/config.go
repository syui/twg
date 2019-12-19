package config

import (
	"fmt"
	"flag"
	"net/url"
	"io/ioutil"
	"encoding/json"
	"github.com/ChimeraCoder/anaconda"
)

type ApiConf struct {
	ConsumerKey       string `json:"consumer_key"`
	ConsumerSecret    string `json:"consumer_secret"`
	AccessToken       string `json:"access_token"`
	AccessTokenSecret string `json:"access_token_secret"`
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetTwitterApi() *anaconda.TwitterApi {
	var apiConf ApiConf
	{
		apiConfPath := flag.String("conf", "config.json", "API Config File")
		flag.Parse()
		data, err_file := ioutil.ReadFile(*apiConfPath)
		Check(err_file)
		err_json := json.Unmarshal(data, &apiConf)
		Check(err_json)
	}
	anaconda.SetConsumerKey(apiConf.ConsumerKey)
	anaconda.SetConsumerSecret(apiConf.ConsumerSecret)
	api := anaconda.NewTwitterApi(apiConf.AccessToken, apiConf.AccessTokenSecret)
	return api
}

func GetTimeLine() {
	api := GetTwitterApi()
	v := url.Values{}
	v.Set("count","10")
	tweets, err := api.GetHomeTimeline(v)
	if err != nil {
	  panic(err)
	}
	for _, tweet := range tweets {
	  fmt.Println(tweet.User.ScreenName, tweet.Text)
	}
	return
}
