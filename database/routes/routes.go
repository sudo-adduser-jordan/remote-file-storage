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
	mux.HandleFunc("/register", middleware.Logger(handlers.Register))

	mux.HandleFunc("/login", middleware.Logger(handlers.Login))
	mux.HandleFunc("/update", middleware.Logger(handlers.Update))
	mux.HandleFunc("/delete", middleware.Logger(handlers.Delete))
	mux.HandleFunc("/upload", middleware.Logger(handlers.Upload))
	mux.HandleFunc("/download", middleware.Logger(handlers.Download))

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
		// IdleTimeout:  time.Minute,
		// ReadTimeout:  10 * time.Second,
		// WriteTimeout: 30 * time.Second,
	}

	fmt.Print(styles.BlueText("\n(∩｀-´)⊃"))
	fmt.Print(styles.YellowText("━ "))
	fmt.Print(styles.PurpleText("☆ﾟ . * ･ ｡ﾟ"))
	fmt.Print(" => http server started on ")
	fmt.Println(styles.GreenText(server.Addr))
	fmt.Println()

	server.ListenAndServe()
}
