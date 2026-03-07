package tools

import (
	"errors"
	"fmt"
	"os"
)

type PathInfo struct {
	ReadOnly bool
	IsDir    bool
}

var ErrPathNotFound = errors.New("path does not exist")
var ErrPathNotReadable = errors.New("path is not readable")

func GetPathInfo(path string) (*PathInfo, error) {
	info, err := os.Stat(path)

	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("%q: %w", path, ErrPathNotFound)
		}

		return nil, fmt.Errorf("Failed to stat %q: %w", path, err)
	}

	perm := info.Mode().Perm() & 0600
	readable := perm&0400 == 0400

	if !readable {
		return nil, fmt.Errorf("%q: %w", path, ErrPathNotReadable)
	}

	editable := perm&0200 == 0200

	return &PathInfo{
		ReadOnly: !editable,
		IsDir:    info.IsDir(),
	}, nil
}

func FindDirFilePaths(dirPath string, extensions []string) ([]string, []PathError) {
	dirWalker := CreateDirWalker(extensions)
	return dirWalker(dirPath)
}

// func PathHasValidExtension(path string, extensions []string) bool {
// 	pathExt := filepath.Ext(path)
// 	return slices.Contains(extensions, pathExt[1:])
// }
