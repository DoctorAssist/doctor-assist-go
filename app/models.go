package app

import "github.com/DoctorAssist/doctor-assist-go/models/patients"

type DBWrapper interface {
	InsertPatientData(patients.Patient) error
	FetchPatient(string) (patients.Patient, error)
	FetchAllPatients() ([]patients.Patient, error)
	UpdatePatient(patients.Patient) error
	RemovePatient(string) error
}
