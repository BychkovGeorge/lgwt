package concurrency

type WebsiteChecker func(string) bool

type mapValue struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	c := make(chan mapValue)
	results := make(map[string]bool)

	for _, url := range urls {
		go func() {
			c <- mapValue{string: url, bool: wc(url)}
		}()
	}

	for range urls {
		newVal := <-c
		results[newVal.string] = newVal.bool
	}
	return results
}
