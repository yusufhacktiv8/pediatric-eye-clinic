package models

import (
	"github.com/jinzhu/gorm"
)

// Role is a model for role
type Role struct {
	gorm.Model
	Code string `json:"code"`
	Name string `json:"name"`
}

/*
// FindRoles to find roles
func FindRoles(db *sql.DB, start, count int, searchText string) ([]Role, error) {
	rows, err := db.Query(
		"SELECT code, name FROM roles WHERE code LIKE $3 OR name LIKE $3  ORDER BY code LIMIT $1 OFFSET $2",
		count, start, "%"+searchText+"%")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	roles := []Role{}

	for rows.Next() {
		var d Role
		if err := rows.Scan(&d.Code, &d.Name); err != nil {
			return nil, err
		}
		roles = append(roles, d)
	}

	return roles, nil
}

func CountRoles(db *sql.DB, searchText string) (int, error) {
	rows, err := db.Query(
		`SELECT
			count(1) AS rowsCount
		FROM
			roles r
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

// FindOne to find one role based on code
func (d *Role) FindOne(db *sql.DB) error {
	return db.QueryRow("SELECT code, name FROM roles WHERE code=$1",
		d.Code).Scan(&d.Code, &d.Name)
}

// Update role
func (d *Role) Update(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE roles SET code=$1, name=$2 WHERE code=$1",
			d.Code, d.Name)

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
		`INSERT INTO
			roles(code, name)
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
*/
