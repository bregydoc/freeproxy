package freeproxy

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func parseDuration(durationString string) time.Duration {
	r := regexp.MustCompile(`\w+ \w+ ago`)
	if !r.MatchString(durationString) {
		return 1 * time.Minute
	}

	chunks := strings.Split(durationString, " ")
	timeNumber := chunks[0]
	sizeDuration := chunks[1]

	finalDuration := 1 * time.Minute

	t, err := strconv.Atoi(timeNumber)
	if err != nil {
		return finalDuration
	}
	switch {
	case sizeDuration == "second" || sizeDuration == "seconds":
		finalDuration = time.Duration(t) * time.Second
		break
	case sizeDuration == "minute" || sizeDuration == "minutes":
		finalDuration = time.Duration(t) * time.Minute
		break
	case sizeDuration == "hour" || sizeDuration == "hours":
		finalDuration = time.Duration(t) * time.Hour
		break
	case sizeDuration == "day" || sizeDuration == "days":
		finalDuration = time.Duration(t) * time.Hour * 24
		break
	}

	return finalDuration
}

func getProxyLinkwithProxyEntry(p *ProxyEntry) string {
	base := "http://%s:%s"
	if p.HTTPS {
		base = "https://%s:%s"
	}
	return fmt.Sprintf(base, p.IP, p.Port)

}
