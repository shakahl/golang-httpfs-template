package vfstemplate

import (
	"io/ioutil"
	"net/http"
	"os"
)

// ReadDir reads the contents of the directory associated with file and
// returns a slice of FileInfo values in directory order.
func ReadDir(fs http.FileSystem, name string) ([]os.FileInfo, error) {
	f, err := fs.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return f.Readdir(-1)
}

// Stat returns the FileInfo structure describing file.
func Stat(fs http.FileSystem, name string) (os.FileInfo, error) {
	f, err := fs.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return f.Stat()
}

// ReadFile reads the file named by path from fs and returns the contents.
func ReadFile(fs http.FileSystem, path string) ([]byte, error) {
	rc, err := fs.Open(path)
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return ioutil.ReadAll(rc)
}

// ReadFileString reads the file named by path from fs and returns the contents.
func ReadFileString(fs http.FileSystem, path string) (string, error) {
	buf, err := ReadFile(fs, path)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

// Exists reports whether the named file or directory exists in http.FileSystem.
func Exists(fs http.FileSystem, name string) bool {
	if _, err := Stat(fs, name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
