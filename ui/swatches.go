package ui

import (
	"fyne.io/fyne/v2"
	"image/color"
	"tayebshahbakhsh/pixel/swatch"
)

func BuildSwatches(app *AppInit) *fyne.Container {
	canvasSwatches := make([]fyne.CanvasObject, 0, 64)
	for i := 0; i < cap(app.Swatches); i++ {
		initialColor := color.NRGBA{255, 255, 255, 255}
		s := swatch.NewSwatch(app.State, initialColor, i, func(s *swatch.Swatch) {
			for j := 0; j < len(app.Swatches); j++ {
				app.Swatches[j].Selected = false
				canvasSwatches[j].Refresh()
			}
			app.State.SwatchSelected = s.SwatchIndex
			app.State.BrushColor = s.Color
		})
	}
}
