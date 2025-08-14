package models

// Tiết học
type TietDay struct {
	MaTietDay        string `json:"ma" label:"Mã tiết dạy"`
	MaPhanCong       string `json:"ma_pc" label:"Mã phân công"`
	TuanHoc          string `json:"tuan" label:"Tuần học"`
	TongTietPhanCong string `json:"tong_tiet" label:"Tổng tiết được phân công"`
	TietSang         string `json:"sang" label:"Tiết buổi sáng"`
	TietChieu        string `json:"chieu" label:"Tiết buổi chiều"`
	TietTuanTruoc    string `json:"tuan_truoc" label:"Tiết còn lại tuần trước"`
}
