package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// type File struct {
// 	Name string
// 	Size int64
// 	Path string
// }

type File struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
	Path string `json:"path"`
}

func ReadStore(username string) []File {
	files := []File{}

	counter := 0
	root := fmt.Sprintf("./store/%s", username)
	err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {

		files = append(files, File{
			Name: f.Name(),
			Size: f.Size(),
			Path: path,
		})

		counter++
		return err
	})

	if err != nil {
		log.Fatal(err)
	}

    return files
}

// TODO
func DeleteStore() {}
