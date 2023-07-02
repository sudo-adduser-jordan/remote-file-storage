package utils

import "net/http"

func SetDefaultHeaders(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
	return w, r
}
