# Free Proxy ![build](https://travis-ci.org/bregydoc/freeproxy.svg?branch=master)
Proxy a la carte

This simple utility was created to use different proxies in your project. The list of proxies is extract from [Free Proxy list](https://free-proxy-list.net/) using scrapping. Freeproxy implements util functions for a rapid use of the scraped proxy list, you can make proxy connections with env variable or http.Client.

## Install
Only use go get

	go get -u -v github.com/bregydoc/freeproxy

(Remove "-v" if you don't need it)

## Use
ProxyEntry struct describe an entry of free-proxy-list.net

```go
type ProxyEntry struct {
	IP          string
	Port        string
	Code        string
	Country     string
	Anonymity   string
	Google      bool
	HTTPS       bool
	LastChecked time.Duration
}
```

with ```freeproxy.GetProxies()``` you can return all list of available proxies in [Free proxy list](https://free-proxy-list.net/), if you want to can to pass a ```freeproxy.Filter``` object to get more control on your proxy list. The form of ```freeproxy.Filter``` is described below.

```go
type Filter struct {
	MaxEntries      int
	SpecificCountry string // You need use country code (e.g. "US")
	OnlyHTTPS       bool
	MaxLastChecked  time.Duration
}
```

Example of use of ```freeproxy.GetProxies()```

```go
import "github.com/bregydoc/freeproxy"
// ...
proxies, err := freeproxy.GetProxies()
// err handle
fmt.Println(proxies[0])
// 201.140.240.9:46338 from Brazil | last checked: 6s
```

```go
proxies, err := freeproxy.GetProxies(freeproxy.Filter{
	SpecificCountry: "US",
	OnlyHTTPS: true,
	MaxLastChecked: 1 * time.Minute,
})
// err handle
fmt.Println(proxies[0])
// 40.117.231.19	:3128 from United States | last checked: 18s
```

Freeproxy has other util functions

```go
freeproxy.GetRandomProxy(filter ...Filter)
freeproxy.SetRandomProxyOnEnvVariable(filter ...Filter)
freeproxy.SetProxyOnEnvVariable(p *ProxyEntry)
freeproxy.GetProxyURL(p *ProxyEntry)
freeproxy.GetRandomProxyURL(filter ...Filter)
freeproxy.GetHTTPClientWithProxy(p *ProxyEntry)
freeproxy.GetHTTPClientWithRandomProxy(filter ...Filter)
```

## Complete Example
```go 
package main

import (
	"fmt"
	"time"

	"github.com/bregydoc/freeproxy"
)

func main() {
	proxies, _ := freeproxy.GetProxies(freeproxy.Filter{
		MaxEntries:     8,
		MaxLastChecked: 1 * time.Minute,
	})
	for _, proxy := range proxies {
		fmt.Println(proxy)
	}
	// 91.219.56.221:8080 from Russian Federation | last checked: 4s
	// 101.255.36.233:55839 from Indonesia | last checked: 1m0s
	// 86.123.166.109:8080 from Romania | last checked: 1m0s
	// 88.204.59.177:32666 from Russian Federation | last checked: 1m0s
	// 116.0.6.254:40723 from Indonesia | last checked: 1m0s
	// 109.86.41.111:45308 from Ukraine | last checked: 1m0s
	// 176.31.141.20:21231 from France | last checked: 1m0s
	// 149.156.171.40:58946 from Poland | last checked: 1m0s
}
```

## TODO
- [ ] Improve documentation
- [ ] Add more util functions

## Contribute
All contribution are welcome, only open an issue.
