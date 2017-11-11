package main

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/go-playground/webhooks.v3"
	"gopkg.in/go-playground/webhooks.v3/github"
)

const (
	path = "/webhooks"
	port = 3016
)

var comments map[string]int

func main() {
	hook := github.New(&github.Config{Secret: os.Getenv("SECRET")})
	hook.RegisterEvents(HandlePullRequestReviewCommentEvent, github.PullRequestReviewCommentEvent)

	comments = make(map[string]int)
	err := webhooks.Run(hook, ":"+strconv.Itoa(port), path)
	if err != nil {
		fmt.Println(err)
	}
}

// HandlePullRequestReviewCommentEvent handles GitHub pull_request events
func HandlePullRequestReviewCommentEvent(payload interface{}, header webhooks.Header) {
	fmt.Println("Handling Pull Request")

	pl := payload.(github.PullRequestReviewCommentPayload)
	user := pl.Sender.Login
	comments[user]++

	// Do whatever you want from here...
	fmt.Printf("%+v", comments)
}
