package handlers

import (
	"fmt"
	"log"
	"main/database"
	"main/utils"
	"net/http"
)

type User struct {
	Username string
	Password string
}

// Login
func Login(w http.ResponseWriter, r *http.Request) {
	user := &User{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}

	password := database.ReadPassword(user.Username)

	if utils.CheckPasswordHash(user.Username, password) {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

}

// Create Account
func Register(w http.ResponseWriter, r *http.Request) {
	user := &User{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}
	hashed_password, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Fatal("Hash password error")
	}

	if database.ReadUser(user.Username) {
		fmt.Println("User exists.")
		w.WriteHeader(http.StatusConflict)
	}

	if database.CreateUser(user.Username, hashed_password) {
		fmt.Println("User created: " + user.Username)
		w.WriteHeader(http.StatusOK)
	}
}

// Upload
func Upload(w http.ResponseWriter, r *http.Request) {
	// return http.StatusOK
}

// Download
func Download(w http.ResponseWriter, r *http.Request) {
	// return http.StatusOK
}
