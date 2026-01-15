package gui

import (
	// "fyne.io/fyne/container"
	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	// "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Run(app fyne.App) {
	w := app.NewWindow("Hello")

	w.SetContent(widget.NewLabel("Hello, Again!"))
	w.ShowAndRun()
}
