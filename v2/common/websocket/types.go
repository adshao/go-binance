package websocket

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/adshao/go-binance/v2/common"
)

// WsApiMethodType define method name for websocket API
type WsApiMethodType string

// WsApiRequest define common websocket API request
type WsApiRequest struct {
	Id     string                 `json:"id"`
	Method WsApiMethodType        `json:"method"`
	Params map[string]interface{} `json:"params"`
}

var (
	// WriteSyncWsTimeout defines timeout for WriteSync method of client_ws
	WriteSyncWsTimeout = 5 * time.Second
)

const (
	// apiKey define key for websocket API parameters
	apiKey = "apiKey"

	// timestampKey define key for websocket API parameters
	timestampKey = "timestamp"

	// signatureKey define key for websocket API parameters
	signatureKey = "signature"

	// SPOT

	// OrderPlaceSpotWsApiMethod define method for creation order via websocket API
	OrderPlaceSpotWsApiMethod WsApiMethodType = "order.place"

	// FUTURES

	// OrderPlaceFuturesWsApiMethod define method for creation order via websocket API
	OrderPlaceFuturesWsApiMethod WsApiMethodType = "order.place"

	// CancelFuturesWsApiMethod define method for cancel order via websocket API
	CancelFuturesWsApiMethod WsApiMethodType = "order.cancel"
)

var (
	// ErrorRequestIDNotSet defines that request ID is not set
	ErrorRequestIDNotSet = errors.New("ws service: request id is not set")

	// ErrorApiKeyIsNotSet defines that ApiKey is not set
	ErrorApiKeyIsNotSet = errors.New("ws service: api key is not set")

	// ErrorSecretKeyIsNotSet defines that SecretKey is not set
	ErrorSecretKeyIsNotSet = errors.New("ws service: secret key is not set")
)

func NewRequestData(
	requestID string,
	apiKey string,
	secretKey string,
	timeOffset int64,
	keyType string,
) RequestData {
	return RequestData{
		requestID:  requestID,
		apiKey:     apiKey,
		secretKey:  secretKey,
		timeOffset: timeOffset,
		keyType:    keyType,
	}
}

type RequestData struct {
	requestID  string
	apiKey     string
	secretKey  string
	timeOffset int64
	keyType    string
}

// CreateRequest creates signed ws request
func CreateRequest(reqData RequestData, method WsApiMethodType, params map[string]interface{}) ([]byte, error) {
	if reqData.requestID == "" {
		return nil, ErrorRequestIDNotSet
	}

	if reqData.apiKey == "" {
		return nil, ErrorApiKeyIsNotSet
	}

	if reqData.secretKey == "" {
		return nil, ErrorSecretKeyIsNotSet
	}

	params[apiKey] = reqData.apiKey
	params[timestampKey] = timestamp(reqData.timeOffset)

	sf, err := common.SignFunc(reqData.keyType)
	if err != nil {
		return nil, err
	}
	signature, err := sf(reqData.secretKey, encodeParams(params))
	if err != nil {
		return nil, err
	}
	params[signatureKey] = signature

	req := WsApiRequest{
		Id:     reqData.requestID,
		Method: method,
		Params: params,
	}

	rawData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	return rawData, nil
}

// encode encodes the parameters to a URL encoded string
func encodeParams(p map[string]interface{}) string {
	queryValues := url.Values{}
	for key, value := range p {
		queryValues.Add(key, fmt.Sprintf("%v", value))
	}
	return queryValues.Encode()
}

func timestamp(offsetMilli int64) int64 {
	return time.Now().UnixMilli() - offsetMilli
}
