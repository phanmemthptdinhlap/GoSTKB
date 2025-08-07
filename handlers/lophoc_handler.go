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

// Danh sách lớp học
type ThaoTac_LopHoc struct {
	DB *sql.DB
}

// Lấy danh sách lớp học
func (h *ThaoTac_LopHoc) DanhSachLopHoc(c *gin.Context) {
	var lophoc []models.LopHoc

	rows, err := h.DB.Query(`SELECT l.ma_lop, l.ten_lop, l.khoi_lop, l.ma_chu_nhiem, g.ho_ten
						FROM lophoc l
						RIGHT OUTER JOIN giaovien g 
						ON l.ma_chu_nhiem=g.ma_giao_vien`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể truy vấn dữ liệu"})
		fmt.Printf("Lỗi phần sql\n")
		return
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		fmt.Printf("Lỗi lấy danh sách cột \n")
		return
	}
	fmt.Printf("Columns: %v", columns)
	var ma, ten, khoi, macn, hoten string
	for rows.Next() {
		if err := rows.Scan(&ma, &ten, &khoi, &macn, &hoten); err != nil {
			fmt.Printf("Lỗi đọc dữ liệu\n")
			return
		}
	}

	/*for rows.Next() {
		var lh models.LopHoc
		if err := rows.Scan(&lh.MaLop, &lh.TenLop, &lh.KhoiLop, &lh.MaChuNhiem, &lh.TenChuNhiem); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi khi đọc dữ liệu"})
			fmt.Printf("Lỗi phần đọc dữ liệu\n")
			return
		}
		lophoc = append(lophoc, lh)
	}*/

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi khi duyệt dữ liệu"})
		fmt.Printf("Lỗi phần duyệt dữ liệu\n")
		return
	}

	if len(lophoc) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Không có lớp học nào"})
		return
	}

	c.JSON(http.StatusOK, lophoc)
}

// Tạo lớp học mới
func (h *ThaoTac_LopHoc) TaoLopHoc(c *gin.Context) {
	var lophoc models.LopHoc
	if err := c.ShouldBindJSON(&lophoc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var sqlcmd string
	var args []interface{}

	if lophoc.MaLop == "" {
		sqlcmd = "INSERT INTO lophoc (ten_lop, khoi_lop, ma_chu_nhiem) VALUES (?, ?,?)"
		args = []interface{}{lophoc.TenLop, lophoc.KhoiLop, lophoc.MaChuNhiem}
	} else {
		sqlcmd = "INSERT INTO lophoc (ma_lop, ten_lop, khoi_lop, ma_chu_nhiem) VALUES (?, ?, ?, ?)"
		args = []interface{}{lophoc.MaLop, lophoc.TenLop, lophoc.KhoiLop, lophoc.MaChuNhiem}
	}

	result, err := h.DB.Exec(sqlcmd, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể thêm lớp học"})
		return
	}

	id, _ := result.LastInsertId()
	lophoc.MaLop = fmt.Sprintf("%d", id)
	c.JSON(http.StatusOK, lophoc)
}

// Cập nhật thông tin lớp học
func (h *ThaoTac_LopHoc) CapNhatLopHoc(c *gin.Context) {
	var lophoc models.LopHoc
	id := c.Param("id")

	// Kiểm tra lớp học tồn tại
	row := h.DB.QueryRow("SELECT ma_lop FROM lophoc WHERE ma_lop = ?", id)
	if err := row.Scan(&lophoc.MaLop); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy thông tin lớp học này"})
		return
	}

	// Đọc dữ liệu mới từ request
	if err := c.ShouldBindJSON(&lophoc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cập nhật thông tin
	_, err := h.DB.Exec("UPDATE lophoc SET ten_lop = ?, khoi_lop = ?, ma_chu_nhiem=? WHERE ma_lop = ?",
		lophoc.TenLop, lophoc.KhoiLop, lophoc.MaChuNhiem, lophoc.MaLop)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể cập nhật lớp học"})
		return
	}

	c.JSON(http.StatusOK, lophoc)
}

// Xóa lớp học
func (h *ThaoTac_LopHoc) XoaLopHoc(c *gin.Context) {
	id := c.Param("id")

	// Kiểm tra lớp học tồn tại
	var lophoc models.LopHoc
	row := h.DB.QueryRow("SELECT ma_lop FROM lophoc WHERE ma_lop = ?", id)
	if err := row.Scan(&lophoc.MaLop); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy thông tin lớp học"})
		return
	}

	// Xóa lớp học
	_, err := h.DB.Exec("DELETE FROM lophoc WHERE ma_lop= ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể xóa lớp học"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Xóa lớp học thành công"})
}

// Xuất danh sách lớp học ra Excel
func (h *ThaoTac_LopHoc) XuatDanhSachLopHoc(c *gin.Context) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}()

	// Lấy tên sheet mặc định
	sheetName := "Danh sách lớp học"
	sheetDefault := f.GetSheetName(0)

	// Đổi tên sheet
	if err := f.SetSheetName(sheetDefault, sheetName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể tạo sheet"})
		return
	}

	// Đặt tiêu đề các cột
	f.SetCellValue(sheetName, "A1", "Mã lớp")
	f.SetCellValue(sheetName, "B1", "Tên lớp")
	f.SetCellValue(sheetName, "C1", "Khối lớp")
	f.SetCellValue(sheetName, "D1", "Mã Chủ Nhiệm")
	f.SetCellValue(sheetName, "E1", "Tên Chủ Nhiệm")

	// Truy vấn dữ liệu
	sqlcmd := `SELECT l.ma_lop, l.ten_lop, l.khoi_lop, l.ma_chu_nhiem, g.ho_ten 
			FROM lophoc l
			INNER JOIN giaovien g on l.ma_chu_nhiem = g.ma_giao_gien`
	rows, err := h.DB.Query(sqlcmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể truy vấn dữ liệu"})
		return
	}
	defer rows.Close()

	// Ghi dữ liệu vào file Excel
	rowIndex := 2
	for rows.Next() {
		var malop int
		var tenlop, khoilop string
		if err := rows.Scan(&malop, &tenlop, &khoilop); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		f.SetCellValue(sheetName, fmt.Sprintf("A%d", rowIndex), malop)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", rowIndex), tenlop)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", rowIndex), khoilop)
		rowIndex++
	}

	// Thiết lập header cho response
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=danh_sach_lop_hoc.xlsx")

	// Ghi file Excel vào response
	if err := f.Write(c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

// Nhập danh sách lớp học từ Excel
func (h *ThaoTac_LopHoc) NhapDanhSachLopHoc(c *gin.Context) {
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

	// Đọc file Excel
	xlsx, err := excelize.OpenReader(f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể đọc file Excel"})
		return
	}
	defer xlsx.Close()

	// Lấy tên sheet đầu tiên
	sheetName := xlsx.GetSheetName(0)

	// Khởi tạo biến đếm
	var countUpdated, countInserted int

	// Đọc từng dòng trong file Excel
	rows, err := xlsx.GetRows(sheetName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể đọc dữ liệu từ sheet"})
		return
	}

	// Bỏ qua dòng tiêu đề
	for i := 1; i < len(rows); i++ {
		row := rows[i]
		if len(row) < 4 {
			continue // Bỏ qua dòng không đủ dữ liệu
		}

		maLop := strings.TrimSpace(row[0])
		tenLop := strings.TrimSpace(row[1])
		khoiLop := strings.TrimSpace(row[2])
		maChuNhiem := strings.TrimSpace(row[3])

		// Kiểm tra dữ liệu bắt buộc
		if tenLop == "" || khoiLop == "" {
			continue
		}

		if maLop != "" {
			// Kiểm tra xem lớp học đã tồn tại chưa
			var exists bool
			err := h.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM lophoc WHERE ma_lop = ?)", maLop).Scan(&exists)
			if err != nil {
				continue
			}

			if exists {
				// Cập nhật nếu đã tồn tại
				_, err = h.DB.Exec("UPDATE lophoc SET ten_lop = ?, khoi_lop = ?,ma_chu_nhiem = ? WHERE ma_lop = ?",
					tenLop, khoiLop, maChuNhiem, maLop)
				if err == nil {
					countUpdated++
				}
			} else {
				// Thêm mới với mã lớp được chỉ định
				_, err = h.DB.Exec("INSERT INTO lophoc (ma_lop, ten_lop, khoi_lop, ma_chu_nhiem) VALUES (?, ?, ?, ?)",
					maLop, tenLop, khoiLop, maChuNhiem)
				if err == nil {
					countInserted++
				}
			}
		} else {
			// Thêm mới không có mã lớp
			_, err = h.DB.Exec("INSERT INTO lophoc (ten_lop, khoi_lop, ma_chu_nhiem) VALUES (?, ?,?)",
				tenLop, khoiLop, maChuNhiem)
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
