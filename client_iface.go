package tavily

import "context"

// Client is the Tavily API client surface (mock-friendly).
type Client interface {
	// Search performs a Tavily web search for the given request.
	Search(ctx context.Context, req *SearchRequest) (*SearchResponse, error)
}

// API is an alias for Client for backward compatibility.
type API = Client
