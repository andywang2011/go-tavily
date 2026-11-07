# go-tavily

Testable Go client for the [Tavily Search API](https://docs.tavily.com/).

## Install

```bash
go get github.com/andywang2011/go-tavily@v0.1.0
```

## Quick start

```go
client, err := tavily.NewClient(tavily.Config{
	Secret:  "tvly-EXAMPLE",
	Timeout: 10 * time.Second,
})
if err != nil {
	// handle
}

resp, err := client.Search(context.Background(), &tavily.SearchRequest{
	Query:         "latest Go release",
	Topic:         tavily.TopicGeneral,
	MaxResults:    5,
	IncludeAnswer: true,
})
_ = resp
```

## Examples

```bash
cd go-tavily/examples/search && go run .
```

## License

MIT — see [LICENSE](LICENSE).
