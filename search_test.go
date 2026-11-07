package tavily

import (
	"context"
	"testing"
)

func TestSearch_TopicConstants(t *testing.T) {
	if TopicGeneral != "general" || TopicNews != "news" {
		t.Error("topic constants mismatch")
	}
	if SearchPath != "/search" {
		t.Errorf("SearchPath = %q", SearchPath)
	}
}

func TestSearchRequest_Fields(t *testing.T) {
	r := &SearchRequest{Query: "test", Topic: TopicGeneral, MaxResults: 5}
	if r.Query != "test" || r.Topic != TopicGeneral || r.MaxResults != 5 {
		t.Errorf("SearchRequest mismatch: %+v", r)
	}
}

func TestSearchResponse_Fields(t *testing.T) {
	cfg := Config{Host: "https://a.com", Secret: "s"}
	c, err := NewClientWithDoer(cfg, &mockDoer{body: `{"query":"q","answer":"a","results":[{"title":"t","url":"u","content":"c","score":0.9,"raw_content":""}]}`})
	if err != nil {
		t.Fatal(err)
	}
	resp, err := c.Search(context.Background(), &SearchRequest{Query: "test", Topic: TopicGeneral})
	if err != nil {
		t.Fatalf("Search err = %v", err)
	}
	if resp.Query != "q" || resp.Answer != "a" || len(resp.Results) != 1 || resp.Results[0].Title != "t" {
		t.Errorf("SearchResponse mismatch: %+v", resp)
	}
}
