package main

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
)

func SpotOrder() {
	binance.UseTestnet = true
	apiKey := ""
	secret := ""
	client := binance.NewClient(apiKey, secret)

	symbol := "BTCUSDT"
	side := binance.SideTypeSell
	orderType := binance.OrderTypeMarket
	timeInForce := binance.TimeInForceTypeGTC
	quantity := "0.001"

	res, err := client.NewCreateOrderService().Symbol(symbol).Side(side).
		Type(orderType).TimeInForce(timeInForce).Quantity(quantity).Do(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}

func FuturesOrder() {
	binance.UseTestnet = true
	apiKey := ""
	secret := ""
	client := binance.NewFuturesClient(apiKey, secret)

	symbol := "BTCUSDT"
	side := futures.SideTypeSell
	orderType := futures.OrderTypeMarket
	timeInForce := futures.TimeInForceTypeGTC
	quantity := "0.001"

	res, err := client.NewCreateOrderService().Symbol(symbol).Side(side).
		Type(orderType).TimeInForce(timeInForce).Quantity(quantity).Do(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)

}
