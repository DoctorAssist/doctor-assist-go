package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/DoctorAssist/doctor-assist-go/core/sqlwrapper"
	"github.com/DoctorAssist/doctor-assist-go/localization"
)

type App struct {
	fyne.App
	db         DBWrapper
	mainWindow fyne.Window
}

func New(dbFilePath string) (newApp *App, err error) {
	newApp = new(App)
	newApp.db, err = sqlwrapper.New("sqlite3", dbFilePath)
	if err != nil {
		return nil, err
	}

	newApp.App = app.New()

	mainWindow := newApp.NewWindow("Doctor Assist")
	mainWindow.Resize(fyne.NewSize(450, 450))

	return newApp, nil
}

func (a *App) SetupUI(uiText []string) {
	var (
		homeButton     *widget.Button
		patientsButton *widget.Button
		languageDialog dialog.Dialog
	)

	text := widget.NewLabel("Doctor Assist")

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
		a.mainWindow.SetContent(container.NewHBox(menuBar, text))
		languageDialog.Hide()
	})

	languageDialog = dialog.NewCustom("Language", "", languageSelect, a.mainWindow)
}

func (app App) Show() {
	app.mainWindow.ShowAndRun()
}
