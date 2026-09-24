package _select

import (
	"errors"
	"net/http"
	"time"
)

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		response, err := http.Get(url)
		if err == nil {
			response.Body.Close()
		}
		close(ch)
	}()
	return ch
}

func Racer(url1, url2 string, timeout time.Duration) (string, error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(timeout):
		return "", errors.New("timeout")
	}
}
