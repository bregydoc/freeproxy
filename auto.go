package freeproxy

import (
	"errors"
	"math/rand"
	"os"
	"time"
)

var errProxiesNotfound = errors.New("proxies filtered not found")

// GetRandomProxy ...
func GetRandomProxy(filter ...Filter) (*ProxyEntry, error) {
	proxies, err := GetProxies(filter...)
	if err != nil {
		return nil, err
	}

	if len(proxies) == 0 {
		return nil, errProxiesNotfound
	}

	rand.Seed(time.Now().Unix())

	i := rand.Intn(len(proxies))

	return proxies[i], nil
}

// SetRandomProxyOnEnvVariable set a random proxy on your env system,
// this function return the proxy setted and you can pass filter params for
// more control
func SetRandomProxyOnEnvVariable(filter ...Filter) (*ProxyEntry, error) {
	p, err := GetRandomProxy(filter...)
	if err != nil {
		return nil, err
	}

	proxyLink := getProxyLinkwithProxyEntry(p)
	os.Setenv("HTTP_PROXY", proxyLink)
	return p, nil
}

// SetProxyOnEnvVariable sets HTTP_PROXY with proxy entry pass as parameter
// with this function you can make all http requests on go with any proxy
func SetProxyOnEnvVariable(p *ProxyEntry) {
	proxyLink := getProxyLinkwithProxyEntry(p)
	os.Setenv("HTTP_PROXY", proxyLink)
}
