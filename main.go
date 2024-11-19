package main

import (
	"encoding/csv"
	"os"
	"strconv"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/layout"
)

type Material struct{
	Name string
	Count int
}

func main(){
	// Set a valid locale
	os.Setenv("LANG", "en_US.UTF-8") 

	// Create new app and window
	counter := app.New()
	window := counter.NewWindow("FGO Material Counter")

	var materials []Material

	// Container to hold all materials rows
	materialRows := container.NewVBox()

	// Input the material name
	materialInput := widget.NewEntry()
	materialInput.SetPlaceHolder("Enter material name")
	materialInputContainer := container.New(layout.NewGridWrapLayout(fyne.NewSize(300, 40)), materialInput) // Proper resizing


	// Add material function
	addMaterialHandler := func() {
		name := materialInput.Text

		// Add the material to the list
		if name != ""{
			// Append added material to the mapping
			newMaterial := Material{Name: name, Count: 0}
			materials = append(materials, newMaterial)

			// Add row
			addMaterial(newMaterial, materialRows, &materials)

			// Reset the input
			materialInput.SetText("")
	}
	}

	// Trigger adding material o pressing enter
	materialInput.OnSubmitted = func(content string) {
		addMaterialHandler()
	}

	// Add buttom
	addMaterialButtom := widget.NewButton("Add Material", func() {
		addMaterialHandler()
	})

	// Export buttom to save data into a csv file
	exportToCSVButtom := widget.NewButton("Export to csv", func() {
		err := exportToCSV("materials.csv", materials)
		if err != nil {
			fmt.Println("Error exporting to CSV:", err)
		} else {
			fmt.Println("Data successfully exported to materials.csv")
		}
	})

	// Layout
	inputArea := container.NewHBox(materialInputContainer, addMaterialButtom)
	mainLayout := container.NewVBox(
		inputArea,
		widget.NewLabel("Materials:"),
		materialRows,
		exportToCSVButtom,
	)


	window.SetContent(mainLayout)
	window.Resize(fyne.NewSize(400, 400))
	window.ShowAndRun()
}

// Helper function to add a material row dynamically
func addMaterial(material Material, rows *fyne.Container, materials *[]Material){
	count := material.Count

	//create widget for the row
	materalLabel := widget.NewLabel(material.Name + ": " + strconv.Itoa(count))
	
	updateCount := func(increment int) {
		// Increment the counter
		count += increment
		// Update the label
		materalLabel.SetText(material.Name + ": " + strconv.Itoa(count))

		// Update the count in the slice
		for i := range *materials{
			if (*materials)[i].Name == material.Name {
				(*materials)[i].Count = count
				break
			} 
		}
	}

	addButton := widget.NewButton("+1", func() { updateCount(1) })
	addFiveButton := widget.NewButton("+5", func() { updateCount(5) })
	addTenButton := widget.NewButton("+10", func() { updateCount(10) })

	// Add the row to the list
	rows.Add(container.NewHBox(
		materalLabel,
		addButton,
		addFiveButton,
		addTenButton,
	))
}

func exportToCSV(filename string, materials []Material)(err error){
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header row
	err = writer.Write([]string{"Material", "Count"})
	if err != nil{
		return err
	}

	// Write data into csv
	for _, material := range materials {
		err := writer.Write([]string{material.Name, strconv.Itoa(material.Count)})
		if err != nil{
			return err
		}
	}

	return nil
}