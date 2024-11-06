package futures

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/adshao/go-binance/v2/common"
)

// WsApiMethodType define method name for websocket API
type WsApiMethodType string

// WsApiRequest define common websocket API request
type WsApiRequest struct {
	Id     string          `json:"id"`
	Method WsApiMethodType `json:"method"`
	Params params          `json:"params"`
}

const (
	// apiKey define key for websocket API parameters
	apiKey = "apiKey"

	// OrderPlaceWsApiMethod define method for creation order via websocket API
	OrderPlaceWsApiMethod WsApiMethodType = "order.place"

	// CancelWsApiMethod define method for cancel order via websocket API
	CancelWsApiMethod WsApiMethodType = "order.cancel"

	// WriteSyncWsTimeout defines timeout for WriteSync method of client_ws
	WriteSyncWsTimeout = 5 * time.Second
)

var (
	// ErrorRequestIDNotSet defines that request ID is not set
	ErrorRequestIDNotSet = errors.New("ws service: request id is not set")

	// ErrorApiKeyIsNotSet defines that ApiKey is not set
	ErrorApiKeyIsNotSet = errors.New("ws service: api key is not set")

	// ErrorSecretKeyIsNotSet defines that SecretKey is not set
	ErrorSecretKeyIsNotSet = errors.New("ws service: secret key is not set")
)

// createWsRequest creates signed ws request
func createWsRequest(requestID string, client wsClient, method WsApiMethodType, params params) ([]byte, error) {
	if requestID == "" {
		return nil, ErrorRequestIDNotSet
	}

	if client.GetApiKey() == "" {
		return nil, ErrorApiKeyIsNotSet
	}

	if client.GetSecretKey() == "" {
		return nil, ErrorSecretKeyIsNotSet
	}

	params[apiKey] = client.GetApiKey()
	params[timestampKey] = currentTimestamp() - client.GetTimeOffset()

	sf, err := common.SignFunc(client.GetKeyType())
	if err != nil {
		return nil, err
	}
	signature, err := sf(client.GetSecretKey(), params.encode())
	if err != nil {
		return nil, err
	}
	params[signatureKey] = signature

	req := WsApiRequest{
		Id:     requestID,
		Method: method,
		Params: params,
	}

	rawData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	return rawData, nil
}
