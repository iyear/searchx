package utils

import (
	"os"
	"path/filepath"
	"strings"
)

type fs struct{}

var FS = fs{}

func (f fs) PathExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// GetFileName without extension
func (f fs) GetFileName(path string) string {
	return strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
}
