package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"main/database"
	"main/utils"
	"net/http"
	"os"
)

type User struct {
	Username string
	Password string
}

// Create Account
func Register(w http.ResponseWriter, r *http.Request) {
	// Check method
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get Values
	user := &User{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}
	hashed_password, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Fatal("Hash password error")
	}

	// Check if values exist
	if database.ReadUser(user.Username) == user.Username {
		// w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.WriteHeader(http.StatusConflict)
		fmt.Fprint(w, "User exists: "+user.Username)
	} else {
		if database.CreateUser(user.Username, hashed_password) {
			// w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w, r = utils.CreateCookie(w, r, user.Username)
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "User created: "+user.Username)
		}
	}
}

// Read Account
func Login(w http.ResponseWriter, r *http.Request) {
	// Check method
	if r.Method != "POST" {
		http.Error(
			w,
			"Method not allowed",
			http.StatusMethodNotAllowed,
		)
		return
	}

	// Get values
	user := &User{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}
	hash := database.ReadPassword(user.Username)

	// Check password
	if utils.CheckPasswordHash(user.Password, hash) {
		// w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w, r = utils.CreateCookie(w, r, user.Username)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "User logged in: "+user.Username)
	} else {
		// w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.WriteHeader(http.StatusBadRequest)
	}
}

// Update Account TODO
func Update(w http.ResponseWriter, r *http.Request) {
	// w, r = utils.ReadCookie(w, r)
	// w, r = utils.UpdateCookie(w, r)

	// Check method
	if r.Method != "PUT" {
		http.Error(
			w,
			"Method not allowed",
			http.StatusMethodNotAllowed,
		)
		return
	}

	// Get values
	user := &User{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}
	new_user := &User{
		Username: r.FormValue("new_username"),
		Password: r.FormValue("new_passwrod"),
	}
	hashed_password, err := utils.HashPassword(new_user.Password)
	if err != nil {
		log.Fatal("Hash password error")
	}

	// Check Password
	hash := database.ReadPassword(user.Username)
	if utils.CheckPasswordHash(user.Password, hash) {
		database.UpdateUser(user.Username, new_user.Username, hashed_password)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "User Updated: "+user.Username)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// Delete Account TODO
func Delete(w http.ResponseWriter, r *http.Request) {
	// w, r = utils.ReadCookie(w, r)
	// w, r = utils.DeleteCookie(w, r)

	if r.Method != "PUT" {
		http.Error(
			w,
			"Method not allowed",
			http.StatusMethodNotAllowed,
		)
		return
	}
	user := &User{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}
	password := database.ReadPassword(user.Username)

	if utils.CheckPasswordHash(user.Password, password) {
		database.DeleteUser(user.Username)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// Upload Files TODO
func Upload(w http.ResponseWriter, r *http.Request) {
	// w, r = utils.ReadCookie(w, r)
	// w, r = utils.UpdateCookie(w, r)

	// Check method
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get values
	username := r.URL.Query().Get("username")

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	form_data := r.MultipartForm
	files := form_data.File["file"]

	// Write files
	for index := range files {

		file, err := files[index].Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		path := fmt.Sprintf("./store/%s/uploads/", username)
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		dst, err := os.Create(path + files[index].Filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

// Download File TODO
func Download(w http.ResponseWriter, r *http.Request) {
	// w, r = utils.ReadCookie(w, r)
	// w, r = utils.UpdateCookie(w, r)

	// Check method
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get values
	username := r.URL.Query().Get("username")
	file := r.URL.Query().Get("file")
	path := fmt.Sprintf("./store/%s/uploads/%s", username, file)

	// Send file
	// ioutil.ReadDir()
	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
}

// Get Download File list TODO
func FileList(w http.ResponseWriter, r *http.Request) {
	// w, r = utils.ReadCookie(w, r)
	// w, r = utils.UpdateCookie(w, r)

	// Check method
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get values
	username := r.URL.Query().Get("username")
	p := utils.ReadStore(username)

	// Send json list
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}