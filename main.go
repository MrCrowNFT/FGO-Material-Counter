package main

import(
	"fyne.io/fyne/v2/app"
    _"fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

func main(){
	// Create new app and window
	counter := app.New()
	window := counter.NewWindow("FGO Material Counter")

	// Mapping to track the material as they get added
	material := make(map[string]int)

	// Will use this widget to display the materials and the count
	materialList := widget.NewLabel("")

	// Input the material name
	materialInput := widget.NewEntry()
	materialInput.SetPlaceHolder("Enter material name")


	window.SetContent(widget.NewLabel("FGO Counter"))
	window.ShowAndRun()
}