package models

// Phân công giảng dạy
type PhanCong struct {
	MaPhanCong string `json:"ma_phan_cong"`
	MaGiaoVien string `json:"ma_giao_vien"`
	MaMonHoc   string `json:"ma_mon_hoc"`
	MaLopHoc   string `json:"ma_lop_hoc"`
}
