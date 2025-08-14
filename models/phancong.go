package models

// Phân công giảng dạy
type PhanCong struct {
	MaPhanCong string `json:"ma" label:"Mã phân công"`
	MaGiaoVien string `json:"ma_gv" label:"Mã giáo viên"`
	MaMonHoc   string `json:"ma_mh" label:"Mã môn học"`
	MaLopHoc   string `json:"ma_lh" label:"Mã lớp học"`
}
