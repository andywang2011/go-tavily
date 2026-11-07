package tavily

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

const searchOKBody = `{"query":"q","answer":"a","results":[{"title":"t","url":"u","content":"c","score":0.9,"raw_content":""}]}`

type mockDoer struct {
	body       string
	statusCode int
	err        error
	lastReq    *http.Request
}

func (m *mockDoer) Do(req *http.Request) (*http.Response, error) {
	m.lastReq = req
	if m.err != nil {
		return nil, m.err
	}
	if m.statusCode == 0 {
		m.statusCode = http.StatusOK
	}
	return &http.Response{
		StatusCode: m.statusCode,
		Body:       io.NopCloser(strings.NewReader(m.body)),
		Header:     make(http.Header),
	}, nil
}

func TestNewClient(t *testing.T) {
	cfg := Config{Host: "https://a.com", Secret: "s", Timeout: 10 * time.Second}
	c, err := NewClient(cfg)
	if err != nil {
		t.Fatalf("NewClient err = %v", err)
	}
	if c == nil {
		t.Fatal("expected non-nil client")
	}
}

func TestNewClient_InvalidConfig(t *testing.T) {
	_, err := NewClient(Config{Host: "https://a.com"})
	if err == nil {
		t.Fatal("expected error when secret missing")
	}
}

func TestNewClientWithDoer(t *testing.T) {
	cfg := Config{Host: "https://a.com", Secret: "s"}
	m := &mockDoer{body: searchOKBody}
	c, err := NewClientWithDoer(cfg, m)
	if err != nil {
		t.Fatalf("NewClientWithDoer err = %v", err)
	}
	if c == nil {
		t.Fatal("expected non-nil client")
	}
}

func TestClient_Search_NilReceiver(t *testing.T) {
	var tc *tavilyClient
	var c Client = tc
	_, err := c.Search(context.Background(), &SearchRequest{Query: "q"})
	if err == nil {
		t.Error("expected error when receiver is nil")
	}
}

func TestClient_Search_NilReq(t *testing.T) {
	cfg := Config{Host: "https://a.com", Secret: "s"}
	c, err := NewClientWithDoer(cfg, &mockDoer{})
	if err != nil {
		t.Fatal(err)
	}
	_, err = c.Search(context.Background(), nil)
	if err == nil {
		t.Error("expected error for nil request")
	}
}

func TestClient_Search_Success(t *testing.T) {
	cfg := Config{Host: "https://a.com", Secret: "s"}
	m := &mockDoer{body: searchOKBody}
	c, err := NewClientWithDoer(cfg, m)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := c.Search(context.Background(), &SearchRequest{Query: "test"})
	if err != nil {
		t.Fatalf("Search err = %v", err)
	}
	if resp.Query != "q" || resp.Answer != "a" || len(resp.Results) != 1 {
		t.Errorf("resp mismatch: %+v", resp)
	}
	if m.lastReq == nil || !strings.Contains(m.lastReq.Header.Get("Authorization"), "Bearer s") {
		t.Error("expected Bearer authorization header")
	}
}

func TestClient_Search_HTTPError(t *testing.T) {
	cfg := Config{Host: "https://a.com", Secret: "s"}
	c, err := NewClientWithDoer(cfg, &mockDoer{statusCode: 401, body: "unauthorized"})
	if err != nil {
		t.Fatal(err)
	}
	_, err = c.Search(context.Background(), &SearchRequest{Query: "q"})
	if err == nil {
		t.Error("expected error for HTTP 401")
	}
}

func TestClient_Search_DoError(t *testing.T) {
	cfg := Config{Host: "https://a.com", Secret: "s"}
	c, err := NewClientWithDoer(cfg, &mockDoer{err: errors.New("do error")})
	if err != nil {
		t.Fatal(err)
	}
	_, err = c.Search(context.Background(), &SearchRequest{Query: "q"})
	if err == nil {
		t.Error("expected error when Do fails")
	}
}

func TestClient_Search_InvalidJSON(t *testing.T) {
	cfg := Config{Host: "https://a.com", Secret: "s"}
	c, err := NewClientWithDoer(cfg, &mockDoer{body: "not json"})
	if err != nil {
		t.Fatal(err)
	}
	_, err = c.Search(context.Background(), &SearchRequest{Query: "q"})
	if err == nil || !strings.Contains(err.Error(), "unmarshal") {
		t.Errorf("expected unmarshal error, got %v", err)
	}
}
