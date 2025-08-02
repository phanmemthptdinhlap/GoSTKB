package models

type LopHoc struct {
	MaLopHoc    string `json:"ma_lop_hoc"`
	TenLop      string `json:"ten_lop"`
	KhoiLop     string `json:"khoi_lop"`
	MaChuNhiem  string `json:"ma_chu_nhiem"`
	TenChuNhiem string `json:"ten_chu_nhiem"`
}
