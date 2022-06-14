package sqlwrapper

import (
	"database/sql"

	"github.com/DoctorAssist/doctor-assist-go/models/patients"
	_ "github.com/mattn/go-sqlite3"
)

type SQLWrapper struct {
	*sql.DB
}

func New(driverName, dataSourceName string) (wrapper *SQLWrapper, err error) {
	wrapper = new(SQLWrapper)
	wrapper.DB, err = sql.Open(driverName, dataSourceName)
	return wrapper, err
}

func (*SQLWrapper) InsertPatientData(patients.Patient) error
func (*SQLWrapper) FetchPatient(string) (patients.Patient, error)
func (*SQLWrapper) FetchAllPatients(string) ([]patients.Patient, error)
func (*SQLWrapper) UpdatePatient(patients.Patient) (error)
func (*SQLWrapper) RemovePatient(string) error
