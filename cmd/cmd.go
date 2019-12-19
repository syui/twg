package cmd

import (
	"os"
	"github.com/urfave/cli/v2"
	"github.com/syui/twg/user"
	"github.com/syui/twg/post"
	"github.com/syui/twg/timeline"
	"github.com/syui/twg/icon"
	"github.com/syui/twg/oauth"
	"github.com/syui/twg/path"
	"github.com/syui/twg/stream"
	"github.com/syui/twg/account"
	"github.com/syui/twg/notify"
	"github.com/syui/twg/search"
)

func Action(c *cli.Context) error {
	if oauth.IconSettingCheck() == true {
		icon.FirstItermCommand()
	} else {
		_, err := os.Stat(path.DirUser)
		if err != nil {
			oauth.RunOAuth()
		} else {
			timeline.GetTimeLine(c)
		}
	}
	return nil
}

func Account(c *cli.Context) error {
	account.GetAccount()
	return nil
}

func Notify(c *cli.Context) error {
	notify.GetNotify(c)
	return nil
}

func Timeline(c *cli.Context) error {
	if oauth.IconSettingCheck() == true {
		icon.ItermGetTimeLineOption(c)
	} else {
		timeline.GetTimeLine(c)
	}
	return nil
}

func User(c *cli.Context) error {
	if oauth.IconSettingCheck() == true {
		icon.ItermUser(c)
	} else {
		user.User(c)
	}
	return nil
}

func Oauth() {
	oauth.RunOAuth()
	return
}

func Stream(c *cli.Context, o string) error {
	if oauth.IconSettingCheck() == true {
		icon.ItermRunStream()
	} else {
		stream.RunStream(c, o)
	}
	return nil
}

func Post(c *cli.Context) error {
	post.Post(c)
	return nil
}

func Search(c *cli.Context) error {
	search.Search(c)
	return nil
}

func SettingMain(c *cli.Context) error {
	oauth.IconSetting(c)
	return nil
}

func SettingCheck(c *cli.Context) error {
	oauth.IconSettingCheckCommand()
	return nil
}

func SettingDelete(c *cli.Context) error {
	oauth.IconSettingDeleteCommand()
	return nil
}
