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
	ApiKey     string
	ApiSecret  string
	Subaccount string
	timestamp  int64
}

// generateReqSignature generate signature for http request
func (header *FtxClientHeader) generateReqSignature(httpMethod string, url string, payload string) string {
	header.timestamp = time.Now().Unix()

	var nonEncodedS string
	if httpMethod == "POST" {
		// FIXME payload should be "Request body (JSON-encoded)"
		nonEncodedS = fmt.Sprintf("%d%s%s%s", header.timestamp, httpMethod, url, payload)
	} else {
		nonEncodedS = fmt.Sprintf("%d%s%s", header.timestamp, httpMethod, url)
	}

	hashedS := hmac.New(sha256.New, []byte(header.ApiSecret))
	hashedS.Write([]byte(nonEncodedS))
	return hex.EncodeToString(hashedS.Sum(nil))

}

// func (header *FtxClientHeader) prepareRequest(req *fasthttp.Request, FtxAPIResource string, httpMethod string, queryMap map[string]string) {
func (header *FtxClientHeader) prepareRequest(req *fasthttp.Request, FtxAPIResource string, httpMethod string) {

	req.Header.SetMethod(httpMethod)
	// _url := generateApiUrl(FtxAPIResource, queryMap)
	_url := generateApiUrl(FtxAPIResource)

	req.SetRequestURI(_url)

	signature := header.generateReqSignature(httpMethod, _url, "")
	req.Header.Add("FTX-SIGN", signature)
	req.Header.Add("FTX-KEY", header.ApiKey)
	req.Header.Add("FTX-TS", string(header.timestamp))
}

// func Do(header *FtxClientHeader, httpMethod string, ftxResource string, queryMap map[string]string) {
func Do(header *FtxClientHeader, httpMethod string, ftxResource string) {

	req := fasthttp.AcquireRequest()
	// header.prepareRequest(req, ftxResource, httpMethod, queryMap)
	header.prepareRequest(req, ftxResource, httpMethod)

	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()

	client := &fasthttp.Client{}
	client.Do(req, resp)

	bodyBytes := resp.Body()
	println(string(bodyBytes))
}
