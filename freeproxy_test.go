package freeproxy

import (
	"testing"
	"time"
)

func TestGetListOfAvailableProxies(t *testing.T) {
	meanDuration := time.Microsecond
	for i := 0; i < 10; i++ {
		t1 := time.Now()
		list, err := getListOfAvailableProxies()
		if err != nil {
			t.Error(err.Error())
		}
		if len(list) == 0 {
			t.Error("invalid list of proxies")
		}
		duration := time.Since(t1)
		meanDuration += duration
		t.Log("request done in ", duration)
	}
	meanDuration = meanDuration / 20
	t.Log("Mean duration:", meanDuration)
}
