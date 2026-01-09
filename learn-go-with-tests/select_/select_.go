package select_

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimoeut = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimoeut)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	// Finding out which channel closes first
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
