package routes

import (
	"fmt"
	"main/handlers"
	"main/middleware"
	"net/http"

	"github.com/sudo-adduser-jordan/Toolchain/Go/styles"
)

func SetupRoutes() {

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./static")))
	mux.HandleFunc("/login", handlers.Login)
	mux.HandleFunc("/register", middleware.Logger(handlers.Register))
	mux.HandleFunc("/delete", handlers.Delete)
	mux.HandleFunc("/upload", handlers.Upload)
	mux.HandleFunc("/download", handlers.Download)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Print("\n(∩｀-´)⊃━ ☆ﾟ . * ･ ｡ﾟ => http server started on ")
	fmt.Println(styles.GreenText(server.Addr))
	server.ListenAndServe()
}
