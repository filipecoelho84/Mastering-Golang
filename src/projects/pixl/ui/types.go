package ui

import (
	"fyne.io/fyne/v2"
	"zerotomastery.io/pixl/apptype"
	"zerotomastery.io/pixl/swatch"
)

type AppInit struct {
	PixWindow fyne.Window
	State *apptype.State
	Swatches []*swatch.Swatch
}