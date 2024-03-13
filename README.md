### go-binance

A Golang SDK for [binance](https://www.binance.com) API.

This fork is created to allow fetching additional information from binance REST APIs, that is not listed in 
[binance API document](https://github.com/binance-exchange/binance-official-api-docs) but represented in official [API definition](https://binance-docs.github.io/apidocs/spot/en/#change-log).

This base Golang SDK is created by [adshao](https://github.com/adshao). To take more information about this SDK and 
implemented REST APIs please visit the [main SDK page](https://github.com/adshao/go-binance).

Fork`s change log:
- Added service `NewFundingRateHistoryService` that allows to fetch funding rate history data using the Rest API.
- Added service `NewFundingRateInfoService` that allows to fetch symbols funding rate info using the Rest API.
