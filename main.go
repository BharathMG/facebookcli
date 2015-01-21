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
		{
			Name:  "me",
			Usage: "Displays your profile info",
			Subcommands: []cli.Command{
				{
					Name:  "info",
					Usage: "Displays your profile information",
					Action: func(c *cli.Context) {
						err := ValidateSession()
						if err != nil {
							fmt.Println(err)
							return
						}
						Me()
					},
				},
				{
					Name:  "feed",
					Usage: "Displays your latest posts",
					Action: func(c *cli.Context) {
						err := ValidateSession()
						if err != nil {
							fmt.Println(err)
							return
						}
						Me()
					},
				},
			},
		},
		{
			Name:  "refresh_access",
			Usage: "Refresh your access token",
			Action: func(c *cli.Context) {
				GetFbAccessToken()
			},
		},
		{
			Name:  "user",
			Usage: "Show user info",
			Action: func(c *cli.Context) {
				err := ValidateSession()
				if err != nil {
					fmt.Println(err)
					return
				}
				if len(c.Args()) == 0 {
					fmt.Println("Specify user. Ex facebookcli user RobPike")
					return
				}
				ShowUser(c.Args()[0])
			},
		},
		{
			Name:  "feed",
			Usage: "Show user info",
			Action: func(c *cli.Context) {
				err := ValidateSession()
				if err != nil {
					fmt.Println(err)
					return
				}
				if len(c.Args()) == 0 {
					fmt.Println("Specify user. Ex facebookcli user RobPike")
					return
				}
				ShowUser(c.Args()[0])
			},
		},
	}
	app.Run(os.Args)
}
