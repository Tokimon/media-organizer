package main

import (
	"context"
	"embed"
	"fmt"
	"media-organizer/backend/api"
	"media-organizer/backend/image_creator"
	"media-organizer/backend/tools"
	"net/http"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
)

//go:embed all:frontend/dist
var assets embed.FS

func Image400(res http.ResponseWriter, err error, path string) {
	res.WriteHeader(http.StatusBadRequest)
	res.Write(fmt.Appendf([]byte{}, "Incorrect image %s", path))
	tools.Logger.Error(err.Error())
}

func Image404(res http.ResponseWriter, err error, path string) {
	res.WriteHeader(http.StatusNotFound)
	res.Write(fmt.Appendf([]byte{}, "Could not load image %s", path))
	tools.Logger.Error(err.Error())
}

func fsFileMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		thumb := req.URL.Query()["thumb"]

		if len(thumb) > 0 {
			path := thumb[0]
			// tools.Logger.Debug(fmt.Sprintf("Serving thumbnail for file : %s", req.URL.String()))
			jpgSrc, err := image_creator.GetJpegSource(path)

			if err != nil {
				Image404(res, err, path)
				next.ServeHTTP(res, req)
				return
			}

			err = image_creator.GenerateThumbnail(res, jpgSrc)

			if err != nil {
				Image400(res, err, path)
				next.ServeHTTP(res, req)
				return
			}

			fmt.Println("ALL OK")
			res.Header().Set("Cache-Control", "max-age=86400")
			res.Header().Set("Content-Type", "image/jpeg")
			res.WriteHeader(http.StatusOK)
		}

		next.ServeHTTP(res, req)
	})
}

func main() {
	// Create an instance of the app structure
	app := NewApp()
	webApi := api.CreateApi()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "media-organizer",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
			// Handler: NewFileLoader(),
			Middleware: fsFileMiddleware,
		},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 0},
		DragAndDrop: &options.DragAndDrop{
			EnableFileDrop: true,
		},
		OnStartup: func(ctx context.Context) {
			app.startup(ctx)
			webApi.Init(ctx)
		},
		Bind: []any{
			app,
			webApi,
		},
		Linux: &linux.Options{
			WindowIsTranslucent: true,
			ProgramName:         "Media Organizer",
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
