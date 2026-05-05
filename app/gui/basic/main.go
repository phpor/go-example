package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello Fyne!")
	var items []string

	// 创建一个list，和两个button，点击button+，list会增加一个item，点击button-，list会减少一个item
	list := widget.NewList(
		func() int {
			return len(items)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(strconv.Itoa(i))
		},
	)
	button1 := widget.NewButton("+", func() {
		items = append(items, strconv.Itoa(len(items)))
	})
	button2 := widget.NewButton("-", func() {
		if len(items) > 0 {
			items = items[1:]
			list.Refresh()
		}
	})

	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
		list,
		button1,
		button2,
	))
	// 让 list的高度为200
	list.Resize(fyne.NewSize(300, 800))
	w.Resize(fyne.NewSize(400, 600))
	w.ShowAndRun()
}
