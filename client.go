package tavily

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/andywang2011/go-httplib"
)

type doer interface {
	Do(req *http.Request) (*http.Response, error)
}

// tavilyClient implements Client.
type tavilyClient struct {
	cfg  Config
	doer doer
}

var _ Client = (*tavilyClient)(nil)

// NewClient creates a Tavily client from Config.
func NewClient(cfg Config) (Client, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	return &tavilyClient{
		cfg:  cfg,
		doer: httplib.NewClient(httplib.WithTimeout(cfg.Timeout)),
	}, nil
}

// NewClientWithDoer injects an HTTP doer for tests.
func NewClientWithDoer(cfg Config, d doer) (Client, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	return &tavilyClient{cfg: cfg, doer: d}, nil
}

func (c *tavilyClient) post(ctx context.Context, path string, body, result any) error {
	if c == nil || c.doer == nil {
		return fmt.Errorf("tavily: client is not initialized")
	}
	payload, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("tavily: marshal body: %w", err)
	}
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		strings.TrimSuffix(c.cfg.Host, "/")+path,
		strings.NewReader(string(payload)),
	)
	if err != nil {
		return fmt.Errorf("tavily: new request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.cfg.Secret)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.doer.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bts, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("tavily: read body: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("tavily: HTTP %d: %s", resp.StatusCode, string(bts))
	}
	if err := json.Unmarshal(bts, result); err != nil {
		return fmt.Errorf("tavily: unmarshal: %w", err)
	}
	return nil
}
