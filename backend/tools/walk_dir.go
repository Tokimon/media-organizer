package tools

import (
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"

	"media-organizer/backend/datatypes"
)

type PathError struct {
	Path string `json:"path"`
	Err  error  `json:"error"`
}

func (e PathError) Error() string {
	return fmt.Sprintf("path %s: %v", e.Path, e.Error)
}

// Creates a method to walk a given diretory path to find the files within
// that has an extension that is within one of the given extensions.
//
// Returns: (Files, ErrorPaths)
func CreateDirWalker(extensions []string) func(string) ([]string, []PathError) {
	extSet := datatypes.NewStringSet()

	for _, ext := range extensions {
		extSet.Add("." + ext)
	}

	return func(dirPath string) ([]string, []PathError) {
		files := make([]string, 0)
		errorPaths := make([]PathError, 0)

		err := filepath.WalkDir(dirPath, func(path string, info fs.DirEntry, err error) error {
			if err != nil {
				errorPaths = append(errorPaths, PathError{Path: path, Err: err})
				return fs.SkipDir
			}

			if !info.IsDir() && extSet.Has(filepath.Ext(path)) {
				files = append(files, path)
			}

			return nil
		})

		if err != nil && !errors.Is(err, fs.SkipDir) {
			errorPaths = append(errorPaths, PathError{Path: dirPath, Err: err})
		}

		return files, errorPaths
	}
}
