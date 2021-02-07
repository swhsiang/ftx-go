package rest

import (
	"fmt"
	"net/url"
	"path"
)

const (
	FtxServer            = "https://ftx.com/api"
	FtxAPISubaccount     = "subaccounts"
	FtxAPIMarkets        = "markets"
	FtxAPIOrderbook      = "orderbook"
	FtxAPITrades         = "trades"
	FtxAPIFutures        = "futures"
	FtxAPIFutureStats    = "stats"
	FtxAPIFundingRates   = "funding_rates"
	FtxAPIIndexes        = "indexes"
	FtxAPIIndexesWeights = "weights"
	FtxAPIExpiredFutures = "expired_futures"
	FtxAPIIndexesCandles = "candles"
	FtxAPIAccount        = "account"
	FtxAPIPositions      = "positions"
	FtxAPIWallet         = "wallet"
	FtxAPICoins          = "coins"
)

var baseUrl *url.URL

func init() {
	var err error
	baseUrl, err = url.Parse(FtxServer)
	if err != nil {
		panic("invalid url")
	}
}

func generateApiUrl(endpoint string, queryMap map[string]string) string {
	u := *baseUrl
	u.Path = path.Join(baseUrl.Path, endpoint)
	q, _ := url.ParseQuery(u.RawQuery)
	for k, v := range queryMap {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}
