package stores

import (
	"image"
	"path/filepath"
	"regexp"
	"strings"

	"media-organizer/backend/tools"
)

type FileEntry struct {
	Path         string      `json:"path"`
	Name         string      `json:"name"`
	OriginalName string      `json:"orgName"`
	Extension    string      `json:"extension"`
	Error        error       `json:"error"`
	ReadOnly     bool        `json:"readOnly"`
	Thumbnail    image.Image `json:"-"`
}

type FileStore struct {
	files map[string]*FileEntry
}

func CreateFileStore() *FileStore {
	return &FileStore{
		files: make(map[string]*FileEntry),
	}
}

func createPathEntry(path string, readOnly bool, err error) *FileEntry {
	ext := filepath.Ext(path)
	baseName := strings.TrimSuffix(filepath.Base(path), ext)

	// Remove the ending from the file name that indicates a duplicate file (eg. "file_2" or "file (copy)")
	duplicateFileEndingExp := regexp.MustCompile(`(?:_\d{1,3}|\s+\((?:\d{1,3}|.*?copy)\))\s*$`)
	name := duplicateFileEndingExp.ReplaceAllString(baseName, "")

	entry := &FileEntry{
		Path:         path,
		Name:         name,
		OriginalName: baseName,
		Extension:    ext[1:],
		ReadOnly:     readOnly,
		Thumbnail:    nil,
		Error:        err,
	}

	return entry
}

func (store *FileStore) addFileEntry(entry *FileEntry) bool {
	_, entryExists := store.files[entry.Path]

	if !entryExists {
		store.files[entry.Path] = entry
	}

	return entryExists
}

func (store *FileStore) AddFile(filePath string) bool {
	inf, err := tools.GetPathInfo(filePath)
	return store.addFileEntry(createPathEntry(filePath, inf.ReadOnly, err))
}

func (store *FileStore) AddErrorPath(error tools.PathError) bool {
	return store.addFileEntry(createPathEntry(error.Path, true, error.Err))
}

func (store *FileStore) GetFiles() []*FileEntry {
	files := make([]*FileEntry, 0)

	for _, entry := range store.files {
		files = append(files, entry)
	}

	return files
}

func (store *FileStore) AddFiles(filePaths []string) []string {
	addedFiles := make([]string, 0)

	for _, path := range filePaths {
		added := store.AddFile(path)

		if added {
			addedFiles = append(addedFiles, path)
		}
	}

	return addedFiles
}

func (store *FileStore) AddErrorPaths(errors []tools.PathError) []string {
	addedFiles := make([]string, 0)

	for _, err := range errors {
		added := store.AddErrorPath(err)

		if added {
			addedFiles = append(addedFiles, err.Path)
		}
	}

	return addedFiles
}

func (store *FileStore) GetOrderedFileList() []*FileEntry {
	paths := make([]*FileEntry, 0, len(store.files))

	for _, entry := range store.files {
		paths = append(paths, entry)
	}

	return paths
}
