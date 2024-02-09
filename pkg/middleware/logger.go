package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// LoggerMW - middleware function for logging requests info
func LoggerMW(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)

			log.Println(
				strings.Join([]string{
					r.Method,
					r.RemoteAddr,
					r.URL.EscapedPath(),
					fmt.Sprint(time.Since(start)),
				}, " "),
			)
		},
	)
}
