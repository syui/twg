package main

import (
	"os"
	"github.com/urfave/cli/v2"
	"github.com/syui/twg/cmd"
)

func App() *cli.App {
	app := cli.NewApp()
	app.Name = "twg"
	app.Usage = "$ twg"
	return app
}

func Action(c *cli.Context) error {
	if c.Args().Get(0) == "" {
		cmd.Action(c)
	}
	return nil
}

func main() {
	app := &cli.App{
		Version: "0.4.6",
		Name: "twg",
		Usage: "$ twg #timeline",
		Action: func(c *cli.Context) error {
			cmd.Timeline(c)
			return nil
		},
	}
	app.Commands = []*cli.Command {
		{
			Name:    "account",
			Aliases: []string{"a"},
			Usage:   "$ twg a",
			Action:  func(c *cli.Context) error {
				cmd.Account(c)
				return nil
			},
		},
		{
			Name:    "timeline",
			Aliases: []string{"t"},
			Usage:   "$ twg t, $ twg t 12",
			Action:  func(c *cli.Context) error {
				cmd.Timeline(c)
				return nil
			},
		},
		{
			Name:    "post",
			Aliases: []string{"p"},
			Usage:   "$ twg p 'message'",
			Action: func(c *cli.Context) error {
				cmd.Post(c)
				return nil
			},
		},
		{
			Name:    "mention",
			Aliases: []string{"m"},
			HideHelp:        false,
			Usage:   "$ twg m",
			Action: func(c *cli.Context) error {
				cmd.MentionTL(c)
				return nil
			},
			Subcommands: []*cli.Command {
				&cli.Command {
					Name:   "user tweet",
					Usage:   "$ twg m u",
					Aliases: []string{"u"},
					Action:  func(c *cli.Context) error {
						cmd.MentionUser(c)
						return nil
					},
				},
				&cli.Command {
					Name:   "notify tweet",
					Usage:   "$ twg m n",
					Aliases: []string{"n"},
					Action:  func(c *cli.Context) error {
						cmd.MentionNotify(c)
						return nil
					},
				},
				&cli.Command {
					Name:   "TL tweet",
					Usage:   "$ twg m t",
					Aliases: []string{"t"},
					Action:  func(c *cli.Context) error {
						cmd.MentionTL(c)
						return nil
					},
				},
			},
		},
		{
			Name:    "mm",
			Aliases: []string{"mm"},
			Usage:   "$ twg mm $tweet_rep_id '@user message';# @userをつけないとmentionにならないので注意",
			Action: func(c *cli.Context) error {
				cmd.Mm(c)
				return nil
			},
		},
		{
			Name:    "img",
			Aliases: []string{"i"},
			Usage:   "$ twg i ~/img/file.jpg 'message'",
			Action: func(c *cli.Context) error {
				cmd.ImgPost(c)
				return nil
			},
		},
		{
			Name:    "fav",
			Aliases: []string{"f"},
			Usage:   "$ twg f $tweet_id",
			Action: func(c *cli.Context) error {
				cmd.Fav(c)
				return nil
			},
		},
		{
			Name:    "retweet",
			Aliases: []string{"r"},
			Usage:   "$ twg r $tweet_id",
			Action: func(c *cli.Context) error {
				cmd.Ret(c)
				return nil
			},
		},
		{
			Name:    "delete",
			Aliases: []string{"d"},
			Usage:   "$ twg d $tweet_id",
			Action: func(c *cli.Context) error {
				cmd.Del(c)
				return nil
			},
		},
		{
			Name:    "search",
			Aliases: []string{"/"},
			Usage:   "$ twg / #twitter",
			Action: func(c *cli.Context) error {
				cmd.Search(c)
				return nil
			},
		},
		{
			Name:    "user",
			Aliases: []string{"u"},
			Usage:   "$ twg u '$screen_name'",
			Action:  func(c *cli.Context) error {
				cmd.User(c)
				return nil
			},
		},
		{
			Name:    "ouath",
			Aliases: []string{"o"},
			Usage:   "$ twg oauth, ~/$USER/.config/twg",
			Action: func(c *cli.Context) error {
				cmd.Oauth()
				return nil
			},
		},
		{
			Name:    "notify",
			Aliases: []string{"n"},
			Usage:   "$ twg n",
			Action:  func(c *cli.Context) error {
				cmd.Notify(c)
				return nil
			},
		},
		{
			Name:    "stream",
			Aliases: []string{"s"},
			HideHelp:        false,
			Usage:   "$ twg s",
			Action: func(c *cli.Context) error {
				o := string("normal")
				cmd.Stream(c, o)
				return nil
			},
			Subcommands: []*cli.Command {
				&cli.Command {
					Name:   "user",
					Usage:   "$ twg s u",
					Aliases: []string{"u"},
					Action:  func(c *cli.Context) error {
						o := string("user")
						cmd.Stream(c, o)
						return nil
					},
				},
				&cli.Command {
					Name:   "site",
					Usage:   "$ twg s s",
					Aliases: []string{"s"},
					Action:  func(c *cli.Context) error {
						o := string("site")
						cmd.Stream(c, o)
						return nil
					},
				},
				&cli.Command {
					Name:   "public",
					Usage:   "$ twg s p",
					Aliases: []string{"p"},
					Action:  func(c *cli.Context) error {
						o := string("public")
						cmd.Stream(c, o)
						return nil
					},
				},
			},
		},
		{
			Name:    "setting",
			Aliases: []string{"set"},
			Usage:   "$ twg set true/false",
			Action:  func(c *cli.Context) error {
				//cli.ShowSubcommandHelp(c)
				cmd.SettingMain(c)
				return nil
			},
			Subcommands: []*cli.Command {
				&cli.Command{
					Name:   "check",
					Usage:   "$ twg set c",
					Aliases: []string{"c"},
					Action:  func(c *cli.Context) error {
						cmd.SettingCheck(c)
						return nil
					},
				},
				&cli.Command {
					Name:   "delete",
					Usage:   "$ twg set delete # user icon clean",
					Aliases: []string{"d"},
					Action:  func(c *cli.Context) error {
						cmd.SettingDelete(c)
						return nil
					},
				},
			},
		},
	}
	app.Run(os.Args)
}
