package rest

import "path"

func (header *FtxClientHeader) GetMarkets() {
	// Do(header, "GET", FtxAPIMarkets, map[string]string{})
	Do(header, "GET", FtxAPIMarkets)

}

func (header *FtxClientHeader) GetMarketOf(marketName string) {
	p := path.Join(FtxAPIMarkets, marketName)
	// Do(header, "GET", p, map[string]string{})
	Do(header, "GET", p)

}
