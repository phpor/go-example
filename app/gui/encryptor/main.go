package main

import (
	"strings"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type FileStatus struct {
	Path   string
	Status string // 状态：等待中/处理中/已完成/错误
}

type GUI struct {
	app        fyne.App
	window     fyne.Window
	list       *widget.List
	fileList   []FileStatus
	mu         sync.Mutex
	processBtn *widget.Button
}

func main() {
	myApp := app.NewWithID("encryptor")
	gui := &GUI{
		app:    myApp,
		window: myApp.NewWindow("CSV文件处理器"),
	}

	gui.createUI()
	gui.window.Resize(fyne.NewSize(600, 400))
	gui.window.ShowAndRun()
}

func (g *GUI) createUI() {
	// 创建操作按钮
	selectBtn := widget.NewButton("选择CSV文件", g.selectFiles)
	g.processBtn = widget.NewButton("开始处理", g.processFiles)
	btnBox := container.NewHBox(selectBtn, g.processBtn)

	// 创建表头
	header := container.NewHBox(
		widget.NewLabel("文件"),
		layout.NewSpacer(),
		widget.NewLabel("状态"),
	)

	// 创建文件列表
	g.list = widget.NewList(
		func() int { return len(g.fileList) },
		func() fyne.CanvasObject {
			return container.NewHBox(
				widget.NewLabel("模板"), // 这一行可以删除，因为我们使用表头
				layout.NewSpacer(),
				widget.NewLabel("状态"), // 这一行可以删除，因为我们使用表头
			)
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			cont := obj.(*fyne.Container)
			label := cont.Objects[0].(*widget.Label)
			status := cont.Objects[1].(*widget.Label) // 修改索引以匹配新的布局

			item := g.fileList[id]
			label.SetText(truncatePath(item.Path, 50))
			status.SetText(item.Status)

			// 根据状态设置颜色
			switch item.Status {
			case "处理中":
				status.TextStyle = fyne.TextStyle{Bold: true}
				status.Importance = widget.HighImportance
			case "已完成":
				status.Importance = widget.SuccessImportance
			case "错误":
				status.Importance = widget.DangerImportance
			default:
				status.Importance = widget.MediumImportance
			}
		},
	)

	// 将表头和列表组合在一起
	listWithHeader := container.NewVBox(
		header,
		g.list,
	)

	// 确保在你的应用中使用 listWithHeader 而不是直接使用 g.list
	content := container.NewBorder(btnBox, nil, nil, nil, listWithHeader)
	g.window.SetContent(content)
}

func (g *GUI) selectFiles() {
	d := dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {
		if err != nil || uc == nil {
			return
		}

		path := uc.URI().Path()
		if !strings.HasSuffix(strings.ToLower(path), ".csv") {
			return
		}

		g.mu.Lock()
		defer g.mu.Unlock()

		// 防止重复添加
		for _, f := range g.fileList {
			if f.Path == path {
				return
			}
		}

		g.fileList = append(g.fileList, FileStatus{
			Path:   path,
			Status: "等待中",
		})
		g.list.Refresh()
	}, g.window)
	d.SetFilter(storage.NewExtensionFileFilter([]string{".csv"}))
	// 允许选择多个文件
	d.Show()
}

func (g *GUI) processFiles() {
	go func() {
		g.mu.Lock()
		println("lock")
		defer func() {
			g.mu.Unlock()
			println("unlock")
		}()
		for i := range g.fileList {
			g.updateStatus(i, "处理中")

			// 模拟处理过程
			time.Sleep(2 * time.Second)

			// 这里添加实际处理逻辑
			// 处理成功时：
			g.updateStatus(i, "已完成")
			// 或处理失败时：
			// g.updateStatus(i, "错误: 原因说明")
		}
	}()
}

func (g *GUI) updateStatus(index int, status string) {
	g.fileList[index].Status = status
	g.list.Refresh()
}

// 缩短长路径显示
func truncatePath(path string, maxLen int) string {
	if len(path) <= maxLen {
		return path
	}
	return "..." + path[len(path)-maxLen+3:]
}
