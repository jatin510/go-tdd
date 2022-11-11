package selected

import (
	"net/http"
)

func Racer(a, b string) (winner string) {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

// func measureResponseTime(url string) (responseDuration time.Duration) {
// 	start := time.Now()
// 	http.Get(url)
// 	responseDuration = time.Since(start)
// 	return
// }

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
