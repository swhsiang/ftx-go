package rest

import (
	"net/url"
	"path"
)

const (
	FtxServer        = "https://ftx.com/api"
	FtxAPISubaccount = "subaccounts"
	FtxAPIMarkets    = "markets"
)

var baseUrl *url.URL

func init() {
	var err error
	baseUrl, err = url.Parse(FtxServer)
	if err != nil {
		panic("invalid url")
	}
}

func generateApiUrl(endpoint string) string {
	u := baseUrl
	u.Path = path.Join(baseUrl.Path, endpoint)
	return u.String()
}
