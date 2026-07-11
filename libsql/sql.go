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
			FOREIGN KEY (giao_vien_id) REFERENCES giaovien(id),
			FOREIGN KEY (lop_id) REFERENCES lophoc(id),
			FOREIGN KEY (mon_hoc_id) REFERENCES monhoc(id)
		);
		CREATE TABLE IF NOT EXISTS phanbotiet(
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
			ROREIGN KEY (phan_cong_id) REFERENCES phancong(id)
		);
		CREATE TABLE IF NOT EXISTS rangbuoc(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			giao_vien_id INTEGER,
			thu INTEGER,
			tiet INTEGER,
			loai_rang_buoc TEXT,
			FOREIGN KEY (giao_vien_id) REFERENCES giaovien(id)
		);`
		sqlInsertGiaoVien = `
		INSERT INTO giaovien(ten_ngan, ho_ten, mon_chinh_id)
		VALUES (?, ?, ?);
		`
		sqlInsertMonHoc = `
		INSERT INTO monhoc(ten_mon, loai_mon)
		VALUES (?, ?);
		`
		sqlInsertLopHoc = `
		INSERT INTO lophoc(ten_lop, khoi_lop)
		VALUES (?, ?);
		`
		sqlInsertPhanCong = `
		INSERT INTO phancong(giao_vien_id, lop_id, mon_hoc_id, tuan, so_tiet)
		VALUES (?, ?, ?, ?, ?);
		`
		sqlInsertPhanbotiet = `
		INSERT INTO phanbotiet(phan_cong_id, tuan, so_tiet)
		VALUES (?, ?, ?);
		`
		sqlInsertTiendo = `
		INSERT INTO tiendo(phan_cong_id, tuan, so_tiet)
		VALUES (?, ?, ?);
		`
		sqlInsertRangBuoc = `
		INSERT INTO rangbuoc(giao_vien_id, thu, tiet, loai_rang_buoc)
		VALUES (?, ?, ?, ?);
		`
		sqlEditGiaoVien = `
		UPDATE giaovien
		SET ten_ngan = ?, ho_ten = ?, mon_chinh_id = ?
		WHERE id = ?;
		`
		sqlEditMonHoc = `
		UPDATE monhoc
		SET ten_mon = ?, loai_mon = ?
		WHERE id = ?;
		`
		sqlEditLopHoc = `
		UPDATE lophoc
		SET ten_lop = ?, khoi_lop = ?
		WHERE id = ?;
		`
		sqlEditPhanCong = `
		UPDATE phancong
		SET tuan = ?, so_tiet = ?
		WHERE id = ?;
		`
		sqlEditPhanbotiet = `
		UPDATE phanbotiet
		SET tuan = ?, so_tiet = ?
		WHERE id = ?;
		`
		sqlEditTiendo = `
		UPDATE tiendo
		SET tuan = ?, so_tiet = ?
		WHERE id = ?;
		`
		sqlEditRangBuoc = `
		UPDATE rangbuoc
		SET thu = ?, tiet = ?, loai_rang_buoc = ?
		WHERE id = ?;
		`
		sqlDeleteGiaoVien = `
		DELETE FROM giaovien
		WHERE id = ?;
		`
		sqlDeleteMonHoc = `
		DELETE FROM monhoc
		WHERE id = ?;
		`
		sqlDeleteLopHoc = `
		DELETE FROM lophoc
		WHERE id = ?;
		`
		sqlDeletePhanCong = `		
		DELETE FROM phancong
		WHERE id = ?;
		`
		sqlDeletePhanbotiet = `
		DELETE FROM phanbotiet
		WHERE id = ?;
		`
		sqlDeleteTiendo = `
		DELETE FROM tiendo
		WHERE id = ?;
		`
		sqlDeleteRangBuoc = `
		DELETE FROM rangbuoc
		WHERE id = ?;
		`
		sqlSelectGiaoVien = `
		SELECT * FROM giaovien
		WHERE id = ?;
		`
		sqlSelectMonHoc = `
		SELECT * FROM monhoc
		WHERE id = ?;
		`
		sqlSelectLopHoc = `
		SELECT * FROM lophoc
		WHERE id = ?;
		`
		sqlSelectPhanCong = `
		SELECT * FROM phancong
		WHERE id = ?;
		`
		sqlSelectPhanbotiet = `
		SELECT * FROM phanbotiet
		WHERE id = ?;
		`
		sqlSelectTiendo = `
		SELECT * FROM tiendo
		WHERE id = ?;
		`
		sqlSelectRangBuoc = `
		SELECT * FROM rangbuoc
		WHERE id = ?;
		`
		sqlSelectAllGiaoVien = `
		SELECT * FROM giaovien;
		`
		sqlSelectAllMonHoc = `
		SELECT * FROM monhoc;
		`
		sqlSelectAllLopHoc = `
		SELECT * FROM lophoc;
		`
		sqlSelectAllPhanCong = `
		SELECT * FROM phancong;
		`
		sqlSelectAllPhanbotiet = `
		SELECT * FROM phanbotiet;
		`
		sqlSelectAllTiendo = `
		SELECT * FROM tiendo;
		`
		sqlSelectAllRangBuoc = `
		SELECT * FROM rangbuoc;
		`
		sqlFindGiaoVien = `
		SELECT * FROM giaovien
		WHERE ten_ngan = ?;
		`
		sqlFindMonHoc = `
		SELECT * FROM monhoc
		WHERE ten_mon = ?;
		`
		sqlFindLopHoc = `
		SELECT * FROM lophoc
		WHERE ten_lop = ?;
		`
		sqlFindPhanCong = `
		SELECT * FROM phancong
		WHERE giao_vien_id = ?;
		`
		sqlFindPhanbotiet = `
		SELECT * FROM phanbotiet
		WHERE phan_cong_id = ?;
		`
		sqlFindTiendo = `
		SELECT * FROM tiendo
		WHERE phan_cong_id = ?;
		`
		sqlFindRangBuoc = `
		SELECT * FROM rangbuoc
		WHERE giao_vien_id = ?;
		`
	)

// define struct Giaovien, MonHoc, LopHoc, PhanCong, RangBuoc, ThuaThieu

type GiaoVien struct {
	ID int `json:"id"`
	TenNgan string `json:"ten_ngan"`
	HoTen string `json:"ho_ten"`
	MonChinhId int `json:"mon_chinh_id"`
	Action string `json:"action,omitempty"`
	
}
type MonHoc struct {
	ID int `json:"id"`
	TenMon string `json:"ten_mon"`
	LoaiMon string `json:"loai_mon"`
	Action string `json:"action,omitempty"`
}
type LopHoc struct {
	ID int `json:"id"`
	TenLop string `json:"ten_lop"`
	KhoiLop string `json:"khoi_lop"`
	Action string `json:"action,omitempty"`
}
type PhanCong struct {
	ID int `json:"id"`
	GiaoVienId int `json:"giao_vien_id"`
	LopId int `json:"lop_id"`
	MonHocId int `json:"mon_hoc_id"`
	Tuan int `json:"tuan"`
	Sotiet int `json:"so_tiet"`
	Action string `json:"action,omitempty"`
}
type PhanBotiet struct {
	ID int `json:"id"`
	PhanCongId int `json:"phan_cong_id"`
	Tuan int `json:"tuan"`
	Sotiet int `json:"so_tiet"`
	Action string `json:"action,omitempty"`
}
type Tiendo struct {
	ID int `json:"id"`
	PhanCongId int `json:"phan_cong_id"`
	Tuan int `json:"tuan"`
	Sotiet int `json:"so_tiet"`
	Action string `json:"action,omitempty"`
}
type RangBuoc struct {
	ID int `json:"id"`
	GiaoVienId int `json:"giao_vien_id"`
	Thu int `json:"thu"`
	Tiet int `json:"tiet"`
	LoaiRangBuoc string `json:"loai_rang_buoc"`
	Action string `json:"action,omitempty"`
}

type SqlTKB struct {
	db *sql.DB
}

// func InsertGiaoVien, InsertMonHoc, InsertLopHoc, InsertPhanCong, InsertRangBuoc, InsertThuaThieu
// return count, error

func (s *SqlTKB) InsertGiaoVien(giaovien []GiaoVien) (int, error) {
	// Bắt đầu một transaction
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("không thể bắt đầu giao dịch: %w", err)
	}

	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // Bắn lại panic sau khi rollback
		} else if err != nil {
			tx.Rollback() // Rollback nếu err != nil
		}
	}()

	var count int
	// Chuẩn bị câu lệnh SQL (Prepared Statement) để tái sử dụng trong vòng lặp
	stmt, err := tx.Prepare(sqlInsertGiaoVien)
	if err != nil {
		return 0, fmt.Errorf("không thể chuẩn bị câu lệnh: %w", err)
	}
	defer stmt.Close() // Đóng statement khi xong

	for _, gv := range giaovien {
		// Dùng statement đã chuẩn bị để thực thi
		_, err = stmt.Exec(gv.TenNgan, gv.HoTen, gv.MonChinhId)
		if err != nil {
			// Lỗi được gán cho biến err, defer sẽ kích hoạt Rollback()
			return count, fmt.Errorf("không thể thêm giáo viên %s (%w)", gv.TenNgan, err)
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
func (s *SqlTKB) InsertMonHoc(monhoc []MonHoc) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("không thể bắt đầu giao dịch: %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // Bắn lại panic sau khi rollback
		} else if err != nil {
			tx.Rollback() // Rollback nếu err != nil
			return // Đừng gọi return khi err != nil
		}
	}()
	var count int
	// Chuẩn bị câu lệnh SQL (Prepared Statement) để tái sử dụng trong vòng lặp
	stmt, err := tx.Prepare(sqlInsertMonHoc)
	if err != nil {
		return 0, fmt.Errorf("không thể chuẩn bị câu lệnh: %w", err)
	}
	defer stmt.Close() // Đóng statement khi xong
	for _, mh := range monhoc {
		// Dùng statement đã chuẩn bị để thực thi
		_, err = stmt.Exec(mh.TenMon, mh.LoaiMon)
		if err != nil {
			// Lỗi được gán cho biến err, defer sẽ kích hoạt Rollback()
			return count, fmt.Errorf("không thể thêm môn học %s (%w)", mh.TenMon, err)
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

func (s *SqlTKB) InsertLopHoc(lophoc []LopHoc) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("không thể bắt đầu giao dịch: %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // Bắn lại panic sau khi rollback
		} else if err != nil {
			tx.Rollback() // Rollback nếu err != nil
			return // Đừng gọi return khi err != nil
		}
	}()
	var count int
	// Chuẩn bị câu lệnh SQL (Prepared Statement) để tái sử dụng trong vòng lặp
	stmt, err := tx.Prepare(sqlInsertLopHoc)
	if err != nil {
		return 0, fmt.Errorf("không thể chuẩn bị câu lệnh: %w", err)
	}
	defer stmt.Close() // Đóng statement khi xong
	for _, lh := range lophoc {
		// Dùng statement đã chuẩn bị để thực thi
		_, err = stmt.Exec(lh.TenLop, lh.KhoiLop)
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

func (s *SqlTKB) InsertPhanCong(phancong []PhanCong) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("không thể bắt đầu giao dịch: %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // Bắn lại panic sau khi rollback
		} else if err != nil {
			tx.Rollback() // Rollback nếu err != nil
			return // Đừng gọi return khi err != nil
		}
	}()
	var count int
	// Chuẩn bị câu lệnh SQL (Prepared Statement) để tái sử dụng trong vòng lặp
	stmt, err := tx.Prepare(sqlInsertPhanCong)
	if err != nil {
		return 0, fmt.Errorf("không thể chuẩn bị câu lệnh: %w", err)
	}
	defer stmt.Close() // Đóng statement khi xong
	for _, pc := range phancong {
		// Dùng statement đã chuẩn bị để thực thi
		_, err = stmt.Exec(pc.GiaoVienId, pc.LopId, pc.MonHocId, pc.Tuan, pc.Sotiet)
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

func (s *SqlTKB) InsertPhanbotiet(phanbotiet []PhanBotiet) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("không thể bắt đầu giao dịch: %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // Bắn lại panic sau khi rollback
		} else if err != nil {
			tx.Rollback() // Rollback nếu err != nil
			return // Đừng gọi return khi err != nil
		}
	}()
	var count int
	// Chuẩn bị câu lệnh SQL (Prepared Statement) để tái sử dụng trong vòng lặp
	stmt, err := tx.Prepare(sqlInsertPhanbotiet)
	if err != nil {
		return 0, fmt.Errorf("không thể chuẩn bị câu lệnh: %w", err)
	}
	defer stmt.Close() // Đóng statement khi xong
	for _, pb := range phanbotiet {
		// Dùng statement đã chuẩn bị để thực thi
		_, err = stmt.Exec(pb.Tuan, pb.Sotiet, pb.ID)
		if err != nil {
			// Lỗi được gán cho biến err, defer sẽ kích hoạt Rollback()
			return count, fmt.Errorf("không thể thêm phanbotiet %d (%w)", pb.PhanCongId, err)
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

func (s *SqlTKB) InsertTiendo(tiendo []Tiendo) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("không thể bắt đầu giao dịch: %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // Bắn lại panic sau khi rollback
		} else if err != nil {
			tx.Rollback() // Rollback nếu err != nil
			return // Đừng gọi return khi err != nil
		}
	}()
	var count int
	// Chuẩn bị câu lệnh SQL (Prepared Statement) để tái sử dụng trong vòng lặp
	stmt, err := tx.Prepare(sqlInsertTiendo)
	if err != nil {
		return 0, fmt.Errorf("không thể chuẩn bị câu lệnh: %w", err)
	}
	defer stmt.Close() // Đóng statement khi xong
	for _, tt := range tiendo {
		// Dùng statement đã chuẩn bị để thực thi
		_, err = stmt.Exec(tt.Tuan, tt.Sotiet, tt.ID)
		if err != nil {
			// Lỗi được gán cho biến err, defer sẽ kích hoạt Rollback()
			return count, fmt.Errorf("không thể thêm tiendo %d (%w)", tt.PhanCongId, err)
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

func (s *SqlTKB) InsertRangBuoc(rangbuoc []RangBuoc) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("không thể bắt đầu giao dịch: %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // Bắn lại panic sau khi rollback
		} else if err != nil {
			tx.Rollback() // Rollback nếu err != nil
			return // Đừng gọi return khi err != nil
		}
	}()
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


// func EditGiaoVien, EditMonHoc, EditLopHoc, EditPhanCong, EditRangBuoc, EditThuaThieu, EditGiaoVien
// return error

func (s *SqlTKB) EditGiaoVien(giaovien []GiaoVien) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("không thể bắt đầu giao dịch: %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // Bắn lại panic sau khi rollback
		} else if err != nil {
			tx.Rollback() // Rollback nếu err != nil
			return // Đừng gọi return khi err != nil
		}
	}()
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

func (s *SqlTKB) EditMonHoc(monhoc []MonHoc) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("không thể bắt đầu giao dịch: %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // Bắn lại panic sau khi rollback
		} else if err != nil {
			tx.Rollback() // Rollback nếu err != nil
			return // Đừng gọi return khi err != nil
		}
	}()
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

func (s *SqlTKB) EditLopHoc(lophoc []LopHoc) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("không thể bắt đầu giao dịch: %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // Bắn lại panic sau khi rollback
		} else if err != nil {
			tx.Rollback() // Rollback nếu err != nil
			return // Đừng gọi return khi err != nil
		}
	}()
	var count int
	// Chuẩn bị câu lệnh SQL (Prepared Statement) để tái sử dụng trong vòng lặp
	stmt, err := tx.Prepare(sqlEditLopHoc)
	if err != nil {
		return 0, fmt.Errorf("không thể chuẩn bị câu lệnh: %w", err)
	}
	defer stmt.Close() // Đóng statement khi xong
	for _, lh := range lophoc {
		// Dùng statement đã chuẩn bị để thực thi
		_, err = stmt.Exec(lh.TenLop, lh.KhoiLop, lh.ID)
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

func (s *SqlTKB) EditPhanCong(phancong []PhanCong) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("không thể bắt đầu giao dịch: %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // Bắn lại panic sau khi rollback
		} else if err != nil {
			tx.Rollback() // Rollback nếu err != nil
			return // Đừng gọi return khi err != nil
		}
	}()
	var count int
	// Chuẩn bị câu lệnh SQL (Prepared Statement) để tái sử dụng trong vòng lặp
	stmt, err := tx.Prepare(sqlEditPhanCong)
	if err != nil {
		return 0, fmt.Errorf("không thể chuẩn bị câu lệnh: %w", err)
	}
	defer stmt.Close() // Đóng statement khi xong
	for _, pc := range phancong {
		// Dùng statement đã chuẩn bị để thực thi
		_, err = stmt.Exec(pc.Tuan, pc.Sotiet, pc.ID)
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

func (s *SqlTKB) EditPhanbotiet(phanbotiet []PhanBotiet) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("không thể bắt đầu giao dịch: %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // Bắn lại panic sau khi rollback
		} else if err != nil {
			tx.Rollback() // Rollback nếu err != nil
			return // Đừng gọi return khi err != nil
		}
	}()
	var count int
	// Chuẩn bị câu lệnh SQL (Prepared Statement) để tái sử dụng trong vòng lặp
	stmt, err := tx.Prepare(sqlEditPhanbotiet)
	if err != nil {
		return 0, fmt.Errorf("không thể chuẩn bị câu lệnh: %w", err)
	}
	defer stmt.Close() // Đóng statement khi xong
	for _, pb := range phanbotiet {
		// Dùng statement đã chuẩn bị để thực thi
		_, err = stmt.Exec(pb.Tuan, pb.Sotiet, pb.ID)
		if err != nil {
			// Lỗi được gán cho biến err, defer sẽ kích hoạt Rollback()
			return count, fmt.Errorf("không thể sửa phanbotiet %d (%w)", pb.PhanCongId, err)
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

func (s *SqlTKB) EditTiendo(tiendo []Tiendo) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("không thể bắt đầu giao dịch: %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // Bắn lại panic sau khi rollback
		} else if err != nil {
			tx.Rollback() // Rollback nếu err != nil
			return // Đừng gọi return khi err != nil
		}
	}()
	var count int
	// Chuẩn bị câu lệnh SQL (Prepared Statement) để tái sử dụng trong vòng lặp
	stmt, err := tx.Prepare(sqlEditTiendo)
	if err != nil {
		return 0, fmt.Errorf("không thể chuẩn bị câu lệnh: %w", err)
	}
	defer stmt.Close() // Đóng statement khi xong
	for _, tt := range tiendo {
		// Dùng statement đã chuẩn bị để thực thi
		_, err = stmt.Exec(tt.Tuan, tt.Sotiet, tt.ID)
		if err != nil {
			// Lỗi được gán cho biến err, defer sẽ kích hoạt Rollback()
			return count, fmt.Errorf("không thể sửa tiendo %d (%w)", tt.PhanCongId, err)
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
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // Bắn lại panic sau khi rollback
		} else if err != nil {
			tx.Rollback() // Rollback nếu err != nil
			return // Đừng gọi return khi err != nil
		}
	}()
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


// func DeleteGiaoVien, DeleteMonHoc, DeleteLopHoc, DeletePhanCong, DeleteRangBuoc, DeleteThuaThieu
// return error

func (s *SqlTKB) DeleteGiaoVien(ids []int) (int, error) {
	tx,err:=s.db.Begin()
	if err!=nil{
		return 0, fmt.Errorf("Không thể bắt đầu giao dịch %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // Bắn lại panic sau khi rollback
		} else if err != nil {
			tx.Rollback() // Rollback nếu err != nil
			return // Đừng gọi return khi err != nil
		}
	}()
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

func (s *SqlTKB) DeleteMonHoc(ids []int) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("Không thể bắt đầu giao dịch %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // Bắn lại panic sau khi rollback
		} else if err != nil {
			tx.Rollback() // Rollback nếu err != nil
			return // Đừng gọi return khi err != nil
		}
	}()
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

func (s *SqlTKB) DeleteLopHoc(ids []int) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("Không thể bắt đầu giao dịch %w", err)
	}
		// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // Bắn lại panic sau khi rollback
		} else if err != nil {
			tx.Rollback() // Rollback nếu err != nil
			return // Đừng gọi return khi err != nil
		}
	}()
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

func (s *SqlTKB) DeletePhanCong(ids []int) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("Không thể bắt đầu giao dịch %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // Bắn lại panic sau khi rollback
		} else if err != nil {
			tx.Rollback() // Rollback nếu err != nil
			return // Đừng gọi return khi err != nil
		}
	}()
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

func (s *SqlTKB) DeletePhanbotiet(ids []int) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("Không thể bắt đầu giao dịch %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // Bắn lại panic sau khi rollback
		} else if err != nil {
			tx.Rollback() // Rollback nếu err != nil
			return // Đừng gọi return khi err != nil
		}
	}()
	var count int
	// Chuẩn bị câu lệnh SQL (Prepared Statement) để tái sử dụng trong vòng lặp
	stmt, err := tx.Prepare(sqlDeletePhanbotiet)
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

func (s *SqlTKB) DeleteTiendo(ids []int) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("Không thể bắt đầu giao dịch %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // Bắn lại panic sau khi rollback
		} else if err != nil {
			tx.Rollback() // Rollback nếu err != nil
			return // Đừng gọi return khi err != nil
		}
	}()
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

func (s *SqlTKB) DeleteRangBuoc(ids []int) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("Không thể bắt đầu giao dịch %w", err)
	}
	// Đảm bảo rollback nếu có lỗi xảy ra (tránh treo CSDL)
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // Bắn lại panic sau khi rollback
		} else if err != nil {
			tx.Rollback() // Rollback nếu err != nil
			return // Đừng gọi return khi err != nil
		}
	}()
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


// func SelectGiaoVien, SelectMonHoc, SelectLopHoc, SelectPhanCong, SelectRangBuoc, SelectThuaThieu
// return data, error

func (s *SqlTKB) SelectGiaoVien(id int) (*GiaoVien, error) {
	var gv GiaoVien
	err := s.db.QueryRow(sqlSelectGiaoVien, id).Scan(&gv.ID, &gv.TenNgan, &gv.HoTen, &gv.MonChinhId)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu từ bảng giaovien : %w", err)
	}
	return &gv, nil
}

func (s *SqlTKB) SelectMonHoc(id int) (*MonHoc, error) {
	var mh MonHoc
	err := s.db.QueryRow(sqlSelectMonHoc, id).Scan(&mh.ID, &mh.TenMon, &mh.LoaiMon)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu từ bảng monhoc : %w", err)
	}
	return &mh, nil
}

func (s *SqlTKB) SelectLopHoc(id int) (*LopHoc, error) {
	var lh LopHoc
	err := s.db.QueryRow(sqlSelectLopHoc, id).Scan(&lh.ID, &lh.TenLop, &lh.KhoiLop)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu từ bảng lophoc : %w", err)
	}
	return &lh, nil
}

func (s *SqlTKB) SelectPhanCong(id int) (*PhanCong, error) {
	var pc PhanCong
	err := s.db.QueryRow(sqlSelectPhanCong, id).Scan(&pc.ID, &pc.GiaoVienId, &pc.LopId, &pc.MonHocId, &pc.Tuan, &pc.Sotiet)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu từ bảng phancong : %w", err)
	}
	return &pc, nil
}

func (s *SqlTKB) SelectPhanbotiet(id int) (*PhanBotiet, error) {
	var pb PhanBotiet
	err := s.db.QueryRow(sqlSelectPhanbotiet, id).Scan(&pb.ID, &pb.PhanCongId, &pb.Tuan, &pb.Sotiet)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu từ bảng phanbotiet : %w", err)
	}
	return &pb, nil
}

func (s *SqlTKB) SelectTiendo(id int) (*Tiendo, error) {
	var td Tiendo
	err := s.db.QueryRow(sqlSelectTiendo, id).Scan(&td.ID, &td.PhanCongId, &td.Tuan, &td.Sotiet)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu từ bảng tiendo : %w", err)
	}
	return &td, nil
}

func (s *SqlTKB) SelectRangBuoc(id int) (*RangBuoc, error) {
	var rb RangBuoc
	err := s.db.QueryRow(sqlSelectRangBuoc, id).Scan(&rb.ID, &rb.GiaoVienId, &rb.Thu, &rb.Tiet, &rb.LoaiRangBuoc)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu từ bảng rangbuoc : %w", err)
	}
	return &rb, nil
}

// func SelectAllGiaoVien, SelectAllMonHoc, SelectAllLopHoc, SelectAllPhanCong, SelectAllRangBuoc, SelectAllThuaThieu
// return data, error

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

func (s *SqlTKB) SelectAllLopHoc() ([]LopHoc, error) {
	var lophoc []LopHoc
	Rows, err := s.db.Query(sqlSelectAllLopHoc)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu từ bảng lophoc : %w", err)
	}
	defer Rows.Close()
	for Rows.Next() {
		var lh LopHoc
		err := Rows.Scan(&lh.ID, &lh.TenLop, &lh.KhoiLop)
		if err != nil {
			return nil, fmt.Errorf("Không thể quét dữ liệu vào biến lophoc : %w", err)
		}
		lophoc = append(lophoc, lh)
	}
	return lophoc, nil
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
		err := Rows.Scan(&pc.ID, &pc.GiaoVienId, &pc.LopId, &pc.MonHocId, &pc.Tuan, &pc.Sotiet)
		if err != nil {
			return nil, fmt.Errorf("Không thể quét dữ liệu vào biến phancong : %w", err)
		}
		phancong = append(phancong, pc)
	}
	return phancong, nil
}

func (s *SqlTKB) SelectAllPhanbotiet() ([]PhanBotiet, error) {
	var phanbotiet []PhanBotiet
	Rows, err := s.db.Query(sqlSelectAllPhanbotiet)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu từ bảng phanbotiet : %w", err)
	}
	defer Rows.Close()
	for Rows.Next() {
		var pb PhanBotiet
		err := Rows.Scan(&pb.ID, &pb.PhanCongId, &pb.Tuan, &pb.Sotiet)
		if err != nil {
			return nil, fmt.Errorf("Không thể quét dữ liệu vào biến phanbotiet : %w", err)
		}
		phanbotiet = append(phanbotiet, pb)
	}
	return phanbotiet, nil
}

func (s *SqlTKB) SelectAllTiendo() ([]Tiendo, error) {
	var tiendo []Tiendo
	Rows, err := s.db.Query(sqlSelectAllTiendo)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu từ bảng tiendo : %w", err)
	}
	defer Rows.Close()
	for Rows.Next() {
		var td Tiendo
		err := Rows.Scan(&td.ID, &td.PhanCongId, &td.Tuan, &td.Sotiet)
		if err != nil {
			return nil, fmt.Errorf("Không thể quét dữ liệu vào biến tiendo : %w", err)
		}
		tiendo = append(tiendo, td)
	}
	return tiendo, nil
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

// func FindGiaoVien, FindMonHoc, FindLopHoc, FindPhanCong, FindRangBuoc, FindThuaThieu
// return data, error	


func (s *SqlTKB) FindGiaoVien(ten_ngan string) (*GiaoVien, error) {
	var gv GiaoVien
	err := s.db.QueryRow(sqlFindGiaoVien, ten_ngan).Scan(&gv.ID, &gv.TenNgan, &gv.HoTen, &gv.MonChinhId)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu từ bảng giaovien : %w", err)	
	}
	return &gv, nil
}

func (s *SqlTKB) FindMonHoc(ten_mon string) (*MonHoc, error) {
	var mh MonHoc
	err := s.db.QueryRow(sqlFindMonHoc, ten_mon).Scan(&mh.ID, &mh.TenMon, &mh.LoaiMon)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu từ bảng monhoc : %w", err)	
	}
	return &mh, nil
}

func (s *SqlTKB) FindLopHoc(ten_lop string) (*LopHoc, error) {
	var lh LopHoc
	err := s.db.QueryRow(sqlFindLopHoc, ten_lop).Scan(&lh.ID, &lh.TenLop, &lh.KhoiLop)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu từ bảng lophoc : %w", err)
	}	
	return &lh, nil
}

func (s *SqlTKB) FindPhanCong(giao_vien_id int) (*PhanCong, error) {	
	var pc PhanCong
	err := s.db.QueryRow(sqlFindPhanCong, giao_vien_id).Scan(&pc.ID, &pc.GiaoVienId, &pc.LopId, &pc.MonHocId, &pc.Tuan, &pc.Sotiet)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu từ bảng phancong : %w", err)
	}
	return &pc, nil
}

func (s *SqlTKB) FindPhanbotiet(phan_cong_id int) (*PhanBotiet, error) {	
	var pb PhanBotiet
	err := s.db.QueryRow(sqlFindPhanbotiet, phan_cong_id).Scan(&pb.ID, &pb.PhanCongId, &pb.Tuan, &pb.Sotiet)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu từ bảng phanbotiet : %w", err)
	}
	return &pb, nil
}

func (s *SqlTKB) FindTiendo(phan_cong_id int) (*Tiendo, error) {	
	var td Tiendo
	err := s.db.QueryRow(sqlFindTiendo, phan_cong_id).Scan(&td.ID, &td.PhanCongId, &td.Tuan, &td.Sotiet)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu từ bảng tiendo : %w", err)
	}
	return &td, nil
}

func (s *SqlTKB) FindRangBuoc(giao_vien_id int) (*RangBuoc, error) {	
	var rb RangBuoc
	err := s.db.QueryRow(sqlFindRangBuoc, giao_vien_id).Scan(&rb.ID, &rb.GiaoVienId, &rb.Thu, &rb.Tiet, &rb.LoaiRangBuoc)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu từ bảng rangbuoc : %w", err)
	}
	return &rb, nil
}

// func Close
// return error

func (s *SqlTKB) Close() error {
	return s.db.Close()
}

// func ConnectSTKB
// return SqlTKB, error

func ConnectSTKB() (*SqlTKB, error) {
	// Kết nối SQLite
	db, err := sql.Open("sqlite3", sqlpath)
	if err != nil {
		return nil, fmt.Errorf("không kết nối được với STKB : %w", err)
	}
	_, err = db.Exec(sqlCreateTables)
	if err != nil {
		return nil, fmt.Errorf("không thể tạo bảng : %w", err)
	}
	return &SqlTKB{db}, nil
}

