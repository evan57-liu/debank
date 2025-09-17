package debank

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/coin50etf/coin-market/internal/dto"
	"github.com/coin50etf/coin-market/internal/pkg/config"
	"github.com/coin50etf/coin-market/internal/pkg/httpclient"
	"github.com/coin50etf/coin-market/internal/pkg/json"
)

const (
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
