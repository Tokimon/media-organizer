package api

import (
	"context"
	"image"
	"media-organizer/backend/tools"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type ApiFile struct {
	Path      string      `json:"path"`
	Name      string      `json:"name"`
	Extension string      `json:"extension"`
	Error     error       `json:"error"`
	ReadOnly  bool        `json:"readOnly"`
	Thumbnail image.Image `json:"-"`
}

type Api struct {
	ctx        context.Context
	Extensions []string
	Files      map[string]ApiFile
}

func CreateApi() *Api {
	return &Api{
		Extensions: slices.Concat(tools.DefaultExtensions.Images(), tools.DefaultExtensions.Video),
		Files:      make(map[string]ApiFile),
	}
}

func initApiFile(path string) ApiFile {
	inf, err := tools.GetPathInfo(path)

	ext := filepath.Ext(path)
	baseName := strings.TrimSuffix(filepath.Base(path), ext)
	exp := regexp.MustCompile(`(?:_\d{1,3}|\s+\((?:\d{1,3}|.*?copy)\))\s*$`)
	name := exp.ReplaceAllString(baseName, "")

	return ApiFile{
		Path:      path,
		Name:      name,
		Extension: ext[1:],
		ReadOnly:  err != nil || inf.ReadOnly,
		Error:     err,
		Thumbnail: nil,
	}
}

func (api *Api) GetApiFile() ApiFile {
	return ApiFile{}
}

func (api *Api) GetExtensions() tools.Extensions {
	return *tools.DefaultExtensions
}

func (api *Api) AddFile(file_path string) {
	_, ok := api.Files[file_path]

	if !ok {
		api.Files[file_path] = initApiFile(file_path)
	}
}

func (api *Api) AddFiles(file_paths []string) {
	for _, path := range file_paths {
		api.AddFile(path)
	}
}

func (api *Api) AddFolder(dir_path string) {
	filePaths := tools.FindDirFilePaths(dir_path, api.Extensions)
	api.AddFiles(filePaths)
}

func (api *Api) Init(ctx context.Context) {
	api.ctx = ctx
	api.AddInitialFiles()
}

func (api *Api) AddInitialFiles() {
	// WARNING: For testing purposes only
	api.AddFolder("/home/toke/Desktop/norway")
}

func (api *Api) GetFiles() map[string]ApiFile {
	return api.Files
}

func (api *Api) SelectFiles() (map[string]ApiFile, error) {
	exts := "*." + strings.Join(api.Extensions, ";*.")

	filePaths, err := runtime.OpenMultipleFilesDialog(api.ctx, runtime.OpenDialogOptions{
		Title: "Select files to add",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Select Files to Add",
				Pattern:     exts,
			},
		},
	})

	if err != nil {
		return api.Files, err
	}

	api.AddFiles(filePaths)

	return api.Files, nil
}

func (api *Api) SelectDirectory() (map[string]ApiFile, error) {
	dirPath, err := runtime.OpenDirectoryDialog(api.ctx, runtime.OpenDialogOptions{
		Title: "Select Folder to add",
	})

	if err != nil {
		return api.Files, err
	}

	api.AddFolder(dirPath)

	return api.Files, nil
}

func (api *Api) SelectFromDrop(x int, y int, paths []string) map[string]ApiFile {
	for _, path := range paths {
		info, err := tools.GetPathInfo(path)

		if err == nil && info.IsDir {
			api.AddFolder(path)
		} else {
			// If file errored or path is not a directory
			api.AddFile(path)
		}
	}

	return api.Files
}
