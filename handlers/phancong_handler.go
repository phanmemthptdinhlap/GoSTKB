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
type ThaoTac_PhanCong struct {
	DB *sql.DB
}

// Lấy danh sách lớp học
func (h *ThaoTac_PhanCong) DanhSachPhanCong(c *gin.Context) {
	var phancong []models.PhanCong

	rows, err := h.DB.Query(`SELECT pc.ma_phan_cong, pc.ma_giao_vien, pc.ma_mon, pc.ma_lop, 
						COALESCE(gv.ho_ten,''), COALESCE(mh.ten_mon,''), COALESCE(lh.ten_lop,'')
						FROM phancong pc
						LEFT OUTER JOIN giaovien gv ON pc.ma_giao_vien = gv.ma_giao_vien
						LEFT OUTER JOIN monhoc mh ON pc.ma_mon=mh.ma_mon
						LEFT OUTER JOIN lophoc lh ON pc.ma_lop=lh.ma_lop`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể truy vấn dữ liệu"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var pc models.PhanCong
		if err := rows.Scan(&pc.MaPhanCong, &pc.MaGiaoVien, &pc.MaMon, &pc.MaLop,
			&pc.TenGiaoVien, &pc.TenMon, &pc.TenLop); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi khi đọc dữ liệu"})
			return
		}
		phancong = append(phancong, pc)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi khi duyệt dữ liệu"})
		fmt.Printf("Lỗi phần duyệt dữ liệu\n")
		return
	}

	if len(phancong) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Không có lớp học nào"})
		return
	}

	c.JSON(http.StatusOK, phancong)
}

// Tạo lớp học mới
func (h *ThaoTac_PhanCong) TaoPhanCong(c *gin.Context) {
	var phancong models.PhanCong
	if err := c.ShouldBindJSON(&phancong); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var sqlcmd string
	var args []interface{}

	if phancong.MaPhanCong == "" {
		sqlcmd = "INSERT INTO phancong (ma_giao_vien, ma_mon, ma_Lop) VALUES (?,?,?)"
		args = []interface{}{phancong.MaGiaoVien, phancong.MaMon, phancong.MaLop}
	} else {
		sqlcmd = "INSERT INTO phancong (ma_phan_cong, ma_giao_vien, ma_mon, ma_Lop) VALUES (?,?,?,?)"
		args = []interface{}{phancong.MaPhanCong, phancong.MaGiaoVien, phancong.MaMon, phancong.MaLop}
	}

	result, err := h.DB.Exec(sqlcmd, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể thêm lớp học"})
		return
	}

	id, _ := result.LastInsertId()
	phancong.MaPhanCong = fmt.Sprintf("%d", id)
	c.JSON(http.StatusOK, phancong)
}

// Cập nhật thông tin phân công
func (h *ThaoTac_PhanCong) CapNhatPhanCong(c *gin.Context) {
	var phancong models.PhanCong
	id := c.Param("id")

	// Kiểm tra lớp học tồn tại
	var exits bool
	row := h.DB.QueryRow("SELECT EXITS(SELECT 1 FROM PhanCong WHERE ma_phan_cong = ?)", id)
	if err := row.Scan(&exits); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lỗi truy cập CSDL"})
		return
	}
	if !exits {
		c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thầy mã phân công"})
		return
	}

	// Đọc dữ liệu mới từ request
	if err := c.ShouldBindJSON(&phancong); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "lỗi lấy dữ liệu từ form"})
		return
	}

	// Cập nhật thông tin
	_, err := h.DB.Exec("UPDATE phancong SET ma_giao_vien = ?, ma_mon = ?, ma_lop=? WHERE ma_phan_cong = ?",
		phancong.MaGiaoVien, phancong.MaMon, phancong.MaLop, phancong.MaPhanCong)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể cập nhật lớp học"})
		return
	}
	c.JSON(http.StatusOK, phancong)
}

// Xóa lớp học
func (h *ThaoTac_PhanCong) XoaPhanCong(c *gin.Context) {
	id := c.Param("id")

	// Kiểm tra lớp học tồn tại
	var exits bool
	sqler := h.DB.QueryRow("SELECT EXITS(SELECT 1 FROM phancong WHERE ma_phan_cong = ?)", id).Scan(&exits)
	if sqler != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi truy cập CSDL"})
		return
	}
	if !exits {
		c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thầy mã phân công"})
		return
	}

	// Xóa phân công
	_, err := h.DB.Exec("DELETE FROM phancong WHERE ma_phan_cong= ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể xóa phân công"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Xóa phân công thành công"})
}

// Xuất danh sách lớp học ra Excel
func (h *ThaoTac_PhanCong) XuatDanhSachPhanCong(c *gin.Context) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi tạo file"})
			return
		}
	}()

	// Lấy tên sheet mặc định
	sheetName := "Danh sách phân công"
	sheetDefault := f.GetSheetName(0)

	// Đổi tên sheet
	if err := f.SetSheetName(sheetDefault, sheetName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể tạo sheet"})
		return
	}

	// Đặt tiêu đề các cột
	f.SetCellValue(sheetName, "A1", "Mã phân công")
	f.SetCellValue(sheetName, "B1", "Mã giáo viên")
	f.SetCellValue(sheetName, "C1", "Tên giáo viên")
	f.SetCellValue(sheetName, "D1", "Mã môn học")
	f.SetCellValue(sheetName, "E1", "Tên môn học")
	f.SetCellValue(sheetName, "F1", "ma lớp học")
	f.SetCellValue(sheetName, "G1", "Tên lớp học")

	// Truy vấn dữ liệu
	rows, err := h.DB.Query(`SELECT pc.ma_phan_cong, pc.ma_giao_vien, pc.ma_mon, pc.ma_lop, 
						COALESCE(gv.ho_ten,''), COALESCE(mh.ten_mon,''), COALESCE(lh.ten_lop,'')
						FROM phancong pc
						LEFT OUTER JOIN giaovien gv ON pc.ma_giao_vien = gv.ma_giao_vien
						LEFT OUTER JOIN monhoc mh ON pc.ma_mon=mh.ma_mon
						LEFT OUTER JOIN lophoc lh ON pc.ma_lop=lh.ma_lop`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể truy vấn CSDL"})
		return
	}
	defer rows.Close()

	// Ghi dữ liệu vào file Excel
	rowIndex := 2
	for rows.Next() {
		var ma_phan_cong, ma_giao_vien, ten_giao_vien, ma_mon, ten_mon, ma_lop, ten_lop string
		if err := rows.Scan(&ma_phan_cong, &ma_giao_vien, &ma_mon, &ma_lop, &ten_giao_vien, &ten_mon, &ten_lop); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Quét dữ liệu không thành công"})
			return
		}

		f.SetCellValue(sheetName, fmt.Sprintf("A%d", rowIndex), ma_phan_cong)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", rowIndex), ma_giao_vien)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", rowIndex), ten_giao_vien)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", rowIndex), ma_mon)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", rowIndex), ten_mon)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", rowIndex), ma_lop)
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", rowIndex), ten_lop)
		rowIndex++
	}

	// Thiết lập header cho response
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=danh_sach_lop_hoc.xlsx")

	// Ghi file Excel vào response
	if err := f.Write(c.Writer); err != nil {
		fmt.Printf("Lỗi ghi file\n")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

// Nhập danh sách lớp học từ Excel
func (h *ThaoTac_PhanCong) NhapDanhSachPhanCong(c *gin.Context) {
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

		ma_phan_cong := strings.TrimSpace(row[0])
		ma_giao_Vien := strings.TrimSpace(row[2])
		ma_mon := strings.TrimSpace(row[4])
		ma_lop := strings.TrimSpace(row[6])

		// Kiểm tra dữ liệu bắt buộc
		if ma_giao_Vien == "" || ma_mon == "" || ma_lop == "" {
			continue
		}

		if ma_phan_cong != "" {
			// Kiểm tra xem lớp học đã tồn tại chưa
			var exists bool
			err := h.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM phancong WHERE ma_phan_cong = ?)", ma_phan_cong).Scan(&exists)
			if err != nil {
				continue
			}

			if exists {
				// Cập nhật nếu đã tồn tại
				_, err = h.DB.Exec("UPDATE phancong SET ma_giao_vien = ?, ma_mon = ?,ma_lop = ? WHERE ma_phan_cong = ?",
					ma_giao_Vien, ma_mon, ma_lop, ma_phan_cong)
				if err == nil {
					countUpdated++
				}
			} else {
				// Thêm mới với mã lớp được chỉ định
				_, err = h.DB.Exec("INSERT INTO phancong (ma_phan_cong, ma_giao_vien, ma_mon, ma_lop) VALUES (?, ?, ?, ?)",
					ma_phan_cong, ma_giao_Vien, ma_mon, ma_lop)
				if err == nil {
					countInserted++
				}
			}
		} else {
			// Thêm mới không có mã lớp
			_, err = h.DB.Exec("INSERT INTO phancong (ma_giao_vien, ma_mon, ma_lop) VALUES (?, ?, ?)",
				ma_giao_Vien, ma_mon, ma_lop)
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
