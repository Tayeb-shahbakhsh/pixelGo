package swatch

import "fyne.io/fyne/v2/driver/desktop"

func (swatch *Swatch) MouseDown(ev *desktop.MouseEvent) {
	defer swatch.Refresh()
	swatch.clickHandler(swatch)
	swatch.Selected = true
}
func (swatch *Swatch) MouseUp(ev *desktop.MouseEvent) {}
