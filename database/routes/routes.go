package routes

import (
	"main/handlers"
	"net/http"
)

func SetupRoutes() {

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./static")))
	mux.HandleFunc("/login", handlers.Login)
	mux.HandleFunc("/upload", handlers.Upload)
	mux.HandleFunc("/download", handlers.Download)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
		// ErrorLog:  nil,
		// TLSConfig: nil,
	}
	server.ListenAndServe()
}
