package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"time"

	"github.com/mrjones/oauth"
)

func main() {
	consumerKey := "gkgkgk"
	consumerSecret := "gkgkgk"
	accessToken := "lddld-gkgkg"
	accessTokenSecret := "ldlld"
	postUpdate := false  // "If true, post a status update to the timeline"


	c := oauth.NewConsumer(
		consumerKey,
		consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		})
	c.Debug(false)

	t := oauth.AccessToken{
		Token: accessToken,
		Secret: accessTokenSecret,
	}

	client, err := c.MakeHttpClient(&t)
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Get(
		"https://api.twitter.com/2/users/2244994945/tweets")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	bits, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("The newest item in your home timeline is: " + string(bits))

	if postUpdate {
		status := fmt.Sprintf("Test post via the API using Go (http://golang.org/) at %s", time.Now().String())

		response, err = client.PostForm(
			"https://api.twitter.com/1.1/statuses/update.json",
			url.Values{"status": []string{status}})

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("%v\n", response)
	}
}