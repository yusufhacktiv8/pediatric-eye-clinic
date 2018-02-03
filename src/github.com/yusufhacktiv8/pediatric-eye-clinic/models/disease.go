package models

import (
	"database/sql"
)

// Disease is a model for disease
type Disease struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

// FindDiseases to find diseases
func FindDiseases(db *sql.DB, start, count int) ([]Disease, error) {
	rows, err := db.Query(
		"SELECT id, code, name FROM diseases LIMIT $1 OFFSET $2",
		count, start)

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

func (d *Disease) getDisease(db *sql.DB) error {
	return db.QueryRow("SELECT code, name FROM diseases WHERE code=$1",
		d.ID).Scan(&d.Code, &d.Name)
}

func (d *Disease) updateDisease(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE diseases SET code=$1, name=$2 WHERE id=$3",
			d.Code, d.Name, d.ID)

	return err
}

func (d *Disease) deleteDisease(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM diseases WHERE id=$1", d.ID)

	return err
}

func (d *Disease) createDisease(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO diseases(name, price) VALUES($1, $2) RETURNING id",
		d.Code, d.Name).Scan(&d.ID)

	if err != nil {
		return err
	}

	return nil
}
