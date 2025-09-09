package httpclient

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/coin50etf/coin-market/internal/pkg/config"
	"github.com/coin50etf/coin-market/internal/pkg/utils/stringutils"
)

const (
	ContentTypeHeader   = "Content-Type"
	AuthorizationHeader = "Authorization"
	ContentTypeJSON     = "application/json"

	logLevelDebug = "DEBUG"
)

var (
	sharedTransport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          500,
		MaxIdleConnsPerHost:   200,
		IdleConnTimeout:       120 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	sharedHTTPClient = &http.Client{
		Timeout:   120 * time.Second,
		Transport: sharedTransport,
	}
)

// Client is a wrapper around http.Client to enforce best practices, like timeout and context usage.
type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: sharedHTTPClient,
	}
}

// DefaultClient creates a new instance of Client with a default request timeout.
func DefaultClient() *Client {
	return NewClient()
}

// Do send an HTTP request and returns an HTTP response, handling context-related cancellation or deadline exceeded errors.
// It automatically handles requests with a given `context.Context`.
func (c *Client) Do(ctx context.Context, method, url string, body io.Reader, header map[string]string) (*http.Response, error) {
	start := time.Now()
	if config.Conf.Log.Level == logLevelDebug {
		log.Println("Sending request to", url)
	}
	// Create an HTTP request with the provided context
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	for k, v := range header {
		req.Header.Set(k, v)
	}

	// Set default Content-Type header if not provided
	if body != nil && req.Header.Get(ContentTypeHeader) == "" {
		req.Header.Set(ContentTypeHeader, ContentTypeJSON)
	}

	if config.Conf.Log.Level == logLevelDebug {
		requestDump, _ := httputil.DumpRequestOut(req, true)
		log.Printf("[HTTP REQUEST] ➜ \n%s", requestDump)
	}

	// Send the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		// Handle specific context-related errors
		if errors.Is(err, context.Canceled) {
			return nil, fmt.Errorf("request canceled due to context cancellation: %w", err)
		}
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, fmt.Errorf("request failed due to context deadline exceeded: %w", err)
		}
		return nil, fmt.Errorf("failed to send HTTP request: %w", err)
	}

	if config.Conf.Log.Level == logLevelDebug {
		responseDump, _ := httputil.DumpResponse(resp, true)
		log.Printf("[HTTP RESPONSE] ➜ \n%s", responseDump)
		defer func() {
			log.Printf("[Duration] %s\n", time.Since(start))
		}()
	}

	// Check for unexpected response statuses (example: return an error for 5xx responses, etc.)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 && resp.StatusCode != http.StatusTooManyRequests {
		defer resp.Body.Close()
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return resp, fmt.Errorf("unexpected HTTP url: %s, request body: %s, status: %d, response body: %s", url, body, resp.StatusCode, stringutils.BytesToString(respBody))
	}

	return resp, nil
}
