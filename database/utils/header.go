package utils

import "net/http"

// TODO
func SetDefaultHeaders(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
	return w, r
}

// TODO
func SetSecurityHeaders(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
	return w, r
}

// TODO
func SetLoginHeaders(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
	return w, r
}
