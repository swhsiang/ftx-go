# ftx-go

Golang package to interact with FTX Exchange APIs. PRs are welcome.

## Current Status

Under development. There are some issues on API authentication. Don't use it for production.

## TODO

* [ ] APIs
    * [x] account
    * [ ] market 
    * [x] wallet 
    * [ ] orders
    * [ ] convert
    * [ ] spot margin 
    * [ ] fills 
    * [ ] funding payments 
    * [ ] leveraged tokens 
    * [ ] options 
    * [ ] srm staking 
* [ ] Authentication. Authenticate all requests.
* [ ] Rate limit to api requests. Two packages at least. one for http and the other one for websocket.
* [ ] pagination for http requests
* [x] Defining data structures / copy data structures from ftxexchange/ftx
* [ ] deciding websocket package
