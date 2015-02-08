package main

import (
	"fmt"
	ct "github.com/daviddengcn/go-colortext"
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

func Me() {
	user := new(User)
	err := FbGet("/me", &user)
	if err != nil {
		fmt.Println(err)
		return
	}
	user.printInfo()
}

func ShowUser(username string) {
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
