package main

import (
	"net/http"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/go-chi/chi/v5"
)

const (
	appID     = "appID"
	apiKey    = "apiKey"
	indexName = "indexName"
)

func main() {
	client := search.NewClient(appID, apiKey)
	_ = client.InitIndex(indexName)

	r := chi.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})
	http.ListenAndServe(":8080", r)
}
