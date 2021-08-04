package twitter

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net/http"

	"github.com/g8rswimmer/go-twitter"
	_ "github.com/joho/godotenv/autoload"
	"github.com/nullaus/go-api/utils"
)

const (
	EnvTwitterBearerToken = "TWITTER_BEARER_TOKEN"
)

// APIAuthorize will set the Bearer Token
type APIAuthorize struct {
	Token string
}

func (a APIAuthorize) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}

// GetRecentTweetsUsingHandle will return the recent tweets for a user handle (`handle`).
func GetRecentTweetsUsingHandle(handle string, maxResults int) (*twitter.TweetRecentSearch, error) {
	flag.Parse()

	tweet := &twitter.Tweet{
		Authorizer: APIAuthorize{
			Token: utils.MustGetEnv(EnvTwitterBearerToken),
		},
		Client: http.DefaultClient,
		Host:   "https://api.twitter.com",
	}
	fieldOpts := twitter.TweetFieldOptions{
		TweetFields: []twitter.TweetField{
			twitter.TweetFieldCreatedAt, twitter.TweetFieldPublicMetrics, twitter.TweetFieldConversationID, twitter.TweetFieldLanguage,
		},
	}
	searchOpts := twitter.TweetRecentSearchOptions{
		MaxResult: maxResults,
	}

	recentSearch, err := tweet.RecentSearch(context.Background(), "from:"+handle, searchOpts, fieldOpts)
	var tweetErr *twitter.TweetErrorResponse
	if errors.As(err, &tweetErr) {
		return nil, printTweetError(tweetErr)
	}
	if err != nil {
		return nil, err
	}

	return recentSearch, err
}
