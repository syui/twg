package oauth

import (
	"flag"
	"fmt"
	"log"
	"os"
	"io/ioutil"
	"encoding/json"
	"github.com/ChimeraCoder/anaconda"
	"github.com/skratchdot/open-golang/open"
	"github.com/mrjones/oauth"
	"github.com/bitly/go-simplejson"
	"github.com/urfave/cli/v2"
	"github.com/syui/twg/path"
	"github.com/syui/twg/o"
	//"github.com/hokaccha/go-prettyjson"
)

var ckey string
var cskey string

func IconSetting(c *cli.Context) error {
	var o o.UserVerifyCredentials
	s := c.Args().First()
	file,err := ioutil.ReadFile(path.DirVerify)
	if err != nil {
		fmt.Printf("$ twg oauth")
		RunOAuth()
	}
	js, err := simplejson.NewJson(file)
	if s == "true" {
		fmt.Println("true : ", path.DirVerify)
		js.Set("twg_icon", true)
	} else if s == "false" || s == "f" {
		fmt.Println("delete key ->  twg_icon : ", path.DirVerify)
		js.Del("twg_icon")
	}
	w, err := os.Create(path.DirVerify)
	defer w.Close()
	out, _ := js.EncodePretty()
	w.Write(out)

	files,err := ioutil.ReadFile(path.DirVerify)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(files, &o)
	fmt.Println("check : ", o.TwgIcon)
	return nil
}

func IconSettingCheck() (check bool){
	var o o.UserVerifyCredentials
	file,err := ioutil.ReadFile(path.DirVerify)
	if err != nil {
		fmt.Printf("$ twg oauth")
		RunOAuth()
	}
	json.Unmarshal(file, &o)
	check = o.TwgIcon
	if check == true {
		return check
	} else {
		return
	}
}

func IconSettingCheckCommand() (check bool){
	check = IconSettingCheck()
	if check == true {
		fmt.Println("iterm-mode/check : ", check)
		return check
	} else {
		fmt.Println("iterm-mode/check : false")
		return
	}
}

func IconSettingDeleteCommand() {
	if err := os.RemoveAll(path.DirImg); err != nil {
	    fmt.Println(err)
	}
}

func GetOAuthApi() *anaconda.TwitterApi {
	var o o.Oauth
	anaconda.SetConsumerKey(ckey)
	anaconda.SetConsumerSecret(cskey)

	_, err := os.Stat(path.DirUser)
	if err := os.MkdirAll(path.Dir, os.ModePerm); err != nil {
		panic(err)
	}
	file,err := ioutil.ReadFile(path.DirUser)
	if err != nil {
		fmt.Printf("$ twg oauth")
	}
	json.Unmarshal(file, &o)
	api := anaconda.NewTwitterApi(o.Token, o.Secret)
	return api
}

func RunOAuth() {
	anaconda.SetConsumerKey(ckey)
	anaconda.SetConsumerSecret(cskey)

	_, err := os.Stat(path.DirUser)
	if err := os.MkdirAll(path.Dir, os.ModePerm); err != nil {
		panic(err)
	}
	flag.Parse()
	c := oauth.NewConsumer(
		string(ckey),
		string(cskey),
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		})

	requestToken, u, err := c.GetRequestTokenAndUrl("oob")
	if err != nil {
	    log.Fatal(err)
	}

	fmt.Print("\nopen url: ", string(u))
	fmt.Print("\ninput pin: ")
	open.Run(u)

	verificationCode := ""
	fmt.Scanln(&verificationCode)
	accessToken, err := c.AuthorizeToken(requestToken, verificationCode)
	if err != nil {
		log.Fatal(err)
	}
	outputJSON, err := json.Marshal(&accessToken)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(path.DirUser, outputJSON, os.ModePerm)
	client, err := c.MakeHttpClient(accessToken)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Get(
		"https://api.twitter.com/1.1/account/verify_credentials.json")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bit, err := ioutil.ReadAll(resp.Body)
	ioutil.WriteFile(path.DirVerify, bit, os.ModePerm)
	return
}
