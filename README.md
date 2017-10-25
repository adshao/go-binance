### go-binance

A Golang SDK for [binance](https://www.binance.com) API.

All the REST APIs listed in [binance API document](https://www.binance.com/restapipub.html) are implemented, as well as the websocket APIs.

For best compatibility, please use Go >= 1.8.

Make sure you have read binance API document before continuing.

### Installation

```shell
go get github.com/adshao/go-binance
```

### Importing

```golang
import (
    "github.com/adshao/go-binance"
)
```

### Setup

Init client for API services. Get APIKey/SecretKey from your binance account.

```golang
var (
    apiKey = "your api key"
    secretKey = "your secret key"
)
client := binance.NewClient(apiKey, secretKey)
// client.Debug = true // enable debug to verbose mode
```

A service instance stands for a REST API endpoint and is initialized by client.NewXXXService function.

Simply call API in chain style. Call Do() in the end to send HTTP request.

### REST API

#### List Ticker Prices

```golang
prices, err := client.NewListPricesService().Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
for _, p := range prices {
    fmt.Println(p)
}
```

#### Show Depth

```golang
// show depth of symbol LTCBTC
res, err := client.NewDepthService().Symbol("LTCBTC").Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(res)
```

#### List Klines

```golang
// list klines of symbol LTCBTC with interval 15m and limit 3
klines, err := client.NewKlinesService().Symbol("LTCBTC").Interval("15m").Limit(3).Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
for _, k := range klines {
    fmt.Println(k)
}
```

#### Create Order

```golang
err := client.NewCreateOrderService().Symbol("BNBETH").
    Side(binance.SideTypeBuy).Type(binance.OrderTypeLimit).
    TimeInForce(binance.TimeInForceGTC).Quantity(100).Price(0.00310000).Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
```

#### Cancel Order

```golang
_, err := client.NewCancelOrderService().Symbol("LTCBTC").OrderID(1).Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
```

#### List Aggregate Trades

Add option binance.WithRecvWindow(recvWindow) if you want to change the default recvWindow.

```golang
trades, err := client.NewAggTradesService().Symbol("LTCBTC").StartTime(1508673256594).EndTime(1508673256594).Do(context.Background(), binance.WithRecvWindow(5000))
if err != nil {
    fmt.Println(err)
    return
}
for _, t := range trades {
    fmt.Println(t)
}
```

#### Get Account

```golang
// show account info
res, err := client.NewGetAccountService().Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(res)
```

Please refer to [godoc](https://godoc.org/github.com/adshao/go-binance) for the full SDK services.

### Websocket

You don't need Client in websocket API. Just call binance.WsXXXServe(args, handler).

#### Depth

```golang
wsDepthHandler := func(event *binance.WsDepthEvent) {
    fmt.Println(event)
}
done, err := binance.WsDepthServe("LTCBTC", wsDepthHandler)
if err != nil {
    fmt.Println(err)
    return
}
// remove this if you do not want to be blocked here
<-done
```

#### Kline

```glang
wsKlineHandler := func(event *binance.WsKlineEvent) {
    fmt.Println(event)
}
done, err := binance.WsKlineServe("LTCBTC", "1m", wsKlineHandler)
if err != nil {
    fmt.Println(err)
}
<-done
```

#### Aggregate

```golang
wsAggTradeHandler := func(event *binance.WsAggTradeEvent) {
    fmt.Println(event)
}
done, err := binance.WsAggTradeServe("LTCBTC", wsAggTradeHandler)
if err != nil {
    fmt.Println(err)
}
<-done
```

### Documentation

https://godoc.org/github.com/adshao/go-binance
