package main

import (
	"fmt"
	"io/ioutil"
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
		true,
	)
	if err != nil {
		fmt.Println(err)
	}

	// response, err := client.GetUserTweets("4284141501")
	response, err := client.CreateTweet("This is a test tweet")

	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	bits, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("The newest item in your home timeline is: " + string(bits))
}

// IMPORTANT: go mod edit -replace github.com/arshamalh/twigo=../twigo
// go mod tidy
