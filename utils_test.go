package freeproxy

import (
	"testing"
	"time"
)

func TestParseDuration(t *testing.T) {
	table := map[string]time.Duration{
		"1 second ago":    1 * time.Second,
		"10 seconds ago":  10 * time.Second,
		"100 seconds ago": 100 * time.Second,
		"1 minute ago":    1 * time.Minute,
		"3 minutes ago":   3 * time.Minute,
		"20 minutes ago":  20 * time.Minute,
		"1 hour ago":      1 * time.Hour,
		"6 hours ago":     6 * time.Hour,
		"15 hours ago":    15 * time.Hour,
		"100 hours ago":   100 * time.Hour,
		// "10 xxxx ago":     1 * time.Minute,
		// "x asd ago":       1 * time.Minute,
	}

	for text, dur := range table {
		result := parseDuration(text)
		if dur != result {
			t.Error("error at try to parse")
			t.Errorf("expected: %s, result: %s", dur, result)
		}
	}
}
