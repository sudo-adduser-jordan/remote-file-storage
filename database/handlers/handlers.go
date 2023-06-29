package handlers

import (
	"fmt"
	"io"
	"log"
	"main/database"
	"main/utils"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type User struct {
	Username string
	Password string
}

// Create Account
func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	user := &User{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}
	hashed_password, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Fatal("Hash password error")
	}

	if database.ReadUser(user.Username) == user.Username {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprint(w, "User exists: "+user.Username)

	} else {
		if database.CreateUser(user.Username, hashed_password) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "User created: "+user.Username)
		}
	}
}

// Read Account - will assign cookie
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
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
	hash := database.ReadPassword(user.Username)

	if utils.CheckPasswordHash(user.Password, hash) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "User logged in: "+user.Username)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// Update Account
func Update(w http.ResponseWriter, r *http.Request) {
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
	new_user := &User{
		Username: r.FormValue("new_username"),
		Password: r.FormValue("new_passwrod"),
	}
	hashed_password, err := utils.HashPassword(new_user.Password)
	if err != nil {
		log.Fatal("Hash password error")
	}

	hash := database.ReadPassword(user.Username)
	if utils.CheckPasswordHash(user.Password, hash) {
		database.UpdateUser(user.Username, new_user.Username, hashed_password)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "User Updated: "+user.Username)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// Delete Account
func Delete(w http.ResponseWriter, r *http.Request) {
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

const MAX_UPLOAD_SIZE = 1024 * 1024 // 1MB
func Upload(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	user := &User{
		Username: r.FormValue("username"),
	}

	file_path := fmt.Sprintf("./store/%s/uploads/", user.Username)

	// 32 MB is the default used by FormFile()
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, "ParseMultipartForm error.", http.StatusInternalServerError)
		return
	}

	// Get a reference to the fileHeaders.
	// They are accessible only after ParseMultipartForm is called
	files := r.MultipartForm.File["file"]

	for _, fileHeader := range files {
		// Restrict the size of each uploaded file to 1MB.
		// To prevent the aggregate size from exceeding
		// a specified value, use the http.MaxBytesReader() method
		// before calling ParseMultipartForm()
		if fileHeader.Size > MAX_UPLOAD_SIZE {
			http.Error(w, fmt.Sprintf("The uploaded file is too big: %s. Please use an file less than 1MB in size", fileHeader.Filename), http.StatusBadRequest)
			return
		}

		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, "FileHeader Open error.", http.StatusInternalServerError)

			return
		}
		defer file.Close()

		buff := make([]byte, 512)
		_, err = file.Read(buff)
		if err != nil {
			http.Error(w, "File Read error.", http.StatusInternalServerError)
			return
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			http.Error(w, "File Seek error.", http.StatusInternalServerError)
			return
		}

		err = os.MkdirAll(file_path, os.ModePerm)
		if err != nil {
			http.Error(w, "Os MkdirAll error", http.StatusInternalServerError)
			return
		}

		array := strings.Split(fileHeader.Filename, ".")
		file_name := fmt.Sprintf(file_path+"%s%s", array[0], filepath.Ext(fileHeader.Filename))
		f, err := os.Create(file_name)
		if err != nil {
			http.Error(w, "Os Create Error", http.StatusBadRequest)
			return
		}
		defer f.Close()

		_, err = io.Copy(f, file)
		if err != nil {
			http.Error(w, "Io Copy Error", http.StatusBadRequest)
			// http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Upload successful")
}

// Fix
// Download
func Download(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	user := &User{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}

	file := fmt.Sprintf("./store/%s/", user.Username)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Download successful")

	http.ServeFile(w, r, file)
	// http.ServeContent(rw, r, "myfile", time.Now(), bytes.NewReader(data))
}
