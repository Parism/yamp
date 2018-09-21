package middleware

import (
	"log"
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
			defer stopTimer(start, w, r)
			h.ServeHTTP(w, r)
		})
	}
}

func stopTimer(start time.Time, w http.ResponseWriter, r *http.Request) {
	total := time.Since(start)
	log.Println(total, r.URL.String())
}
