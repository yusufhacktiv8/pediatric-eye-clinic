package models

import (
	"database/sql"
	"fmt"
)

// Insurance is a model for insurance
type Insurance struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

// FindInsurances to find insurances
func FindInsurances(db *sql.DB, start, count int, searchText string) ([]Insurance, error) {
	rows, err := db.Query(
		"SELECT id, code, name FROM insurances WHERE code LIKE $3 OR name LIKE $3  ORDER BY code LIMIT $1 OFFSET $2",
		count, start, "%"+searchText+"%")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	insurances := []Insurance{}

	for rows.Next() {
		var d Insurance
		if err := rows.Scan(&d.ID, &d.Code, &d.Name); err != nil {
			return nil, err
		}
		insurances = append(insurances, d)
	}

	return insurances, nil
}

func CountInsurances(db *sql.DB, searchText string) (int, error) {
	rows, err := db.Query(
		`SELECT
			count(1) AS rowsCount
		FROM
			insurances r
		WHERE r.code LIKE $1 or r.name LIKE $1`,
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

func FindAllInsurances(db *sql.DB) ([]Insurance, error) {
	rows, err := db.Query("SELECT id, code, name FROM insurances")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	insurances := []Insurance{}

	for rows.Next() {
		var d Insurance
		if err := rows.Scan(&d.ID, &d.Code, &d.Name); err != nil {
			return nil, err
		}
		insurances = append(insurances, d)
	}

	return insurances, nil
}

// FindOne to find one insurance based on code
func (d *Insurance) FindOne(db *sql.DB) error {
	return db.QueryRow("SELECT code, name FROM insurances WHERE code=$1",
		d.Code).Scan(&d.Code, &d.Name)
}

// Update insurance
func (d *Insurance) Update(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE insurances SET code=$1, name=$2 WHERE code=$3",
			d.Code, d.Name, d.Code)

	return err
}

// Delete insurance
func (d *Insurance) Delete(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM insurances WHERE code=$1", d.Code)

	return err
}

// Create create insurance
func (d *Insurance) Create(db *sql.DB) error {
	err := db.QueryRow(
		`INSERT INTO
			insurances(code, name)
		VALUES($1, $2)
		ON CONFLICT (code) DO UPDATE
		SET name=$2
		RETURNING id`,
		d.Code, d.Name).Scan(&d.ID)

	if err != nil {
		return err
	}

	return nil
}
