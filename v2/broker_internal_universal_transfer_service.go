package binance

import "context"

// BrokerInternalUniversalTransferService Universal Transfer (For Master Account)
type BrokerInternalUniversalTransferService struct {
	c               *Client
	fromID          *string
	toID            *string
	fromAccountType *string
	toAccountType   *string
	clientTranId    *string
	asset           string
	amount          float64
	recvWindow      *int64
	timestamp       int64
}

func (s *BrokerInternalUniversalTransferService) RecvWindow(v int64) *BrokerInternalUniversalTransferService {
	s.recvWindow = &v
	return s
}

func (s *BrokerInternalUniversalTransferService) Timestamp(v int64) *BrokerInternalUniversalTransferService {
	s.timestamp = v
	return s
}

func (s *BrokerInternalUniversalTransferService) FromID(v string) *BrokerInternalUniversalTransferService {
	s.fromID = &v
	return s
}

func (s *BrokerInternalUniversalTransferService) ToID(v string) *BrokerInternalUniversalTransferService {
	s.toID = &v
	return s
}

func (s *BrokerInternalUniversalTransferService) FromAccountType(v string) *BrokerInternalUniversalTransferService {
	s.fromAccountType = &v
	return s
}

func (s *BrokerInternalUniversalTransferService) ToAccountType(v string) *BrokerInternalUniversalTransferService {
	s.toAccountType = &v
	return s
}

func (s *BrokerInternalUniversalTransferService) Asset(v string) *BrokerInternalUniversalTransferService {
	s.asset = v
	return s
}

func (s *BrokerInternalUniversalTransferService) Amount(v float64) *BrokerInternalUniversalTransferService {
	s.amount = v
	return s
}

func (s *BrokerInternalUniversalTransferService) ClientTranId(v string) *BrokerInternalUniversalTransferService {
	s.clientTranId = &v
	return s
}

func (s *BrokerInternalUniversalTransferService) Do(ctx context.Context, opts ...RequestOption) (*BrokerInternalUniversalTransferResponse, error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/broker/universalTransfer",
		secType:  secTypeSigned,
	}
	if v := s.fromID; v != nil {
		r.setParam("fromId", *v)
	}
	if v := s.toID; v != nil {
		r.setParam("toId", *v)
	}
	r.setParam("asset", s.asset)
	r.setParam("amount", s.amount)
	r.setParam("timestamp", s.timestamp)
	if v := s.fromAccountType; v != nil {
		r.setParam("fromAccountType", *v)
	}
	if v := s.toAccountType; v != nil {
		r.setParam("toAccountType", *v)
	}
	if v := s.clientTranId; v != nil {
		r.setParam("clientTranId", *v)
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := &BrokerInternalUniversalTransferResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

type BrokerInternalUniversalTransferResponse struct {
	ID           int64  `json:"tranId"`
	ClientTranID string `json:"clientTranId"`
}
