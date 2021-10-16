package selects

import (
	"fmt"
	"net/http"
	"time"
)

const tenSecondTimeout = 10 * time.Second

func Racer(urla, urlb string) (winner string, err error) {
	return ConfigurableRacer(urla, urlb, tenSecondTimeout)
}

func ConfigurableRacer(urla, urlb string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(urla):
		return urla, nil
	case <-ping(urlb):
		return urlb, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", urla, urlb)
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
