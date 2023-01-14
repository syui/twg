package cmd

import (
	"os"
	"github.com/urfave/cli/v2"
	"github.com/syui/twg/user"
	"github.com/syui/twg/post"
	"github.com/syui/twg/postp"
	"github.com/syui/twg/timeline"
	"github.com/syui/twg/icon"
	"github.com/syui/twg/oauth"
	"github.com/syui/twg/path"
	"github.com/syui/twg/stream"
	"github.com/syui/twg/account"
	"github.com/syui/twg/notify"
	"github.com/syui/twg/search"
	"github.com/syui/twg/mention"
	"github.com/syui/twg/mm"
	"github.com/syui/twg/fav"
	"github.com/syui/twg/ret"
	"github.com/syui/twg/del"
	"github.com/syui/twg/img"
	"github.com/syui/twg/imgp"
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

func Postp(c *cli.Context) error {
	postp.Postp(c)
	return nil
}

func ImgPost(c *cli.Context) error {
	imgpost.ImgPost(c)
	return nil
}

func ImgPostp(c *cli.Context) error {
	imgpostp.ImgPostp(c)
	return nil
}

func Fav(c *cli.Context) error {
	fav.Fav(c)
	return nil
}

func Ret(c *cli.Context) error {
	ret.Ret(c)
	return nil
}

func Del(c *cli.Context) error {
	del.Del(c)
	return nil
}

func MentionNotify(c *cli.Context) error {
	mention.GetMentionId(c)
	return nil
}

func MentionTL(c *cli.Context) error {
	mention.GetTimeLineId(c)
	return nil
}

func MentionUser(c *cli.Context) error {
	mention.GetUserTimeLineId(c)
	return nil
}

func Mm(c *cli.Context) error {
	mm.Mm(c)
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
