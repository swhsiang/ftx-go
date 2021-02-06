package rest

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

// FtxClientHeader header object
type FtxClientHeader struct {
	apiKey     string
	apiSecret  string
	subaccount string
	timestamp  int64
}

// GenerateReqSignature generate signature for http request
func (header *FtxClientHeader) GenerateReqSignature(httpMethod string, url string, payload string) string {
	header.timestamp = time.Now().Unix()

	var nonEncodedS string
	if httpMethod == "POST" {
		// FIXME payload should be "Request body (JSON-encoded)"
		nonEncodedS = fmt.Sprintf("%d%s%s%s", header.timestamp, httpMethod, url, payload)
	} else {
		nonEncodedS = fmt.Sprintf("%d%s%s", header.timestamp, httpMethod, url)
	}

	hashedS := hmac.New(sha256.New, []byte(header.apiSecret))
	hashedS.Write(nonEncodedS)
	return hex.EncodeToString(hashedS.Sum(nil))

}

func (header *FtxClientHeader) GetSubaccounts() {
	_url := GenerateApiUrl(FtxAPISubaccount)
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(_url)
	signature := header.GenerateReqSignature("GET", _url, "")
	req.Header.Add("FTX-SIGN", signature)
	req.Header.Add("FTX-KEY", header.apiKey)
	req.Header.Add("FTX-TS", string(header.timestamp))

	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	client.Do(req, resp)

	bodyBytes := resp.Body()
	println(string(bodyBytes))
	// // User-Agent: fasthttp
	// // Body:
}

func (header *FtxClientHeader) GetMarkets() {
	_url := GenerateApiUrl(FtxAPIMarkets)
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(_url)
	signature := header.GenerateReqSignature("GET", _url, "")
	req.Header.Add("FTX-SIGN", signature)
	req.Header.Add("FTX-KEY", header.apiKey)
	req.Header.Add("FTX-TS", string(header.timestamp))

	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	client.Do(req, resp)

	bodyBytes := resp.Body()
	println(string(bodyBytes))
}
