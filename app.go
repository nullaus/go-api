package main

import (
	"net/http"

	rest "github.com/ant0ine/go-json-rest/rest"
	"github.com/nullaus/go-api/twitter"
)

const maxResults = 100

func getTwitterRecentSearch(w rest.ResponseWriter, r *rest.Request) {
	handle := r.PathParam("handle")
	tweets, err := twitter.GetRecentTweetsUsingHandle(handle, maxResults)

	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteJson(tweets)
}
