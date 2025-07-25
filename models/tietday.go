package models

// Tiết học
type TietDay struct {
	MaTietDay        string `json:"ma_tiet_day"`
	MaPhanCong       string `json:"ma_phan_cong"`
	TuanHoc          string `json:"tuan_hoc"`
	TongTietPhanCong string `json:"tong_tiet_phan_cong"`
	TietSang         string `json:"tiet_sang"`
	TietChieu        string `json:"tiet_chieu"`
	TietTuanTruoc    string `json:"tiet_tuan_truoc"`
}
