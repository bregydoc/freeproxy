package main

import (
	"fmt"
	"time"

	"github.com/bregydoc/freeproxy"
)

func main() {
	proxies, _ := freeproxy.GetProxies(freeproxy.Filter{
		MaxEntries:     30,
		MaxLastChecked: 10 * time.Second,
	})
	fmt.Println(len(proxies))
}
