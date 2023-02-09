package main

import (
	"fyne.io/fyne/v2/app"
	"image/color"
	"tayebshahbakhsh/pixel/apptype"
	"tayebshahbakhsh/pixel/swatch"
	"tayebshahbakhsh/pixel/ui"
)

func main() {
	pixelApp := app.New()
	pixelWindow := pixelApp.NewWindow("pixel")

	state := apptype.State{
		BrushColor:     color.NRGBA{R: 255, G: 255, B: 255, A: 255},
		SwatchSelected: 0,
	}

	appInit := ui.AppInit{
		PixelWindow: pixelWindow,
		State:       &state,
		Swatches:    make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)
	appInit.PixelWindow.ShowAndRun()
}
