package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// AppMenu := menu.NewMenu()

	// newAppMenu := *menu.Menu{
	// 	Items: ,
	// }
	// var item *menu.MenuItem
	// item = &menu.MenuItem{
	// 	Label: "退出",
	// }
	// AppMenu.Append(item)

	otherApp := NewOtherApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "机选大乐透",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Menu:             app.GetNewApp(),
		Bind: []interface{}{
			app,
			otherApp,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
