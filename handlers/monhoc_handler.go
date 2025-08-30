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
type ThaoTac_MonHoc struct {
	DB *sql.DB
}

// Lấy danh sách giáo viên
func (h *ThaoTac_MonHoc) DanhSachMonHoc(c *gin.Context) {
	var monhoc []models.MonHoc
	rows, err := h.DB.Query("SELECT ma_mon, ten_mon FROM monhoc")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể truy vấn dữ liệu"})
		fmt.Printf("Lỗi sql\n")
		return
	}
	defer rows.Close()
	for rows.Next() {
		var mh models.MonHoc
		if err := rows.Scan(&mh.MaMonHoc, &mh.TenMonHoc); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi khi đọc dữ liệu"})
			fmt.Printf("Lỗi Scan\n")
			return
		}
		monhoc = append(monhoc, mh)
	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi khi duyệt dữ liệu"})
		fmt.Printf("Lỗi rows error\n")
		return
	}
	if len(monhoc) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Không có môn học nào"})
		fmt.Printf("Lỗi không có môn học\n")
		return
	}
	c.JSON(http.StatusOK, monhoc)
}

// Tạo giáo viên
func (h *ThaoTac_MonHoc) TaoMonHoc(c *gin.Context) {
	var monhoc models.MonHoc
	if err := c.ShouldBindJSON(&monhoc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Printf("Lỗi lấy thông tin từ server \n")
		return
	}
	var sqlcmd string
	var args []interface{}
	if monhoc.MaMonHoc == "" {
		sqlcmd = "INSERT INTO monhoc (ten_mon) VALUES (?)"
		args = []interface{}{monhoc.TenMonHoc}
	} else {

		sqlcmd = "INSERT INTO monhoc (ma_mon, ten_mon) VALUES (?, ?)"
		args = []interface{}{monhoc.MaMonHoc, monhoc.TenMonHoc}
	}
	result, err := h.DB.Exec(sqlcmd, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể thêm môn học"})
		fmt.Printf("Lỗi SQL \n")
		return
	}
	id, _ := result.LastInsertId()
	monhoc.MaMonHoc = fmt.Sprintf("%d", id)
	c.JSON(http.StatusOK, monhoc)
}

// Cập nhật giáo viên
func (h *ThaoTac_MonHoc) CapNhatMonHoc(c *gin.Context) {
	var monhoc models.MonHoc
	id := c.Param("id")
	row := h.DB.QueryRow("SELECT ma_mon FROM monhoc WHERE ma_mon = ?", id)
	if err := row.Scan(&monhoc.MaMonHoc); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Không tìm thấy thông tin môn học này"})
		return
	}
	if err := c.ShouldBindJSON(&monhoc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	_, err := h.DB.Exec("UPDATE monhoc SET ten_mon = ? WHERE ma_mon = ?", monhoc.TenMonHoc, monhoc.MaMonHoc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể cập nhật môn học"})
		fmt.Printf("Lỗi SQL \n")
		return
	}
	c.JSON(http.StatusOK, monhoc)
}

// Xóa giáo viên khỏi danh sách
func (h *ThaoTac_MonHoc) XoaMonHoc(c *gin.Context) {
	id := c.Param("id")
	fmt.Printf("ID: %s", id)
	var monhoc models.MonHoc
	row := h.DB.QueryRow("SELECT ma_mon, ten_mon FROM monhoc WHERE ma_mon = ?", id)
	if err := row.Scan(&monhoc.MaMonHoc, &monhoc.TenMonHoc); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Không tìm thấy Môn học"})
		return
	}
	_, err := h.DB.Exec("DELETE FROM MonHoc WHERE ma_mon = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Không thể xóa giáo viên"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Xóa giáo viên thành công"})
}

// Thêm vào file hiện tại

func (h *ThaoTac_MonHoc) NhapDanhSachMonHoc(c *gin.Context) {
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

		ma := strings.TrimSpace(row[0])
		ten := strings.TrimSpace(row[1])
		fmt.Printf("Mã Môn học: %s, Tên môn học: %s\n", ma, ten)

		// Kiểm tra dữ liệu bắt buộc
		if ten == "" {
			fmt.Printf("Thiếu tên môn\n")
			continue // Bỏ qua nếu thiếu thông tin bắt buộc
		}

		if ma != "" {
			// Kiểm tra xem giáo viên đã tồn tại chưa
			var exists bool
			err := h.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM monhoc WHERE ma_mon = ?)", ma).Scan(&exists)
			if err != nil {
				fmt.Printf("Lỗi SQL: %v\n", err)
				continue
			}

			if exists {
				// Cập nhật nếu đã tồn tại
				_, err = h.DB.Exec("UPDATE monhoc SET ten_mon = ? WHERE ma_mon = ?", ten, ma)
				if err == nil {
					countUpdated++
				}
			} else {
				// Thêm mới với mã giáo viên được chỉ định
				_, err = h.DB.Exec("INSERT INTO monhoc (ma_mon, ten_mon) VALUES (?, ?)", ma, ten)
				if err == nil {
					countInserted++
				}
			}
		} else {
			// Thêm mới không có mã giáo viên
			_, err = h.DB.Exec("INSERT INTO monhoc (ten_mon) VALUES (?)", ten)
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

func (h *ThaoTac_MonHoc) XuatDanhSachMonHoc(c *gin.Context) {
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
	sheetName := "Danh sách Môn học"
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
	f.SetCellValue(sheetName, "A1", "Mã môn học")
	f.SetCellValue(sheetName, "B1", "Tên môn học")

	// Truy vấn dữ liệu từ database
	rows, err := h.DB.Query("SELECT ma_mon, ten_mon FROM monhoc")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể truy vấn dữ liệu"})
		return
	}
	defer rows.Close()

	// Ghi dữ liệu vào file Excel
	rowIndex := 2
	for rows.Next() {
		var id int
		var ten string
		if err := rows.Scan(&id, &ten); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", rowIndex), id)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", rowIndex), ten)
		rowIndex++
	}

	f.SetActiveSheet(index)

	// Thiết lập header cho response
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=danh_sach_môn_hoc.xlsx")

	// Ghi file Excel vào response
	if err := f.Write(c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
