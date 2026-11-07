// Tavily search with IncludeRawContent (different from README quick start).
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/andywang2011/go-tavily"
)

func main() {
	secret := os.Getenv("TAVILY_API_KEY")
	if secret == "" {
		log.Fatal("set TAVILY_API_KEY")
	}
	client, err := tavily.NewClient(tavily.Config{
		Secret:  secret,
		Timeout: 15 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Search(context.Background(), &tavily.SearchRequest{
		Query:              "site reliability practices",
		Topic:              tavily.TopicGeneral,
		MaxResults:         3,
		IncludeRawContent:  true,
		SearchDepth:        "advanced",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("answer=%q results=%d\n", resp.Answer, len(resp.Results))
}
