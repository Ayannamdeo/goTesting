package concurrency

type (
	WebsiteChecker func(string) bool
	Result         struct {
		string
		bool
	}
)

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan Result)

	for _, url := range urls {
		go func() {
			resultChannel <- Result{url, wc(url)}
		}()
	}

	for range urls {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
