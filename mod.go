package mod

import (
	"errors"
	"os"
	"path/filepath"
)

var (
	ErrNotFound = errors.New("not found")
)

// ModFilePath look for go.mod at dir and parent dirs
// ModFilePath("/a/b") will check
// - /a/b/go.mod
// - /a/go.mod
// - /go.mod
// and return first found file path or ErrNotFound if not found
func ModFilePath(startingDir string) (string, error) {
	return modFilePath(startingDir, startingDir, true)
}

func modFilePath(dir, prevDir string, first bool) (string, error) {
	if !first && dir == prevDir {
		return "", ErrNotFound
	}

	_, err := os.Stat(filepath.Join(dir, "go.mod"))
	if os.IsNotExist(err) {
		return modFilePath(filepath.Dir(dir), dir, false)
	}
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, "go.mod"), nil
}
