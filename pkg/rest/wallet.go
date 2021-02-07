package rest

import "path"

func (header *FtxClientHeader) GetCoins() {
	p := path.Join(FtxAPIWallet, FtxAPICoins)
	Do(header, "GET", p, map[string]string{})
}
