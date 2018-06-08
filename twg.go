package main

import (
	"os"
	"github.com/urfave/cli"
	"gitlab.com/syui/twg/oauth"
	"gitlab.com/syui/twg/post"
	"gitlab.com/syui/twg/user"
	"gitlab.com/syui/twg/timeline"
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
	app.Version = "0.1.2"
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
			Usage:   "twg t",
			Action:  func(c *cli.Context) error {
				timeline.GetTimeLine(c)
				//oauth.GetOAuthTimeLine()
				return nil
			},
		},
		{
			Name:    "post",
			Aliases: []string{"p"},
			Usage:   "twg p 'message'",
			Action: func(c *cli.Context) error {
				post.Post(c)
				return nil
			},
		},
		{
			Name:    "user",
			Aliases: []string{"u"},
			Usage:   "twg u '$screen_name'",
			Action:  func(c *cli.Context) error {
				user.User(c)
				return nil
			},
		},
		{
			Name:    "ouath",
			Aliases: []string{"o"},
			Usage:   "twg oauth, ~/$USER/.config/twg",
			Action: func(c *cli.Context) error {
				oauth.RunOAuth()
				return nil
			},
		},
		{
			Name:    "stream",
			Aliases: []string{"s"},
			Usage:   "twg s",
			Action: func(c *cli.Context) error {
				oauth.RunStream()
				return nil
			},
		},
	}
	app.Run(os.Args)
}
