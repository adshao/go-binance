package binance

import (
	"context"
)

type CustomRequestParams struct {
	Method   string
	Endpoint string
	Params   map[string]interface{}
}

// CustomRequestService implemented to send custom request
//
// See https://binance-docs.github.io/apidocs/spot/en/#deposit-history-user_data
type CustomRequestService struct {
	c *Client
}

// Do send the request.
func (s *CustomRequestService) Do(ctx context.Context, params *CustomRequestParams) (res interface{}, err error) {
	r := &request{
		method:   params.Method,
		endpoint: params.Endpoint,
		secType:  secTypeSigned,
	}
	for k, v := range params.Params {
		r.setParam(k, v)
	}

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	return res, nil
}
