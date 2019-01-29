package freeproxy

import (
	"fmt"
	"time"
)

// ProxyEntry describe an entry from https://free-proxy-list.net/
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

// Filter is an filter struct to get more accurate proxies
type Filter struct {
	MaxEntries      int
	SpecificCountry string
	OnlyHTTPS       bool
	MaxLastChecked  time.Duration
}

func (p *ProxyEntry) String() string {
	return fmt.Sprintf("%s:%s from %s | last checked: %s", p.IP, p.Port, p.Country, p.LastChecked)
}
