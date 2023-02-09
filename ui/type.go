package ui

import (
	"fyne.io/fyne/v2"
	"tayebshahbakhsh/pixel/apptype"
	"tayebshahbakhsh/pixel/swatch"
)

type AppInit struct {
	PixelWindow fyne.Window
	State       apptype.State
	Swatches    []*swatch.Swatch
}
