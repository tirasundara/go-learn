package concurrency

type WebsiteChecker func(url string) bool

type result struct {
	url  string
	isUp bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{url: u, isUp: wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.url] = result.isUp
	}

	return results
}
