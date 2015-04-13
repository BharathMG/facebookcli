package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	ct "github.com/daviddengcn/go-colortext"
)

var (
	Me = cli.Command{
		Name:  "me",
		Usage: "Displays your profile info",
		Subcommands: []cli.Command{
			{
				Name:  "info",
				Usage: "Displays your profile information",
				Action: func(c *cli.Context) {
					err := HandleSession()
					if err != nil {
						fmt.Println(err)
						return
					}
					printCurrentUserInfo()
				},
			},
		},
	}
	ShowUser = cli.Command{
		Name:  "user",
		Usage: "Show user info",
		Action: func(c *cli.Context) {
			err := HandleSession()
			if err != nil {
				fmt.Println(err)
				return
			}
			if len(c.Args()) == 0 {
				fmt.Println("Specify user. Ex facebookcli user RobPike")
				return
			}
			printUserInfo(c.Args()[0])
		},
	}
)

func (user *User) printInfo() {
	ct.ChangeColor(ct.Green, false, ct.None, false)
	fmt.Printf("Name : %v", user.Name)
	fmt.Println()
	fmt.Printf("Email : %v", user.Email)
	fmt.Println()
	fmt.Printf("About : %v", user.Bio)
	fmt.Println()
	ct.ResetColor()
}

func print(data []FacebookFeed) {
	fmt.Printf("%#v", data)
}

func printCurrentUserInfo() {
	user := new(User)
	err := FbGet("/me", &user)
	if err != nil {
		fmt.Println(err)
		return
	}
	user.printInfo()
}

func printUserInfo(username string) {
	user := new(User)
	err := FbPublicGet("/"+username, &user)
	if err != nil {
		fmt.Println(err)
		return
	}
	user.printInfo()
}

func ShowWall(args []string) {
	results, err := FbPagingGet("/me/home?fields=status_type,message,from,to")
	feeds := make([]FacebookFeed, len(results))
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, val := range results {
		val.Decode(&feeds[i])
	}
	for _, val := range feeds {
		ct.ChangeColor(ct.Green, false, ct.None, false)
		fmt.Printf("%v : %v From %v", val.StatusType, val.Message, val.FeedFrom.Name)
		fmt.Println()
		ct.ResetColor()
	}

}
