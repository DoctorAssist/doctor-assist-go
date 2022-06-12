package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/DoctorAssist/doctor-assist-go/localization"
)

func main() {
	mainApp := app.New()
	mainWindow := mainApp.NewWindow("Doctor Assist")
	mainWindow.Resize(fyne.NewSize(450, 450))
	text := widget.NewLabel("Doctor Assist")
	var homeButton *widget.Button
	var patientsButton *widget.Button
	var languageDialog dialog.Dialog

	languageSelect := widget.NewSelect([]string{"Spanish", "English"}, func(option string) {
		var language []string
		switch option {
		case "Spanish":
			language = localization.Spanish
		case "English":
			language = localization.English
		}

		homeButton = widget.NewButton(language[localization.HomeButtonText], func() {})
		patientsButton = widget.NewButton(language[localization.PatientsButtonText], func() {})

		menuBar := container.NewVBox(homeButton, patientsButton)
		mainWindow.SetContent(container.NewHBox(menuBar, text))
		languageDialog.Hide()
	})

	languageDialog = dialog.NewCustom("Language", "", languageSelect, mainWindow)
	languageDialog.Show()
	mainWindow.ShowAndRun()
}
