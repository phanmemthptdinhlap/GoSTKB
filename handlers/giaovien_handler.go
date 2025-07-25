package handlers

import (
	"GoSTKB/models"
	"fmt"
	"net/http"

	"database/sql"

	"github.com/gin-gonic/gin"
)

// Danh sách giáo viên
type ThaoTac_GiaoVien struct {
	DB *sql.DB
}

// Lấy danh sách giáo viên
func (h *ThaoTac_GiaoVien) DanhSachGiaoVien(c *gin.Context) {
	var giaovien []models.GiaoVien
	rows, err := h.DB.Query("SELECT ma_giao_vien, ho_ten, ten_tkb FROM giaovien")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể truy vấn dữ liệu"})
		return
	}
	defer rows.Close()
	for rows.Next() {
		var gv models.GiaoVien
		if err := rows.Scan(&gv.MaGiaoVien, &gv.HoTen, &gv.TenTKB); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi khi đọc dữ liệu"})
			return
		}
		giaovien = append(giaovien, gv)
	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi khi duyệt dữ liệu"})
		return
	}
	if len(giaovien) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Không có giáo viên nào"})
		return
	}
	c.JSON(http.StatusOK, giaovien)
}

// Tạo giáo viên
func (h *ThaoTac_GiaoVien) TaoGiaoVien(c *gin.Context) {
	var giaovien models.GiaoVien
	if err := c.ShouldBindJSON(&giaovien); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var sqlcmd string
	var args []interface{}
	if giaovien.MaGiaoVien == "" {
		sqlcmd = "INSERT INTO giaovien (ho_ten, ten_tkb) VALUES (?, ?)"
		args = []interface{}{giaovien.HoTen, giaovien.TenTKB}
	} else {
		sqlcmd = "INSERT INTO giaovien (ma_giao_vien, ho_ten, ten_tkb) VALUES (?, ?, ?)"
		args = []interface{}{giaovien.MaGiaoVien, giaovien.HoTen, giaovien.TenTKB}
	}
	result, err := h.DB.Exec(sqlcmd, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể thêm giáo viên"})
		return
	}
	id, _ := result.LastInsertId()
	giaovien.MaGiaoVien = fmt.Sprintf("%d", id)
	c.JSON(http.StatusOK, giaovien)
}

// Cập nhật giáo viên
func (h *ThaoTac_GiaoVien) CapNhatGiaoVien(c *gin.Context) {
	var giaovien models.GiaoVien
	id := c.Param("id")
	row := h.DB.QueryRow("SELECT ma_giao_vien, ho_ten, ten_tkb FROM giaovien WHERE ma_giao_vien = ?", id)
	if err := row.Scan(&giaovien.MaGiaoVien, &giaovien.HoTen, &giaovien.TenTKB); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Không tìm thấy thông tin giáo viên này"})
		return
	}
	if err := c.ShouldBindJSON(&giaovien); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	_, err := h.DB.Exec("UPDATE giaovien SET ho_ten = ?, ten_tkb = ? WHERE ma_giao_vien = ?", giaovien.HoTen, giaovien.TenTKB, giaovien.MaGiaoVien)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể cập nhật giáo viên"})
		return
	}
	c.JSON(http.StatusOK, giaovien)
}

// Xóa giáo viên khỏi danh sách
func (h *ThaoTac_GiaoVien) XoaGiaoVien(c *gin.Context) {
	id := c.Param("id")
	fmt.Printf("ID: %s", id)
	var giaovien models.GiaoVien
	row := h.DB.QueryRow("SELECT ma_giao_vien, ho_ten, ten_tkb FROM giaovien WHERE ma_giao_vien = ?", id)
	if err := row.Scan(&giaovien.MaGiaoVien, &giaovien.HoTen, &giaovien.TenTKB); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Không tìm thấy thông tin giáo viên"})
		return
	}
	_, err := h.DB.Exec("DELETE FROM giaovien WHERE ma_giao_vien = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Không thể xóa giáo viên"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Xóa giáo viên thành công"})
}
func (h *ThaoTac_GiaoVien) TestGiaoVien(c *gin.Context) {
	fmt.Print("test giao vien")
	c.JSON(http.StatusNotFound, gin.H{"message": "Test GIao VIên"})
}
