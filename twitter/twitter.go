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

type ResponseRecentTweets struct {
	Latest   twitter.TweetLookup   `json:"latest"`
	New      []twitter.TweetLookup `json:"new"`
	Previous []twitter.TweetLookup `json:"previous"`
}

// GetRecentTweetsUsingHandle will return the recent tweets for a user handle (`handle`).
func GetRecentTweetsUsingHandle(handle string, maxResults int) (*ResponseRecentTweets, error) {
	flag.Parse()

	tweet := &twitter.Tweet{
		Authorizer: APIAuthorize{
			Token: utils.MustGetEnv(EnvTwitterBearerToken),
		},
		Client: http.DefaultClient,
		Host:   "https://api.twitter.com",
	}
	fieldOpts := twitter.TweetFieldOptions{
		Expansions: []twitter.Expansion{twitter.ExpansionAuthorID},
		TweetFields: []twitter.TweetField{
			twitter.TweetFieldAttachments,
			twitter.TweetFieldCreatedAt,
			twitter.TweetFieldConversationID,
			twitter.TweetFieldLanguage,
			twitter.TweetFieldPublicMetrics,
		},
		UserFields: []twitter.UserField{
			twitter.UserFieldName,
			twitter.UserFieldLocation,
			twitter.UserFieldProfileImageURL,
			twitter.UserFieldPublicMetrics,
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

	response := &ResponseRecentTweets{
		Latest: recentSearch.LookUps[recentSearch.Meta.NewestID],
	}

	return response, err
}

func StoreRecentTweetsUsingHandle(handle string, recent *twitter.TweetRecentSearch) error {
	return nil
}
