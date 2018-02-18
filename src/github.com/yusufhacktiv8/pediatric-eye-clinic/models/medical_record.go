package models

import (
	"database/sql"
	"fmt"
	"strconv"
)

// MedicalRecord is a model for patient
type MedicalRecord struct {
	ID                  int     `json:"id"`
	Code                string  `json:"code"`
	Patient             Patient `json:"patient"`
	CornealDiameter     string  `json:"cornealDiameter"`
	IntraocularPressure float32 `json:"intraocularPressure"`
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
func FindMedicalRecords(db *sql.DB, start, count int, searchText string) ([]MedicalRecord, error) {
	rows, err := db.Query(
		`SELECT
			m.id,
			m.code,
			m.corneal_diameter,
			m.intraocular_pressure,
			m.axial_length,
			m.refraksi,
			m.axis,
			m.iol_type,
			m.iol_power,
			m.keratometri,
			m.pre_op_visual_acuity,
			m.post_op_visual_acuity,
			p.id as patient_id,
			p.code as patient_code,
			p.name as patient_name
		FROM medical_records m
		LEFT JOIN patients p ON m.patient = p.id
		WHERE m.code LIKE $3 ORDER BY m.code
		LIMIT $1 OFFSET $2`,
		count, start, "%"+searchText+"%")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	medicalRecords := []MedicalRecord{}
	var patientID sql.NullString
	var patientCode sql.NullString
	var patientName sql.NullString

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
			&d.PostOpVisualAcuity,
			&patientID,
			&patientCode,
			&patientName); err != nil {
			return nil, err
		}
		if patientCode.Valid {
			d.Patient.ID, _ = strconv.Atoi(patientID.String)
			d.Patient.Code = patientCode.String
			d.Patient.Name = patientName.String
		}
		medicalRecords = append(medicalRecords, d)
	}

	return medicalRecords, nil
}

func CountMedicalRecords(db *sql.DB, searchText string) (int, error) {
	rows, err := db.Query(
		`SELECT
			count(1) AS rowsCount
		FROM
			medical_records r
		WHERE r.code LIKE $1`,
		"%"+searchText+"%")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	defer rows.Close()

	rowsCount := 0
	for rows.Next() {
		if err := rows.Scan(&rowsCount); err != nil {
			return 0, err
		}
	}

	return rowsCount, nil
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
			post_op_visual_acuity,
			patient) VALUES
			(	$1,
				$2,
			 	$3,
				$4,
				$5,
				$6,
				$7,
				$8,
				$9,
				$10,
				$11,
				$12)
			ON CONFLICT (code) DO UPDATE
			SET
				corneal_diameter=$2,
				intraocular_pressure=$3,
				axial_length=$4,
				refraksi=$5,
				axis=$6,
				iol_type=$7,
				iol_power=$8,
				keratometri=$9,
				pre_op_visual_acuity=$10,
				post_op_visual_acuity=$11,
				patient=$12
			RETURNING id`,
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
		d.Patient.ID).Scan(&d.ID)

	if err != nil {
		return err
	}

	return nil
}
