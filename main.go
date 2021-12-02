// Package main provides various examples of Fyne API capabilities
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"ra2_auxiliary_tool/config"
	"ra2_auxiliary_tool/tool"
)

var (
	ra2ToolApp    fyne.App
	ra2ToolWindow fyne.Window
)

func init() {}
func main() {

	ra2ToolApp = app.New()
	ra2ToolWindow = ra2ToolApp.NewWindow(config.AppName)
	ra2ToolWindow.SetFixedSize(true)
	ra2ToolWindow.Resize(fyne.Size{
		Width:  400,
		Height: 650,
	})
	tool.InitUI(ra2ToolApp, ra2ToolWindow)

	ra2ToolWindow.ShowAndRun()
}
