package imgpostp

import (
	"os"
	"fmt"
	"net/url"
	"encoding/base64"
	"github.com/urfave/cli/v2"
	"github.com/syui/twg/oauth"
)

func ImgPostp(c *cli.Context) error {
	if c.NArg() == 0 {
		fmt.Println("twg ii ~/file/img.jpg $tw_id")
		os.Exit(0)
	}

	f := c.Args().First()
	file, _ := os.Open(f)
	defer file.Close()
	fi, _ := file.Stat()
	size := fi.Size()
	data := make([]byte, size)
	file.Read(data)
	base64String := base64.StdEncoding.EncodeToString(data)

	api := oauth.GetOAuthApi()
	id := c.Args().Get(1)
	mes := c.Args().Get(2)
	media, _ := api.UploadMedia(base64String)
	v := url.Values{}
	v.Add("in_reply_to_status_id", id)
	v.Add("media_ids", media.MediaIDString)
	v.Set("tweet_mode", "extended")
	tweet, err := api.PostTweet(mes,v)
	if err != nil {
		panic(err)
	}
	fmt.Println(tweet.Id)
	return nil
}
