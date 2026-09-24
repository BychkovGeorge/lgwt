package concurrency

type WebsiteChecker func(string) bool

type mapValue struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	c := make(chan mapValue)
	results := make(map[string]bool)

	for i := range urls {
		go func() {
			c <- mapValue{string: urls[i], bool: wc(urls[i])}
		}()
	}

	for range urls {
		newVal := <-c
		results[newVal.string] = newVal.bool
	}
	return results
}
