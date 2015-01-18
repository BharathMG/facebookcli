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
