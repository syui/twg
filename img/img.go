package imgpost

import (
	"os"
	"fmt"
	"net/url"
	"encoding/base64"
	"github.com/urfave/cli/v2"
	"github.com/syui/twg/oauth"
	"github.com/syui/twg/color"
)


func ImgPost(c *cli.Context) error {

	f := c.Args().First()
	file, _ := os.Open(f)
	defer file.Close()
	fi, _ := file.Stat()
	size := fi.Size()
	data := make([]byte, size)
	file.Read(data)
	base64String := base64.StdEncoding.EncodeToString(data)

	if c.NArg() > 0 {
		api := oauth.GetOAuthApi()
		mes := c.Args().Get(1)
		media, _ := api.UploadMedia(base64String)
		v := url.Values{}
		v.Add("media_ids", media.MediaIDString)
		v.Set("tweet_mode", "extended")
		tweet, err := api.PostTweet(mes, v)
		if err != nil {
			panic(err)
		}
		fmt.Println(color.Cyan(tweet.User.ScreenName), tweet.FullText)
	} else {
		fmt.Println("twg i 'message'")
	}
	return nil
}
