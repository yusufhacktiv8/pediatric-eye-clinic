package models

import (
	"database/sql"
)

// User is a model for user
type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Role     Role   `json:"role"`
}

// FindUsers to find users
func FindUsers(db *sql.DB, start, count int) ([]User, error) {
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
		LIMIT $1 OFFSET $2`,
		count, start)

	if err != nil {
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

func CountUsers(db *sql.DB) (int, error) {
	rows, err := db.Query(
		`SELECT
			count(1) AS rowsCount
		FROM
			users u`)

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
			role) VALUES($1, $2, $3, $4) RETURNING id`,
		d.Email, d.Password, d.Name, d.Role.ID).Scan(&d.ID)

	if err != nil {
		return err
	}

	return nil
}
