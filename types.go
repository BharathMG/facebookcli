package main

type User struct {
	Id    string
	Bio   string
	Name  string
	Email string
}

type FacebookFeed struct {
	Id          string `facebook:",required"`
	Story       string
	FeedFrom    *FacebookFeedFrom `facebook:"from"`
	CreatedTime string
}

type FacebookFeedFrom struct {
	Name, Id string
}
