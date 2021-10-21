package context

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
	Cancle()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		store.Cancle()
		fmt.Fprint(w, store.Fetch())
	}
}
