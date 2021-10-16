package selects

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(urla, urlb string) (winner string, err error) {
	select {
	case <-ping(urla):
		return urla, nil
	case <-ping(urlb):
		return urlb, nil
	case <-time.After(10 * time.Second):
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
