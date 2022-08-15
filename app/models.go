package app

import (
	"context"

	"github.com/DoctorAssist/doctor-assist-go/models/patients"
)

type DBWrapper interface {
	InsertPatientData(context.Context, patients.Patient) error
	FetchPatient(context.Context, string) (patients.Patient, error)
	FetchAllPatients(context.Context) ([]patients.Patient, error)
	UpdatePatient(context.Context, patients.Patient) error
	RemovePatient(context.Context, string) error
}
