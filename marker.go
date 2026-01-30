package main

import (
	"image"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

const (
	Width  = 800
	Height = 600
)

func main() {
	a := app.NewWithID("com.example.imagehighlighter")
	w := a.NewWindow("Image Highlighter")
	w.Resize(fyne.NewSize(Width, Height))
	ov := NewOverlay()

	img := &canvas.Image{}
	var big image.Image
	var imgPath string

	if len(os.Args) == 2 {
		imgPath = os.Args[1]
		f, err := os.Open(imgPath)
		if err != nil {
			panic(err)
		}
		big, img = ov.loadImage(f)
	}

	stack := container.NewStack(img, ov)
	w.Canvas().SetOnTypedKey(func(ev *fyne.KeyEvent) {
		key := string(ev.Name)
		switch key {
		case "Q":
			os.Exit(0)
		}
	})

	openBtn := widget.NewButton("Open Image", func() {
		dialog.NewFileOpen(func(r fyne.URIReadCloser, err error) {
			if err != nil || r == nil {
				return
			}
			defer r.Close()
			big, img = ov.loadImage(r)
			stack.Objects[0] = img
			stack.Refresh()
		}, w).Show()
	})

	saveBtn := widget.NewButton("Save", func() {
		ov.SaveBig(big, imgPath)
	})

	quitBtn := widget.NewButton("Quit", func() {
		os.Exit(0)
	})

	buttons := container.NewHBox(openBtn, saveBtn, quitBtn)

	w.SetContent(container.NewVBox(buttons, stack))
	w.ShowAndRun()
}
