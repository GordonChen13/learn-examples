package main

type WebsiteChecker func(string) bool

func checkWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	// fatal error: concurrent map writes
	results := make(map[string]bool)

	for _, url := range urls {
		go func(u string) {
			results[u] = wc(u)
		}(url)
	}

	return results
}
