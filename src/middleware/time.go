package middleware

import (
	"log"
	"logger"
	"net/http"
	"time"
)

/*
Time function
is a middleware to measure execution times
will be deleted in production
*/
func Time() Middleware {
	return func(h http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			go logger.IncrementRequests()
			defer stopTimer(start, w, r)
			h.ServeHTTP(w, r)
		})
	}
}

func stopTimer(start time.Time, w http.ResponseWriter, r *http.Request) {
	total := time.Since(start)
	go logger.UpdateLatency(total)
	log.Println(total, r.URL.String())
}
