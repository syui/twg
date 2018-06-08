package main

import (
	"os"
	"github.com/urfave/cli"
	"gitlab.com/syui/twg/oauth"
	"gitlab.com/syui/twg/post"
	"gitlab.com/syui/twg/user"
	"gitlab.com/syui/twg/timeline"
	"gitlab.com/syui/twg/icon"
	_ "reflect"
)

func Action(c *cli.Context) {
	if c.Args().Get(0) == "" {
		oauth.FirstRunOAuth()
		return
	}
	return
}

func App() *cli.App {
	app := cli.NewApp()
	app.Name = "twg"
	app.Usage = "$ twg"
	app.Version = "0.1.6"
	app.Author = "syui"
	return app
}

func main() {

	app := App()
	app.Action = Action
	app.Commands = []cli.Command{
		{
			Name:    "timeline",
			Aliases: []string{"t"},
			Usage:   "$ twg t, $ twg t 12",
			Action:  func(c *cli.Context) error {
				if oauth.IconSettingCheck() == true {
					icon.ItermGetTimeLine()
				} else {
					timeline.GetTimeLine(c)
				}
				return nil
			},
			Subcommands: cli.Commands{
				cli.Command{
					Name:   "icon",
					Usage:   "$ twg t i (mac iterm)",
					Aliases: []string{"i"},
					Action:  func(c *cli.Context) error {
						icon.ItermGetTimeLine()
						return nil
					},
				},
			},
		},
		{
			Name:    "post",
			Aliases: []string{"p"},
			Usage:   "$ twg p 'message'",
			Action: func(c *cli.Context) error {
				post.Post(c)
				return nil
			},
		},
		{
			Name:    "user",
			Aliases: []string{"u"},
			Usage:   "$ twg u '$screen_name'",
			Action:  func(c *cli.Context) error {
				if oauth.IconSettingCheck() == true {
					icon.ItermUser(c)
				} else {
					user.User(c)
				}
				return nil
			},
			Subcommands: cli.Commands{
				cli.Command{
					Name:   "icon",
					Usage:   "$ twg u i (mac iterm)",
					Aliases: []string{"i"},
					Action:  func(c *cli.Context) error {
						icon.ItermUser(c)
						return nil
					},
				},
			},
		},
		{
			Name:    "ouath",
			Aliases: []string{"o"},
			Usage:   "$ twg oauth, ~/$USER/.config/twg",
			Action: func(c *cli.Context) error {
				oauth.RunOAuth()
				return nil
			},
		},
		{
			Name:    "stream",
			Aliases: []string{"s"},
			Usage:   "$ twg s",
			Action: func(c *cli.Context) error {
				if oauth.IconSettingCheck() == true {
					icon.ItermRunStream()
				} else {
					oauth.RunStream()
				}
				return nil
			},
			Subcommands: cli.Commands{
				cli.Command{
					Name:   "icon",
					Usage:   "$ twg s i (mac iterm)",
					Aliases: []string{"i"},
					Action:  func(c *cli.Context) error {
						icon.ItermRunStream()
						return nil
					},
				},
			},
		},
		{
			Name:    "setting",
			Aliases: []string{"set"},
			Usage:   "$ twg set true",
			Action:  func(c *cli.Context) error {
				oauth.IconSetting(c)
				return nil
			},
			Subcommands: cli.Commands{
				cli.Command{
					Name:   "check",
					Usage:   "$ twg set c",
					Aliases: []string{"c"},
					Action:  func(c *cli.Context) error {
						oauth.IconSettingCheckCommand()
						return nil
					},
				},
			},
		},

	}
	app.Run(os.Args)
}
