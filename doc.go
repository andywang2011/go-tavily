// Package tavily is a testable client for the Tavily Search API.
//
// Design constraints:
//   - Configuration via Config (no viper, no global singletons, no log.Fatal).
//   - NewClient returns Client (client_iface.go); use ConfigDefault() for host/timeout defaults.
//   - Injectable HTTP doer via NewClientWithDoer for unit tests.
//   - Bearer token auth against https://api.tavily.com.
//
// Doc: https://docs.tavily.com/api-reference/endpoint/search
package tavily
