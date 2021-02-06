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

func GenerateApiUrl(endpoint string) string {
	u, err := url.Parse(FtxServer)
	if err != nil {
		panic("invalid url")
	}

	u.Path = path.Join(u.Path, endpoint)
	return u.String()
}
