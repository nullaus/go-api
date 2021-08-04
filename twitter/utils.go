package twitter

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/g8rswimmer/go-twitter"
)

func printTweetError(tweetErr *twitter.TweetErrorResponse) error {
	enc, err := json.MarshalIndent(tweetErr, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Errorf(string(enc))
}
