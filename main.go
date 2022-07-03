package main

import (
	"os"

	"github.com/DoctorAssist/doctor-assist-go/app"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	mainApp, err := app.New("./file.db")
	if err != nil {
		log.Error().AnErr("Could not open DB", err)
		os.Exit(1)
	}

	localizationDialog, localizationChannel, err := mainApp.NewLocalizationDialog()
	if err != nil {
		log.Error().AnErr("Could not launch language dialog", err)
		os.Exit(1)
	}

	go func() {
		localization := <-localizationChannel
		mainApp.SetupUI(localization)
		localizationDialog.Hide()
	}()

	localizationDialog.Show()

	mainApp.Show()
}
