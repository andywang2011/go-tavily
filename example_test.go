package tavily_test

import (
	"context"
	"time"

	"github.com/andywang2011/go-tavily"
)

func ExampleNewClient_search() {
	client, err := tavily.NewClient(tavily.Config{
		Secret:  "tvly-EXAMPLE",
		Timeout: 10 * time.Second,
	})
	if err != nil {
		return
	}

	resp, err := client.Search(context.Background(), &tavily.SearchRequest{
		Query:         "latest Go release",
		Topic:         tavily.TopicGeneral,
		MaxResults:    5,
		IncludeAnswer: true,
	})
	if err != nil {
		return
	}
	_ = resp
}
