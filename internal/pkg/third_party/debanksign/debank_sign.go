package debanksign

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/coin50etf/coin-market/internal/dto"
	"github.com/coin50etf/coin-market/internal/pkg/httpclient"
)

type Client struct {
	httpClient *httpclient.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: httpclient.DefaultClient(),
	}
}

func (c *Client) GetSignature(ctx context.Context, walletAddress string) (*dto.DebankSignDto, error) {
	const maxRetries = 3

	var (
		resp *dto.DebankSignDto
		err  error
	)

	for attempt := 1; attempt <= maxRetries; attempt++ {
		time.Sleep(time.Second)

		// 请求参数
		params := map[string]string{"address": walletAddress}

		// 发起请求
		var body []byte
		body, err = get(ctx, c.httpClient, "/sign", params)
		if err != nil {
			// 打印日志可选
			log.Printf("Attempt %d: request failed: %v", attempt, err)
			continue
		}

		// 解析结果
		if err = json.Unmarshal(body, &resp); err != nil {
			log.Printf("Attempt %d: failed to unmarshal response: %v", attempt, err)
			continue
		}

		// 检查 resp 是否为空（根据你的业务判断）
		if resp == nil || resp.XApiSign == "" {
			log.Printf("Attempt %d: empty response", attempt)
			err = fmt.Errorf("debank sign empty response")
			continue
		}

		// 成功则返回
		return resp, nil
	}

	// 最后一次仍然失败
	return nil, fmt.Errorf("failed to get signature after %d attempts: %w", maxRetries, err)
}

func get(ctx context.Context, client *httpclient.Client, endpoint string, params map[string]string) ([]byte, error) {
	if params == nil {
		params = make(map[string]string)
	}
	//baseURL, err := url.Parse("http://localhost:8899")
	baseURL, err := url.Parse("http://ec2-3-237-93-119.compute-1.amazonaws.com:8899")
	if err != nil {
		return nil, err
	}

	fullURL, err := baseURL.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	query := url.Values{}
	for key, value := range params {
		query.Set(key, value)
	}
	fullURL.RawQuery = query.Encode()

	resp, err := client.Do(ctx, http.MethodGet, fullURL.String(), nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests {
		return nil, errors.New("debank Apis rate limit exceeded from response")
	}

	return io.ReadAll(resp.Body)
}
