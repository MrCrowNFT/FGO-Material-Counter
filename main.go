package main

import (
	"os"
	"strconv"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main(){
	// Set a valid locale
	os.Setenv("LANG", "en_US.UTF-8") 

	// Create new app and window
	counter := app.New()
	counter.Settings().SetTheme(theme.DarkTheme()) // Apply dark theme
	window := counter.NewWindow("FGO Material Counter")

	// Container to hold all materials rows
	materialRows := container.NewVBox()

	// Input the material name
	materialInput := widget.NewEntry()
	materialInput.SetPlaceHolder("Enter material name")

	// Add buttom
	addMaterialButtom := widget.NewButton("Add Material", func() {
		name := materialInput.Text

		// Add the material to the list
		if name != ""{
			addMaterial(name, materialRows)
			// Reset the input
			materialInput.SetText("")
		}
	})

	// Layout
	mainLayout := container.NewVBox(
		container.NewHBox(materialInput, addMaterialButtom),
		widget.NewLabel("Materials:"),
		materialRows,
	)


	window.SetContent(mainLayout)
	window.Resize(fyne.NewSize(400, 400))
	window.ShowAndRun()
}

// Helper function to add a material row dynamically
func addMaterial(material string, rows *fyne.Container){
	count := 0

	//create widget for the row
	materalLabel := widget.NewLabel(material + ": " + strconv.Itoa(count))
	addButtom := widget.NewButton("Add", func() {
		// Increment the counter
		count ++
		// Update the label
		materalLabel.SetText(material + ": " + strconv.Itoa(count))
	})

	// Add the row to the list
	rows.Add(container.NewHBox(
		materalLabel,
		addButtom,
	))
}