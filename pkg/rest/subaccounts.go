package rest

func (header *FtxClientHeader) GetSubaccounts() {
	Do(header, "GET", FtxAPISubaccount, map[string]string{})
}
