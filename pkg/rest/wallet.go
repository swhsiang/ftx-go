package rest

import "path"

func (header *FtxClientHeader) GetCoins() {
	p := path.Join(FtxAPIWallet, FtxAPIWalletCoins)
	Do(header, "GET", p, map[string]string{})
}

func (header *FtxClientHeader) GetBalances() {
	p := path.Join(FtxAPIWallet, FtxAPIWalletBalances)
	Do(header, "GET", p, map[string]string{})
}

func (header *FtxClientHeader) GetAllBalances() {
	p := path.Join(FtxAPIWallet, FtxAPIWalletAllBalances)
	Do(header, "GET", p, map[string]string{})
}

// GetDepositAddressOf
// https://docs.ftx.com/#get-deposit-address
func (header *FtxClientHeader) GetDepositAddressOf(coin string, method ...string) {
	p := path.Join(FtxAPIWallet, FtxAPIWalletDepositAddress, coin)
	queryMap := map[string]string{}
	if len(method) > 0 {
		queryMap["method"] = method[0]
	}
	Do(header, "GET", p, queryMap)
}

// GetDeposits
// params 1 element: index 0 == limit
// params 2 elements: index 0 == start_time, index 1 == end_time
// params 3 elements: index 0 == limit, index 1 == start_time, index 2 == end_time
// https://docs.ftx.com/#get-deposit-history
func (header *FtxClientHeader) GetDeposits(params ...string) {
	queryMap := map[string]string{}
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
	p := path.Join(FtxAPIWallet, FtxAPIWalletDeposits)
	Do(header, "GET", p, queryMap)
}

// GetWithdrawals
// params 1 element: index 0 == limit
// params 2 elements: index 0 == start_time, index 1 == end_time
// params 3 elements: index 0 == limit, index 1 == start_time, index 2 == end_time
// https://docs.ftx.com/#get-deposit-history
func (header *FtxClientHeader) GetWithdrawals(params ...string) {
	queryMap := map[string]string{}
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
	p := path.Join(FtxAPIWallet, FtxAPIWalletWithdrawals)
	Do(header, "GET", p, queryMap)
}

// https://docs.ftx.com/#request-withdrawal

// GetAirdrops
// params 1 element: index 0 == limit
// params 2 elements: index 0 == start_time, index 1 == end_time
// params 3 elements: index 0 == limit, index 1 == start_time, index 2 == end_time
// https://docs.ftx.com/#get-deposit-history
func (header *FtxClientHeader) GetAirdrops(params ...string) {
	queryMap := map[string]string{}
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
	p := path.Join(FtxAPIWallet, FtxAPIWalletAirdrops)
	Do(header, "GET", p, queryMap)
}

// GetSavedAddresses
// https://docs.ftx.com/#get-saved-addresses
func (header *FtxClientHeader) GetSavedAddresses(coin ...string) {
	p := path.Join(FtxAPIWallet, FtxAPIWalletSavedAddresses)
	queryMap := map[string]string{}
	if len(method) > 0 {
		queryMap["coin"] = method[0]
	}
	Do(header, "GET", p, queryMap)
}

// https://docs.ftx.com/#create-saved-addresses
// https://docs.ftx.com/#delete-saved-addresses
