package handlers

import (
	"GoSTKB/models"
	"fmt"
	"net/http"
	"strings"

	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
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

// Thêm vào file hiện tại

func (h *ThaoTac_GiaoVien) NhapDanhSachGiaoVien(c *gin.Context) {
	// Nhận file từ form
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Vui lòng chọn file Excel"})
		return
	}

	// Mở file Excel
	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể mở file Excel"})
		return
	}
	defer f.Close()

	fmt.Printf("Đã nhận file: %s\n", file.Filename)
	// Đọc file Excel
	xlsx, err := excelize.OpenReader(f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể đọc file Excel"})
		return
	}
	defer xlsx.Close()

	// Lấy tên sheet đầu tiên
	sheetName := xlsx.GetSheetName(0)
	fmt.Printf("Tên sheet: %s\n", sheetName)
	// Khởi tạo biến đếm số lượng cập nhật và thêm mới
	var countUpdated, countInserted int

	// Đọc từng dòng trong file Excel
	rows, err := xlsx.GetRows(sheetName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể đọc dữ liệu từ sheet"})
		return
	}
	// Kiểm tra nếu không có dòng nào
	fmt.Printf("Số dòng trong sheet: %d\n", len(rows))
	// Bỏ qua dòng tiêu đề
	for i := 1; i < len(rows); i++ {
		row := rows[i]
		if len(row) < 3 {
			continue // Bỏ qua dòng không đủ dữ liệu
		}

		maGiaoVien := strings.TrimSpace(row[0])
		hoTen := strings.TrimSpace(row[1])
		tenTKB := strings.TrimSpace(row[2])
		fmt.Printf("Mã giáo viên: %s, Họ tên: %s, Tên TKB: %s\n", maGiaoVien, hoTen, tenTKB)

		// Kiểm tra dữ liệu bắt buộc
		if hoTen == "" || tenTKB == "" {
			continue // Bỏ qua nếu thiếu thông tin bắt buộc
		}

		if maGiaoVien != "" {
			// Kiểm tra xem giáo viên đã tồn tại chưa
			var exists bool
			err := h.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM giaovien WHERE ma_giao_vien = ?)", maGiaoVien).Scan(&exists)
			if err != nil {
				continue
			}

			if exists {
				// Cập nhật nếu đã tồn tại
				_, err = h.DB.Exec("UPDATE giaovien SET ho_ten = ?, ten_tkb = ? WHERE ma_giao_vien = ?",
					hoTen, tenTKB, maGiaoVien)
				if err == nil {
					countUpdated++
				}
			} else {
				// Thêm mới với mã giáo viên được chỉ định
				_, err = h.DB.Exec("INSERT INTO giaovien (ma_giao_vien, ho_ten, ten_tkb) VALUES (?, ?, ?)",
					maGiaoVien, hoTen, tenTKB)
				if err == nil {
					countInserted++
				}
			}
		} else {
			// Thêm mới không có mã giáo viên
			_, err = h.DB.Exec("INSERT INTO giaovien (ho_ten, ten_tkb) VALUES (?, ?)",
				hoTen, tenTKB)
			if err == nil {
				countInserted++
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Nhập dữ liệu thành công",
		"updated":  countUpdated,
		"inserted": countInserted,
	})
}

func (h *ThaoTac_GiaoVien) XuatDanhSachGiaoVien(c *gin.Context) {
	// Tạo file Excel mới
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}()
	shellnameold := f.GetSheetName(0)
	// Tạo sheet mới
	sheetName := "Danh sách giáo viên"
	if err := f.SetSheetName(shellnameold, sheetName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể tạo sheet"})
		return
	}
	index, err := f.NewSheet(sheetName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Đặt tiêu đề các cột
	f.SetCellValue(sheetName, "A1", "Mã giáo viên")
	f.SetCellValue(sheetName, "B1", "Họ và tên")
	f.SetCellValue(sheetName, "C1", "Tên trên TKB")

	// Truy vấn dữ liệu từ database
	rows, err := h.DB.Query("SELECT ma_giao_vien, ho_ten, ten_tkb FROM giaovien")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể truy vấn dữ liệu"})
		return
	}
	defer rows.Close()

	// Ghi dữ liệu vào file Excel
	rowIndex := 2
	for rows.Next() {
		var id int
		var hoten, tentkb string
		if err := rows.Scan(&id, &hoten, &tentkb); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		f.SetCellValue(sheetName, fmt.Sprintf("A%d", rowIndex), id)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", rowIndex), hoten)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", rowIndex), tentkb)
		rowIndex++
	}

	f.SetActiveSheet(index)

	// Thiết lập header cho response
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=danh_sach_giao_vien.xlsx")

	// Ghi file Excel vào response
	if err := f.Write(c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
