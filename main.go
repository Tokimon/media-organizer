package main

import (
	"context"
	"embed"
	"media-organizer/backend/api"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
)

//go:embed all:frontend/dist
var assets embed.FS

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
		},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 0},
		DragAndDrop: &options.DragAndDrop{
			EnableFileDrop: true,
		},
		OnStartup: func(ctx context.Context) {
			app.startup(ctx)
			webApi.Init(ctx)
		},
		Bind: []interface{}{
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
