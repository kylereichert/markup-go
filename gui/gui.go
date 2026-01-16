package gui

import (
	"log"
	// "time"
	// "fyne.io/fyne/container"
	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Size struct {
	Width  float32
	Height float32
}

var mainSize Size = Size{
	Width:  1000,
	Height: 800,
}

func Run(app fyne.App) {
	w := app.NewWindow("Hello World")
	w.SetContent(widget.NewLabel("Hello World"))
	w.Resize(fyne.NewSize(mainSize.Width, mainSize.Height))

	usjInput := widget.NewEntry()
	usjInput.SetPlaceHolder("USJ: ")

	content := container.NewVBox(usjInput, widget.NewButton("Save", func() {
		log.Println("Content was:", usjInput.Text)
	}))

	w.SetContent(content)
	w.ShowAndRun()
	// app.Run()
}
