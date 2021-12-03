package tool

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"ra2_auxiliary_tool/utils"
)

var IsFullScreen bool = false

func InitUI(a fyne.App, w fyne.Window) {

	// 显示全图
	btn1 := widget.NewButton("full map", func() {
		//fmt.Println("显示全图", utils.GetProcess("WXWorkWeb.exe"))

		pid := utils.GetProcess("WXWorkWeb.exe")
		_, _ = utils.OpenProcess(false, uint32(pid))
		fmt.Println(pid)
		//winlogonProcess
	})
	btn2 := widget.NewButton("full keji", func() {
		fmt.Println("全部科技")
	})
	btn3 := widget.NewButton("full map", func() {
		fmt.Println("显示全图")
	})
	btn4 := widget.NewButton("full map", func() {
		fmt.Println("显示全图")
	})
	btn5 := widget.NewButton("full map", func() {
		fmt.Println("显示全图")
	})
	btn6 := widget.NewButton("full map", func() {
		fmt.Println("显示全图")
	})
	btn7 := widget.NewButton("full map", func() {
		fmt.Println("显示全图")
	})
	w.SetContent(
		container.NewVBox(btn1, btn2, btn3, btn4, btn5, btn6, btn7),
	)
}
