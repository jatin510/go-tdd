package selected

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a
	}

	return b
}

func measureResponseTime(url string) (responseDuration time.Duration) {
	start := time.Now()
	http.Get(url)
	responseDuration = time.Since(start)
	return
}
