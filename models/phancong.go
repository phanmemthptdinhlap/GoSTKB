package models

import "gorm.io/gorm"

// Phân công giảng dạy
type PhanCong struct {
	gorm.Model
	MaPhanCong string   `json:"ma_phan_cong" gorm:"unique;not null"`
	MaGiaoVien string   `json:"ma_giao_vien" gorm:"not null"`
	MaLopHoc   string   `json:"ma_lop_hoc" gorm:"not null"`
	MaMonHoc   string   `json:"ten_mon_hoc" gorm:"not null"`
	GiaoVien   GiaoVien `gorm:"foreignKey:MaGiaoVien"`
	LopHoc     LopHoc   `gorm:"foreignKey:MaLopHoc"`
	MonHoc     MonHoc   `gorm:"foreignKey:MaMonHoc"`
}
