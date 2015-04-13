package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Facebook CLI"
	app.Usage = "Command Line Interface for Facebook"
	app.Author = "Email = mgbharath@ymail.com Github = BharathMG"
	app.Action = func(c *cli.Context) {
	}

	app.Commands = []cli.Command{
		Me,
		RefreshAccessToken,
		ShowUser,
		{
			Name:  "wall",
			Usage: "Show my timeline posts",
			Action: func(c *cli.Context) {
				err := HandleSession()
				if err != nil {
					fmt.Println(err)
					return
				}
				ShowWall(c.Args())
			},
		},
	}
	app.Run(os.Args)
}
