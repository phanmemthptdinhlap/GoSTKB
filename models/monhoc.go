package models

import "gorm.io/gorm"

type MonHoc struct {
	gorm.Model
	MaMonHoc  string `json:"ma_mon_hoc" gorm:"unique;not null"`
	TenMonHoc string `json:"ten_mon_hoc" gorm:"not null"`
}
