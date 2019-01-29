package freeproxy

import (
	"errors"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"time"
)

var errProxiesNotfound = errors.New("proxies filtered not found")

// GetRandomProxy returns a random proxy, you can pass filter options
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
// this function return the proxy setted and you can pass filter options for
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

// GetProxyURL return a parsed URL native Object with proxyEntry struct
func GetProxyURL(p *ProxyEntry) (*url.URL, error) {
	u := getProxyLinkwithProxyEntry(p)
	return url.Parse(u)
}

// GetRandomProxyURL return a random url go native
func GetRandomProxyURL(filter ...Filter) (*url.URL, error) {
	p, err := GetRandomProxy(filter...)
	if err != nil {
		return nil, err
	}

	return GetProxyURL(p)
}

// GetHTTPClientWithProxy return a http client with specific proxy entry
func GetHTTPClientWithProxy(p *ProxyEntry) (*http.Client, error) {
	proxyURL, err := GetProxyURL(p)
	if err != nil {
		return nil, err
	}
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
	return client, nil
}

// GetHTTPClientWithRandomProxy return a http client bsaed on random proxy
func GetHTTPClientWithRandomProxy(filter ...Filter) (*http.Client, *ProxyEntry, error) {
	p, err := GetRandomProxy(filter...)
	if err != nil {
		return nil, nil, err
	}

	client, err := GetHTTPClientWithProxy(p)
	if err != nil {
		return nil, nil, err
	}
	return client, p, err
}
