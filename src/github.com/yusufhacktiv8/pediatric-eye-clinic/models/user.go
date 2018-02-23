package models

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
)

// User is a model for user
type User struct {
	gorm.Model
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Role     Role   `json:"role"`
}

// FindUsers to find users
func FindUsers(db *sql.DB, start, count int, searchText string) ([]User, error) {
	rows, err := db.Query(
		`SELECT
			u.id,
			u.email,
			u.name,
			r.code AS role_code,
			r.name AS role_name
		FROM
			users u
		LEFT JOIN roles r ON u.role = r.id
		WHERE u.email LIKE $3 or u.name LIKE $3
		LIMIT $1 OFFSET $2`,
		count, start, "%"+searchText+"%")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var roleCode sql.NullString
		var roleName sql.NullString
		var d User
		if err := rows.Scan(&d.ID, &d.Email, &d.Name, &roleCode, &roleName); err != nil {
			return nil, err
		}

		if roleCode.Valid {
			d.Role.Code = roleCode.String
			d.Role.Name = roleName.String
		}
		users = append(users, d)
	}

	return users, nil
}

func CountUsers(db *sql.DB, searchText string) (int, error) {
	rows, err := db.Query(
		`SELECT
			count(1) AS rowsCount
		FROM
			users u
		WHERE u.email LIKE $1 or u.name LIKE $1`,
		"%"+searchText+"%")

	if err != nil {
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

func (d *User) FindOne(db *sql.DB) error {
	return db.QueryRow("SELECT email, name, password FROM users WHERE email=$1 AND password=$2",
		d.Email, d.Password).Scan(&d.Email, &d.Name, &d.Password)
}

// Update user
func (d *User) Update(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE users SET name=$1, role=$2 WHERE email=$3",
			d.Name, d.Role.ID, d.Email)

	return err
}

// Delete user
func (d *User) Delete(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM users WHERE email=$1", d.Email)

	return err
}

// Create create user
func (d *User) Create(db *sql.DB) error {
	err := db.QueryRow(
		`INSERT INTO users
		(email,
			password,
			name,
			role) VALUES($1, $2, $3, $4)
			ON CONFLICT (email) DO UPDATE
			SET name=$5 RETURNING id`,
		d.Email, d.Password, d.Name, d.Role.ID, d.Name).Scan(&d.ID)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
