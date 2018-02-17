package models

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

// Patient is a model for patient
type Patient struct {
	ID               int        `json:"id"`
	Code             string     `json:"code"`
	Name             string     `json:"name"`
	DateOfBirth      time.Time  `json:"dateOfBirth"`
	Address          string     `json:"address"`
	FatherName       string     `json:"fatherName"`
	MotherName       string     `json:"motherName"`
	FatherOccupation Occupation `json:"fatherOccupation"`
	MotherOccupation Occupation `json:"motherOccupation"`
	ReferralOrigin   string     `json:"referralOrigin"`
	Insurance        Insurance  `json:"insurance"`
}

// FindPatients to find patients
func FindPatients(db *sql.DB, start, count int, searchText string) ([]Patient, error) {
	rows, err := db.Query(
		`SELECT
			p.id,
			p.code,
			p.name,
			p.date_of_birth,
			p.address,
			p.father_name,
			p.mother_name,
			o1.id as father_occupation_id,
			o1.code as father_occupation_code,
			o1.name as father_occupation_name,
			o2.code as mother_occupation_code,
			o2.name as mother_occupation_name,
			p.referral_origin,
			i.code as insurance_code,
			i.name as insurance_name
		FROM
			patients p
		LEFT JOIN occupations o1 ON p.father_occupation = o1.id
		LEFT JOIN occupations o2 ON p.mother_occupation = o2.id
		LEFT JOIN insurances i ON p.insurance = i.id
		WHERE p.code LIKE $3 OR p.name LIKE $3  ORDER BY p.code
		LIMIT $1 OFFSET $2`,
		count, start, "%"+searchText+"%")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	patients := []Patient{}
	var fatherOccupationId sql.NullString
	var fatherOccupationCode sql.NullString
	var fatherOccupationName sql.NullString
	var motherOccupationCode sql.NullString
	var motherOccupationName sql.NullString
	var insuranceCode sql.NullString
	var insuranceName sql.NullString

	for rows.Next() {
		var d Patient
		if err := rows.Scan(
			&d.ID,
			&d.Code,
			&d.Name,
			&d.DateOfBirth,
			&d.Address,
			&d.FatherName,
			&d.MotherName,
			&fatherOccupationId,
			&fatherOccupationCode,
			&fatherOccupationName,
			&motherOccupationCode,
			&motherOccupationName,
			&d.ReferralOrigin,
			&insuranceCode,
			&insuranceName); err != nil {
			return nil, err
		}

		if fatherOccupationCode.Valid {
			d.FatherOccupation.ID, _ = strconv.Atoi(fatherOccupationId.String)
			d.FatherOccupation.Code = fatherOccupationCode.String
			d.FatherOccupation.Name = fatherOccupationName.String
		}

		if motherOccupationCode.Valid {
			d.MotherOccupation.Code = motherOccupationCode.String
			d.MotherOccupation.Name = motherOccupationName.String
		}

		if insuranceCode.Valid {
			d.Insurance.Code = insuranceCode.String
			d.Insurance.Name = insuranceName.String
		}
		patients = append(patients, d)
	}

	return patients, nil
}

func CountPatients(db *sql.DB, searchText string) (int, error) {
	rows, err := db.Query(
		`SELECT
			count(1) AS rowsCount
		FROM
			patients p
		WHERE p.code LIKE $1 or p.name LIKE $1`,
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
func (d *Patient) FindOne(db *sql.DB) error {
	var fatherOccupationCode sql.NullString
	var fatherOccupationName sql.NullString
	var motherOccupationCode sql.NullString
	var motherOccupationName sql.NullString
	var insuranceCode sql.NullString
	var insuranceName sql.NullString
	err := db.QueryRow(
		`SELECT
			p.id,
			p.code,
			p.name,
			p.date_of_birth,
			p.address,
			p.father_name,
			p.mother_name,
			o1.code as father_occupation_code,
			o1.name as father_occupation_name,
			o2.code as mother_occupation_code,
			o2.name as mother_occupation_name,
			p.referral_origin,
			i.code as insurance_code,
			i.name as insurance_name
		FROM
			patients p
		LEFT JOIN occupations o1 ON p.father_occupation = o1.id
		LEFT JOIN occupations o2 ON p.mother_occupation = o2.id
		LEFT JOIN insurances i ON p.insurance = i.id
		WHERE p.code=$1`,
		d.Code).Scan(&d.ID,
		&d.Code,
		&d.Name,
		&d.DateOfBirth,
		&d.Address,
		&d.FatherName,
		&d.MotherName,
		&fatherOccupationCode,
		&fatherOccupationName,
		&motherOccupationCode,
		&motherOccupationName,
		&d.ReferralOrigin,
		&insuranceCode,
		&insuranceName)

	if fatherOccupationCode.Valid {
		d.FatherOccupation.Code = fatherOccupationCode.String
		d.FatherOccupation.Name = fatherOccupationName.String
	}

	if motherOccupationCode.Valid {
		d.MotherOccupation.Code = motherOccupationCode.String
		d.MotherOccupation.Name = motherOccupationName.String
	}

	if insuranceCode.Valid {
		d.Insurance.Code = insuranceCode.String
		d.Insurance.Name = insuranceName.String
	}

	return err
}

// Update patient
func (d *Patient) Update(db *sql.DB) error {
	_, err :=
		db.Exec(
			`UPDATE patients SET
				code=$1,
				name=$2,
				date_of_birth=$3,
				address=$4,
				father_name=$5,
				mother_name=$6,
				father_occupation=$7,
				mother_occupation=$8,
				referral_origin=$9,
				insurance=$10
			WHERE
				code=$11`,
			d.Code,
			d.Name,
			d.DateOfBirth,
			d.Address,
			d.FatherName,
			d.MotherName,
			d.FatherOccupation.ID,
			d.MotherOccupation.ID,
			d.ReferralOrigin,
			d.Insurance.ID,
			d.Code)
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
		`INSERT INTO patients
		(code,
			name,
			date_of_birth,
			address,
			father_name,
			mother_name,
			father_occupation,
			mother_occupation,
			referral_origin,
			insurance)
		VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		ON CONFLICT (code) DO UPDATE
		SET
			name=$2,
			date_of_birth=$3,
			address=$4,
			father_name=$5,
			mother_name=$6,
			father_occupation=$7,
			mother_occupation=$8,
			referral_origin=$9,
			insurance=$10
		RETURNING id`,
		d.Code,
		d.Name,
		d.DateOfBirth,
		d.Address,
		d.FatherName,
		d.MotherName,
		d.FatherOccupation.ID,
		d.MotherOccupation.ID,
		d.ReferralOrigin,
		d.Insurance.ID).Scan(&d.ID)

	if err != nil {
		return err
	}

	return nil
}
