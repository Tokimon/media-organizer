package tools

import (
	"errors"
	"media-organizer/backend/datatypes"
	"os"
	"path/filepath"
)

type PathInfo struct {
	ReadOnly bool
	IsDir    bool
}

func GetPathInfo(path string) (*PathInfo, error) {
	info, err := os.Stat(path)

	if err != nil {
		return nil, err
	}

	perm := info.Mode().Perm() & 0600
	readable := perm&0400 == 0400

	if !readable {
		return nil, errors.New("Path not readable")
	}

	editable := perm&0200 == 0200

	return &PathInfo{
		ReadOnly: !editable,
		IsDir:    info.IsDir(),
	}, nil
}

func FindDirFiles(dirPath string, extensions []string) []string {
	set := datatypes.NewStringSet()

	for _, ext := range extensions {
		set.Add("." + ext)
	}

	walker := DirWalker{
		files:  []string{},
		extSet: *set,
	}

	walker.walk(dirPath)

	return walker.files
}

func PathHasValidExtension(path string, extensions []string) bool {
	pathExt := filepath.Ext(path)
	pathExt = pathExt[1:]

	for _, ext := range extensions {
		if ext == pathExt {
			return true
		}
	}

	return false
}
