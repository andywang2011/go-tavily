package tavily

import (
	"context"
	"fmt"
)

const (
	// SearchPath POST /search
	// Doc: https://docs.tavily.com/api-reference/endpoint/search
	SearchPath = "/search"

	// TopicGeneral searches general web content.
	TopicGeneral = "general"
	// TopicNews restricts results to news sources.
	TopicNews = "news"
)

// SearchRequest is the Tavily search API request body.
type SearchRequest struct {
	// Query is the natural language search query.
	Query string `json:"query,omitempty"`
	// Topic is general or news (TopicGeneral, TopicNews).
	Topic string `json:"topic,omitempty"`
	// SearchDepth is basic or advanced crawl depth.
	SearchDepth string `json:"search_depth,omitempty"`
	// MaxResults limits the number of returned sources.
	MaxResults int `json:"max_results,omitempty"`
	// TimeRange filters by recency (day, week, month, year).
	TimeRange string `json:"time_range,omitempty"`
	// Days limits news results to the last N days when Topic is news.
	Days int `json:"days,omitempty"`
	// IncludeAnswer asks Tavily to synthesize a short answer.
	IncludeAnswer bool `json:"include_answer,omitempty"`
	// IncludeRawContent returns full page content when true.
	IncludeRawContent bool `json:"include_raw_content,omitempty"`
	// IncludeImages returns image URLs when true.
	IncludeImages bool `json:"include_images,omitempty"`
	// IncludeImageDescriptions adds image captions when IncludeImages is true.
	IncludeImageDescriptions bool `json:"include_image_descriptions,omitempty"`
	// IncludeDomains restricts results to these domains.
	IncludeDomains []string `json:"include_domains,omitempty"`
	// ExcludeDomains omits results from these domains.
	ExcludeDomains []string `json:"exclude_domains,omitempty"`
}

// SearchResult is one search hit.
type SearchResult struct {
	// Title is the page title.
	Title string `json:"title"`
	// URL is the result link.
	URL string `json:"url"`
	// Content is the extracted snippet.
	Content string `json:"content"`
	// Score is the relevance score in [0, 1].
	Score float64 `json:"score"`
	// RawContent is the full page text when requested.
	RawContent string `json:"raw_content"`
}

// SearchResponse is the Tavily search API response.
type SearchResponse struct {
	// Query echoes the submitted query.
	Query string `json:"query"`
	// Answer is the optional synthesized answer.
	Answer string `json:"answer"`
	// Results lists matching sources.
	Results []SearchResult `json:"results"`
}

// Search calls POST /search.
func (c *tavilyClient) Search(ctx context.Context, req *SearchRequest) (*SearchResponse, error) {
	if c == nil || c.doer == nil {
		return nil, fmt.Errorf("tavily: client is not initialized")
	}
	if req == nil {
		return nil, fmt.Errorf("tavily: search request is required")
	}
	var resp SearchResponse
	if err := c.post(ctx, SearchPath, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
