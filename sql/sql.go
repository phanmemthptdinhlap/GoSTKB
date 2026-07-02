package SQL

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func ConnectSTKB() (*sql.DB, error) {
	// Kết nối SQLite
	db, err := sql.Open("sqlite3", "./sql/STKB.db")
	if err != nil {
		return nil, fmt.Errorf("không kết nối được với CSDL : %w", err)
	}
	return db, nil
}
func CreateTable(db *sql.DB) error {
	// Tạo bảng giáo viên nếu chưa tồn tại
	sqlCreateTables := `
	CREATE TABLE IF NOT EXISTS giaovien(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		ten_ngan TEXT,
		ho_ten TEXT,
		mon_chinh_id INTEGER
		FOREIGN KEY (mon_chinh_id) REFERENCES monhoc(id)
	);
	CREATE TABLE IF NOT EXISTS monhoc(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		ten_mon TEXT,
		loai_mon TEXT
	);
	CREATE TABLE IF NOT EXISTS lophoc(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		ten_lop TEXT,
		khoi_lop TEXT
	);
	CREATE TABLE IF NOT EXISTS phancong(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		giaovien_id INTEGER,
		lop_id INTEGER,
		monhoc_id INTEGER,
		tuan INTEGER,
		sotiet INTEGER,
		FOREIGN KEY (giaovien_id) REFERENCES giaovien(id),
		FOREIGN KEY (lop_id) REFERENCES lophoc(id),
		FOREIGN KEY (monhoc_id) REFERENCES monhoc(id)
	);
	CREATE TABLE IF NOT EXISTS rangbuoc(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		giaovien_id INTEGER,
		thu INTEGR,
		tiet INTEGER,
		loai_rang_buoc TEXT,
		FOREIGN KEY (giaovien_id) REFERENCES giaovien(id)
	);
	`
	_, err := db.Exec(sqlCreateTables)
	if err != nil {
		return fmt.Errorf("khởi tạo bảng thất bại: %w", err)
	}
	return nil
}
