package freeproxy

import (
	"github.com/gocolly/colly"
)

const freeProxyLink = "https://free-proxy-list.net"

// GetProxies return all proxies if filter is nil
func GetProxies(filter ...Filter) ([]*ProxyEntry, error) {
	if len(filter) == 0 {
		return getListOfAvailableProxies()
	}

	f := filter[0]
	return filteredProxies(f)
}

func filteredProxies(f Filter) ([]*ProxyEntry, error) {
	allProxies, err := getListOfAvailableProxies()
	if err != nil {
		return nil, err
	}

	finalFiltered := make([]*ProxyEntry, 0)
	for _, entry := range allProxies {
		if f.MaxEntries != 0 {
			if len(finalFiltered)+1 > f.MaxEntries {
				break
			}
		}
		if f.MaxLastChecked != 0 {
			if f.MaxLastChecked < entry.LastChecked {
				continue
			}
		}
		if f.OnlyHTTPS {
			if !entry.HTTPS {
				continue
			}
		}
		if f.SpecificCountry != "" {
			if f.SpecificCountry != entry.Code {
				continue
			}
		}
		finalFiltered = append(finalFiltered, entry)
	}

	return finalFiltered, nil
}

func getListOfAvailableProxies() ([]*ProxyEntry, error) {
	c := colly.NewCollector()

	allEntries := make([]*ProxyEntry, 0)
	c.OnHTML("table#proxylisttable", func(e *colly.HTMLElement) {
		e.ForEachWithBreak("tbody", func(i int, body *colly.HTMLElement) bool {
			// Only one time
			body.ForEach("tr", func(j int, tr *colly.HTMLElement) {
				// cols := tr.DOM.Find("td")
				// ip := cols.First().Text()
				// port := cols.
				// fmt.Println(ip, port)
				entry := new(ProxyEntry)
				tr.ForEach("td", func(q int, td *colly.HTMLElement) {
					switch q {
					case 0: // IP
						entry.IP = td.Text
					case 1: // Port
						entry.Port = td.Text
					case 2: // Code
						entry.Code = td.Text
					case 3: // Country
						entry.Country = td.Text
					case 4: // Anonymity
						entry.Anonymity = td.Text
					case 5: // Google
						entry.Google = td.Text == "yes"
					case 6: // HTTPS
						entry.HTTPS = td.Text == "yes"
					case 7: // LastChecked
						d := parseDuration(td.Text)
						entry.LastChecked = d
					}
				})
				allEntries = append(allEntries, entry)
			})
			return false
		})

	})

	c.Visit(freeProxyLink)

	return allEntries, nil
}
