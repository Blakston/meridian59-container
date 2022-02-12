package patcher

import (
	"io/ioutil"
	"path/filepath"
)

// GetFiles ...
func GetFiles(path string) (files []string) {

	fileInfo, err := ioutil.ReadDir(path)
	if err != nil {
		return files
	}

	for _, fi := range fileInfo {
		if fi.IsDir() {
			files = append(files, GetFiles(filepath.Join(path, fi.Name()))...)
		} else {
			file := filepath.Join(path, fi.Name())
			files = append(files, file)
		}
	}
	return
}
