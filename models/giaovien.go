package models

import "gorm.io/gorm"

// GiaoVien represents a teacher in the system.
type GiaoVien struct {
	gorm.Model
	MaGiaoVien string `json:"ma_giao_vien" gorm:"unique;not null"`
	HoTen      string `json:"ho_ten" gorm:"not null"`
	TenTKB     string `json:"ten_tkb" gorm:"not null"`
}
