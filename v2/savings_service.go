package binance

import (
	"context"
	"net/http"
)

// ListSavingsFlexibleProductsService https://binance-docs.github.io/apidocs/spot/en/#get-flexible-product-list-user_data
type ListSavingsFlexibleProductsService struct {
	c        *Client
	status   string
	featured string
	current  int64
	size     int64
}

// Status represent the product status ("ALL", "SUBSCRIBABLE", "UNSUBSCRIBABLE") - Default: "ALL"
func (s *ListSavingsFlexibleProductsService) Status(status string) *ListSavingsFlexibleProductsService {
	s.status = status
	return s
}

// Featured ("ALL", "TRUE") - Default: "ALL"
func (s *ListSavingsFlexibleProductsService) Featured(featured string) *ListSavingsFlexibleProductsService {
	s.featured = featured
	return s
}

// Current query page. Default: 1, Min: 1
func (s *ListSavingsFlexibleProductsService) Current(current int64) *ListSavingsFlexibleProductsService {
	s.current = current
	return s
}

// Size Default: 50, Max: 100
func (s *ListSavingsFlexibleProductsService) Size(size int64) *ListSavingsFlexibleProductsService {
	s.size = size
	return s
}

// Do send request
func (s *ListSavingsFlexibleProductsService) Do(ctx context.Context, opts ...RequestOption) ([]*SavingsFlexibleProduct, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/lending/daily/product/list",
		secType:  secTypeSigned,
	}
	m := params{}
	if s.status != "" {
		m["status"] = s.status
	}
	if s.featured != "" {
		m["featured"] = s.featured
	}
	if s.current != 0 {
		m["current"] = s.current
	}
	if s.size != 0 {
		m["size"] = s.size
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res []*SavingsFlexibleProduct
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SavingsFlexibleProduct define a flexible product (Savings)
type SavingsFlexibleProduct struct {
	Asset                    string `json:"asset"`
	AvgAnnualInterestRate    string `json:"avgAnnualInterestRate"`
	CanPurchase              bool   `json:"canPurchase"`
	CanRedeem                bool   `json:"canRedeem"`
	DailyInterestPerThousand string `json:"dailyInterestPerThousand"`
	Featured                 bool   `json:"featured"`
	MinPurchaseAmount        string `json:"minPurchaseAmount"`
	ProductId                string `json:"productId"`
	PurchasedAmount          string `json:"purchasedAmount"`
	Status                   string `json:"status"`
	UpLimit                  string `json:"upLimit"`
	UpLimitPerUser           string `json:"upLimitPerUser"`
}

// PurchaseSavingsFlexibleProductService https://binance-docs.github.io/apidocs/spot/en/#purchase-flexible-product-user_data
type PurchaseSavingsFlexibleProductService struct {
	c         *Client
	productId string
	amount    float64
}

// ProductId represent the id of the flexible product to purchase
func (s *PurchaseSavingsFlexibleProductService) ProductId(productId string) *PurchaseSavingsFlexibleProductService {
	s.productId = productId
	return s
}

// Amount is the quantity of the product to purchase
func (s *PurchaseSavingsFlexibleProductService) Amount(amount float64) *PurchaseSavingsFlexibleProductService {
	s.amount = amount
	return s
}

// Do send request
func (s *PurchaseSavingsFlexibleProductService) Do(ctx context.Context, opts ...RequestOption) (uint64, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/lending/daily/purchase",
		secType:  secTypeSigned,
	}
	m := params{
		"productId": s.productId,
		"amount":    s.amount,
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return 0, err
	}

	var res *PurchaseSavingsFlexibleProductResponse
	if err = json.Unmarshal(data, &res); err != nil {
		return 0, err
	}

	return res.PurchaseId, nil
}

type PurchaseSavingsFlexibleProductResponse struct {
	PurchaseId uint64 `json:"purchaseId"`
}

// RedeemSavingsFlexibleProductService https://binance-docs.github.io/apidocs/spot/en/#redeem-flexible-product-user_data
type RedeemSavingsFlexibleProductService struct {
	c          *Client
	productId  string
	amount     float64
	redeemType string
}

// ProductId represent the id of the flexible product to redeem
func (s *RedeemSavingsFlexibleProductService) ProductId(productId string) *RedeemSavingsFlexibleProductService {
	s.productId = productId
	return s
}

// Amount is the quantity of the product to redeem
func (s *RedeemSavingsFlexibleProductService) Amount(amount float64) *RedeemSavingsFlexibleProductService {
	s.amount = amount
	return s
}

// Type ("FAST", "NORMAL")
func (s *RedeemSavingsFlexibleProductService) Type(redeemType string) *RedeemSavingsFlexibleProductService {
	s.redeemType = redeemType
	return s
}

// Do send request
func (s *RedeemSavingsFlexibleProductService) Do(ctx context.Context, opts ...RequestOption) error {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/lending/daily/redeem",
		secType:  secTypeSigned,
	}
	m := params{
		"productId": s.productId,
		"amount":    s.amount,
	}
	if s.redeemType != "" {
		m["type"] = s.redeemType
	}
	r.setParams(m)
	_, err := s.c.callAPI(ctx, r, opts...)

	return err
}

// ListSavingsFixedAndActivityProductsService https://binance-docs.github.io/apidocs/spot/en/#get-fixed-and-activity-project-list-user_data
type ListSavingsFixedAndActivityProductsService struct {
	c           *Client
	asset       string
	projectType string
	status      string
	isSortAsc   bool
	sortBy      string
	current     int64
	size        int64
}

// Asset desired asset
func (s *ListSavingsFixedAndActivityProductsService) Asset(asset string) *ListSavingsFixedAndActivityProductsService {
	s.asset = asset
	return s
}

// Type set project type ("ACTIVITY", "CUSTOMIZED_FIXED")
func (s *ListSavingsFixedAndActivityProductsService) Type(projectType string) *ListSavingsFixedAndActivityProductsService {
	s.projectType = projectType
	return s
}

// IsSortAsc default "true"
func (s *ListSavingsFixedAndActivityProductsService) IsSortAsc(isSortAsc bool) *ListSavingsFixedAndActivityProductsService {
	s.isSortAsc = isSortAsc
	return s
}

// Status ("ALL", "SUBSCRIBABLE", "UNSUBSCRIBABLE") - default "ALL"
func (s *ListSavingsFixedAndActivityProductsService) Status(status string) *ListSavingsFixedAndActivityProductsService {
	s.status = status
	return s
}

// SortBy ("START_TIME", "LOT_SIZE", "INTEREST_RATE", "DURATION") - default "START_TIME"
func (s *ListSavingsFixedAndActivityProductsService) SortBy(sortBy string) *ListSavingsFixedAndActivityProductsService {
	s.sortBy = sortBy
	return s
}

// Current Currently querying page. Start from 1. Default:1
func (s *ListSavingsFixedAndActivityProductsService) Current(current int64) *ListSavingsFixedAndActivityProductsService {
	s.current = current
	return s
}

// Size Default:10, Max:100
func (s *ListSavingsFixedAndActivityProductsService) Size(size int64) *ListSavingsFixedAndActivityProductsService {
	s.size = size
	return s
}

// Do send request
func (s *ListSavingsFixedAndActivityProductsService) Do(ctx context.Context, opts ...RequestOption) ([]*SavingsFixedProduct, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/lending/project/list",
		secType:  secTypeSigned,
	}
	m := params{
		"type": s.projectType,
	}
	if s.asset != "" {
		m["asset"] = s.asset
	}
	if s.status != "" {
		m["status"] = s.status
	}
	if s.isSortAsc != true {
		m["isSortAsc"] = s.isSortAsc
	}
	if s.sortBy != "" {
		m["sortBy"] = s.sortBy
	}
	if s.current != 1 {
		m["current"] = s.current
	}
	if s.size != 10 {
		m["size"] = s.size
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res []*SavingsFixedProduct
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

// SavingsFixedProduct define a fixed product (Savings)
type SavingsFixedProduct struct {
	Asset              string `json:"asset"`
	DisplayPriority    int    `json:"displayPriority"`
	Duration           int    `json:"duration"`
	InterestPerLot     string `json:"interestPerLot"`
	InterestRate       string `json:"interestRate"`
	LotSize            string `json:"lotSize"`
	LotsLowLimit       int    `json:"lotsLowLimit"`
	LotsPurchased      int    `json:"lotsPurchased"`
	LotsUpLimit        int    `json:"lotsUpLimit"`
	MaxLotsPerUser     int    `json:"maxLotsPerUser"`
	NeedKyc            bool   `json:"needKyc"`
	ProjectId          string `json:"projectId"`
	ProjectName        string `json:"projectName"`
	Status             string `json:"status"`
	Type               string `json:"type"`
	WithAreaLimitation bool   `json:"withAreaLimitation"`
}

// SavingFlexibleProductPositionsService fetches the saving flexible product positions
type SavingFlexibleProductPositionsService struct {
	c     *Client
	asset string
}

// Asset sets the asset parameter.
func (s *SavingFlexibleProductPositionsService) Asset(asset string) *SavingFlexibleProductPositionsService {
	s.asset = asset
	return s
}

// Do send request
func (s *SavingFlexibleProductPositionsService) Do(ctx context.Context, opts ...RequestOption) ([]*SavingFlexibleProductPosition, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/lending/daily/token/position",
		secType:  secTypeSigned,
	}
	m := params{}
	if s.asset != "" {
		m["asset"] = s.asset
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res []*SavingFlexibleProductPosition
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

// SavingFlexibleProductPosition represents a saving flexible product position.
type SavingFlexibleProductPosition struct {
	Asset                 string `json:"asset"`
	ProductId             string `json:"productId"`
	ProductName           string `json:"productName"`
	AvgAnnualInterestRate string `json:"avgAnnualInterestRate"`
	AnnualInterestRate    string `json:"annualInterestRate"`
	DailyInterestRate     string `json:"dailyInterestRate"`
	TotalInterest         string `json:"totalInterest"`
	TotalAmount           string `json:"totalAmount"`
	TotalPurchasedAmount  string `json:"todayPurchasedAmount"`
	RedeemingAmount       string `json:"redeemingAmount"`
	FreeAmount            string `json:"freeAmount"`
	FreezeAmount          string `json:"freezeAmount,omitempty"`
	LockedAmount          string `json:"lockedAmount,omitempty"`
	CanRedeem             bool   `json:"canRedeem"`
}

// SavingFixedProjectPositionsService fetches the saving flexible product positions
type SavingFixedProjectPositionsService struct {
	c         *Client
	asset     string
	status    string
	projectId string
}

// Asset sets the asset parameter.
func (s *SavingFixedProjectPositionsService) Asset(asset string) *SavingFixedProjectPositionsService {
	s.asset = asset
	return s
}

// Status ("HOLDING", "REDEEMED"), default will fetch all
func (s *SavingFixedProjectPositionsService) Status(status string) *SavingFixedProjectPositionsService {
	s.status = status
	return s
}

// Project ID of the fixed project/activity
func (s *SavingFixedProjectPositionsService) ProjectID(projectId string) *SavingFixedProjectPositionsService {
	s.projectId = projectId
	return s
}

// Do send request
func (s *SavingFixedProjectPositionsService) Do(ctx context.Context, opts ...RequestOption) ([]*SavingFixedProjectPosition, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/lending/project/position/list",
		secType:  secTypeSigned,
	}
	m := params{}
	if s.asset != "" {
		m["asset"] = s.asset
	}
	if s.status != "" {
		m["status"] = s.status
	}
	if s.projectId != "" {
		m["projectId"] = s.projectId
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res []*SavingFixedProjectPosition
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

// SavingFixedProjectPosition represents a saving flexible product position.
type SavingFixedProjectPosition struct {
	Asset           string `json:"asset"`
	CanTransfer     bool   `json:"canTransfer"`
	CreateTimestamp int64  `json:"createTimestamp"`
	Duration        int64  `json:"duration"`
	StartTime       int64  `json:"startTime"`
	EndTime         int64  `json:"endTime"`
	PurchaseTime    int64  `json:"purchaseTime"`
	RedeemDate      string `json:"redeemDate"`
	Interest        string `json:"interest"`
	InterestRate    string `json:"interestRate"`
	Lot             int32  `json:"lot"`
	PositionId      int64  `json:"positionId"`
	Principal       string `json:"principal"`
	ProjectId       string `json:"projectId"`
	ProjectName     string `json:"projectName"`
	Status          string `json:"status"`
	ProjectType     string `json:"type"`
}

// GetFlexibleProductPositionService: https://binance-docs.github.io/apidocs/spot/en/#get-flexible-product-position-user_data
type GetFlexibleProductPositionService struct {
	c         *Client
	productId string
	current   int64
	size      int64
}

func (s *GetFlexibleProductPositionService) ProductId(productId string) *GetFlexibleProductPositionService {
	s.productId = productId
	return s
}

// Currently querying the page. Start from 1. Default:1
func (s *GetFlexibleProductPositionService) Current(current int64) *GetFlexibleProductPositionService {
	s.current = current
	return s
}

// Default:10, Max:100
func (s *GetFlexibleProductPositionService) Size(size int64) *GetFlexibleProductPositionService {
	s.size = size
	return s
}

func (s *GetFlexibleProductPositionService) Do(ctx context.Context, opts ...RequestOption) (GetFlexibleProductPositionResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/simple-earn/flexible/position",
		secType:  secTypeSigned,
	}
	m := params{}
	if s.productId != "" {
		m["productId"] = s.productId
	}
	if s.current != 0 {
		m["current"] = s.current
	}
	if s.size != 0 {
		m["size"] = s.size
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return GetFlexibleProductPositionResponse{}, err
	}
	var res GetFlexibleProductPositionResponse
	if err = json.Unmarshal(data, &res); err != nil {
		return GetFlexibleProductPositionResponse{}, err
	}
	return res, nil
}

type GetFlexibleProductPositionResponse struct {
	Rows  []FlexibleProductPosition `json:"rows,omitempty"`
	Total int                       `json:"total,omitempty"`
}

type FlexibleProductPosition struct {
	TotalAmount                    string             `json:"totalAmount,omitempty"`
	TierAnnualPercentageRate       map[string]float64 `json:"tierAnnualPercentageRate,omitempty"`
	LatestAnnualPercentageRate     string             `json:"latestAnnualPercentageRate,omitempty"`
	YesterdayAirdropPercentageRate string             `json:"yesterdayAirdropPercentageRate,omitempty"`
	Asset                          string             `json:"asset,omitempty"`
	AirDropAsset                   string             `json:"airDropAsset,omitempty"`
	CanRedeem                      bool               `json:"canRedeem,omitempty"`
	CollateralAmount               string             `json:"collateralAmount,omitempty"`
	ProductID                      string             `json:"productId,omitempty"`
	YesterdayRealTimeRewards       string             `json:"yesterdayRealTimeRewards,omitempty"`
	CumulativeBonusRewards         string             `json:"cumulativeBonusRewards,omitempty"`
	CumulativeRealTimeRewards      string             `json:"cumulativeRealTimeRewards,omitempty"`
	CumulativeTotalRewards         string             `json:"cumulativeTotalRewards,omitempty"`
	AutoSubscribe                  bool               `json:"autoSubscribe,omitempty"`
}

// https://binance-docs.github.io/apidocs/spot/en/#get-locked-product-position-user_data
type GetLockedProductPositionService struct {
	c          *Client
	asset      string
	positionId string
	productId  string
	current    int64
	size       int64
}

func (s *GetLockedProductPositionService) Asset(asset string) *GetLockedProductPositionService {
	s.asset = asset
	return s
}

func (s *GetLockedProductPositionService) PositionID(positionId string) *GetLockedProductPositionService {
	s.positionId = positionId
	return s
}

func (s *GetLockedProductPositionService) ProductId(productId string) *GetLockedProductPositionService {
	s.productId = productId
	return s
}

// Currently querying the page. Start from 1. Default:1
func (s *GetLockedProductPositionService) Current(current int64) *GetLockedProductPositionService {
	s.current = current
	return s
}

// Default:10, Max:100
func (s *GetLockedProductPositionService) Size(size int64) *GetLockedProductPositionService {
	s.size = size
	return s
}

func (s *GetLockedProductPositionService) Do(ctx context.Context, opts ...RequestOption) (GetLockedProductPositionResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/simple-earn/locked/position",
		secType:  secTypeSigned,
	}
	m := params{}
	if s.asset != "" {
		m["asset"] = s.asset
	}
	if s.positionId != "" {
		m["positionId"] = s.positionId
	}
	if s.productId != "" {
		m["productId"] = s.productId
	}
	if s.current != 0 {
		m["current"] = s.current
	}
	if s.size != 0 {
		m["size"] = s.size
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return GetLockedProductPositionResponse{}, err
	}
	var res GetLockedProductPositionResponse
	if err = json.Unmarshal(data, &res); err != nil {
		return GetLockedProductPositionResponse{}, err
	}
	return res, nil
}

type GetLockedProductPositionResponse struct {
	Rows  []LockedProductPosition `json:"rows,omitempty"`
	Total int                     `json:"total,omitempty"`
}
type LockedProductPosition struct {
	PositionID   string `json:"positionId,omitempty"`
	ProjectID    string `json:"projectId,omitempty"`
	Asset        string `json:"asset,omitempty"`
	Amount       string `json:"amount,omitempty"`
	PurchaseTime string `json:"purchaseTime,omitempty"`
	Duration     string `json:"duration,omitempty"`
	AccrualDays  string `json:"accrualDays,omitempty"`
	RewardAsset  string `json:"rewardAsset,omitempty"`
	Apy          string `json:"APY,omitempty"`
	IsRenewable  bool   `json:"isRenewable,omitempty"`
	IsAutoRenew  bool   `json:"isAutoRenew,omitempty"`
	RedeemDate   string `json:"redeemDate,omitempty"`
}

// https://binance-docs.github.io/apidocs/spot/en/#get-flexible-rewards-history-user_data
type FlexibleRewardHistoryService struct {
	c         *Client
	productID string
	asset     string
	startTime int64
	endTime   int64
	typ       string
	current   int64
	size      int64
}

func (s *FlexibleRewardHistoryService) ProductID(productID string) *FlexibleRewardHistoryService {
	s.productID = productID
	return s
}

func (s *FlexibleRewardHistoryService) Asset(asset string) *FlexibleRewardHistoryService {
	s.asset = asset
	return s
}

func (s *FlexibleRewardHistoryService) StartTime(startTime int64) *FlexibleRewardHistoryService {
	s.startTime = startTime
	return s
}

func (s *FlexibleRewardHistoryService) EndTime(endTime int64) *FlexibleRewardHistoryService {
	s.endTime = endTime
	return s
}

func (s *FlexibleRewardHistoryService) Typ(typ string) *FlexibleRewardHistoryService {
	s.typ = typ
	return s
}

func (s *FlexibleRewardHistoryService) Current(current int64) *FlexibleRewardHistoryService {
	s.current = current
	return s
}

func (s *FlexibleRewardHistoryService) Size(size int64) *FlexibleRewardHistoryService {
	s.size = size
	return s
}

func (s *FlexibleRewardHistoryService) Do(ctx context.Context, opts ...RequestOption) (FlexibleRewardHistoryResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/simple-earn/flexible/history/rewardsRecord",
		secType:  secTypeSigned,
	}
	m := params{}
	if s.asset != "" {
		m["asset"] = s.asset
	}
	if s.startTime != 0 {
		m["startTime"] = s.startTime
	}
	if s.endTime != 0 {
		m["endTime"] = s.startTime
	}
	if s.productID != "" {
		m["productId"] = s.productID
	}
	if s.current != 0 {
		m["current"] = s.current
	}
	if s.size != 0 {
		m["size"] = s.size
	}
	if s.typ != "" {
		m["type"] = s.typ
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return FlexibleRewardHistoryResponse{}, err
	}

	var res FlexibleRewardHistoryResponse
	if err = json.Unmarshal(data, &res); err != nil {
		return FlexibleRewardHistoryResponse{}, err
	}
	return res, nil

}

type FlexibleRewardHistoryResponse struct {
	Rows  []FlexibleRewardHistory `json:"rows,omitempty"`
	Total int                     `json:"total,omitempty"`
}

type FlexibleRewardHistory struct {
	Asset     string `json:"asset,omitempty"`
	Rewards   string `json:"rewards,omitempty"`
	ProjectID string `json:"projectId,omitempty"`
	Typ       string `json:"type,omitempty"`
	Time      int64  `json:"time,omitempty"`
}
