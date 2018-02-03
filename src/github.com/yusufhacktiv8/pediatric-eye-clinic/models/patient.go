package models

import (
	"database/sql"
)

// Patient is a model for patient
type Patient struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

// FindPatients to find patients
func FindPatients(db *sql.DB, start, count int) ([]Patient, error) {
	rows, err := db.Query(
		"SELECT id, code, name FROM patients LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	patients := []Patient{}

	for rows.Next() {
		var d Patient
		if err := rows.Scan(&d.ID, &d.Code, &d.Name); err != nil {
			return nil, err
		}
		patients = append(patients, d)
	}

	return patients, nil
}

// FindOne to find one patient based on code
func (d *Patient) FindOne(db *sql.DB) error {
	return db.QueryRow("SELECT code, name FROM patients WHERE code=$1",
		d.Code).Scan(&d.Code, &d.Name)
}

// Update patient
func (d *Patient) Update(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE patients SET code=$1, name=$2 WHERE code=$3",
			d.Code, d.Name, d.Code)

	return err
}

// Delete patient
func (d *Patient) Delete(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM patients WHERE code=$1", d.Code)

	return err
}

// Create create patient
func (d *Patient) Create(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO patients(code, name) VALUES($1, $2) RETURNING id",
		d.Code, d.Name).Scan(&d.ID)

	if err != nil {
		return err
	}

	return nil
}
