package binance

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type AllCoinService struct {
	c *Client
}

type AllCoinResponse struct {
	Code          string     `json:"code"`
	Message       string     `json:"message"`
	MessageDetail string     `json:"messageDetail"`
	Data          []CoinInfo `json:"data"`
	Success       bool       `json:"success"`
}

func (s *AllCoinService) Do(ctx context.Context) ([]CoinInfo, error) {
	var endpoint string = "https://www.binance.com/bapi/capital/v1/public/capital/getNetworkCoinAll"
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bb, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	res := AllCoinResponse{}
	if err = json.Unmarshal(bb, &res); err != nil {
		return nil, err
	}
	if !res.Success {
		return nil, fmt.Errorf("%s", res.MessageDetail)
	}
	return res.Data, nil
}
