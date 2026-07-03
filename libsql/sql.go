package libsql

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)
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
func (s *SqlTKB) Init(db *sql.DB) error {
	s.db = db
	return nil
}

func (s *SqlTKB) InsertGiaoVien(data GiaoVien) error {
	sqlInsert := `
	INSERT INTO giaovien(ten_ngan, ho_ten, mon_chinh_id)
	VALUES (?, ?, ?);
	`
	_, err := s.db.Exec(sqlInsert, data.TenNgan, data.HoTen, data.MonChinhId)
	if err != nil {
		return fmt.Errorf("không thêm dữ liệu vào bảng : %w", err)
	}
	return nil
}

func (s *SqlTKB) InsertMonHoc(data MonHoc) error {
	sqlInsert := `
	INSERT INTO monhoc(ten_mon, loai_mon)
	VALUES (?, ?);
	`
	_, err := s.db.Exec(sqlInsert, data.TenMon, data.LoaiMon)
	if err != nil {
		return fmt.Errorf("không thêm dữ liệu vào bảng : %w", err)
	}
	return nil
}

func (s *SqlTKB) InsertLopHoc(data LopHoc) error {
	sqlInsert := `
	INSERT INTO lophoc(ten_lop, khoi_lop)
	VALUES (?, ?);
	`
	_, err := s.db.Exec(sqlInsert, data.TenLop, data.KhoiLop)
	if err != nil {
		return fmt.Errorf("không thêm dữ liệu vào bảng : %w", err)
	}
	return nil
}

func (s *SqlTKB) InsertPhanCong(data PhanCong) error {
	sqlInsert := `
	INSERT INTO phancong(giao_vien_id, lop_id, mon_hoc_id, tuan, so_tiet)
	VALUES (?, ?, ?, ?, ?);
	`
	_, err := s.db.Exec(sqlInsert, data.GiaoVienId, data.LopId, data.MonHocId, data.Tuan, data.Sotiet)
	if err != nil {
		return fmt.Errorf("không thêm dữ liệu vào bảng : %w", err)
	}
	return nil
}

func (s *SqlTKB) InsertRangBuoc(data RangBuoc) error {
	sqlInsert := `
	INSERT INTO rangbuoc(giao_vien_id, thu, tiet, loai_rang_buoc)
	VALUES (?, ?, ?, ?);
	`
	_, err := s.db.Exec(sqlInsert, data.GiaoVienId, data.Thu, data.Tiet, data.LoaiRangBuoc)
	if err != nil {
		return fmt.Errorf("không thêm dữ liệu vào bảng : %w", err)
	}
	return nil
}

func (s *SqlTKB) InsertThuaThieu(data ThuaThieu) error {
	sqlInsert := `
	INSERT INTO thuathieu(lop_id, mon_hoc_id, tiet_thieu)
	VALUES (?, ?, ?);
	`
	_, err := s.db.Exec(sqlInsert, data.LopId, data.MonhocId, data.TietThieu)
	if err != nil {
		return fmt.Errorf("không thêm dữ liệu vào bảng : %w", err)
	}
	return nil
}
func (s *SqlTKB) EditGiaoVien(data GiaoVien) error {
	sqlEdit := `
	UPDATE giaovien
	SET ten_ngan = ?, ho_ten = ?, mon_chinh_id = ?
	WHERE id = ?;
	`
	_, err := s.db.Exec(sqlEdit, data)
	if err != nil {
		return fmt.Errorf("không thêm dữ liệu vào bảng : %w", err)
	}
	return nil
}

func (s *SqlTKB) EditMonHoc(data MonHoc) error {
	sqlEdit := `
	UPDATE monhoc
	SET ten_mon = ?, loai_mon = ?
	WHERE id = ?;
	`
	_, err := s.db.Exec(sqlEdit, data.TenMon, data.LoaiMon, data.ID)
	if err != nil {
		return fmt.Errorf("không thêm dữ liệu vào bảng : %w", err)
	}
	return nil
}

func (s *SqlTKB) EditLopHoc(data LopHoc) error {
	sqlEdit := `
	UPDATE lophoc
	SET ten_lop = ?, khoi_lop = ?
	WHERE id = ?;
	`
	_, err := s.db.Exec(sqlEdit, data.TenLop, data.KhoiLop, data.ID)
	if err != nil {
		return fmt.Errorf("không thêm dữ liệu vào bảng : %w", err)
	}
	return nil
}

func (s *SqlTKB) EditPhanCong(data PhanCong) error {
	sqlEdit := `
	UPDATE phancong
	SET tuan = ?, so_tiet = ?
	WHERE id = ?;
	`
	_, err := s.db.Exec(sqlEdit, data.Tuan, data.Sotiet, data.ID)
	if err != nil {
		return fmt.Errorf("không thêm dữ liệu vào bảng : %w", err)
	}
	return nil
}

func (s *SqlTKB) EditRangBuoc(data RangBuoc) error {
	sqlEdit := `
	UPDATE rangbuoc
	SET thu = ?, tiet = ?, loai_rang_buoc = ?
	WHERE id = ?;
	`
	_, err := s.db.Exec(sqlEdit, data.Thu, data.Tiet, data.LoaiRangBuoc, data.ID)
	if err != nil {
		return fmt.Errorf("không thêm dữ liệu vào bảng : %w", err)
	}
	return nil
}

func (s *SqlTKB) EditThuaThieu(data ThuaThieu) error {
	sqlEdit := `
	UPDATE thuathieu
	SET tiet_thieu = ?
	WHERE id = ?;
	`
	_, err := s.db.Exec(sqlEdit, data.TietThieu, data.ID)
	if err != nil {
		return fmt.Errorf("không thêm dữ liệu vào bảng : %w", err)
	}
	return nil
}
func (s *SqlTKB) DeleteGiaoVien(id int) error {
	sqlDelete := `
	DELETE FROM giaovien
	WHERE id = ?;
	`
	_, err := s.db.Exec(sqlDelete, id)
	if err != nil {
		return fmt.Errorf("không thêm dữ liệu vào bảng : %w", err)
	}
	return nil
}

func (s *SqlTKB) DeleteMonHoc(id int) error {
	sqlDelete := `
	DELETE FROM monhoc
	WHERE id = ?;
	`
	_, err := s.db.Exec(sqlDelete, id)
	if err != nil {
		return fmt.Errorf("không thêm dữ liệu vào bảng : %w", err)
	}
	return nil
}

func (s *SqlTKB) DeleteLopHoc(id int) error {
	sqlDelete := `
	DELETE FROM lophoc
	WHERE id = ?;
	`
	_, err := s.db.Exec(sqlDelete, id)
	if err != nil {
		return fmt.Errorf("không thêm dữ liệu vào bảng : %w", err)
	}
	return nil
}

func (s *SqlTKB) DeletePhanCong(id int) error {
	sqlDelete := `
	DELETE FROM phancong
	WHERE id = ?;
	`
	_, err := s.db.Exec(sqlDelete, id)
	if err != nil {
		return fmt.Errorf("không thêm dữ liệu vào bảng : %w", err)
	}
	return nil
}

func (s *SqlTKB) DeleteRangBuoc(id int) error {
	sqlDelete := `
	DELETE FROM rangbuoc
	WHERE id = ?;
	`
	_, err := s.db.Exec(sqlDelete, id)
	if err != nil {
		return fmt.Errorf("không thêm dữ liệu vào bảng : %w", err)
	}
	return nil
}

func (s *SqlTKB) DeleteThuaThieu(id int) error {
	sqlDelete := `
	DELETE FROM thuathieu
	WHERE id = ?;
	`
	_, err := s.db.Exec(sqlDelete, id)
	if err != nil {
		return fmt.Errorf("không thêm dữ liệu vào bảng : %w", err)
	}
	return nil
}
func (s *SqlTKB) SelectGiaoVien(id int) (*GiaoVien, error) {
	sqlSelect := `
	SELECT * FROM giaovien
	WHERE id = ?;
	`
	var data GiaoVien
	err := s.db.QueryRow(sqlSelect, id).Scan(&data.ID, &data.TenNgan, &data.HoTen, &data.MonChinhId)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu : %w", err)
	}
	return &data, nil
}

func (s *SqlTKB) SelectMonHoc(id int) (*MonHoc, error) {
	sqlSelect := `
	SELECT * FROM monhoc
	WHERE id = ?;
	`
	var data MonHoc
	err := s.db.QueryRow(sqlSelect, id).Scan(&data.ID, &data.TenMon, &data.LoaiMon)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu : %w", err)
	}
	return &data, nil
}

func (s *SqlTKB) SelectLopHoc(id int) (*LopHoc, error) {
	sqlSelect := `
	SELECT * FROM lophoc
	WHERE id = ?;
	`
	var data LopHoc
	err := s.db.QueryRow(sqlSelect, id).Scan(&data.ID, &data.TenLop, &data.KhoiLop)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu : %w", err)
	}
	return &data, nil
}

func (s *SqlTKB) SelectPhanCong(id int) (*PhanCong, error) {
	sqlSelect := `
	SELECT * FROM phancong
	WHERE id = ?;
	`
	var data PhanCong
	err := s.db.QueryRow(sqlSelect, id).Scan(&data.ID, &data.GiaoVienId, &data.LopId, &data.MonHocId, &data.Tuan, &data.Sotiet)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu : %w", err)
	}
	return &data, nil
}

func (s *SqlTKB) SelectRangBuoc(id int) (*RangBuoc, error) {
	sqlSelect := `
	SELECT * FROM rangbuoc
	WHERE id = ?;
	`
	var data RangBuoc
	err := s.db.QueryRow(sqlSelect, id).Scan(&data.ID, &data.GiaoVienId, &data.Thu, &data.Tiet, &data.LoaiRangBuoc)	
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu : %w", err)
	}
	return &data, nil
}

func (s *SqlTKB) SelectThuaThieu(id int) (*ThuaThieu, error) {
	sqlSelect := `
	SELECT * FROM thuathieu
	WHERE id = ?;
	`
	var data ThuaThieu
	err := s.db.QueryRow(sqlSelect, id).Scan(&data.ID, &data.LopId, &data.MonhocId, &data.TietThieu)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu : %w", err)
	}
	return &data, nil
}
func (s *SqlTKB) SelectAllGiaoVien() ([]GiaoVien, error) {
	sqlSelect := `
	SELECT * FROM giaovien;
	`
	var data []GiaoVien
	Rows, err := s.db.Query(sqlSelect)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu : %w", err)
	}
	defer Rows.Close()
	for Rows.Next() {
		var gv GiaoVien
		err := Rows.Scan(&gv.ID, &gv.TenNgan, &gv.HoTen, &gv.MonChinhId)
		if err != nil {
			return nil, fmt.Errorf("không thể lấy dữ liệu : %w", err)
		}
		data = append(data, gv)
	}
	return data, nil
}

func (s *SqlTKB) SelectAllMonHoc() ([]MonHoc, error) {
	sqlSelect := `
	SELECT * FROM monhoc;
	`
	var data []MonHoc
	 Rows, err := s.db.Query(sqlSelect, &data)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu : %w", err)
	}
	defer Rows.Close()
	for Rows.Next() {
		var mh MonHoc
		err := Rows.Scan(&mh.ID, &mh.TenMon, &mh.LoaiMon)
		if err != nil {
			return nil, fmt.Errorf("không thể lấy dữ liệu : %w", err)
		}
		data = append(data, mh)
	}
	return data, nil
}

func (s *SqlTKB) SelectAllLopHoc() ([]LopHoc, error) {
	sqlSelect := `
	SELECT * FROM lophoc;
	`
	var data []LopHoc
	Rows, err := s.db.Query(sqlSelect)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu : %w", err)
	}
	defer Rows.Close()
	for Rows.Next() {
		var lh LopHoc
		err := Rows.Scan(&lh.ID, &lh.TenLop, &lh.KhoiLop)
		if err != nil {
			return nil, fmt.Errorf("không thể lấy dữ liệu : %w", err)
		}
		data = append(data, lh)
	}
	return data, nil
}

func (s *SqlTKB) SelectAllPhanCong() ([]PhanCong, error) {
	sqlSelect := `
	SELECT * FROM phancong;
	`
	var data []PhanCong
	Rows, err := s.db.Query(sqlSelect)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu : %w", err)
	}
	defer Rows.Close()
	for Rows.Next() {
		var pc PhanCong
		err := Rows.Scan(&pc.ID, &pc.GiaoVienId, &pc.LopId, &pc.MonHocId, &pc.Tuan, &pc.Sotiet)
		if err != nil {
			return nil, fmt.Errorf("không thể lấy dữ liệu : %w", err)
		}
		data = append(data, pc)
	}
	return data, nil
}

func (s *SqlTKB) SelectAllRangBuoc() ([]RangBuoc, error) {
	sqlSelect := `
	SELECT * FROM rangbuoc;
	`
	var data []RangBuoc
	Rows, err := s.db.Query(sqlSelect)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu : %w", err)
	}
	defer Rows.Close()
	for Rows.Next() {
		var rb RangBuoc
		err := Rows.Scan(&rb.ID, &rb.GiaoVienId, &rb.Thu, &rb.Tiet, &rb.LoaiRangBuoc)
		if err != nil {
			return nil, fmt.Errorf("không thể lấy dữ liệu : %w", err)
		}
		data = append(data, rb)
	}
	return data, nil
}

func (s *SqlTKB) SelectAllThuaThieu() ([]ThuaThieu, error) {
	sqlSelect := `
	SELECT * FROM thuathieu;
	`
	var data []ThuaThieu
	Rows, err := s.db.Query(sqlSelect)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu : %w", err)
	}
	defer Rows.Close()
	for Rows.Next() {
		var tt ThuaThieu
		err := Rows.Scan(&tt.ID, &tt.LopId, &tt.MonhocId, &tt.TietThieu)
		if err != nil {
			return nil, fmt.Errorf("không thể lấy dữ liệu : %w", err)
		}
		data = append(data, tt)
	}
	return data, nil
}

func (s *SqlTKB) FindGiaoVien(ten_ngan string) (*GiaoVien, error) {
	sqlSelect := `
	SELECT * FROM giaovien
	WHERE ten_ngan = ?;
	`
	var data GiaoVien
	err := s.db.QueryRow(sqlSelect, ten_ngan).Scan(&data.ID, &data.TenNgan, &data.HoTen, &data.MonChinhId)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu : %w", err)
	}
	return &data, nil
}

func (s *SqlTKB) FindMonHoc(ten_mon string) (*MonHoc, error) {
	sqlSelect := `
	SELECT * FROM monhoc
	WHERE ten_mon = ?;
	`
	var data MonHoc
	err := s.db.QueryRow(sqlSelect, ten_mon).Scan(&data.ID, &data.TenMon, &data.LoaiMon)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu : %w", err)
	}
	return &data, nil
}

func (s *SqlTKB) FindLopHoc(ten_lop string) (*LopHoc, error) {
	sqlSelect := `
	SELECT * FROM lophoc
	WHERE ten_lop = ?;
	`
	var data LopHoc
	err := s.db.QueryRow(sqlSelect, ten_lop).Scan(&data.ID, &data.TenLop, &data.KhoiLop)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu : %w", err)
	}
	return &data, nil
}

func (s *SqlTKB) FindPhanCong(giao_vien_id int) (*PhanCong, error) {
	sqlSelect := `
	SELECT * FROM phancong
	WHERE giao_vien_id = ?;
	`
	var data PhanCong
	err := s.db.QueryRow(sqlSelect, giao_vien_id).Scan(&data.ID, &data.GiaoVienId, &data.LopId, &data.MonHocId, &data.Tuan, &data.Sotiet)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu : %w", err)
	}
	return &data, nil
}

func (s *SqlTKB) FindRangBuoc(giao_vien_id int) (*RangBuoc, error) {
	sqlSelect := `
	SELECT * FROM rangbuoc
	WHERE giao_vien_id = ?;
	`
	var data RangBuoc
	err := s.db.QueryRow(sqlSelect, giao_vien_id).Scan(&data.ID, &data.GiaoVienId, &data.Thu, &data.Tiet, &data.LoaiRangBuoc)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu : %w", err)
	}
	return &data, nil
}

func (s *SqlTKB) FindThuaThieu(mon_hoc_id int) (*ThuaThieu, error) {
	sqlSelect := `
	SELECT * FROM thuathieu
	WHERE mon_hoc_id = ?;
	`
	var data ThuaThieu
	err := s.db.QueryRow(sqlSelect, mon_hoc_id).Scan(&data.ID, &data.LopId, &data.MonhocId, &data.TietThieu)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy dữ liệu : %w", err)
	}
	return &data, nil
}
func (s *SqlTKB) Close() error {
	return s.db.Close()
}

func ConnectSTKB() (*SqlTKB, error) {
	// Kết nối SQLite
	db, err := sql.Open("sqlite3", "./stkb.db")
	if err != nil {
		return nil, fmt.Errorf("không kết nối được với STKB : %w", err)
	}
	sqlCreateTables := `
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
	_, err = db.Exec(sqlCreateTables)
	if err != nil {
		return nil, fmt.Errorf("không thể tạo bảng : %w", err)
	}
	return &SqlTKB{db}, nil
}

