package debank

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/coin50etf/coin-market/internal/dto"
	"github.com/coin50etf/coin-market/internal/pkg/config"
	"github.com/coin50etf/coin-market/internal/pkg/httpclient"
	"github.com/coin50etf/coin-market/internal/pkg/json"
	"github.com/coin50etf/coin-market/internal/pkg/logger"
)

const (
	getUserAllSimpleProtocolList = "/v1/user/all_simple_protocol_list"
	// See: https://docs.cloud.debank.com/en/readme/api-pro-reference/user#get-user-protocol
	getUserProtocol = "/v1/user/protocol"
	// See: https://docs.cloud.debank.com/en/readme/api-pro-reference/user#get-user-token-list
	getUserTokenList = "/v1/user/token_list"
	// See: https://docs.cloud.debank.com/en/readme/api-pro-reference/user#get-user-used-chain
	getUserChainList = "/v1/user/used_chain_list"
	// See:https://docs.cloud.debank.com/en/readme/api-pro-reference/user#get-user-total-balance-on-all-supported-chains
	getUserTotalBalance = "/v1/user/total_balance"
)

type Client struct {
	httpClient *httpclient.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: httpclient.DefaultClient(),
	}
}

func (c *Client) GetUserAllSimpleProtocolList(ctx context.Context, walletAddress string) ([]*dto.ProtocolDto, error) {
	params := map[string]string{"id": walletAddress}
	body, err := get(ctx, c.httpClient, getUserAllSimpleProtocolList, params)
	if err != nil {
		return nil, err
	}

	var resp []*dto.ProtocolDto
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetUserProtocol(ctx context.Context, walletAddress, protocolID string) (*dto.ProtocolDto, error) {
	params := map[string]string{"id": walletAddress, "protocol_id": protocolID}
	body, err := get(ctx, c.httpClient, getUserProtocol, params)
	if err != nil {
		return nil, err
	}

	resp := new(dto.ProtocolDto)
	if err := json.Unmarshal(body, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetUserTokenList(ctx context.Context, walletAddress, chainID string) ([]*dto.UserTokenDto, error) {
	params := map[string]string{
		"id":       walletAddress,
		"chain_id": chainID,
	}
	body, err := get(ctx, c.httpClient, getUserTokenList, params)
	if err != nil {
		return nil, err
	}

	var resp []*dto.UserTokenDto
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetUserChainList(ctx context.Context, walletAddress string) ([]*dto.UserUsedChainDto, error) {
	params := map[string]string{
		"id": walletAddress,
	}
	body, err := get(ctx, c.httpClient, getUserChainList, params)
	if err != nil {
		return nil, err
	}

	var resp []*dto.UserUsedChainDto
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetUserTotalBalance(ctx context.Context, walletAddress string) (*dto.UserTotalBalanceDto, error) {
	params := map[string]string{
		"id": walletAddress,
	}
	body, err := get(ctx, c.httpClient, getUserTotalBalance, params)
	if err != nil {
		return nil, err
	}

	resp := new(dto.UserTotalBalanceDto)
	if err := json.Unmarshal(body, resp); err != nil {
		return nil, err
	}

	return resp, nil

}

func get(ctx context.Context, client *httpclient.Client, endpoint string, params map[string]string) ([]byte, error) {
	if params == nil {
		params = make(map[string]string)
	}
	debankConfig := config.Conf.ThirdParty.Debank
	baseURL, err := url.Parse(debankConfig.BaseUrl)
	if err != nil {
		return nil, err
	}

	header := map[string]string{
		"AccessKey": debankConfig.AccessKey,
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

	resp, err := client.Do(ctx, http.MethodGet, fullURL.String(), nil, header)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests {
		return nil, errors.New("debank Apis rate limit exceeded from response")
	}

	return io.ReadAll(resp.Body)
}

func (c *Client) GetAllTransactions(ctx context.Context, walletAddress string, sign *dto.DebankSignDto) (*dto.DebankResponse, error) {
	const maxRetries = 5

	var err error
	for attempt := 1; attempt <= maxRetries; attempt++ {
		time.Sleep(time.Second)

		url := fmt.Sprintf("https://api.debank.com/history/all_list?user_id=%s", walletAddress)
		header := map[string]string{
			"accept":             "*/*",
			"accept-language":    "en,zh-CN;q=0.9,zh;q=0.8",
			"origin":             "https://debank.com",
			"referer":            "https://debank.com/",
			"sec-ch-ua-platform": "\"macOS\"",
			"sec-fetch-dest":     "empty",
			"sec-fetch-mode":     "cors",
			"sec-fetch-site":     "same-site",
			"source":             "web",
			"user-agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/140.0.0.0 Safari/537.36",
			"x-api-nonce":        sign.XApiNonce,
			"x-api-sign":         sign.XApiSign,
			"x-api-ts":           sign.XApiTs,
			"x-api-ver":          sign.XApiVer,
		}
		var resp *http.Response
		resp, err = c.httpClient.Do(ctx, http.MethodGet, url, nil, header)
		if err != nil {
			logger.Error(ctx, "GetAllTransactions httpclient.Do failed", "error", err, "attempt", attempt)
			continue
		}
		defer resp.Body.Close()

		var body []byte
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			logger.Error(ctx, "GetAllTransactions io.ReadAll failed", "error", err, "attempt", attempt)
			continue
		}

		if resp.StatusCode == http.StatusTooManyRequests {
			return nil, errors.New("debank Apis rate limit exceeded from response")
		}

		response := &dto.DebankResponse{}
		if err := json.Unmarshal(body, response); err != nil {
			logger.Error(ctx, "GetAllTransactions json.Unmarshal failed", "error", err, "attempt", attempt)
			continue
		}

		if response == nil || response.Data.Result == nil {
			logger.Error(ctx, "GetAllTransactions empty response", "attempt", attempt)
			err = fmt.Errorf("empty response")
			continue
		}

		return response, nil
	}

	return nil, fmt.Errorf("failed to get transactions after %d attempts: %w", maxRetries, err)
}
