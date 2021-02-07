package rest

import (
	"log"
	"path"
	"strconv"
)

func (header *FtxClientHeader) GetMarkets() {
	Do(header, "GET", FtxAPIMarkets, map[string]string{})
}

func (header *FtxClientHeader) GetMarketOf(marketName string) {
	p := path.Join(FtxAPIMarkets, marketName)
	Do(header, "GET", p, map[string]string{})
}

func (header *FtxClientHeader) GetOrderbookOf(marketName string, depth string) {
	p := path.Join(FtxAPIMarkets, marketName, FtxAPIOrderbook)
	queryMap := map[string]string{"depth": depth}
	Do(header, "GET", p, queryMap)
}

func (header *FtxClientHeader) GetTrades(marketName string, params ...int) {
	if len(params) > 3 || len(params) == 2 {
		log.Fatal("wrong length of params")
		return
	}

	queryMap := map[string]string{}
	limit := 20

	var startTime, endTime int
	if len(params) == 1 {
		if 100 >= params[0] && 20 <= params[0] {
			limit = params[0]
		}
	} else if len(params) == 3 {
		if 100 >= params[0] && 20 <= params[0] {
			limit = params[0]
		}
		startTime, endTime = params[1], params[2]
		queryMap["start_time"] = strconv.Itoa(startTime)
		queryMap["end_time"] = strconv.Itoa(endTime)
	}
	queryMap["limit"] = strconv.Itoa(limit)
	p := path.Join(FtxAPIMarkets, marketName, FtxAPITrades)
	Do(header, "GET", p, queryMap)
}

func (header *FtxClientHeader) GetFutures() {
	Do(header, "GET", FtxAPIFutures, map[string]string{})
}

func (header *FtxClientHeader) GetFutureOf(futureName string) {
	p := path.Join(FtxAPIFutures, futureName)
	Do(header, "GET", p, map[string]string{})
}

func (header *FtxClientHeader) GetFutureStatsOf(futureName string) {
	p := path.Join(FtxAPIFutures, futureName, FtxAPIFutureStats)
	Do(header, "GET", p, map[string]string{})
}

func (header *FtxClientHeader) GetFundingRates() {
	Do(header, "GET", FtxAPIFundingRates, map[string]string{})
}

func (header *FtxClientHeader) GetIndexWeightOf(index string) {
	p := path.Join(FtxAPIIndexes, index, FtxAPIIndexesWeights)
	Do(header, "GET", p, map[string]string{})
}

func (header *FtxClientHeader) GetExpiredFutures() {
	Do(header, "GET", FtxAPIExpiredFutures, map[string]string{})
}

// GetHistoricalIndexOf /indexes/{market_name}/candles?resolution={resolution}&limit={limit}&start_time={start_time}&end_time={end_time}
// marketName, resolution: required
// params 1 element: index 0 == limit
// params 2 elements: index 0 == start_time, index 1 == end_time
// params 3 elements: index 0 == limit, index 1 == start_time, index 2 == end_time
func (header *FtxClientHeader) GetHistoricalIndexOf(marketName string, resolution string, params ...string) {
	queryMap := map[string]string{"resolution": resolution}
	if len(params) == 1 {
		queryMap["limit"] = params[0]
	} else if len(params) == 2 {
		queryMap["start_time"] = params[0]
		queryMap["end_time"] = params[1]
	} else if len(params) == 3 {
		queryMap["limit"] = params[0]
		queryMap["start_time"] = params[1]
		queryMap["end_time"] = params[2]
	}
	p := path.Join(FtxAPIIndexes, marketName, FtxAPIIndexesCandles)
	Do(header, "GET", p, queryMap)
}
