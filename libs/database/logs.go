package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTableLogs(db *sql.DB) error {
	// Tạo bảng nếu chưa tồn tại
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS logs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		action TEXT,
		timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)`)
	if err != nil {
		return fmt.Errorf("failed to create logs table: %w", err)
	}
	return nil
}
func LogAction(db *sql.DB, userID int, action string) error {
	// Ghi log hành động
	_, err := db.Exec(`INSERT INTO logs (user_id, action) VALUES (?, ?)`, userID, action)
	if err != nil {
		return fmt.Errorf("failed to log action: %w", err)
	}
	return nil
}
func GetLogs(db *sql.DB) ([]map[string]interface{}, error) {
	// Lấy tất cả logs
	rows, err := db.Query(`SELECT id, user_id, action, timestamp FROM logs`)
	if err != nil {
		return nil, fmt.Errorf("failed to get logs: %w", err)
	}
	defer rows.Close()

	var logs []map[string]interface{}
	for rows.Next() {
		var id, userID int
		var action, timestamp string
		if err := rows.Scan(&id, &userID, &action, &timestamp); err != nil {
			return nil, fmt.Errorf("failed to scan log: %w", err)
		}
		log := map[string]interface{}{
			"id":        id,
			"user_id":   userID,
			"action":    action,
			"timestamp": timestamp,
		}
		logs = append(logs, log)
	}
	return logs, nil
}
