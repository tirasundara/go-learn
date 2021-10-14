package concurrency

import (
	"fmt"
	"net/http"
)

// CheckWebsite returns true if the URL returns a 200 status code, false otherwise.
func CheckWebsite(url string) bool {
	res, err := http.Head(url)

	if err != nil {
		fmt.Println("err: ", err)
		return false
	}

	if res.StatusCode != http.StatusOK {
		fmt.Println("StatusCode: ", res.StatusCode)
		return false
	}

	return true
}
