package tools

import (
	"media-organizer/backend/datatypes"
	"os"
	"path/filepath"
)

type DirWalker struct {
	files  []string
	extSet datatypes.StringSet
}

func (walker *DirWalker) walk(dirPath string) {
	filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// // Only match files, not directories
		if info.IsDir() {
			if path != dirPath {
				walker.walk(path)
			}
		} else if walker.extSet.Has(filepath.Ext(path)) {
			walker.files = append(walker.files, path)
		}

		return nil
	})
}
