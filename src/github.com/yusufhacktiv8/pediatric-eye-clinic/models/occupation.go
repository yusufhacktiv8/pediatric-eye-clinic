package models

import (
	"database/sql"
)

// Occupation is a model for occupation
type Occupation struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

// FindOccupations to find occupations
func FindOccupations(db *sql.DB, start, count int) ([]Occupation, error) {
	rows, err := db.Query(
		"SELECT id, code, name FROM occupations LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	occupations := []Occupation{}

	for rows.Next() {
		var d Occupation
		if err := rows.Scan(&d.ID, &d.Code, &d.Name); err != nil {
			return nil, err
		}
		occupations = append(occupations, d)
	}

	return occupations, nil
}

// FindOne to find one occupation based on code
func (d *Occupation) FindOne(db *sql.DB) error {
	return db.QueryRow("SELECT code, name FROM occupations WHERE code=$1",
		d.Code).Scan(&d.Code, &d.Name)
}

// Update occupation
func (d *Occupation) Update(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE occupations SET code=$1, name=$2 WHERE code=$3",
			d.Code, d.Name, d.Code)

	return err
}

// Delete occupation
func (d *Occupation) Delete(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM occupations WHERE code=$1", d.Code)

	return err
}

// Create create occupation
func (d *Occupation) Create(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO occupations(code, name) VALUES($1, $2) RETURNING id",
		d.Code, d.Name).Scan(&d.ID)

	if err != nil {
		return err
	}

	return nil
}
