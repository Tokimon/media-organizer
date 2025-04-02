package api

import (
	"context"
	"media-organizer/backend/tools"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type ApiFile struct {
	Path      string `json:"path"`
	Name      string `json:"name"`
	Extension string `json:"extension"`
	Error     error  `json:"error"`
	ReadOnly  bool   `json:"readOnly"`
}

type Api struct {
	ctx        context.Context
	Extensions []string
}

func CreateApi() *Api {
	return &Api{
		Extensions: slices.Concat(tools.DefaultExtensions.Images(), tools.DefaultExtensions.Video),
	}
}

func pathToApiFile(path string) ApiFile {
	inf, err := tools.GetPathInfo(path)

	ext := filepath.Ext(path)
	baseName := strings.TrimSuffix(filepath.Base(path), ext)
	exp := regexp.MustCompile(`(?:_\d{1,3}|\s+\((?:\d{1,3}|.*?copy)\))\s*$`)
	name := exp.ReplaceAllString(baseName, "")

	return ApiFile{
		Path:      path,
		Name:      name,
		Extension: ext,
		ReadOnly:  err != nil || inf.ReadOnly,
		Error:     err,
	}
}

func pathsToApiFiles(paths []string) []ApiFile {
	files := make([]ApiFile, len(paths))

	for i, path := range paths {
		files[i] = pathToApiFile(path)
	}

	return files
}

func (api *Api) Init(ctx context.Context) {
	api.ctx = ctx
}

func (api *Api) SelectFiles() ([]ApiFile, error) {
	exts := "*." + strings.Join(api.Extensions, ";*.")

	paths, err := runtime.OpenMultipleFilesDialog(api.ctx, runtime.OpenDialogOptions{
		Title: "Select files to add",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Select Files to Add",
				Pattern:     exts,
			},
		},
	})

	if err != nil {
		return nil, err
	}

	return pathsToApiFiles(paths), nil
}

func (api *Api) SelectDirectory() ([]ApiFile, error) {
	path, err := runtime.OpenDirectoryDialog(api.ctx, runtime.OpenDialogOptions{
		Title: "Select Folder to add",
	})

	if err != nil {
		return nil, err
	}

	paths := tools.FindDirFiles(path, api.Extensions)

	return pathsToApiFiles(paths), nil
}

func (api *Api) SelectFromDrop(x int, y int, paths []string) ([]ApiFile, error) {
	allFiles := make([]ApiFile, 0)

	for _, path := range paths {
		info, err := tools.GetPathInfo(path)

		if err == nil && info.IsDir {
			filePaths := tools.FindDirFiles(path, api.Extensions)
			allFiles = append(allFiles, pathsToApiFiles(filePaths)...)
		} else {
			// If file errored or path is not a directory
			allFiles = append(allFiles, pathToApiFile(path))
		}
	}

	return allFiles, nil
}
