package handlers

import (
	"GoSTKB/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Danh sách giáo viên
type ThaoTac_GiaoVien struct {
	DB *gorm.DB
}

// Lấy danh sách giáo viên
func (h *ThaoTac_GiaoVien) DanhSachGiaoVien(c *gin.Context) {
	var giaovien []models.GiaoVien
	h.DB.Find(&giaovien)
	c.JSON(http.StatusOK, giaovien)
}

// Tạo giáo viên
func (h *ThaoTac_GiaoVien) TaoGiaoVien(c *gin.Context) {
	var giaovien models.GiaoVien
	if err := c.ShouldBindJSON(&giaovien); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Create(&giaovien)
	c.JSON(http.StatusOK, giaovien)
}

// Cập nhật giáo viên
func (h *ThaoTac_GiaoVien) CapNhatGiaoVien(c *gin.Context) {
	var giaovien models.GiaoVien
	id := c.Param("id")
	if err := h.DB.First(&giaovien, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Không tìm thấy thông tin giáo viên này"})
		return
	}
	if err := c.ShouldBindJSON(&giaovien); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	h.DB.Save(&giaovien)
	c.JSON(http.StatusOK, giaovien)
}

// Xóa giáo viên khỏi danh sách
func (h *ThaoTac_GiaoVien) XoaGiaoVien(c *gin.Context) {
	id := c.Param("id")
	fmt.Printf("ID: %s", id)
	var giaovien models.GiaoVien
	if err := h.DB.First(&giaovien, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Không tìm thấy thông tin giáo viên"})
		return
	}
	h.DB.Delete(&giaovien)
	c.JSON(http.StatusOK, gin.H{
		"message": "Xóa giáo viên thành công"})
}
func (h *ThaoTac_GiaoVien) TestGiaoVien(c *gin.Context) {
	fmt.Print("test giao vien")
	c.JSON(http.StatusNotFound, gin.H{"message": "Test GIao VIên"})
}
