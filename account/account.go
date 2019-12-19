package account

import (
	"fmt"
	"net/url"
	"encoding/json"
	"github.com/syui/twg/o"
	"github.com/syui/twg/oauth"
	"github.com/syui/twg/color"
)

func CheckAccount() (account bool){
	api := oauth.GetOAuthApi()
	account, err := api.VerifyCredentials()
	if err != nil {
		panic(err)
	}
	if account == true {
		fmt.Println("account", color.Blue(account))
	} else if account == false {
		fmt.Println("account", color.Red(account))
	}
	return account
}

func GetAccount() (user string) {
	var o o.Account
	if CheckAccount() == true {
		api := oauth.GetOAuthApi()
		v := url.Values{}
		v.Set("include_entities", "false")
		v.Set("skip_status", "true")
		account, err := api.GetSelf(v)
		if err != nil {
			panic(err)
		}
		fmt.Println(account)
		out, err := json.Marshal(account)
		if err != nil {
			panic(err)
		}
		json.Unmarshal(out, &o)
		user := o.ScreenName
		return user
	} else {
		fmt.Println("account", color.Red(CheckAccount()))
	}
	return
}

func GetAccountSearch(user string) {
		api := oauth.GetOAuthApi()
		user = GetAccount()
		name := "to:" + user
		v := url.Values{}
		v.Set("count", "30")
		searchResult, _ := api.GetSearch(name, v)
		for _ , tweet := range searchResult.Statuses {
			fmt.Println(tweet.Text)
		}
	return
}
