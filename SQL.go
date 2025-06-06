package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectSTKB() (*sql.DB, error) {
	// Kết nối SQLite
	db, err := sql.Open("sqlite3", "./database/STKB.db")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db, nil
}
func CreateTable(db *sql.DB) error {
	// Tạo bảng giáo viên nếu chưa tồn tại
	sqlCreateTables := `
	CREATE TABLE IF NOT EXISTS giaovien(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		tkb_name TEXT,
		name TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE IF NOT EXISTS monhoc(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE IF NOT EXISTS lophoc(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		khoi TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE IF NOT EXISTS phancong(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		giaovien_id INTEGER,
		monhoc_id INTEGER,
		tuan INTEGER,
		sotiet INTEGER,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (giaovien_id) REFERENCES giaovien(id),
		FOREIGN KEY (monhoc_id) REFERENCES monhoc(id)
	);
	`
	_, err := db.Exec(sqlCreateTables)
	if err != nil {
		return fmt.Errorf("khởi tạo bảng thất bại: %w", err)
	}
	return nil
}
func AddTeacher(db *sql.DB, name, tkb_name string) error {
	// Thêm giáo viên mới
	sqlInsert := `INSERT INTO giaovien (tkb_name, name) VALUES (?,?)`
	_, err := db.Exec(sqlInsert, tkb_name, name)
	if err != nil {
		return fmt.Errorf("thêm giáo viên thất bại: %w", err)
	}
	return nil
}
func DelTeacher(db *sql.DB, id int) error {
	// Xóa giáo viên theo ID
	sqlDelete := `DELETE FROM giaovien WHERE id = ?`
	_, err := db.Exec(sqlDelete, id)
	if err != nil {
		return fmt.Errorf("xóa giáo viên thất bại: %w", err)
	}
	return nil
}

// UpdateTeacher cập nhật thông tin giáo viên
func UpdateTeacher(db *sql.DB, id int, name, tkb_name string) error {
	// Cập nhật thông tin giáo viên
	sqlUpdate := `UPDATE giaovien SET name = ?, tkb_name = ? WHERE id = ?`
	_, err := db.Exec(sqlUpdate, name, tkb_name, id)
	if err != nil {
		return fmt.Errorf("cập nhật giáo viên thất bại: %w", err)
	}
	return nil
}

// GetTeachers lấy danh sách giáo viên
func GetTeachers(db *sql.DB) ([]map[string]interface{}, error) {
	// Lấy danh sách giáo viên
	sqlSelect := `SELECT id, tkb_name, name, created_at FROM giaovien`
	rows, err := db.Query(sqlSelect)
	if err != nil {
		return nil, fmt.Errorf("lấy danh sách giáo viên thất bại: %w", err)
	}
	defer rows.Close()

	var teachers []map[string]interface{}
	for rows.Next() {
		var id int
		var tkbName, name string
		var createdAt string
		if err := rows.Scan(&id, &tkbName, &name, &createdAt); err != nil {
			return nil, fmt.Errorf("lỗi khi quét hàng: %w", err)
		}
		teachers = append(teachers, map[string]interface{}{
			"id":         id,
			"tkb_name":   tkbName,
			"name":       name,
			"created_at": createdAt,
		})
	}
	return teachers, nil
}
func AddSubject(db *sql.DB, name string) error {
	// Thêm môn học mới
	sqlInsert := `INSERT INTO monhoc (name) VALUES (?)`
	_, err := db.Exec(sqlInsert, name)
	if err != nil {
		return fmt.Errorf("thêm môn học thất bại: %w", err)
	}
	return nil
}

// DelSubject xóa môn học theo ID
func DelSubject(db *sql.DB, id int) error {
	sqlDelete := `DELETE FROM monhoc WHERE id = ?`
	_, err := db.Exec(sqlDelete, id)
	if err != nil {
		return fmt.Errorf("xóa môn học thất bại: %w", err)
	}
	return nil
}

// UpdateSubject cập nhật thông tin môn học
func UpdateSubject(db *sql.DB, id int, name string) error {
	sqlUpdate := `UPDATE monhoc SET name = ? WHERE id = ?`
	_, err := db.Exec(sqlUpdate, name, id)
	if err != nil {
		return fmt.Errorf("cập nhật môn học thất bại: %w", err)
	}
	return nil
}

// GetSubjects lấy danh sách môn học
func GetSubjects(db *sql.DB) ([]map[string]interface{}, error) {
	// Lấy danh sách môn học
	sqlSelect := `SELECT id, name, created_at FROM monhoc`
	rows, err := db.Query(sqlSelect)
	if err != nil {
		return nil, fmt.Errorf("lấy danh sách môn học thất bại: %w", err)
	}
	defer rows.Close()

	var subjects []map[string]interface{}
	for rows.Next() {
		var id int
		var name, createdAt string
		if err := rows.Scan(&id, &name, &createdAt); err != nil {
			return nil, fmt.Errorf("lỗi khi quét hàng: %w", err)
		}
		subjects = append(subjects, map[string]interface{}{
			"id":         id,
			"name":       name,
			"created_at": createdAt,
		})
	}
	return subjects, nil
}

// Lớp học
func AddClass(db *sql.DB, name, khoi string) error {
	// Thêm lớp học mới
	sqlInsert := `INSERT INTO lophoc (name, khoi) VALUES (?, ?)`
	_, err := db.Exec(sqlInsert, name, khoi)
	if err != nil {
		return fmt.Errorf("thêm lớp học thất bại: %w", err)
	}
	return nil
}

// DelClass xóa lớp học theo ID
func DelClass(db *sql.DB, id int) error {
	sqlDelete := `DELETE FROM lophoc WHERE id = ?`
	_, err := db.Exec(sqlDelete, id)
	if err != nil {
		return fmt.Errorf("xóa lớp học thất bại: %w", err)
	}
	return nil
}

// UpdateClass cập nhật thông tin lớp học
func UpdateClass(db *sql.DB, id int, name, khoi string) error {
	sqlUpdate := `UPDATE lophoc SET name = ?, khoi = ? WHERE id = ?`
	_, err := db.Exec(sqlUpdate, name, khoi, id)
	if err != nil {
		return fmt.Errorf("cập nhật lớp học thất bại: %w", err)
	}
	return nil
}

// GetClasses lấy danh sách lớp học
func GetClasses(db *sql.DB) ([]map[string]interface{}, error) {
	// Lấy danh sách lớp học
	sqlSelect := `SELECT id, name, khoi, created_at FROM lophoc`
	rows, err := db.Query(sqlSelect)
	if err != nil {
		return nil, fmt.Errorf("lấy danh sách lớp học thất bại: %w", err)
	}
	defer rows.Close()

	var classes []map[string]interface{}
	for rows.Next() {
		var id int
		var name, khoi, createdAt string
		if err := rows.Scan(&id, &name, &khoi, &createdAt); err != nil {
			return nil, fmt.Errorf("lỗi khi quét hàng: %w", err)
		}
		classes = append(classes, map[string]interface{}{
			"id":         id,
			"name":       name,
			"khoi":       khoi,
			"created_at": createdAt,
		})
	}
	return classes, nil
}

// Phân công giáo viên cho môn học
func AssignTeacherToSubject(db *sql.DB, teacherID, subjectID, week, hours int) error {
	sqlInsert := `INSERT INTO phancong (giaovien_id, monhoc_id, tuan, sotiet) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(sqlInsert, teacherID, subjectID, week, hours)
	if err != nil {
		return fmt.Errorf("phân công giáo viên cho môn học thất bại: %w", err)
	}
	return nil
}

// DelAssignment xóa phân công giáo viên theo ID
func DelAssignment(db *sql.DB, id int) error {
	sqlDelete := `DELETE FROM phancong WHERE id = ?`
	_, err := db.Exec(sqlDelete, id)
	if err != nil {
		return fmt.Errorf("xóa phân công giáo viên thất bại: %w", err)
	}
	return nil
}

// UpdateAssignment cập nhật phân công giáo viên
func UpdateAssignment(db *sql.DB, id, teacherID, subjectID, week, hours int) error {
	sqlUpdate := `UPDATE phancong SET giaovien_id = ?, monhoc_id = ?, tuan = ?, sotiet = ? WHERE id = ?`
	_, err := db.Exec(sqlUpdate, teacherID, subjectID, week, hours, id)
	if err != nil {
		return fmt.Errorf("cập nhật phân công giáo viên thất bại: %w", err)
	}
	return nil
}

// GetAssignments lấy danh sách phân công giáo viên
func GetAssignments(db *sql.DB) ([]map[string]interface{}, error) {
	// Lấy danh sách phân công giáo viên
	sqlSelect := `
	SELECT p.id, g.name AS teacher_name, m.name AS subject_name, p.tuan, p.sotiet, p.created_at 
	FROM phancong p 
	JOIN giaovien g ON p.giaovien_id = g.id 
	JOIN monhoc m ON p.monhoc_id = m.id`
	rows, err := db.Query(sqlSelect)
	if err != nil {
		return nil, fmt.Errorf("lấy danh sách phân công giáo viên thất bại: %w", err)
	}
	defer rows.Close()

	var assignments []map[string]interface{}
	for rows.Next() {
		var id int
		var teacherName, subjectName string
		var week, hours int
		var createdAt string
		if err := rows.Scan(&id, &teacherName, &subjectName, &week, &hours, &createdAt); err != nil {
			return nil, fmt.Errorf("lỗi khi quét hàng: %w", err)
		}
		assignments = append(assignments, map[string]interface{}{
			"id":           id,
			"teacher_name": teacherName,
			"subject_name": subjectName,
			"week":         week,
			"hours":        hours,
			"created_at":   createdAt,
		})
	}
	return assignments, nil
}

func Run() {
	// Kết nối đến cơ sở dữ liệu SQLite
	db, err := ConnectSTKB()
	if err != nil {
		fmt.Println("Lỗi kết nối cơ sở dữ liệu:", err)
		return
	}
	defer db.Close()
	fmt.Println("Cơ sở dữ liệu đã được khởi tạo thành công!")

	// Thêm giáo viên và môn học mẫu

	var tearchers = []map[string]interface{}{}
	tearchers, err = GetTeachers(db)
	if err != nil {
		fmt.Println("Lỗi lấy danh sách giáo viên:", err)
		return
	}
	for _, teacher := range tearchers {
		fmt.Printf("Giáo viên: %d %s, TKB_Name: %s\n", teacher["id"], teacher["name"], teacher["tkb_name"])
	}
	subjects, err := GetSubjects(db)
	if err != nil {
		fmt.Println("Lỗi lấy danh sách môn học:", err)
		return
	}
	for _, subject := range subjects {
		fmt.Printf("Môn học: %d %s\n", subject["id"], subject["name"])
	}
	var classes []map[string]interface{}
	classes, err = GetClasses(db)
	if err != nil {
		fmt.Println("Lỗi lấy danh sách lớp học:", err)
		return
	}
	for _, class := range classes {
		fmt.Printf("Lớp học: %d %s, Khối: %s\n", class["id"], class["name"], class["khoi"])
	}

}
