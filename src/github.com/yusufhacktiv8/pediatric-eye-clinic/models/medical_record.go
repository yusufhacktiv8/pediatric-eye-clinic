package models

import (
	"database/sql"
)

// MedicalRecord is a model for patient
type MedicalRecord struct {
	ID                  int     `json:"id"`
	Code                string  `json:"code"`
	CornealDiameter     string  `json:"cornealDiameter"`
	IntraocularPressure string  `json:"intraocularPressure"`
	AxialLength         float32 `json:"axialLength"`
	Refraksi            string  `json:"refraksi"`
	Axis                float32 `json:"axis"`
	IOLType             string  `json:"iOLType"`
	IOLPower            float32 `json:"iOLPower"`
	Keratometri         string  `json:"keratometri"`
	PreOpVisualAcuity   string  `json:"preOpVisualAcuity"`
	PostOpVisualAcuity  string  `json:"postOpVisualAcuity"`
}

// FindMedicalRecords to find medicalRecords
func FindMedicalRecords(db *sql.DB, start, count int) ([]MedicalRecord, error) {
	rows, err := db.Query(
		`SELECT
			id,
			code,
			corneal_diameter,
			intraocular_pressure,
			axial_length,
			refraksi,
			axis,
			iol_type,
			iol_power,
			keratometri,
			pre_op_visual_acuity,
			post_op_visual_acuity
		FROM medical_records LIMIT $1 OFFSET $2`,
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	medicalRecords := []MedicalRecord{}

	for rows.Next() {
		var d MedicalRecord
		if err := rows.Scan(
			&d.ID,
			&d.Code,
			&d.CornealDiameter,
			&d.IntraocularPressure,
			&d.AxialLength,
			&d.Refraksi,
			&d.Axis,
			&d.IOLType,
			&d.IOLPower,
			&d.Keratometri,
			&d.PreOpVisualAcuity,
			&d.PostOpVisualAcuity); err != nil {
			return nil, err
		}
		medicalRecords = append(medicalRecords, d)
	}

	return medicalRecords, nil
}

// FindOne to find one patient based on code
func (d *MedicalRecord) FindOne(db *sql.DB) error {
	return db.QueryRow(
		`SELECT
			id,
			code,
			corneal_diameter,
			intraocular_pressure,
			axial_length,
			refraksi,
			axis,
			iol_type,
			iol_power,
			keratometri,
			pre_op_visual_acuity,
			post_op_visual_acuity
		FROM medical_records WHERE code=$1`,
		d.Code).Scan(
		&d.ID,
		&d.Code,
		&d.CornealDiameter,
		&d.IntraocularPressure,
		&d.AxialLength,
		&d.Refraksi,
		&d.Axis,
		&d.IOLType,
		&d.IOLPower,
		&d.Keratometri,
		&d.PreOpVisualAcuity,
		&d.PostOpVisualAcuity)
}

// Update patient
func (d *MedicalRecord) Update(db *sql.DB) error {
	_, err :=
		db.Exec(
			`UPDATE medical_records SET
				code=$1,
				corneal_diameter=$2,
				intraocular_pressure=$3,
				axial_length=$4,
				refraksi=$5,
				axis=$6,
				iol_type=$7,
				iol_power=$8,
				keratometri=$9,
				pre_op_visual_acuity=$10,
				post_op_visual_acuity=$11
			WHERE
				code=$12`,
			d.Code,
			d.CornealDiameter,
			d.IntraocularPressure,
			d.AxialLength,
			d.Refraksi,
			d.Axis,
			d.IOLType,
			d.IOLPower,
			d.Keratometri,
			d.PreOpVisualAcuity,
			d.PostOpVisualAcuity,
			d.Code)

	return err
}

// Delete patient
func (d *MedicalRecord) Delete(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM medical_records WHERE code=$1", d.Code)

	return err
}

// Create create patient
func (d *MedicalRecord) Create(db *sql.DB) error {
	err := db.QueryRow(
		`INSERT INTO medical_records
			(code,
			corneal_diameter,
			intraocular_pressure,
			axial_length,
			refraksi,
			axis,
			iol_type,
			iol_power,
			keratometri,
			pre_op_visual_acuity,
			post_op_visual_acuity) VALUES($1, $2) RETURNING id`,
		d.Code,
		d.CornealDiameter,
		d.IntraocularPressure,
		d.AxialLength,
		d.Refraksi,
		d.Axis,
		d.IOLType,
		d.IOLPower,
		d.Keratometri,
		d.PreOpVisualAcuity,
		d.PostOpVisualAcuity).Scan(&d.ID)

	if err != nil {
		return err
	}

	return nil
}
