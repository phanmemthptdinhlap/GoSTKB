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
			ten_ngan TEXT,
			ho_ten TEXT,
			mon_chinh_id INTEGER,
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
			giao_vien_id INTEGER,
			lop_id INTEGER,
			mon_hoc_id INTEGER,
			tuan INTEGER,
			so_tiet INTEGER,
			FOREIGN KEY (giao_vien_id) REFERENCES giaovien(id),
			FOREIGN KEY (lop_id) REFERENCES lophoc(id),
			FOREIGN KEY (mon_hoc_id) REFERENCES monhoc(id)
		);
		CREATE TABLE IF NOT EXISTS rangbuoc(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			giao_vien_id INTEGER,
			thu INTEGR,
			tiet INTEGER,
			loai_rang_buoc TEXT,
			FOREIGN KEY (giao_vien_id) REFERENCES giaovien(id)
		);
		CREATE TABLE IF NOT EXISTS thuathieu(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			lop_id INTEGER,
			mon_hoc_id INTEGER,
			tiet_thieu INTEGER,
			FOREIGN KEY (lop_id) REFERENCES lophoc(id),
			FOREIGN KEY (mon_hoc_id) REFERENCES monhoc(id)
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
		sqlInsertRangBuoc = `
		INSERT INTO rangbuoc(giao_vien_id, thu, tiet, loai_rang_buoc)
		VALUES (?, ?, ?, ?);
		`
		sqlInsertThuaThieu = `
		INSERT INTO thuathieu(lop_id, mon_hoc_id, tiet_thieu)
		VALUES (?, ?, ?);
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
		sqlEditRangBuoc = `
		UPDATE rangbuoc
		SET thu = ?, tiet = ?, loai_rang_buoc = ?
		WHERE id = ?;
		`
		sqlEditThuaThieu = `
		UPDATE thuathieu
		SET tiet_thieu = ?
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
		sqlDeleteRangBuoc = `
		DELETE FROM rangbuoc
		WHERE id = ?;
		`
		sqlDeleteThuaThieu = `
		DELETE FROM thuathieu
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
		sqlSelectRangBuoc = `
		SELECT * FROM rangbuoc
		WHERE id = ?;
		`
		sqlSelectThuaThieu = `
		SELECT * FROM thuathieu
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
		sqlSelectAllRangBuoc = `
		SELECT * FROM rangbuoc;
		`
		sqlSelectAllThuaThieu = `
		SELECT * FROM thuathieu;
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
		sqlFindRangBuoc = `
		SELECT * FROM rangbuoc
		WHERE giao_vien_id = ?;
		`
		sqlFindThuaThieu = `
		SELECT * FROM thuathieu
		WHERE mon_hoc_id = ?;
		`
	)

// define struct Giaovien, MonHoc, LopHoc, PhanCong, RangBuoc, ThuaThieu

type GiaoVien struct {
	ID int `json:"id"`
	TenNgan string `json:"ten_ngan"`
	HoTen string `json:"ho_ten"`
	MonChinhId int `json:"mon_chinh_id"`
}
type MonHoc struct {
	ID int `json:"id"`
	TenMon string `json:"ten_mon"`
	LoaiMon string `json:"loai_mon"`
}
type LopHoc struct {
	ID int `json:"id"`
	TenLop string `json:"ten_lop"`
	KhoiLop string `json:"khoi_lop"`
}
type PhanCong struct {
	ID int `json:"id"`
	GiaoVienId int `json:"giao_vien_id"`
	LopId int `json:"lop_id"`
	MonHocId int `json:"mon_hoc_id"`
	Tuan int `json:"tuan"`
	Sotiet int `json:"so_tiet"`
}
type RangBuoc struct {
	ID int `json:"id"`
	GiaoVienId int `json:"giao_vien_id"`
	Thu int `json:"thu"`
	Tiet int `json:"tiet"`
	LoaiRangBuoc string `json:"loai_rang_buoc"`
}
type ThuaThieu struct {
	ID int `json:"id"`
	LopId int `json:"lop_id"`
	MonhocId int `json:"mon_hoc_id"`
	TietThieu int `json:"tiet_thieu"`
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
			return count, fmt.Errorf("không thể thêm phần cống %s (%w)", pc.GiaoVienId, err)
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
			return count, fmt.Errorf("không thể thêm rang bước %s (%w)", rb.GiaoVienId, err)
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

func (s *SqlTKB) InsertThuaThieu(thuathieu []ThuaThieu) (int, error) {
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
	stmt, err := tx.Prepare(sqlInsertThuaThieu)
	if err != nil {
		return 0, fmt.Errorf("không thể chuẩn bị câu lệnh: %w", err)
	}
	defer stmt.Close() // Đóng statement khi xong
	for _, tt := range thuathieu {
		// Dùng statement đã chuẩn bị để thực thi
		_, err = stmt.Exec(tt.LopId, tt.MonhocId, tt.TietThieu)
		if err != nil {
			// Lỗi được gán cho biến err, defer sẽ kích hoạt Rollback()
			return count, fmt.Errorf("không thể thêm thuật thiếu %s (%w)", tt.LopId, err)
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


// func EditGiaoVien, EditMonHoc, EditLopHoc, EditPhanCong, EditRangBuoc, EditThuaThieu
// return error

func (s *SqlTKB) EditGiaoVien(gv GiaoVien) error {
	_, err := s.db.Exec(sqlEditGiaoVien, gv.TenNgan, gv.HoTen, gv.MonChinhId, gv.ID)
	if err != nil {
		return fmt.Errorf("Không thể chỉnh sửa giáo viên %s (%w )", gv.TenNgan, err)
	}
	return nil
}

func (s *SqlTKB) EditMonHoc(mh MonHoc) error {
	_, err := s.db.Exec(sqlEditMonHoc, mh.TenMon, mh.LoaiMon, mh.ID)
	if err != nil {
		return fmt.Errorf("Không thể chỉnh sửa môn học %s (%w )", mh.TenMon, err)
	}
	return nil
}

func (s *SqlTKB) EditLopHoc(lh LopHoc) error {
	_, err := s.db.Exec(sqlEditLopHoc, lh.TenLop, lh.KhoiLop, lh.ID)
	if err != nil {
		return fmt.Errorf("Không thể chỉnh sửa lớp học %s (%w )", lh.TenLop, err)
	}
	return nil
}

func (s *SqlTKB) EditPhanCong(pc PhanCong) error {
	_, err := s.db.Exec(sqlEditPhanCong, pc.Tuan, pc.Sotiet, pc.ID)
	if err != nil {
		return fmt.Errorf("Không thể chỉnh sửa phần cống %s (%w )", pc.GiaoVienId, err)
	}
	return nil
}

func (s *SqlTKB) EditRangBuoc(rb RangBuoc) error {
	_, err := s.db.Exec(sqlEditRangBuoc, rb.Thu, rb.Tiet, rb.LoaiRangBuoc, rb.ID)
	if err != nil {
		return fmt.Errorf("Không thể chỉnh sửa rang bước %s (%w )", rb.GiaoVienId, err)
	}
	return nil
}

func (s *SqlTKB) EditThuaThieu(tt ThuaThieu) error {
	_, err := s.db.Exec(sqlEditThuaThieu, tt.TietThieu, tt.ID)
	if err != nil {
		return fmt.Errorf("Không thể chỉnh sửa thuật thiếu %s (%w )", tt.LopId, err)
	}
	return nil
}

// func DeleteGiaoVien, DeleteMonHoc, DeleteLopHoc, DeletePhanCong, DeleteRangBuoc, DeleteThuaThieu
// return error

func (s *SqlTKB) DeleteGiaoVien(id int) error {
	_, err := s.db.Exec(sqlDeleteGiaoVien, id)
	if err != nil {
		return fmt.Errorf("Không thể xóa giáo viên %s (%w )", id, err)
	}
	return nil
}

func (s *SqlTKB) DeleteMonHoc(id int) error {
	_, err := s.db.Exec(sqlDeleteMonHoc, id)
	if err != nil {
		return fmt.Errorf("Không thể xóa môn học %s (%w )", id, err)
	}
	return nil
}

func (s *SqlTKB) DeleteLopHoc(id int) error {
	_, err := s.db.Exec(sqlDeleteLopHoc, id)
	if err != nil {
		return fmt.Errorf("Không thể xóa lớp học %s (%w )", id, err)
	}
	return nil
}

func (s *SqlTKB) DeletePhanCong(id int) error {
	_, err := s.db.Exec(sqlDeletePhanCong, id)
	if err != nil {
		return fmt.Errorf("Không thể xóa phần cống %s (%w )", id, err)
	}
	return nil
}

func (s *SqlTKB) DeleteRangBuoc(id int) error {
	_, err := s.db.Exec(sqlDeleteRangBuoc, id)
	if err != nil {
		return fmt.Errorf("Không thể xóa rang bước %s (%w )", id, err)
	}
	return nil
}

func (s *SqlTKB) DeleteThuaThieu(id int) error {
	_, err := s.db.Exec(sqlDeleteThuaThieu, id)
	if err != nil {
		return fmt.Errorf("Không thể xóa thuật thiếu %s (%w )", id, err)
	}
	return nil
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

func (s *SqlTKB) SelectRangBuoc(id int) (*RangBuoc, error) {
	var rb RangBuoc
	err := s.db.QueryRow(sqlSelectRangBuoc, id).Scan(&rb.ID, &rb.GiaoVienId, &rb.Thu, &rb.Tiet, &rb.LoaiRangBuoc)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu từ bảng rangbuoc : %w", err)
	}
	return &rb, nil
}

func (s *SqlTKB) SelectThuaThieu(id int) (*ThuaThieu, error) {
	var tt ThuaThieu
	err := s.db.QueryRow(sqlSelectThuaThieu, id).Scan(&tt.ID, &tt.LopId, &tt.MonhocId, &tt.TietThieu)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu từ bảng thuathieu : %w", err)
	}
	return &tt, nil
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

func (s *SqlTKB) SelectAllThuaThieu() ([]ThuaThieu, error) {
	var thuathieu []ThuaThieu
	Rows, err := s.db.Query(sqlSelectAllThuaThieu)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu từ bảng thuathieu : %w", err)
	}
	defer Rows.Close()
	for Rows.Next() {
		var tt ThuaThieu		
		err := Rows.Scan(&tt.ID, &tt.LopId, &tt.MonhocId, &tt.TietThieu)	
		if err != nil {			
			return nil, fmt.Errorf("Không thể quét dữ liệu vào biến thuathieu : %w", err)
		}		
		thuathieu = append(thuathieu, tt)
	}	
	return thuathieu, nil
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

func (s *SqlTKB) FindRangBuoc(giao_vien_id int) (*RangBuoc, error) {	
	var rb RangBuoc
	err := s.db.QueryRow(sqlFindRangBuoc, giao_vien_id).Scan(&rb.ID, &rb.GiaoVienId, &rb.Thu, &rb.Tiet, &rb.LoaiRangBuoc)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu từ bảng rangbuoc : %w", err)
	}
	return &rb, nil
}

func (s *SqlTKB) FindThuaThieu(mon_hoc_id int) (*ThuaThieu, error) {	
	var tt ThuaThieu
	err := s.db.QueryRow(sqlFindThuaThieu, mon_hoc_id).Scan(&tt.ID, &tt.LopId, &tt.MonhocId, &tt.TietThieu)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu từ bảng thuathieu : %w", err)
	}
	return &tt, nil
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

