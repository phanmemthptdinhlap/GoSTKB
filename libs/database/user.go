package database

import (
	"GoSTKB/libs/myauth"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTableUsers(db *sql.DB) error {
	// Tạo bảng nếu chưa tồn tại
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT UNIQUE,
		password TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}
	return nil
}
func CreateUser(db *sql.DB, name, email, password string) error {
	// Mã hóa mật khẩu
	hashedPassword, err := myauth.HasPassword(password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	// Thêm người dùng mới
	_, err = db.Exec(`INSERT INTO users (name, email, password) VALUES (?, ?, ?)`, name, email, hashedPassword)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}
func GetUserByEmail(db *sql.DB, email string) (string, error) {
	// Lấy người dùng theo email
	var password string
	err := db.QueryRow(`SELECT password FROM users WHERE email = ?`, email).Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("user not found")
		}
		return "", fmt.Errorf("failed to get user: %w", err)
	}
	return password, nil
}
func GetAllUsers(db *sql.DB) ([]map[string]interface{}, error) {
	// Lấy tất cả người dùng
	rows, err := db.Query(`SELECT id, name, email, created_at FROM users`)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}
	defer rows.Close()

	var users []map[string]interface{}
	for rows.Next() {
		var id int
		var name, email string
		var createdAt string
		if err := rows.Scan(&id, &name, &email, &createdAt); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		user := map[string]interface{}{
			"id":         id,
			"name":       name,
			"email":      email,
			"created_at": createdAt,
		}
		users = append(users, user)
	}
	return users, nil
}
func DeleteUser(db *sql.DB, id int) error {
	// Xóa người dùng theo ID
	_, err := db.Exec(`DELETE FROM users WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
func UpdateUser(db *sql.DB, id int, name, email, password string) error {
	// Mã hóa mật khẩu mới nếu có
	var hashedPassword string
	if password != "" {
		var err error
		hashedPassword, err = myauth.HasPassword(password)
		if err != nil {
			return fmt.Errorf("failed to hash password: %w", err)
		}
	}

	// Cập nhật thông tin người dùng
	query := `UPDATE users SET name = ?`
	args := []interface{}{name}
	if email != "" {
		query += `, email = ?`
		args = append(args, email)
	}
	if hashedPassword != "" {
		query += `, password = ?`
		args = append(args, hashedPassword)
	}
	query += ` WHERE id = ?`
	args = append(args, id)

	_, err := db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}
