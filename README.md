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

### Usage

#### REST API

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

```golang
// show account info
res, err := client.NewGetAccountService().Do(context.Background())
if err != nil {
    fmt.Printf("error: %s\n", err)
    return
}
fmt.Printf("%#v\n", res)

// show depth of symbol LTCBTC
res, err := client.NewDepthService().Symbol("LTCBTC").Do(context.Background())
if err != nil {
    fmt.Printf("err: %s\n", err)
    return
}
fmt.Printf("%#v\n", res)

// list klines of symbol LTCBTC with interval 15m and limit 3
klines, err := client.NewKlinesService().Symbol("LTCBTC").Interval("15m").Limit(3).Do(context.Background())
if err != nil {
    fmt.Printf("err: %s\n", err)
    return
}
for _, kline := range klines {
    fmt.Printf("%#v\n", kline)
}
```

Furthermore, add option binance.WithRecvWindow(recvWindow) if you want to change the default recvWindow.

```golang
trades, err := client.NewAggregateTradesService().Symbol("LTCBTC").StartTime(1508673256594).EndTime(1508673256594).Do(context.Background(), binance.WithRecvWindow(5000))
if err != nil {
    fmt.Printf("err: %s\n", err)
    return
}
for _, t := range trades {
    fmt.Printf("%#v\n", t)
}
```

#### Websocket

You don't need Client in websocket API. Just call binance.WsXXXServe(args, handler).

```golang
wsDepthHandler := func(event *binance.WsDepthEvent) {
    fmt.Printf("%#v\n", event)
}
done, err := binance.WsDepthServe("LTCBTC", wsDepthHandler)
if err != nil {
    fmt.Printf("%s\n", err)
}
// remove this if you do not want to be blocked here
<-done

wsKlineHandler := func(event *binance.WsKlineEvent) {
    fmt.Printf("%#v\n", event)
}
done, err := binance.WsKlineServe("LTCBTC", "1m", wsKlineHandler)
if err != nil {
    fmt.Printf("error: %s\n", err)
}
<-done

wsAggTradeHandler := func(event *binance.WsAggTradeEvent) {
    fmt.Printf("%#v\n\n", event)
}
done, err := binance.WsAggTradeServe("LTCBTC", wsAggTradeHandler)
if err != nil {
    fmt.Printf("error: %s\n", err)
}
<-done
```

### Documentation

### TODO

1. Godoc support
