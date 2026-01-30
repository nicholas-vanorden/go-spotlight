package main

import (
	"image"
	"image/draw"
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/disintegration/imaging"
)

type Overlay struct {
	widget.BaseWidget
	con      *fyne.Container
	marker   *canvas.Rectangle
	rect     *Rect
	inMotion bool
	zoom     float32
}

func NewOverlay() *Overlay {
	over := &Overlay{}
	over.ExtendBaseWidget(over)
	over.con = container.NewWithoutLayout()
	over.rect = NewRect()
	return over
}

func (t *Overlay) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(t.con)
}

func (t *Overlay) Dragged(e *fyne.DragEvent) {
	if t.inMotion == false {
		t.inMotion = true
		t.rect.From = e.Position
		return
	}

	t.rect.To = e.Position
	pos, size := t.rect.Dims()
	t.DrawMarker(pos, size)
}

func (t *Overlay) DragEnd() {
	t.inMotion = false
}

func (t *Overlay) DrawMarker(pos fyne.Position, size fyne.Size) {
	rect := canvas.NewRectangle(t.rect.Color())
	rect.Resize(size)
	rect.Move(pos)

	if t.marker == nil {
		t.marker = rect
		return
	}

	t.con.Remove(t.marker)
	t.marker = rect
	t.con.Add(rect)
	t.con.Refresh()
	return
}

func (t *Overlay) SaveBig(big image.Image, path string) error {
	dimg := imaging.Clone(big)
	r := t.rect.AsImage(t.zoom)
	draw.Draw(dimg, r, &image.Uniform{t.rect.Color()}, r.Min, draw.Over)
	err := imaging.Save(dimg, path)
	if err != nil {
		return err
	}
	return nil
}

func (t *Overlay) loadImage(r io.Reader) (image.Image, *canvas.Image) {
	big, err := imaging.Decode(r, imaging.AutoOrientation(true))
	if err != nil {
		panic(err)
	}
	shrunk := imaging.Resize(big, Width, 0, imaging.Lanczos)
	t.zoom = float32(float64(big.Bounds().Dx()) / float64(Width))
	img := canvas.NewImageFromImage(shrunk)
	img.FillMode = canvas.ImageFillOriginal
	return big, img
}
