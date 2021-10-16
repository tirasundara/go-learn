package selects

import (
	"net/http"
)

func Racer(urla, urlb string) (winner string) {
	select {
	case <-ping(urla):
		return urla
	case <-ping(urlb):
		return urlb
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
