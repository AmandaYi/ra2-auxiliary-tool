package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New();
	win := myApp.NewWindow("Ra2_tool");
	btn1 := widget.NewLabel("btn1");
	var btn2 *widget.Button;
	btn2 = widget.NewButton("btn2", func() {
		//按钮做一些事情
		//this.SetText("I am change!")
		btn2.SetText("I am click");
		btn1.TextStyle.Bold = true;
	})

	btn3 := widget.NewButton("btn3", func() {
		btn2.SetText("haha")
	})
	win.SetContent(container.NewVBox(
		btn1,
		btn2,
		btn3,

	))

	//win.SetContent(widget.NewLabel("list button~"));

	win.Resize(fyne.NewSize(200, 200));
	win.ShowAndRun();
}
