package main

type User struct {
	Id    string
	Bio   string
	Name  string
	Email string
}

type FacebookFeed struct {
	Story       string
	FeedFrom    *FacebookFeedFrom `facebook:"from"`
	CreatedTime string
	Message     string
	FeedTo      *FacebookFeedFrom `facebook:"to"`
	StatusType  string
}

type FacebookFeedFrom struct {
	Name, Id string
}
