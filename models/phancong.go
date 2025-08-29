package models

// Phân công giảng dạy
type PhanCong struct {
	MaPhanCong string `json:"ma_phan_cong" label:"Mã phân công"`
	MaGiaoVien string `json:"ma_giao_vien" label:"Mã giáo viên"`
	MaMonHoc   string `json:"ma_mon" label:"Mã môn học"`
	MaLopHoc   string `json:"ma_lop" label:"Mã lớp học"`
}
