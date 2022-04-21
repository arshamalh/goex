package main

import (
	"fmt"
	"os"

	"github.com/arshamalh/twigo"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	client, err := twigo.NewClient(
		os.Getenv("ConsumerKey"),
		os.Getenv("ConsumerSecret"),
		os.Getenv("AccessToken"),
		os.Getenv("AccessTokenSecret"),
		os.Getenv("BearerToken"),
	)
	if err != nil {
		fmt.Println(err)
	}
	// "This is a test tweet" id => 1516784368601153548
	// @arshamalh => 1216345203453452289
	// @arshamsishemi => 1466354024274345986

	// response, err := client.GetUserTweets(
	// 	"1216345203453452289",
	// 	false,
	// 	map[string]interface{}{
	// 		"max_results": 5,
	// 	},
	// )
	// response, err := client.CreateTweet("This is a test tweet", nil)
	// response, err := client.Retweet("1431751228145426438")
	// response, err := client.Like("1431751228145426438")
	// client.UnRetweet("1516784368601153548")
	// client.Unlike("1516784368601153548")
	// response, err := client.GetUsersByUsernames([]string{"arshamalh"}, false, nil)
	response, err := client.GetLikingUsers(
		"1431751228145426438",
		false,
		map[string]interface{}{
			"max_results": 5,
		})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", response)
}

// IMPORTANT: go mod edit -replace github.com/arshamalh/twigo=../twigo
// go mod tidy
