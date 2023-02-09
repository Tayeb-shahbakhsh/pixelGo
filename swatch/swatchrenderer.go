package swatch

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

type SwatchRenderer struct {
	Square  canvas.Rectangle
	objects []fyne.CanvasObject
	Parent  *Swatch
}

func (renderer *SwatchRenderer) MinSize() fyne.Size {
	return renderer.Square.MinSize()
}

func (renderer *SwatchRenderer) Layout(size fyne.Size) {
	renderer.objects[0].Resize(size)
}

func (renderer *SwatchRenderer) Refresh() {
	defer canvas.Refresh(renderer.Parent)
	renderer.Layout(fyne.NewSize(20, 20))
	renderer.Square.FillColor = renderer.Parent.Color
	if renderer.Parent.Selected {
		renderer.Square.StrokeWidth = 3
		renderer.Square.StrokeColor = color.NRGBA{255, 255, 255, 255}
		renderer.objects[0] = &renderer.Square
	} else {
		renderer.Square.StrokeWidth = 0
		renderer.objects[0] = &renderer.Square
	}
}

func (renderer *SwatchRenderer) Objects() []fyne.CanvasObject {
	return renderer.Objects()
}

func (renderer *SwatchRenderer) Destroy() {}
