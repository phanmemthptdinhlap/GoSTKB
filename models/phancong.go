package models

// Phân công giảng dạy
type PhanCong struct {
	MaPhanCong  string `json:"ma_phan_cong"`
	MaGiaoVien  string `json:"ma_giao_vien"`
	MaMon       string `json:"ma_mon"`
	MaLop       string `json:"ma_lop"`
	TenGiaoVien string `json:"ten_giao_vien"`
	TenMon      string `json:"ten_mon"`
	TenLop      string `json:"ten_lop"`
}
