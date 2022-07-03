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
	newApp.mainWindow = mainWindow

	return newApp, nil
}

func (a *App) NewLocalizationDialog() (dialog.Dialog, chan []string, error) {
	localizationChannel := make(chan []string)
	languageSelect := widget.NewSelect([]string{"Spanish", "English"},
		func(option string) {
			var language []string
			switch option {
			case "Spanish":
				language = localization.Spanish
			case "English":
				language = localization.English
			}

			localizationChannel <- language
		})

	languageDialog := dialog.NewCustom("Language", "", languageSelect, a.mainWindow)
	return languageDialog, localizationChannel, nil

}

func (a *App) SetupUI(uiText []string) {
	var (
		homeButton     *widget.Button
		patientsButton *widget.Button
	)

	text := widget.NewLabel("Doctor Assist")

	homeButton = widget.NewButton(uiText[localization.HomeButtonText], func() {})
	patientsButton = widget.NewButton(uiText[localization.PatientsButtonText], func() {})

	menuBar := container.NewVBox(homeButton, patientsButton)
	a.mainWindow.SetContent(container.NewHBox(menuBar, text))
}

func (a *App) Show() {
	a.mainWindow.ShowAndRun()
}
