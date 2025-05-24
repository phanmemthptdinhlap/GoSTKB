package database

import (
	"GoSTKB/libs/myauth"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func Connect() (*sql.DB, error) {
	// Kết nối SQLite
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db, nil
}
func CreateTable_admin(db *sql.DB) error {
	// Tạo bảng nếu chưa tồn tại
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS admin(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user TEXT,
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
func CreateTable_users(db *sql.DB) error {
	// Tạo bảng nếu chưa tồn tại
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}
	return nil
}
