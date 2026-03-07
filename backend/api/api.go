package api

import (
	"context"
	"slices"
	"strings"
	"sync"

	"media-organizer/backend/events"
	"media-organizer/backend/stores"
	"media-organizer/backend/tools"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Api struct {
	ctx        context.Context
	extensions []string
	files      stores.FileStore
	events     events.BackToFront
	loadWg     sync.WaitGroup
}

func CreateApi() *Api {
	return &Api{
		extensions: slices.Concat(tools.DefaultExtensions.Images(), tools.DefaultExtensions.Video),
		files:      *stores.CreateFileStore(),
	}
}

func (api *Api) GetFiles() []*stores.FileEntry {
	return api.files.GetFiles()
}

func (api *Api) GetExtensions() []string {
	// TODO: This should be be extended to take in a list of extensions to use
	return api.extensions
}

func (api *Api) Init(ctx context.Context) {
	api.ctx = ctx
	api.events = *events.CreateBackToFrontEvents(ctx)

	// TODO: This has to come from the input when the application is called (like "mora /some/path")
	api.addFolder("/home/toke/Desktop/norway")
}

func (api *Api) addFolder(dirPath string) {
	api.loadWg.Add(1)

	go func(p string) {
		defer api.loadWg.Done()

		filePaths, errors := tools.FindDirFilePaths(dirPath, api.extensions)
		api.files.AddFiles(filePaths)
		api.files.AddErrorPaths(errors)
	}(dirPath)
}

func (api *Api) addPath(path string) {
	info, err := tools.GetPathInfo(path)

	if err != nil {
		api.loadWg.Add(1)

		go func(p string, err error) {
			defer api.loadWg.Done()
			api.files.AddErrorPath(tools.PathError{Path: p, Err: err})
		}(path, err)
	} else if info.IsDir {
		api.addFolder(path)
	} else {
		api.loadWg.Add(1)

		go func(p string) {
			defer api.loadWg.Done()
			api.files.AddFile(path)
		}(path)
	}
}

func (api *Api) addPaths(paths []string) {
	for _, path := range paths {
		api.addPath(path)
	}
}

func (api *Api) endLoading() {
	api.loadWg.Wait()
	api.events.FileLoadingEnd(api.files.GetFiles())
}

func (api *Api) RequestAllFiles() []*stores.FileEntry {
	api.loadWg.Wait()
	return api.files.GetFiles()
}

func (api *Api) SelectFiles() {
	exts := "*." + strings.Join(api.extensions, ";*.")

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
		api.events.DialogSelectFailed(err)
		return
	}

	api.events.FileLoadingBegin()
	api.addPaths(filePaths)
	api.endLoading()
}

func (api *Api) SelectDirectory() {
	dirPath, err := runtime.OpenDirectoryDialog(api.ctx, runtime.OpenDialogOptions{
		Title: "Select Folder to add",
	})

	if err != nil {
		api.events.DialogSelectFailed(err)
		return
	}

	api.events.FileLoadingBegin()
	api.addFolder(dirPath)
	api.endLoading()
}

func (api *Api) SelectFromDrop(x int, y int, paths []string) {
	api.events.FileLoadingBegin()
	api.addPaths(paths)
	api.endLoading()
}
