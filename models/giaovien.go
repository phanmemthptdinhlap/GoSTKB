package models

// GiaoVien represents a teacher in the system.
type GiaoVien struct {
	MaGiaoVien string `json:"ma" label:"Mã giáo viên"`
	HoTen      string `json:"ten" label:"Họ tên giáo viên"`
	TenTKB     string `json:"ten_tkb" label:"Tên trên TKB"`
}
