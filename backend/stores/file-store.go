package stores

import (
	"errors"
	"fmt"
	"media-organizer/backend/datatypes"
	"media-organizer/backend/tools"
)

type FileStore struct {
	files      *datatypes.StringSet
	Extensions []string
	errors     map[string]error
}

func CreateFileStore(extensions []string) *FileStore {
	return &FileStore{
		files:      datatypes.NewStringSet(),
		Extensions: extensions,
		errors:     make(map[string]error),
	}
}

func (store *FileStore) addPath(path string) {
	pathReadable, err := tools.IsReadablePath(path)

	if err != nil {
		store.errors[path] = fmt.Errorf("Failed to validate path: %w", err)
	} else if !pathReadable {
		store.errors[path] = errors.New("Path is not a readable path")
	} else if !tools.PathHasValidExtension(path, store.Extensions) {
		store.errors[path] = errors.New("Path does not have the right extension")
	} else {
		store.files.Add(path)
	}
}

func (store *FileStore) addDirectoryFiles(path string) {
	files := tools.FindDirFiles(path, store.Extensions)

	if len(files) > 0 {
		store.AddFiles(files)
	}
}

func (store *FileStore) AddFiles(paths []string) []string {
	for _, path := range paths {
		store.addPath(path)
	}

	return store.Paths()
}

func (store *FileStore) AddDirectory(path string) []string {
	pathOk, validatePathErr := tools.IsReadablePath(path)

	if validatePathErr != nil {
		store.errors[path] = fmt.Errorf("Failed to validate path: %w", validatePathErr)
	} else if !pathOk {
		store.errors[path] = errors.New("Path is not a readable path")
	} else {
		store.addDirectoryFiles(path)
	}

	return store.Paths()
}

func (store *FileStore) Paths() []string {
	return store.files.Entries()
}

func (store *FileStore) Errors() map[string]string {
	stringErrorMap := make(map[string]string)

	for key, value := range store.errors {
		stringErrorMap[key] = value.Error()
	}

	return stringErrorMap
}
