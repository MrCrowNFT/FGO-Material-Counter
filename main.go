package main

import (
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	_ "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main(){
	// Create new app and window
	counter := app.New()
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

	mainLayout := container.NewVBox(
		container.NewHBox(materialInput, addMaterialButtom),
		widget.NewLabel("Materials:"),
		materialRows,
	)


	window.SetContent(mainLayout)
	window.ShowAndRun()
}

// Helper function to add a material row dynamically
func addMaterial(material string, rows *fyne.Container){

}