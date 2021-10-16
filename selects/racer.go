package selects

import (
	"net/http"
	"time"
)

func Racer(urla, urlb string) (winner string) {
	timeStartURLa := time.Now()
	http.Get(urla)
	durationURLa := time.Since(timeStartURLa)

	timeStartURLb := time.Now()
	http.Get(urlb)
	durationURLb := time.Since(timeStartURLb)

	if durationURLa < durationURLb {
		return urla
	}

	return urlb
}
