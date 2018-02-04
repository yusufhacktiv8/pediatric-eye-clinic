package models

import (
	"database/sql"
)

// Role is a model for role
type Role struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

// FindRoles to find roles
func FindRoles(db *sql.DB, start, count int) ([]Role, error) {
	rows, err := db.Query(
		"SELECT id, code, name FROM roles LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	roles := []Role{}

	for rows.Next() {
		var d Role
		if err := rows.Scan(&d.ID, &d.Code, &d.Name); err != nil {
			return nil, err
		}
		roles = append(roles, d)
	}

	return roles, nil
}

// FindOne to find one role based on code
func (d *Role) FindOne(db *sql.DB) error {
	return db.QueryRow("SELECT code, name FROM roles WHERE code=$1",
		d.Code).Scan(&d.Code, &d.Name)
}

// Update role
func (d *Role) Update(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE roles SET code=$1, name=$2 WHERE code=$3",
			d.Code, d.Name, d.Code)

	return err
}

// Delete role
func (d *Role) Delete(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM roles WHERE code=$1", d.Code)

	return err
}

// Create create role
func (d *Role) Create(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO roles(code, name) VALUES($1, $2) RETURNING id",
		d.Code, d.Name).Scan(&d.ID)

	if err != nil {
		return err
	}

	return nil
}
