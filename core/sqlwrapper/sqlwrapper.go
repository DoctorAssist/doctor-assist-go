package sqlwrapper

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/DoctorAssist/doctor-assist-go/models/patients"
)

type SQLWrapper struct {
	*sqlx.DB
}

func New(driverName, dataSourceName string) (wrapper *SQLWrapper, err error) {
	wrapper = new(SQLWrapper)
	wrapper.DB, err = sqlx.Open(driverName, dataSourceName)
	return wrapper, err
}

func (db *SQLWrapper) InsertPatientData(ctx context.Context, p patients.Patient) error {
	_, err := db.NamedExecContext(
		ctx,
		"INSERT INTO patient (name, sex, age, id, phone, email, address) "+
			"VALUES (:name, :sex, :age, :id, :phone, :email, :address)",
		&p,
	)
	if err != nil {
		return err
	}

	return nil
}

func (db *SQLWrapper) FetchPatient(ctx context.Context,
	id string) (patients.Patient, error) {
	p := patients.Patient{}
	err := db.GetContext(ctx, &p, "SELECT * FROM patients WHERE id = $1", id)
	if err != nil {
		return patients.Patient{}, err
	}

	return p, nil
}

func (db *SQLWrapper) FetchAllPatients(ctx context.Context) ([]patients.Patient, error) {
	ps := []patients.Patient{}
	err := db.SelectContext(ctx, &p, "SELECT * FROM patients WHERE id = $1", id)
	if err != nil {
		return patients.Patient{}, err
	}

	return p, nil
}

func (db *SQLWrapper) UpdatePatient(patients.Patient) error {

}

func (db *SQLWrapper) RemovePatient(string) error {

}
