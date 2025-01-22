package api

import (
	"context"
	"fmt"
	"media-organizer/backend/stores"
	"media-organizer/backend/tools"
	"slices"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type ApiReturnValue struct {
	Files  []string
	Errors map[string]string
}

type Temp struct {
	Value string
}

type Api struct {
	ctx   context.Context
	files stores.FileStore
}

func CreateApi() *Api {
	extensions := slices.Concat(tools.DefaultExtensions.Images(), tools.DefaultExtensions.Video)

	return &Api{
		files: *stores.CreateFileStore(extensions),
	}
}

func (api *Api) Init(ctx context.Context) {
	api.ctx = ctx
}

func (api *Api) AddFiles() (*ApiReturnValue, error) {
	exts := "*." + strings.Join(api.files.Extensions, ";*.")

	paths, err := runtime.OpenMultipleFilesDialog(api.ctx, runtime.OpenDialogOptions{
		Title: "Select files to add",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Images or Videos",
				Pattern:     exts,
			},
		},
	})

	if err != nil {
		return nil, err
	}

	api.files.AddFiles(paths)

	fmt.Println(api.GetData())

	return api.GetData(), nil
}

func (api *Api) AddDirectory() (*ApiReturnValue, error) {
	path, err := runtime.OpenDirectoryDialog(api.ctx, runtime.OpenDialogOptions{
		Title: "Select folder to add",
	})

	if err != nil {
		return nil, err
	}

	api.files.AddDirectory(path)

	return api.GetData(), nil
}

func (api *Api) GetData() *ApiReturnValue {
	return &ApiReturnValue{
		Files:  api.files.Paths(),
		Errors: api.files.Errors(),
	}
}

func (api *Api) OnFileDrop(x int, y int, paths []string) {
	for _, path := range paths {
		isDir, err := tools.IsDir(path)

		if err != nil {
			fmt.Println("Failed to stat:", path)
		} else if !isDir {
			// Path is a file
			fmt.Println(path)
		} else {
			// Path is a dir
			for _, filePath := range tools.FindDirFiles(path, api.files.Extensions) {
				fmt.Println(filePath)
			}
		}
	}
}
