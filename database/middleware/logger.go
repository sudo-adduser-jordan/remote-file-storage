package middleware

import (
	"log"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()

			f(w, r)
		}
	}
}

// func logging(f http.HandlerFunc) http.HandlerFunc {
//     return func(w http.ResponseWriter, r *http.Request) {
//         log.Println(r.URL.Path)
//         f(w, r)
//     }
// }
