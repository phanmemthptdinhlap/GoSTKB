package models

import "gorm.io/gorm"

type LopHoc struct {
	gorm.Model
	MaLopHoc string `json:"ma_lop_hoc" gorm:"unique;not null"`
	TenLop      string `json:"ten_lop" gorm:"not null"`
	Khoi     init `json:"khoi" gorm:"not null"`
}