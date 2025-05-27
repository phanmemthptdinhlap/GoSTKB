package database

import (
	"GoSTKB/libs/myauth"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func Connect() (*sql.DB, error) {
	// Kết nối SQLite
	db, err := sql.Open("sqlite3", "./database/database.db")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db, nil
}
func CreateTableAdmin(db *sql.DB) error {
	// Tạo bảng nếu chưa tồn tại
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS admin(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user TEXT UNIQUE,
		password TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}
	// Tạo tài khoản admin mặc định
	hashedPassword, _ := myauth.HasPassword("admin1234")
	_, err = db.Exec(`INSERT INTO admin (user, password) VALUES (?, ?)`, "admin", hashedPassword)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}
	return nil
}
func CreateAdmin(db *sql.DB, user, password string) error {
	// Mã hóa mật khẩu
	hashedPassword, err := myauth.HasPassword(password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	// Thêm admin mới
	_, err = db.Exec(`INSERT INTO admin (user, password) VALUES (?, ?)`, user, hashedPassword)
	if err != nil {
		return fmt.Errorf("failed to create admin: %w", err)
	}
	return nil
}
func GetAdminByUser(db *sql.DB, user string) (string, error) {
	// Lấy admin theo tên người dùng
	var password string
	err := db.QueryRow(`SELECT password FROM admin WHERE user = ?`, user).Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("admin not found")
		}
		return "", fmt.Errorf("failed to get admin: %w", err)
	}
	return password, nil
}
func GetAdminByID(db *sql.DB, id int) (string, error) {
	// Lấy admin theo ID
	var user string
	err := db.QueryRow(`SELECT user FROM admin WHERE id = ?`, id).Scan(&user)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("admin not found")
		}
		return "", fmt.Errorf("failed to get admin: %w", err)
	}
	return user, nil
}
