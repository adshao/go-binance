package binance

import (
	"fmt"
	"net/url"
)

type secType int

const (
	secTypeNone secType = iota
	secTypeAPIKey
	secTypeSigned
)

type params map[string]interface{}

// request define an API request
type request struct {
	method     string
	endpoint   string
	query      url.Values
	form       url.Values
	recvWindow int64
	secType    secType
}

func (r *request) SetParam(key string, value interface{}) *request {
	if r.query == nil {
		r.query = url.Values{}
	}
	r.query.Set(key, fmt.Sprintf("%v", value))
	return r
}

func (r *request) SetParams(m params) *request {
	for k, v := range m {
		r.SetParam(k, v)
	}
	return r
}

func (r *request) SetFormParam(key string, value interface{}) *request {
	if r.form == nil {
		r.form = url.Values{}
	}
	r.form.Set(key, fmt.Sprintf("%v", value))
	return r
}

func (r *request) SetFormParams(m params) *request {
	for k, v := range m {
		r.SetFormParam(k, v)
	}
	return r
}

func (r *request) validate() (err error) {
	if r.query == nil {
		r.query = url.Values{}
	}
	if r.form == nil {
		r.form = url.Values{}
	}
	return nil
}

// RequestOption define option type for request
type RequestOption func(*request)

// WithRecvWindow set recvWindow param for the request
func WithRecvWindow(recvWindow int64) RequestOption {
	return func(r *request) {
		r.recvWindow = recvWindow
	}
}
