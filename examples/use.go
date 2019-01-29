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
