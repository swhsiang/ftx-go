package rest

func (header *FtxClientHeader) GetAccount() {
	Do(header, "GET", FtxAPIAccount, map[string]string{})
}

func (header *FtxClientHeader) GetPositions() {
	Do(header, "GET", FtxAPIPositions, map[string]string{})
}
