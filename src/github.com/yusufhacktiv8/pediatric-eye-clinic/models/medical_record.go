package models

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

// MedicalRecord is a model for patient
type MedicalRecord struct {
	ID                  int     	`json:"id"`
	Code                string  	`json:"code"`
	Patient             Patient 	`json:"patient"`
	RecordDate      		time.Time `json:"recordDate"`
	CornealDiameter     string  	`json:"cornealDiameter"`
	IntraocularPressure float32 	`json:"intraocularPressure"`
	AxialLength         float32 	`json:"axialLength"`
	Refraksi            string  	`json:"refraksi"`
	Axis                float32 	`json:"axis"`
	IOLType             string  	`json:"iOLType"`
	IOLPower            float32 	`json:"iOLPower"`
	Keratometri         string  	`json:"keratometri"`
	PreOpVisualAcuity   string  	`json:"preOpVisualAcuity"`
	PostOpVisualAcuity  string  	`json:"postOpVisualAcuity"`
	CornealDiameter2     string  	`json:"cornealDiameter2"`
	IntraocularPressure2 float32 	`json:"intraocularPressure2"`
	AxialLength2         float32 	`json:"axialLength2"`
	Refraksi2            string  	`json:"refraksi2"`
	Axis2                float32 	`json:"axis2"`
	IOLType2             string  	`json:"iOLType2"`
	IOLPower2            float32 	`json:"iOLPower2"`
	Keratometri2         string  	`json:"keratometri2"`
	PreOpVisualAcuity2   string  	`json:"preOpVisualAcuity2"`
	PostOpVisualAcuity2  string  	`json:"postOpVisualAcuity2"`
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
			p.name as patient_name,
			m.record_date,
			m.corneal_diameter2,
			m.intraocular_pressure2,
			m.axial_length2,
			m.refraksi2,
			m.axis2,
			m.iol_type2,
			m.iol_power2,
			m.keratometri2,
			m.pre_op_visual_acuity2,
			m.post_op_visual_acuity2
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
			&patientName,
			&d.RecordDate,
			&d.CornealDiameter2,
			&d.IntraocularPressure2,
			&d.AxialLength2,
			&d.Refraksi2,
			&d.Axis2,
			&d.IOLType2,
			&d.IOLPower2,
			&d.Keratometri2,
			&d.PreOpVisualAcuity2,
			&d.PostOpVisualAcuity2); err != nil {
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
			patient,
			record_date,
			corneal_diameter2,
			intraocular_pressure2,
			axial_length2,
			refraksi2,
			axis2,
			iol_type2,
			iol_power2,
			keratometri2,
			pre_op_visual_acuity2,
			post_op_visual_acuity2) VALUES
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
				$12,
				$13,
				$14,
				$15,
				$16,
				$17,
				$18,
				$19,
				$20,
				$21,
				$22,
				$23)
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
				patient=$12,
				record_date=$13,
				corneal_diameter2=$14,
				intraocular_pressure2=$15,
				axial_length2=$16,
				refraksi2=$17,
				axis2=$18,
				iol_type2=$19,
				iol_power2=$20,
				keratometri2=$21,
				pre_op_visual_acuity2=$22,
				post_op_visual_acuity2=$23
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
		d.Patient.ID,
		d.RecordDate,
		d.CornealDiameter2,
		d.IntraocularPressure2,
		d.AxialLength2,
		d.Refraksi2,
		d.Axis2,
		d.IOLType2,
		d.IOLPower2,
		d.Keratometri2,
		d.PreOpVisualAcuity2,
		d.PostOpVisualAcuity2).Scan(&d.ID)

	if err != nil {
		return err
	}

	return nil
}
