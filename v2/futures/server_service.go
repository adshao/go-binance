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

// Do send request
func (s *ServerTimeService) Do(ctx context.Context, opts ...RequestOption) (serverTime int64, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/time",
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return 0, err
	}
	j, err := newJSON(data)
	if err != nil {
		return 0, err
	}
	serverTime = j.Get("serverTime").MustInt64()
	return serverTime, nil
}

// SetServerTimeService set server time
type SetServerTimeService struct {
	c *Client
}

// Do send request
func (s *SetServerTimeService) Do(ctx context.Context, opts ...RequestOption) (timeOffset int64, err error) {
	serverTime, err := s.c.NewServerTimeService().Do(ctx)
	if err != nil {
		return 0, err
	}
	timeOffset = currentTimestamp() - serverTime
	s.c.TimeOffset = timeOffset
	return timeOffset, nil
}
