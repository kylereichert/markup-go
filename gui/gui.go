package gui

import (
	// "image/color"
	// "log"

	// "time"
	"fyne.io/fyne/v2"
	// "golang.org/x/net/route"
	// "fyne.io/fyne/v2/app"
	// "fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/layout"
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
	w := app.NewWindow("Grades")
	// w.SetContent(widget.NewLabel("Hello World"))
	w.Resize(fyne.NewSize(mainSize.Width, mainSize.Height))

	type ElevationEntry struct {
		label     *widget.Entry
		elevation *widget.Entry
		container *fyne.Container
	}

	plotLineEntries := []*ElevationEntry{}
	gradeEntries := []*ElevationEntry{}

	plotLineContainer := container.NewVBox()
	gradeContainer := container.NewVBox()

	resultsText := widget.NewLabel("Results \n\nWall Drop: \nWall Height: ")
	resultsText.Wrapping = fyne.TextWrapWord
	resultsScroll := container.NewScroll(resultsText)

	// Creates
	createEntry := func(entries *[]*ElevationEntry, parentContainer *fyne.Container) {
		labelEntry := widget.NewEntry()
		labelEntry.SetPlaceHolder("Label placeholder; remove later")

		elevationEntry := widget.NewEntry()
		elevationEntry.SetPlaceHolder("Elevation")

		deleteButton := widget.NewButton("x", nil)

		rowContainer := container.NewBorder(
			nil, nil, nil, deleteButton,
			container.NewHBox(
				container.NewGridWithColumns(2,
					labelEntry,
					elevationEntry,
				),
			),
		)

		entry := &ElevationEntry{
			label:     labelEntry,
			elevation: elevationEntry,
			container: rowContainer,
		}

		deleteButton.OnTapped = func() {
			for i, e := range *entries {
				if e == entry {
					*entries = append((*entries)[:i], (*entries)[i+1:]...)
					break
				}
			}

			parentContainer.Remove(rowContainer)
			parentContainer.Refresh()
		}

		*entries = append(*entries, entry)
		parentContainer.Add(rowContainer)
		parentContainer.Refresh()

	}

	addPlotLineBtn := widget.NewButton("+ Add Plot Line", func() {
		createEntry(&plotLineEntries, plotLineContainer)
	})

	addGradeBtn := widget.NewButton("+ Add Grade Point", func() {
		createEntry(&gradeEntries, gradeContainer)
	})

	calculateBtn := widget.NewButton("Calculate", func() {
		// This is where you'd plug in your actual calculation logic
		resultsText.SetText("Results:\n\nWall drop: [calculated]\nWall height: [calculated]\n\n(Your logic goes here)")
	})

	plotLineSection := container.NewBorder(
		widget.NewLabel("Plot Line Elevations"),
		addPlotLineBtn,
		nil, nil,
		container.NewScroll(plotLineContainer),
	)

	gradeSection := container.NewBorder(
		widget.NewLabel("Grade Elevations"),
		addGradeBtn,
		nil, nil,
		container.NewScroll(gradeContainer),
	)

	leftSide := container.NewBorder(
		nil,
		calculateBtn,
		nil, nil,
		container.NewVSplit(plotLineSection, gradeSection),
	)

	// Right side layout
	rightSide := container.NewBorder(
		widget.NewLabel("Results"),
		nil, nil, nil,
		resultsScroll,
	)

	// Main split container
	split := container.NewHSplit(leftSide, rightSide)
	split.Offset = 0.5

	// Add a couple default entries to show the layout
	createEntry(&plotLineEntries, plotLineContainer)
	createEntry(&gradeEntries, gradeContainer)

	w.SetContent(split)
	w.ShowAndRun()
	// app.Run()
}
