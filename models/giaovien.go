package models

// GiaoVien represents a teacher in the system.
type GiaoVien struct {
	MaGiaoVien string `json:"ma_giao_vien"`
	HoTen      string `json:"ho_ten"`
	TenTKB     string `json:"ten_tkb"`
}
