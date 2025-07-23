package models
import "gorm.io/gorm"

// Tiết học
type TietDay struct {
	gorm.Model
	Tuan int    `json:"tuan" gorm:"not null"`
	MaPhanCong string `json:"ma_phan_cong" gorm:"not null"`
	TongSoTiet	 int    `json:"so_tiet" gorm:"not null"`
	SoTietSang int    `json:"so_tiet_sang" gorm:"not null"`
	SoTietNo int	`json:"so_tiet_no" gorm:"not null"`
}