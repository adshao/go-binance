package binance

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/bitly/go-simplejson"
)

type SideType string
type OrderType string
type TimeInForce string

const (
	SideTypeBuy  SideType = "BUY"
	SideTypeSell SideType = "SELL"

	OrderTypeLimit  OrderType = "LIMIT"
	OrderTypeMarket OrderType = "MARKET"

	TimeInForceGTC TimeInForce = "GTC"
	TimeInForceIOC TimeInForce = "IOC"

	timestampKey  = "timestamp"
	signatureKey  = "signature"
	recvWindowKey = "recvWindow"
)

func currentTimestamp() int64 {
	return int64(time.Nanosecond) * time.Now().UnixNano() / int64(time.Millisecond)
}

func newJSON(data []byte) (j *simplejson.Json, err error) {
	j, err = simplejson.NewJson(data)
	if err != nil {
		return
	}
	return
}

// NewClient initialize API client
func NewClient(apiKey, secretKey string) *Client {
	return &Client{
		APIKey:     apiKey,
		SecretKey:  secretKey,
		BaseURL:    "https://www.binance.com",
		UserAgent:  "Binance/golang",
		HTTPClient: http.DefaultClient,
		Logger:     log.New(os.Stderr, "Binance-golang ", log.LstdFlags),
	}
}

type doFunc func(req *http.Request) (*http.Response, error)

type Client struct {
	APIKey     string
	SecretKey  string
	BaseURL    string
	UserAgent  string
	HTTPClient *http.Client
	Debug      bool
	Logger     *log.Logger
	do         doFunc
}

func (c *Client) debug(format string, v ...interface{}) {
	if c.Debug {
		c.Logger.Printf(format, v...)
	}
}

func (c *Client) callAPI(ctx context.Context, r *request, opts ...RequestOption) (data []byte, err error) {
	// set request options from user
	for _, opt := range opts {
		opt(r)
	}
	err = r.validate()
	if err != nil {
		return
	}

	endpoint := fmt.Sprintf("%s%s", c.BaseURL, r.endpoint)
	if r.recvWindow > 0 {
		r.SetParam(recvWindowKey, r.recvWindow)
	}
	if r.secType == secTypeSigned {
		r.SetParam(timestampKey, currentTimestamp())
	}
	queryString := r.query.Encode()
	body := &bytes.Buffer{}
	bodyString := r.form.Encode()
	header := http.Header{}
	if bodyString != "" {
		header.Set("Context-Type", "application/x-www-form-urlencoded")
		body = bytes.NewBufferString(bodyString)
	}
	if r.secType == secTypeAPIKey || r.secType == secTypeSigned {
		header.Set("X-MBX-APIKEY", c.APIKey)
	}

	if r.secType == secTypeSigned {
		raw := fmt.Sprintf("%s%s", queryString, bodyString)
		mac := hmac.New(sha256.New, []byte(c.SecretKey))
		mac.Write([]byte(raw))
		v := url.Values{}
		v.Set(signatureKey, fmt.Sprintf("%x", (mac.Sum(nil))))
		if queryString == "" {
			queryString = v.Encode()
		} else {
			queryString = fmt.Sprintf("%s&%s", queryString, v.Encode())
		}
	}
	if queryString != "" {
		endpoint = fmt.Sprintf("%s?%s", endpoint, queryString)
	}
	c.debug("endpoint: %s, body: %s", endpoint, bodyString)

	req, err := http.NewRequest(r.method, endpoint, body)
	if err != nil {
		return
	}
	req = req.WithContext(ctx)
	req.Header = header
	c.debug("request: %#v", req)
	f := c.do
	if f == nil {
		f = c.HTTPClient.Do
	}
	res, err := f(req)
	if err != nil {
		return
	}
	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	defer res.Body.Close()
	c.debug("response: %#v", res)
	c.debug("response body: %s", string(data))

	if res.StatusCode >= 400 {
		apiErr := new(APIError)
		e := json.Unmarshal(data, apiErr)
		if e != nil {
			c.debug("failed to unmarshal json: %s", err)
		}
		return nil, apiErr
	}
	return
}

// NewPingService initialize ping service
func (c *Client) NewPingService() *PingService {
	return &PingService{c: c}
}

// NewServerTimeService initialize server time service
func (c *Client) NewServerTimeService() *ServerTimeService {
	return &ServerTimeService{c: c}
}

// NewDepthService initialize depth service
func (c *Client) NewDepthService() *DepthService {
	return &DepthService{c: c}
}

// NewAggregateTradesService initialize aggregate trades service
func (c *Client) NewAggregateTradesService() *AggregateTradesService {
	return &AggregateTradesService{c: c}
}

// NewKlinesService initialize klines service
func (c *Client) NewKlinesService() *KlinesService {
	return &KlinesService{c: c}
}

func (c *Client) NewPriceChangeStatsService() *PriceChangeStatsService {
	return &PriceChangeStatsService{c: c}
}

func (c *Client) NewListPricesService() *ListPricesService {
	return &ListPricesService{c: c}
}

func (c *Client) NewListBookTickersService() *ListBookTickersService {
	return &ListBookTickersService{c: c}
}

func (c *Client) NewCreateOrderService() *CreateOrderService {
	return &CreateOrderService{c: c}
}

func (c *Client) NewGetOrderService() *GetOrderService {
	return &GetOrderService{c: c}
}

func (c *Client) NewCancelOrderService() *CancelOrderService {
	return &CancelOrderService{c: c}
}

func (c *Client) NewListOpenOrdersService() *ListOpenOrdersService {
	return &ListOpenOrdersService{c: c}
}

func (c *Client) NewListOrdersService() *ListOrdersService {
	return &ListOrdersService{c: c}
}

func (c *Client) NewGetAccountService() *GetAccountService {
	return &GetAccountService{c: c}
}

func (c *Client) NewListTradesService() *ListTradesService {
	return &ListTradesService{c: c}
}

func (c *Client) NewListDepositsService() *ListDepositsService {
	return &ListDepositsService{c: c}
}

func (c *Client) NewCreateWithdrawService() *CreateWithdrawService {
	return &CreateWithdrawService{c: c}
}

func (c *Client) NewListWithdrawsService() *ListWithdrawsService {
	return &ListWithdrawsService{c: c}
}

func (c *Client) NewStartUserStreamService() *StartUserStreamService {
	return &StartUserStreamService{c: c}
}

func (c *Client) NewKeepaliveUserStreamService() *KeepaliveUserStreamService {
	return &KeepaliveUserStreamService{c: c}
}

func (c *Client) NewCloseUserStreamService() *CloseUserStreamService {
	return &CloseUserStreamService{c: c}
}
