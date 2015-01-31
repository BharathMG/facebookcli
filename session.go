package main

import (
	"errors"
	"fmt"
	fb "github.com/huandu/facebook"
	"github.com/skratchdot/open-golang/open"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var access_token string
var done chan bool
var session *fb.Session

var (
	CLIENT_ID     = os.Getenv("FB_CLIENT_ID")
	CLIENT_SECRET = os.Getenv("FB_CLIENT_SECRET")
)

const (
	REDIRECT_URI        = "http://localhost:3000/"
	FB_DIALOG_AUTH_URL  = "https://www.facebook.com/dialog/oauth?client_id=%v&redirect_uri=%v&scope=%v"
	FB_ACCESS_TOKEN_URL = "https://graph.facebook.com/oauth/access_token?client_id=%v&client_secret=%v&redirect_uri=%v&code=%v"
)

func ValidateSession() error {
	err := checkSession()
	if err != nil {
		fmt.Println(err)
		GetFbAccessToken()
	}
	return err
}

// Pluck code from Facebook Redirect request and fetch Access Token with the code.
func AccessTokenHandler(w http.ResponseWriter, req *http.Request) {
	code := req.URL.Query().Get("code")
	if code == "" {
		fmt.Println("Access code could not be obtained. Please raise an issue in github.com/BharathMG/facebookcli")
		done <- true
		return
	}
	url := fmt.Sprintf(FB_ACCESS_TOKEN_URL, CLIENT_ID, CLIENT_SECRET, REDIRECT_URI, code)
	resp, err := http.Get(url)
	if err != nil {
		println(err)
		done <- true
		return
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	if strings.Contains(bodyString, "access_token") {
		temp := strings.Split(bodyString, "=")[1]
		access_token = strings.Split(temp, "&")[0]
		fmt.Println("Please do the following.")
		fmt.Println("1) Run this in your terminal. 'export ACCESS_TOKEN_FBCLI=" + access_token + "'")
		fmt.Println()
		fmt.Println("2) You can copy the above export command to your bash_profile to preserve your facebook session.")
		fmt.Println()
		fmt.Println("3) All set! Please verify if env variable is set by 'echo $ACCESS_TOKEN_FBCLI' in your terminal. It should display the access token.")
		fmt.Println()
		fmt.Println("Now you can rerun the facebookcli command!")
		w.Write([]byte("You are all set. Please check your terminal!"))
		done <- true
		return
	} else {
		fmt.Println("Access code could not be obtained. Please raise an issue in github.com/BharathMG/facebookcli")
		done <- true
		return
	}
}

func validateToken() error {
	access_token := os.Getenv("ACCESS_TOKEN_FBCLI")
	globalApp := fb.New(CLIENT_ID, CLIENT_SECRET)
	session = globalApp.Session(access_token)
	return session.Validate()
}

func checkSession() error {
	if os.Getenv("ACCESS_TOKEN_FBCLI") != "" {
		return validateToken()
	}
	return errors.New("Access Token not set")
}

func GetFbAccessToken() {
	done = make(chan bool)
	permissions := []string{"read_stream", "user_about_me", "email"}
	scope := strings.Join(permissions, ",")
	go open.Run(fmt.Sprintf(FB_DIALOG_AUTH_URL, CLIENT_ID, REDIRECT_URI, scope))
	go func() {
		http.HandleFunc("/", AccessTokenHandler)
		http.HandleFunc("/favicon.ico", http.NotFound)
		http.ListenAndServe(":3000", nil)
	}()
	<-done
}

func GetSession() *fb.Session {
	return session
}
