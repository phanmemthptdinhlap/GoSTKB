package models

type LopHoc struct {
	MaLop       string `json:"ma" label:"Mã lớp học"`
	TenLop      string `json:"ten" label:"Tên lớp học"`
	KhoiLop     string `json:"khoi" label:"Khối lớp"`
	MaChuNhiem  string `json:"ma_cn" label:"Mã chủ nhiệm"`
	TenChuNhiem string `json:"ten_cn" label:"Tên chủ nhiệm"`
}
