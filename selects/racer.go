package selects

import (
	"net/http"
	"time"
)

func Racer(urla, urlb string) (winner string) {
	durationURLa := measureResponseTime(urla)
	durationURLb := measureResponseTime(urlb)

	if durationURLa < durationURLb {
		return urla
	}

	return urlb
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
