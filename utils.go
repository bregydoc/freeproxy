package freeproxy

import (
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

	finalDuration := time.Minute

	t, err := strconv.Atoi(timeNumber)
	if err != nil {
		return finalDuration
	}
	switch sizeDuration {
	case "second":
	case "seconds":
		finalDuration = time.Duration(t) * time.Second
		break
	case "minute":
	case "minutes":
		finalDuration = time.Duration(t) * time.Minute
		break
	case "hour":
	case "hours":
		finalDuration = time.Duration(t) * time.Hour
		break
	case "day":
	case "days":
		finalDuration = time.Duration(t) * time.Hour * 24
		break
	}

	return finalDuration
}
