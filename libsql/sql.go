package libsql

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// sql cmd
const (
	  sqlpath = "./stkb.db"
		sqlCreateTables = `
		CREATE TABLE IF NOT EXISTS giaovien(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			ten_ngan TEXT UNIQUE,
			ho_ten TEXT,
			mon_chinh_id INTEGER,
			FOREIGN KEY (mon_chinh_id) REFERENCES monhoc(id)
		);
		CREATE TABLE IF NOT EXISTS monhoc(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			ten_mon TEXT UNIQUE,
			loai_mon TEXT
		);
		CREATE TABLE IF NOT EXISTS lophoc(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			ten_lop TEXT UNIQUE,
			khoi_lop TEXT,
			gvcn_id INTEGER,
			FOREIGN KEY (gvcn_id) REFERENCES giaovien(id)
		);
		CREATE TABLE IF NOT EXISTS phancong(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			giao_vien_id INTEGER,
			lop_id INTEGER,
			mon_hoc_id INTEGER,
			tong_tiet INTEGER,
			FOREIGN KEY (giao_vien_id) REFERENCES giaovien(id),
			FOREIGN KEY (lop_id) REFERENCES lophoc(id),
			FOREIGN KEY (mon_hoc_id) REFERENCES monhoc(id)
		);
		CREATE TABLE IF NOT EXISTS chitiet(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			phan_cong_id INTEGER,
			tuan INTEGER,
			so_tiet INTEGER,
			FOREIGN KEY (phan_cong_id) REFERENCES phancong(id)
		);
		CREATE TABLE IF NOT EXISTS tiendo(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			phan_cong_id INTEGER,
			tuan INTEGER,
			so_tiet INTEGER,
			FOREIGN KEY (phan_cong_id) REFERENCES phancong(id)
		);
		CREATE TABLE IF NOT EXISTS rangbuoc(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			giao_vien_id INTEGER,
			thu INTEGER,
			tiet INTEGER,
			loai_rang_buoc TEXT,
			FOREIGN KEY (giao_vien_id) REFERENCES giaovien(id)
		);`
	)
type SqlTKB struct {
	db *sql.DB
}

func ConnectSTKB() (*SqlTKB, error) {
	// Kết nối SQLite
	db, err := sql.Open("sqlite3", sqlpath)
	if err != nil {
		return nil, fmt.Errorf("không kết nối được với STKB : %w", err)
	}
	sqloptions:=`
	PRAGMA foreign_keys = ON;
	PRAGMA journal_mode = WAL;
	PRAGMA synchronous = NORMAL;
	`
	_, err = db.Exec(sqloptions)
	if err != nil {
		return nil, fmt.Errorf("không thể thiết lập các tùy chọn : %w", err)
	}
	_, err = db.Exec(sqlCreateTables)
	if err != nil {
		return nil, fmt.Errorf("không thể tạo bảng : %w", err)
	}
	return &SqlTKB{db}, nil
}

func (s *SqlTKB) Close() error {
	return s.db.Close()
}

// Giáo viên

const (
	sqlSelectGiaoVien = `
	SELECT * FROM giaovien
	WHERE id = ?;
	`
	sqlSelectAllGiaoVien = `
	SELECT * FROM giaovien;
	`
	sqlFindGiaoVien = `
	SELECT * FROM giaovien
	WHERE ten_ngan = ?;
	`
	sqlInsertGiaoVien = `
	INSERT INTO giaovien(ten_ngan, ho_ten, mon_chinh_id)
	VALUES (?, ?, ?);
	`
	sqlEditGiaoVien = `
	UPDATE giaovien
	SET ten_ngan = ?, ho_ten = ?, mon_chinh_id = ?
	WHERE id = ?;
	`
	sqlDeleteGiaoVien = `
	DELETE FROM giaovien
	WHERE id = ?;
	`
)

type GiaoVien struct {
	ID int `json:"id"`
	TenNgan string `json:"ten_ngan"`
	HoTen string `json:"ho_ten"`
	MonChinhId int `json:"mon_chinh_id"`
	Action string `json:"action,omitempty"`
}

func (s *SqlTKB) SelectAllGiaoVien() ([]GiaoVien, error) {
	var giaovien []GiaoVien
	Rows, err := s.db.Query(sqlSelectAllGiaoVien)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu từ bảng giaovien : %w", err)
	}
	defer Rows.Close()
	for Rows.Next() {
		var gv GiaoVien
		err := Rows.Scan(&gv.ID, &gv.TenNgan, &gv.HoTen, &gv.MonChinhId)
		if err != nil {
			return nil, fmt.Errorf("Không thể quét dữ liệu vào biến giaovien : %w", err)
		}
		giaovien = append(giaovien, gv)
	}
	return giaovien, nil
}

func (s *SqlTKB) SelectGiaoVien(id int) (*GiaoVien, error) {
	var gv GiaoVien
	err := s.db.QueryRow(sqlSelectGiaoVien, id).Scan(&gv.ID, &gv.TenNgan, &gv.HoTen, &gv.MonChinhId)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu từ bảng giaovien : %w", err)
	}
	return &gv, nil
}

func (s *SqlTKB) InsertGiaoVien(giaovien []GiaoVien) (int, error) {
	// Bắt đầu một transaction
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("không thể bắt đầu giao dịch: %w", err)
	}

	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer tx.Rollback()

	var count int
	// Chuẩn bị câu lệnh SQL (Prepared Statement) để tái sử dụng trong vòng lặp
	stmt, err := tx.Prepare(sqlInsertGiaoVien)
	if err != nil {
		return 0, fmt.Errorf("không thể chuẩn bị câu lệnh: %w", err)
	}
	defer stmt.Close() // Đóng statement khi xong
	fmt.Printf("Thêm mới %d giáo viên: ", len(giaovien))
	for _, gv := range giaovien {
		// Dùng statement đã chuẩn bị để thực thi
		_, err = stmt.Exec(gv.TenNgan, gv.HoTen, gv.MonChinhId)
		if err != nil {
			// Lỗi được gán cho biến err, defer sẽ kích hoạt Rollback()
			//fmt.Printf("Lỗi thêm giáo viên: (%v)\n", err)
			return count, fmt.Errorf("không thể thêm giáo viên %s (%w)", gv.TenNgan, err)
		}
		//fmt.Printf("Đã thêm giáo viên: %s \n", gv.TenNgan)
		count++
	}

	// Nếu mọi thứ ổn, commit giao dịch
	err = tx.Commit()
	if err != nil {
		return count, fmt.Errorf("không thể chốt giao dịch: %w", err)
	}

	return count, nil
}

func (s *SqlTKB) EditGiaoVien(giaovien []GiaoVien) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("không thể bắt đầu giao dịch: %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer tx.Rollback()

	var count int
	// Chuẩn bị câu lệnh SQL (Prepared Statement) để tái sử dụng trong vòng lặp
	stmt, err := tx.Prepare(sqlEditGiaoVien)
	if err != nil {
		return 0, fmt.Errorf("không thể chuẩn bị câu lệnh: %w", err)
	}
	defer stmt.Close() // Đóng statement khi xong
	for _, gv := range giaovien {
		// Dùng statement đã chuẩn bị để thực thi
		_, err = stmt.Exec(gv.TenNgan, gv.HoTen, gv.MonChinhId, gv.ID)
		if err != nil {
			// Lỗi được gán cho biến err, defer sẽ kích hoạt Rollback()
			return count, fmt.Errorf("không thể sửa giáo viên %s (%w)", gv.TenNgan, err)
		}
		count++
	}
	// Nếu mọi thứ ổn, commit giao dịch
	err = tx.Commit()
	if err != nil {
		return count, fmt.Errorf("không thể chốt giao dịch: %w",	err)
	}
	return count, nil
}

func (s *SqlTKB) DeleteGiaoVien(ids []int) (int, error) {
	tx,err:=s.db.Begin()
	if err!=nil{
		return 0, fmt.Errorf("Không thể bắt đầu giao dịch %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer tx.Rollback()

	var count int
	// Chuẩn bị câu lệnh SQL (Prepared Statement) để tái sử dụng trong vòng lặp
	stmt, err := tx.Prepare(sqlDeleteGiaoVien)
	if err != nil {
		return 0, fmt.Errorf("Không thể chuẩn bị câu lệnh %w", err)
	}
	defer stmt.Close() // Đóng statement khi xong
	for _, id := range ids {
		// Dùng statement đã chuẩn bị để thực thi
		_, err = stmt.Exec(id)
		if err != nil {
			// Lỗi được gán cho biến err, defer sẽ kích hoạt Rollback()
			return count, fmt.Errorf("Không thể xóa giáo viên %d (%w)", id, err)
		}
		count++
	}
	// Nếu mọi thứ ổn, commit giao dịch
	err = tx.Commit()
	if err != nil {
		return count, fmt.Errorf("Không thể chốt giao dịch %w", err)
	}
	return count, nil
}

// Mon hoc

const (
	sqlSelectMonHoc = `
	SELECT * FROM monhoc
	WHERE id = ?;
	`
	sqlSelectAllMonHoc = `
	SELECT * FROM monhoc;
	`
	sqlFindMonHoc = `
	SELECT * FROM monhoc
	WHERE ten_mon = ?;
	`
	sqlInsertMonHoc = `
	INSERT INTO monhoc(ten_mon, loai_mon)
	VALUES (?, ?);
	`
	sqlEditMonHoc = `
	UPDATE monhoc
	SET ten_mon = ?, loai_mon = ?
	WHERE id = ?;
	`
	sqlDeleteMonHoc = `
	DELETE FROM monhoc
	WHERE id = ?;
	`
)

type MonHoc struct {
	ID int `json:"id"`
	TenMon string `json:"ten_mon"`
	LoaiMon string `json:"loai_mon"`
	Action string `json:"action,omitempty"`
}

func (s *SqlTKB) SelectAllMonHoc() ([]MonHoc, error) {
	var monhoc []MonHoc
	Rows, err := s.db.Query(sqlSelectAllMonHoc)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu từ bảng monhoc : %w", err)
	}
	defer Rows.Close()
	for Rows.Next() {
		var mh MonHoc
		err := Rows.Scan(&mh.ID, &mh.TenMon, &mh.LoaiMon)
		if err != nil {
			return nil, fmt.Errorf("Không thể quét dữ liệu vào biến monhoc : %w", err)
		}
		monhoc = append(monhoc, mh)
	}
	return monhoc, nil
}

func (s *SqlTKB) SelectMonHoc(id int) (*MonHoc, error) {
	var mh MonHoc
	err := s.db.QueryRow(sqlSelectMonHoc, id).Scan(&mh.ID, &mh.TenMon, &mh.LoaiMon)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu từ bảng monhoc : %w", err)
	}
	return &mh, nil
}

func (s *SqlTKB) InsertMonHoc(monhoc []MonHoc) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("không thể bắt đầu giao dịch: %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer tx.Rollback()

	var count int
	// Chuẩn bị câu lệnh SQL (Prepared Statement) để tái sử dụng trong vòng lặp
	stmt, err := tx.Prepare(sqlInsertMonHoc)
	if err != nil {
		return 0, fmt.Errorf("không thể chuẩn bị câu lệnh: %w", err)
	}
	defer stmt.Close() // Đóng statement khi xong
	fmt.Printf("Thêm mới %d môn học:\n ", len(monhoc))
	for _, mh := range monhoc {
		// Dùng statement đã chuẩn bị để thực thi
		_, err = stmt.Exec(mh.TenMon, mh.LoaiMon)
		if err != nil {
			// Lỗi được gán cho biến err, defer sẽ kích hoạt Rollback()
			//fmt.Printf("Lỗi thêm môn học: (%v)\n", err)
			return count, fmt.Errorf("không thể thêm môn học %s (%w)", mh.TenMon, err)
		}
		//fmt.Printf("Đã thêm môn học: %s \n", mh.TenMon)
		count++
	}
	// Nếu mọi thứ ổn, commit giao dịch
	err = tx.Commit()
	if err != nil {
		return count, fmt.Errorf("không thể chốt giao dịch: %w", err)
	}
	return count, nil
}

func (s *SqlTKB) EditMonHoc(monhoc []MonHoc) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("không thể bắt đầu giao dịch: %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer tx.Rollback()

	var count int
	// Chuẩn bị câu lệnh SQL (Prepared Statement) để tái sử dụng trong vòng lặp
	stmt, err := tx.Prepare(sqlEditMonHoc)
	if err != nil {
		return 0, fmt.Errorf("không thể chuẩn bị câu lệnh: %w", err)
	}
	defer stmt.Close() // Đóng statement khi xong
	for _, mh := range monhoc {
		// Dùng statement đã chuẩn bị để thực thi
		_, err = stmt.Exec(mh.TenMon, mh.LoaiMon, mh.ID)
		if err != nil {
			// Lỗi được gán cho biến err, defer sẽ kích hoạt Rollback()
			return count, fmt.Errorf("không thể sửa môn học %s (%w)", mh.TenMon, err)
		}
		count++
	}
	// Nếu mọi thứ ổn, commit giao dịch
	err = tx.Commit()
	if err != nil {
		return count, fmt.Errorf("không thể chốt giao dịch: %w",	err)
	}
	return count, nil
}

func (s *SqlTKB) DeleteMonHoc(ids []int) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("Không thể bắt đầu giao dịch %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer tx.Rollback()

	var count int
	// Chuẩn bị câu lệnh SQL (Prepared Statement) để tái sử dụng trong vòng lặp
	stmt, err := tx.Prepare(sqlDeleteMonHoc)
	if err != nil {
		return 0, fmt.Errorf("Không thể chuẩn bị câu lệnh %w", err)
	}
	defer stmt.Close() // Đóng statement khi xong
	for _, id := range ids {
			// Dùng statement đã chuẩn bị để thực thi
		_, err = stmt.Exec(id)
		if err != nil {
				// Lỗi được gán cho biến err, defer sẽ kích hoạt Rollback()
			return count, fmt.Errorf("Không thể xóa môn học %d (%w)", id, err)
		}
		count++
	}
			// Nếu mọi thứ ổn, commit giao dịch
	err = tx.Commit()
		if err != nil {
			return count, fmt.Errorf("Không thể chốt giao dịch %w", err)
		}
	return count, nil
}

// Lớp học

const (
	sqlSelectLopHoc = `
	SELECT * FROM lophoc
	WHERE id = ?;
	`
	sqlSelectAllLopHoc = `
	SELECT * FROM lophoc;
	`
	sqlFindLopHoc = `
	SELECT * FROM lophoc
	WHERE ten_lop = ?;
	`
	sqlInsertLopHoc = `
	INSERT INTO lophoc(ten_lop, khoi_lop, gvcn_id)
	VALUES (?, ?, ?);
	`
	sqlEditLopHoc = `
	UPDATE lophoc
	SET ten_lop = ?, khoi_lop = ?, gvcn_id = ?
	WHERE id = ?;
	`
	sqlDeleteLopHoc = `
	DELETE FROM lophoc
	WHERE id = ?;
	`
)

type LopHoc struct {
	ID int `json:"id"`
	TenLop string `json:"ten_lop"`
	KhoiLop string `json:"khoi_lop"`
	GvcnId int `json:"gvcn_id"`
	Action string `json:"action,omitempty"`
}

func (s *SqlTKB) SelectAllLopHoc() ([]LopHoc, error) {
	var lophoc []LopHoc
	Rows, err := s.db.Query(sqlSelectAllLopHoc)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu từ bảng lophoc : %w", err)
	}
	defer Rows.Close()
	for Rows.Next() {
		var lh LopHoc
		err := Rows.Scan(&lh.ID, &lh.TenLop, &lh.KhoiLop, &lh.GvcnId)
		if err != nil {
			return nil, fmt.Errorf("Không thể quét dữ liệu vào biến lophoc : %w", err)
		}
		lophoc = append(lophoc, lh)
	}
	return lophoc, nil
}

func (s *SqlTKB) SelectLopHoc(id int) (*LopHoc, error) {
	var lh LopHoc
	err := s.db.QueryRow(sqlSelectLopHoc, id).Scan(&lh.ID, &lh.TenLop, &lh.KhoiLop, &lh.GvcnId)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu từ bảng lophoc : %w", err)
	}
	return &lh, nil
}

func (s *SqlTKB) InsertLopHoc(lophoc []LopHoc) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("không thể bắt đầu giao dịch: %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer tx.Rollback()

	var count int
	// Chuẩn bị câu lệnh SQL (Prepared Statement) để tái sử dụng trong vòng lặp
	stmt, err := tx.Prepare(sqlInsertLopHoc)
	if err != nil {
		return 0, fmt.Errorf("không thể chuẩn bị câu lệnh: %w", err)
	}
	defer stmt.Close() // Đóng statement khi xong
	for _, lh := range lophoc {
		// Dùng statement đã chuẩn bị để thực thi
		_, err = stmt.Exec(lh.TenLop, lh.KhoiLop, lh.GvcnId)
		if err != nil {
			// Lỗi được gán cho biến err, defer sẽ kích hoạt Rollback()
			return count, fmt.Errorf("không thể thêm lớp học %s (%w)", lh.TenLop, err)
		}
		count++
	}
	// Nếu mọi thứ ổn, commit giao dịch
	err = tx.Commit()
	if err != nil {
		return count, fmt.Errorf("không thể chốt giao dịch: %w", err)
	}
	return count, nil
}

func (s *SqlTKB) EditLopHoc(lophoc []LopHoc) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("không thể bắt đầu giao dịch: %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer tx.Rollback()

	var count int
	// Chuẩn bị câu lệnh SQL (Prepared Statement) để tái sử dụng trong vòng lặp
	stmt, err := tx.Prepare(sqlEditLopHoc)
	if err != nil {
		return 0, fmt.Errorf("không thể chuẩn bị câu lệnh: %w", err)
	}
	defer stmt.Close() // Đóng statement khi xong
	for _, lh := range lophoc {
		// Dùng statement đã chuẩn bị để thực thi
		_, err = stmt.Exec(lh.TenLop, lh.KhoiLop, lh.GvcnId, lh.ID)
		if err != nil {
			// Lỗi được gán cho biến err, defer sẽ kích hoạt Rollback()
			return count, fmt.Errorf("không thể sửa lớp học %s (%w)", lh.TenLop, err)
		}
		count++
	}
	// Nếu mọi thứ ổn, commit giao dịch
	err = tx.Commit()
	if err != nil {
		return count, fmt.Errorf("không thể chốt giao dịch: %w",	err)
	}
	return count, nil
}

func (s *SqlTKB) DeleteLopHoc(ids []int) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("Không thể bắt đầu giao dịch %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer tx.Rollback()

	var count int
	// Chuẩn bị câu lệnh SQL (Prepared Statement) để tái sử dụng trong vòng lặp
	stmt, err := tx.Prepare(sqlDeleteLopHoc)
	if err != nil {
		return 0, fmt.Errorf("Không thể chuẩn bị câu lệnh %w", err)
	}
	defer stmt.Close() // Đóng statement khi xong
	for _, id := range ids {
			// Dùng statement đã chuẩn bị để thực thi
		_, err = stmt.Exec(id)
		if err != nil {
				// Lỗi được gán cho biến err, defer sẽ kích hoạt Rollback()
			return count, fmt.Errorf("Không thể xóa lớp học %d (%w)", id, err)
		}
		count++
	}
			// Nếu mọi thứ ổn, commit giao dịch
	err = tx.Commit()
	if err != nil {
		return count, fmt.Errorf("Không thể chốt giao dịch %w", err)
	}
	return count, nil
}

// Phần cống

const (
	sqlSelectPhanCong = `
	SELECT * FROM phancong
	WHERE id = ?;
	`
	sqlSelectAllPhanCong = `
	SELECT * FROM phancong;
	`
	sqlFindPhanCong = `
	SELECT * FROM phancong
	WHERE giao_vien_id = ?;
	`
	sqlInsertPhanCong = `
	INSERT INTO phancong(giao_vien_id, lop_id, mon_hoc_id, tong_tiet)
	VALUES (?, ?, ?, ?);
	`
	sqlEditPhanCong = `
	UPDATE phancong
	SET giao_vien_id = ?, lop_id = ?, mon_hoc_id = ?, tong_tiet = ?
	WHERE id = ?;
	`
	sqlDeletePhanCong = `
	DELETE FROM phancong
	WHERE id = ?;
	`
)

type PhanCong struct {
	ID int `json:"id"`
	GiaoVienId int `json:"giao_vien_id"`
	LopId int `json:"lop_id"`
	MonHocId int `json:"mon_hoc_id"`
	TongTiet int `json:"tong_tiet"`
	Action string `json:"action,omitempty"`
}

func (s *SqlTKB) SelectAllPhanCong() ([]PhanCong, error) {
	var phancong []PhanCong
	Rows, err := s.db.Query(sqlSelectAllPhanCong)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu từ bảng phancong : %w", err)
	}
	defer Rows.Close()
	for Rows.Next() {
		var pc PhanCong
		err := Rows.Scan(&pc.ID, &pc.GiaoVienId, &pc.LopId, &pc.MonHocId)
		if err != nil {
			return nil, fmt.Errorf("Không thể quét dữ liệu vào biến phancong : %w", err)
		}
		phancong = append(phancong, pc)
	}
	return phancong, nil
}

func (s *SqlTKB) SelectPhanCong(id int) (*PhanCong, error) {
	var pc PhanCong
	err := s.db.QueryRow(sqlSelectPhanCong, id).Scan(&pc.ID, &pc.GiaoVienId, &pc.LopId, &pc.MonHocId)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu từ bảng phancong : %w", err)
	}
	return &pc, nil
}

func (s *SqlTKB) InsertPhanCong(phancong []PhanCong) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("không thể bắt đầu giao dịch: %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer tx.Rollback()

	var count int
	// Chuẩn bị câu lệnh SQL (Prepared Statement) để tái sử dụng trong vòng lặp
	stmt, err := tx.Prepare(sqlInsertPhanCong)
	if err != nil {
		return 0, fmt.Errorf("không thể chuẩn bị câu lệnh: %w", err)
	}
	defer stmt.Close() // Đóng statement khi xong
	for _, pc := range phancong {
		// Dùng statement đã chuẩn bị để thực thi
		_, err = stmt.Exec(pc.GiaoVienId, pc.LopId, pc.MonHocId)
		if err != nil {
			// Lỗi được gán cho biến err, defer sẽ kích hoạt Rollback()
			return count, fmt.Errorf("không thể thêm phần cống %d (%w)", pc.GiaoVienId, err)
		}
		count++
	}
	// Nếu mọi thứ ổn, commit giao dịch
	err = tx.Commit()
	if err != nil {
		return count, fmt.Errorf("không thể chốt giao dịch: %w", err)
	}
	return count, nil
}

func (s *SqlTKB) EditPhanCong(phancong []PhanCong) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("không thể bắt đầu giao dịch: %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer tx.Rollback()

	var count int
	// Chuẩn bị câu lệnh SQL (Prepared Statement) để tái sử dụng trong vòng lặp
	stmt, err := tx.Prepare(sqlEditPhanCong)
	if err != nil {
		return 0, fmt.Errorf("không thể chuẩn bị câu lệnh: %w", err)
	}
	defer stmt.Close() // Đóng statement khi xong
	for _, pc := range phancong {
		// Dùng statement đã chuẩn bị để thực thi
		_, err = stmt.Exec(pc.GiaoVienId, pc.LopId, pc.MonHocId,pc.TongTiet, pc.ID)
		if err != nil {
			// Lỗi được gán cho biến err, defer sẽ kích hoạt Rollback()
			return count, fmt.Errorf("không thể sửa phần cống %d (%w)", pc.GiaoVienId, err)
		}
		count++
	}
	// Nếu mọi thứ ổn, commit giao dịch
	err = tx.Commit()
	if err != nil {
		return count, fmt.Errorf("không thể chốt giao dịch: %w",	err)
	}
	return count, nil
}

func (s *SqlTKB) DeletePhanCong(ids []int) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("Không thể bắt đầu giao dịch %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer tx.Rollback()

	var count int
	// Chuẩn bị câu lệnh SQL (Prepared Statement) để tái sử dụng trong vòng lặp
	stmt, err := tx.Prepare(sqlDeletePhanCong)
	if err != nil {
		return 0, fmt.Errorf("Không thể chuẩn bị câu lệnh %w", err)
	}
	defer stmt.Close() // Đóng statement khi xong
	for _, id := range ids {
		// Dùng statement đã chuẩn bị để thực thi
		_, err = stmt.Exec(id)
		if err != nil {
			// Lỗi được gán cho biến err, defer sẽ kích hoạt Rollback()
			return count, fmt.Errorf("Không thể xóa phần cống %d (%w)", id, err)
		}
		count++
	}
	// Nếu mọi thứ ổn, commit giao dịch
	err = tx.Commit()
	if err != nil {
		return count, fmt.Errorf("Không thể chốt giao dịch %w", err)
	}
	return count, nil
}	

// Chi tiết

const (
	sqlSelectChiTiet = `
	SELECT * FROM chitiet
	WHERE id = ?;
	`
	sqlSelectAllChiTiet = `
	SELECT * FROM chitiet;
	`
	sqlFindChiTiet = `
	SELECT * FROM chitiet
	WHERE phan_cong_id = ?;
	`
	sqlInsertChiTiet = `
	INSERT INTO chitiet(phan_cong_id, tuan, so_tiet)
	VALUES (?, ?, ?);
	`
	sqlEditChiTiet = `
	UPDATE chitiet
	SET tuan = ?, so_tiet = ?
	WHERE id = ?;
	`
	sqlDeleteChiTiet = `
	DELETE FROM chitiet
	WHERE id = ?;
	`
	sqlSelectAllChiTietTheoLop=`
		SELECT 
			l.id, l.ten_lop, m.id, 
			COALESCE(pc.id, 0) as phan_cong_id,
			COALESCE(ct.id, 0) as chi_tiet_id,
			COALESCE(ct.so_tiet, 0) as so_tiet
		FROM lophoc l
		CROSS JOIN monhoc m
		LEFT JOIN phancong pc ON pc.lop_id = l.id AND pc.mon_hoc_id = m.id
		LEFT JOIN chitiet ct ON ct.phan_cong_id = pc.id
		WHERE ct.tuan = ?
	`
)

type ChiTiet struct {
	ID int `json:"id"`
	PhanCongId int `json:"phan_cong_id"`
	Tuan int `json:"tuan"`
	Sotiet int `json:"so_tiet"`
	Action string `json:"action,omitempty"`
}
type ChiTietTheoLop struct {
	LopId int `json:"lop_id"`
	TenLop string `json:"ten_lop"`
	KhoiLop string `json:"khoi_lop"`
	MonHoc map[int]*struct {
		PhanCongId int `json:"phan_cong_id"`
		ChiTietId int `json:"chi_tiet_id"`
		SoTiet int `json:"so_tiet"`
	} `json:"mon_hoc"`
}

func (s *SqlTKB) SelectAllChiTietTheoLop(tuan int) ([]ChiTietTheoLop, error) {

	Rows, err := s.db.Query(sqlSelectAllChiTietTheoLop,tuan)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu từ bảng chitiet : %w", err)
	}
	fmt.Println(Rows)
	defer Rows.Close()
	chitiet:=make(map[int]*ChiTietTheoLop)
	for Rows.Next() {
		var lopID, monID, phancongID, chitietID, soTiet int
		var tenLop string
		err := Rows.Scan(&lopID, &tenLop, &monID, &phancongID, &chitietID, &soTiet)
		if err != nil {
			return nil, fmt.Errorf("Không thể quét dữ liệu vào biến chitiet : %w", err)
		}
		if _, ok := chitiet[lopID]; !ok {
			chitiet[lopID] = &ChiTietTheoLop{
				LopId: lopID,
				TenLop: tenLop,
				MonHoc: make(map[int]*struct {
					PhanCongId int `json:"phan_cong_id"`
					ChiTietId int `json:"chi_tiet_id"`
					SoTiet int `json:"so_tiet"`
				}),
			}
		}
		chitiet[lopID].MonHoc[monID] = &struct {
			PhanCongId int `json:"phan_cong_id"`
			ChiTietId int `json:"chi_tiet_id"`
			SoTiet int `json:"so_tiet"`
		}{
			PhanCongId: phancongID,
			ChiTietId: chitietID,
			SoTiet: soTiet,
		}
	}
	var result []ChiTietTheoLop
	for _, v := range chitiet {
		result = append(result, *v)
	}
	return result, nil
}

func (s *SqlTKB) SelectAllChiTiet() ([]ChiTiet, error) {
	var chitiet []ChiTiet
	Rows, err := s.db.Query(sqlSelectAllChiTiet)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu từ bảng phanbotiet : %w", err)
	}
	defer Rows.Close()
	for Rows.Next() {
		var ct ChiTiet
		err := Rows.Scan(&ct.ID, &ct.PhanCongId, &ct.Tuan, &ct.Sotiet)
		if err != nil {
			return nil, fmt.Errorf("Không thể quét dữ liệu vào biến chitiet : %w", err)
		}
		chitiet = append(chitiet, ct)
	}
	return chitiet, nil
}

func (s *SqlTKB) SelectChiTiet(id int) (*ChiTiet, error) {
	var ct ChiTiet
	err := s.db.QueryRow(sqlSelectChiTiet, id).Scan(&ct.ID, &ct.PhanCongId, &ct.Tuan, &ct.Sotiet)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu từ bảng phanbotiet : %w", err)
	}
	return &ct, nil
}

func (s *SqlTKB) InsertChiTiet(chitiet []ChiTiet) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("không thể bắt đầu giao dịch: %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer tx.Rollback()

	var count int
	// Chuẩn bị câu lệnh SQL (Prepared Statement) để tái sử dụng trong vòng lặp
	stmt, err := tx.Prepare(sqlInsertChiTiet)
	if err != nil {
		return 0, fmt.Errorf("không thể chuẩn bị câu lệnh: %w", err)
	}
	defer stmt.Close() // Đóng statement khi xong
	for _, ct := range chitiet {
		// Dùng statement đã chuẩn bị để thực thi
		_, err = stmt.Exec(ct.PhanCongId, ct.Tuan, ct.Sotiet)
		if err != nil {
			// Lỗi được gán cho biến err, defer sẽ kích hoạt Rollback()
			return count, fmt.Errorf("không thể thêm phanbotiet %d (%w)", ct.PhanCongId, err)
		}
		count++
	}
	// Nếu mọi thứ ổn, commit giao dịch
	err = tx.Commit()
	if err != nil {
		return count, fmt.Errorf("không thể chốt giao dịch: %w", err)
	}
	return count, nil
}

func (s *SqlTKB) EditChiTiet(chitiet []ChiTiet) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("không thể bắt đầu giao dịch: %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer tx.Rollback()

	var count int
	// Chuẩn bị câu lệnh SQL (Prepared Statement) để tái sử dụng trong vòng lặp
	stmt, err := tx.Prepare(sqlEditChiTiet)
	if err != nil {
		return 0, fmt.Errorf("không thể chuẩn bị câu lệnh: %w", err)
	}
	defer stmt.Close() // Đóng statement khi xong
	for _, ct := range chitiet {
		// Dùng statement đã chuẩn bị để thực thi
		_, err = stmt.Exec(ct.PhanCongId, ct.Tuan, ct.Sotiet)
		if err != nil {
			// Lỗi được gán cho biến err, defer sẽ kích hoạt Rollback()
			return count, fmt.Errorf("không thể sửa phanbotiet %d (%w)", ct.PhanCongId, err)
		}
		count++
	}
	// Nếu mọi thứ ổn, commit giao dịch
	err = tx.Commit()
	if err != nil {
		return count, fmt.Errorf("không thể chốt giao dịch: %w", err)
	}
	return count, nil
}

func (s *SqlTKB) DeleteChiTiet(ids []int) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("Không thể bắt đầu giao dịch %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer tx.Rollback()

	var count int
	// Chuẩn bị câu lệnh SQL (Prepared Statement) để tái sử dụng trong vòng lặp
	stmt, err := tx.Prepare(sqlDeleteChiTiet)
	if err != nil {
		return 0, fmt.Errorf("Không thể chuẩn bị câu lệnh %w", err)
	}
	defer stmt.Close() // Đóng statement khi xong
	for _, id := range ids {
		// Dùng statement đã chuẩn bị để thực thi
		_, err = stmt.Exec(id)
		if err != nil {
			// Lỗi được gán cho biến err, defer sẽ kích hoạt Rollback()
			return count, fmt.Errorf("Không thể xóa phanbotiet %d (%w)", id, err)
		}
		count++
	}
	// Nếu mọi thứ ổn, commit giao dịch
	err = tx.Commit()
	if err != nil {
		return count, fmt.Errorf("Không thể chốt giao dịch %w", err)
	}
	return count, nil
}

// Tien do

const (
	sqlSelectTiendo = `
	SELECT * FROM tiendo
	WHERE id = ?;
	`
	sqlSelectAllTiendo = `
	SELECT * FROM tiendo;
	`
	sqlFindTiendo = `
	SELECT * FROM tiendo
	WHERE phan_cong_id = ?;
	`
	sqlInsertTiendo = `
	INSERT INTO tiendo(phan_cong_id, tuan, so_tiet)
	VALUES (?, ?, ?);
	`
	sqlEditTiendo = `
	UPDATE tiendo
	SET tuan = ?, so_tiet = ?
	WHERE id = ?;
	`
	sqlDeleteTiendo = `
	DELETE FROM tiendo
	WHERE id = ?;
	`
)

type TienDo struct {
	ID int `json:"id"`
	PhanCongId int `json:"phan_cong_id"`
	Tuan int `json:"tuan"`
	Sotiet int `json:"so_tiet"`
	Action string `json:"action,omitempty"`
}

func (s *SqlTKB) SelectAllTiendo() ([]TienDo, error) {
	var tiendo []TienDo
	Rows, err := s.db.Query(sqlSelectAllTiendo)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu từ bảng tiendo : %w", err)
	}
	defer Rows.Close()
	for Rows.Next() {
		var td TienDo
		err := Rows.Scan(&td.ID, &td.PhanCongId, &td.Tuan, &td.Sotiet)
		if err != nil {
			return nil, fmt.Errorf("Không thể quét dữ liệu vào biến tiendo : %w", err)
		}
		tiendo = append(tiendo, td)
	}
	return tiendo, nil
}

func (s *SqlTKB) SelectTiendo(id int) (*TienDo, error) {
	var td TienDo
	err := s.db.QueryRow(sqlSelectTiendo, id).Scan(&td.ID, &td.PhanCongId, &td.Tuan, &td.Sotiet)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu từ bảng tiendo : %w", err)
	}
	return &td, nil
}

func (s *SqlTKB) InsertTiendo(tiendo []TienDo) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("không thể bắt đầu giao dịch: %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer tx.Rollback()

	var count int
	// Chuẩn bị câu lệnh SQL (Prepared Statement) để tái sử dụng trong vòng lặp
	stmt, err := tx.Prepare(sqlInsertTiendo)
	if err != nil {
		return 0, fmt.Errorf("không thể chuẩn bị câu lệnh: %w", err)
	}
	defer stmt.Close() // Đóng statement khi xong
	for _, td := range tiendo {
		// Dùng statement đã chuẩn bị để thực thi
		_, err = stmt.Exec(td.Tuan, td.Sotiet, td.ID)
		if err != nil {
			// Lỗi được gán cho biến err, defer sẽ kích hoạt Rollback()
			return count, fmt.Errorf("không thể thêm tiendo %d (%w)", td.PhanCongId, err)
		}
		count++
	}
	// Nếu mọi thứ ổn, commit giao dịch
	err = tx.Commit()
	if err != nil {
		return count, fmt.Errorf("không thể chốt giao dịch: %w", err)
	}
	return count, nil
}

func (s *SqlTKB) EditTiendo(tiendo []TienDo) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("không thể bắt đầu giao dịch: %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer tx.Rollback()

	var count int
	// Chuẩn bị câu lệnh SQL (Prepared Statement) để tái sử dụng trong vòng lặp
	stmt, err := tx.Prepare(sqlEditTiendo)
	if err != nil {
		return 0, fmt.Errorf("không thể chuẩn bị câu lệnh: %w", err)
	}
	defer stmt.Close() // Đóng statement khi xong
	for _, td := range tiendo {
		// Dùng statement đã chuẩn bị để thực thi
		_, err = stmt.Exec(td.Tuan, td.Sotiet, td.ID)
		if err != nil {
			// Lỗi được gán cho biến err, defer sẽ kích hoạt Rollback()
			return count, fmt.Errorf("không thể sửa tiendo %d (%w)", td.PhanCongId, err)
		}
		count++
	}
	// Nếu mọi thứ ổn, commit giao dịch
	err = tx.Commit()
	if err != nil {
		return count, fmt.Errorf("không thể chốt giao dịch: %w", err)
	}
	return count, nil
}

func (s *SqlTKB) DeleteTiendo(ids []int) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("Không thể bắt đầu giao dịch %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer tx.Rollback()

	var count int
	// Chuẩn bị câu lệnh SQL (Prepared Statement) để tái sử dụng trong vòng lặp
	stmt, err := tx.Prepare(sqlDeleteTiendo)
	if err != nil {
		return 0, fmt.Errorf("Không thể chuẩn bị câu lệnh %w", err)
	}	
	defer stmt.Close() // Đóng statement khi xong
	for _, id := range ids {
		// Dùng statement đã chuẩn bị để thực thi
		_, err = stmt.Exec(id)
		if err != nil {
			// Lỗi được gán cho biến err, defer sẽ kích hoạt Rollback()
			return count, fmt.Errorf("Không thể xóa tiendo %d (%w)", id, err)
		}
		count++
	}
	// Nếu mọi thứ ổn, commit giao dịch
	err = tx.Commit()
	if err != nil {
		return count, fmt.Errorf("Không thể chốt giao dịch %w", err)
	}
	return count, nil
}

// Rang bước

const (
	sqlSelectRangBuoc = `
	SELECT * FROM rangbuoc
	WHERE id = ?;
	`
	sqlSelectAllRangBuoc = `
	SELECT * FROM rangbuoc;
	`
	sqlFindRangBuoc = `
	SELECT * FROM rangbuoc
	WHERE giao_vien_id = ?;
	`
	sqlInsertRangBuoc = `
	INSERT INTO rangbuoc(giao_vien_id, thu, tiet, loai_rang_buoc)
	VALUES (?, ?, ?, ?);
	`
	sqlEditRangBuoc = `
	UPDATE rangbuoc
	SET thu = ?, tiet = ?, loai_rang_buoc = ?
	WHERE id = ?;
	`
	sqlDeleteRangBuoc = `
	DELETE FROM rangbuoc
	WHERE id = ?;
	`
)

type RangBuoc struct {
	ID int `json:"id"`
	GiaoVienId int `json:"giao_vien_id"`
	Thu int `json:"thu"`
	Tiet int `json:"tiet"`
	LoaiRangBuoc string `json:"loai_rang_buoc"`
	Action string `json:"action,omitempty"`
}

func (s *SqlTKB) SelectAllRangBuoc() ([]RangBuoc, error) {
	var rangbuoc []RangBuoc
	Rows, err := s.db.Query(sqlSelectAllRangBuoc)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu từ bảng rangbuoc : %w", err)
	}
	defer Rows.Close()
	for Rows.Next() {
		var rb RangBuoc
		err := Rows.Scan(&rb.ID, &rb.GiaoVienId, &rb.Thu, &rb.Tiet, &rb.LoaiRangBuoc)
		if err != nil {
			return nil, fmt.Errorf("Không thể quét dữ liệu vào biến rangbuoc : %w", err)
		}
		rangbuoc = append(rangbuoc, rb)
	}
	return rangbuoc, nil
}	

func (s *SqlTKB) SelectRangBuoc(id int) (*RangBuoc, error) {
	var rb RangBuoc
	err := s.db.QueryRow(sqlSelectRangBuoc, id).Scan(&rb.ID, &rb.GiaoVienId, &rb.Thu, &rb.Tiet, &rb.LoaiRangBuoc)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu từ bảng rangbuoc : %w", err)
	}
	return &rb, nil
}

func (s *SqlTKB) InsertRangBuoc(rangbuoc []RangBuoc) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("không thể bắt đầu giao dịch: %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer tx.Rollback()

	var count int
	// Chuẩn bị câu lệnh SQL (Prepared Statement) để tái sử dụng trong vòng lặp
	stmt, err := tx.Prepare(sqlInsertRangBuoc)
	if err != nil {
		return 0, fmt.Errorf("không thể chuẩn bị câu lệnh: %w", err)
	}
	defer stmt.Close() // Đóng statement khi xong
	for _, rb := range rangbuoc {
		// Dùng statement đã chuẩn bị để thực thi
		_, err = stmt.Exec(rb.GiaoVienId, rb.Thu, rb.Tiet, rb.LoaiRangBuoc)
		if err != nil {
			// Lỗi được gán cho biến err, defer sẽ kích hoạt Rollback()
			return count, fmt.Errorf("không thể thêm rang bước %d (%w)", rb.GiaoVienId, err)
		}
		count++
	}
	// Nếu mọi thứ ổn, commit giao dịch
	err = tx.Commit()
	if err != nil {
		return count, fmt.Errorf("không thể chốt giao dịch: %w", err)
	}
	return count, nil
}

func (s *SqlTKB) EditRangBuoc(rangbuoc []RangBuoc) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("không thể bắt đầu giao dịch: %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer tx.Rollback()

	var count int
	// Chuẩn bị câu lệnh SQL (Prepared Statement) để tái sử dụng trong vòng lặp
	stmt, err := tx.Prepare(sqlEditRangBuoc)
	if err != nil {
		return 0, fmt.Errorf("không thể chuẩn bị câu lệnh: %w", err)
	}
	defer stmt.Close() // Đóng statement khi xong
	for _, rb := range rangbuoc {
		// Dùng statement đã chuẩn bị để thực thi
		_, err = stmt.Exec(rb.Thu, rb.Tiet, rb.LoaiRangBuoc, rb.ID)
		if err != nil {
			// Lỗi được gán cho biến err, defer sẽ kích hoạt Rollback()
			return count, fmt.Errorf("không thể sửa rang bước %d (%w)", rb.GiaoVienId, err)
		}
		count++
	}
	// Nếu mọi thứ ổn, commit giao dịch
	err = tx.Commit()
	if err != nil {
		return count, fmt.Errorf("không thể chốt giao dịch: %w",	err)
	}
	return count, nil
}

func (s *SqlTKB) DeleteRangBuoc(ids []int) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("Không thể bắt đầu giao dịch %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer tx.Rollback()

	var count int
	// Chuẩn bị câu lệnh SQL (Prepared Statement) để tái sử dụng trong vòng lặp
	stmt, err := tx.Prepare(sqlDeleteRangBuoc)
	if err != nil {
		return 0, fmt.Errorf("Không thể chuẩn bị câu lệnh %w", err)
	}
	defer stmt.Close() // Đóng statement khi xong
	for _, id := range ids {
		// Dùng statement đã chuẩn bị để thực thi
		_, err = stmt.Exec(id)
		if err != nil {
			// Lỗi được gán cho biến err, defer sẽ kích hoạt Rollback()
			return count, fmt.Errorf("Không thể xóa rang bước %d (%w)", id, err)
		}
		count++
	}
	// Nếu mọi thứ ổn, commit giao dịch
	err = tx.Commit()
	if err != nil {
		return count, fmt.Errorf("Không thể chốt giao dịch %w", err)
	}
	return count, nil
}


