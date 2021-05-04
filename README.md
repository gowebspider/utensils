# utensils

## Demand

must

- go version >= 1.10

## Install

```bash
go get -u github.com/gowebspider/utensils/...
```

## Example

```go
package main

import (
	"fmt"
	"github.com/gowebspider/utensils/Head"
)

func main() {
	fmt.Println(Head.RandomMobileUserAgent())
	// output
	//Mozilla/5.0 (Linux; Android 9; Nexus 10 Build/NOF27C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.93 Safari/534.30
}


```

## TODO

[comment]: <> (-[x] Auto Encode)

[comment]: <> (-[x] fake user Agent)

[comment]: <> (-[x] simple pc requests)

[comment]: <> (-[ ] Proxy access)

[comment]: <> (-[ ] more...)

- [x] fake userAgent 
 
- [x] simple pc requests

- [ ] proxyRequests

- [ ] Async Requester

- [ ] url parse

- [ ] goroutine work pool

- [ ] ...


## Contact me

e-mail: wuzhipeng1289690157@mail.com

Wechat: chenxi-0719-chenxi

Personal website: blogs-payne.github.io

other: 公众号：积跬Coder

## Other

Welcome to issue star or PR, thank you