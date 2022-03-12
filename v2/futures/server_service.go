package futures

import (
	"context"
	"net/http"
)

// PingService ping server
type PingService struct {
	c *Client
}

// Do send request
func (s *PingService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/ping",
	}
	_, _, err = s.c.callAPI(ctx, r, opts...)
	return err
}

// ServerTimeService get server time
type ServerTimeService struct {
	c *Client
}

type ServerTimeResponse struct {
	Time              int64
	RateLimitWeight1m string `json:"rateLimitWeight1m,omitempty"`
}

// Do send request
func (s *ServerTimeService) Do(ctx context.Context, opts ...RequestOption) (res *ServerTimeResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/time",
	}
	res = new(ServerTimeResponse)
	var header *http.Header
	data, header, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	j, err := newJSON(data)
	if err != nil {
		return nil, err
	}
	res.Time = j.Get("serverTime").MustInt64()
	res.RateLimitWeight1m = header.Get("X-Mbx-Used-Weight-1m")

	return res, nil
}

// SetServerTimeService set server time
type SetServerTimeService struct {
	c *Client
}

// Do send request
func (s *SetServerTimeService) Do(ctx context.Context, opts ...RequestOption) (res *ServerTimeResponse, err error) {
	res, err = s.c.NewServerTimeService().Do(ctx)
	if err != nil {
		return nil, err
	}
	res.Time = currentTimestamp() - res.Time
	s.c.TimeOffset = res.Time

	return res, nil
}
