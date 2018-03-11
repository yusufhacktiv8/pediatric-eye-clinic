package models

import (
	"database/sql"
	"fmt"
)

// Disease is a model for disease
type Disease struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

// FindDiseases to find diseases
func FindDiseases(db *sql.DB, start, count int, searchText string) ([]Disease, error) {
	rows, err := db.Query(
		"SELECT id, code, name FROM diseases WHERE code LIKE $3 OR name LIKE $3  ORDER BY code LIMIT $1 OFFSET $2",
		count, start, "%"+searchText+"%")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	diseases := []Disease{}

	for rows.Next() {
		var d Disease
		if err := rows.Scan(&d.ID, &d.Code, &d.Name); err != nil {
			return nil, err
		}
		diseases = append(diseases, d)
	}

	return diseases, nil
}

func CountDiseases(db *sql.DB, searchText string) (int, error) {
	rows, err := db.Query(
		`SELECT
			count(1) AS rowsCount
		FROM
			diseases d
		WHERE d.code LIKE $1 or d.name LIKE $1`,
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

// FindOne to find one disease based on code
func (d *Disease) FindOne(db *sql.DB) error {
	return db.QueryRow("SELECT code, name FROM diseases WHERE code=$1",
		d.Code).Scan(&d.Code, &d.Name)
}

// Update disease
func (d *Disease) Update(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE diseases SET code=$1, name=$2 WHERE code=$3",
			d.Code, d.Name, d.Code)

	return err
}

// Delete disease
func (d *Disease) Delete(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM diseases WHERE code=$1", d.Code)

	return err
}

// Create disease
func (d *Disease) Create(db *sql.DB) error {
	err := db.QueryRow(
		`INSERT INTO
			diseases(code, name)
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
