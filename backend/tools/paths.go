package tools

import (
	"media-organizer/backend/datatypes"
	"os"
	"path/filepath"
)

func IsReadablePath(path string) (bool, error) {
	info, err := os.Stat(path)

	if err != nil {
		return false, err
	}

	return info.Mode().Perm()&0400 != 0, nil
}

func IsDir(path string) (bool, error) {
	info, err := os.Stat(path)

	if err != nil {
		return false, err
	}

	return info.IsDir(), nil
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

	for _, ext := range extensions {
		if ext == pathExt {
			return true
		}
	}

	return false
}
