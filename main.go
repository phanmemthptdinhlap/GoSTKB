package main

import (
	"GoSTKB/handlers"
	"fmt"
	"log"
	"net/http"
	"os"

	// Adjust the import path as necessary
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//Kết nối CSDL
	db, err := sql.Open("sqlite3", "tkb.db")
	if err != nil {
		log.Fatalf("Không kết nối được với cơ sở dữ liệu: %v", err)
	}
	defer db.Close()
	//Khởi tạo cấu trúc CSDL
	sqlcmd, err := os.ReadFile("database.sql")
	if err != nil {
		log.Fatalf("Lỗi đọc file cấu trúc CSDL: %v", err)
	}
	_, err = db.Exec(string(sqlcmd))
	if err != nil {
		log.Fatalf("Lỗi khởi tạo CSDL: %v", err)
	}
	//Khởi tạo trình quản lý truy vấn web
	r := gin.Default()
	//Tải template
	r.LoadHTMLFiles(
		"templates/components/footer.html",
		"templates/components/header.html",
		"templates/index.html",
		"templates/giaovien.html",
		"templates/lophoc.html",
		"templates/monhoc.html",
		"templates/phancong.html",
		"templates/test.html",
	)
	//Cấu hình file tĩnh
	r.Static("static", "./static")
	r.StaticFile("/favicon.ico", "static/images/favicon.ico")
	//khỏi tạo các thao tác
	thaotacgiaovien := &handlers.ThaoTac_GiaoVien{DB: db}
	thaotaclophoc := &handlers.ThaoTac_LopHoc{DB: db}
	thaotacmonhoc := &handlers.ThaoTac_MonHoc{DB: db}
	thaotacphancong := &handlers.ThaoTac_PhanCong{DB: db}

	//Điều phối truy vấn trang HTML
	r.GET("/", func(c *gin.Context) {
		fmt.Println("readering index.html")
		c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title": "Trang chủ",
		})
	})

	r.GET("/giaovien", func(c *gin.Context) {
		fmt.Println("readering giaovien.html")
		c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		c.HTML(http.StatusOK, "giaovien.html", gin.H{
			"Title": "Quản lý giáo viên",
		})
	})

	r.GET("/lophoc", func(c *gin.Context) {
		c.HTML(http.StatusOK, "lophoc.html", gin.H{
			"Title": "Quản lý lớp học",
		})
	})

	r.GET("/monhoc", func(c *gin.Context) {
		c.HTML(http.StatusOK, "monhoc.html", gin.H{
			"Title": "Quản lý môn học",
		})
	})
	r.GET("/phancong", func(c *gin.Context) {
		c.HTML(http.StatusOK, "phancong.html", gin.H{
			"Title": "Quản lý phân công giảng dạy",
		})
	})
	r.GET("/test", func(c *gin.Context) {
		c.HTML(http.StatusOK, "test.html", gin.H{
			"Title": "Test trang web",
		})
	})
	//Điều phối cổng dịch vụ
	api := r.Group("/api")

	{
		//Giao viên
		api.POST("/giaovien", thaotacgiaovien.TaoGiaoVien)
		api.GET("/giaovien", thaotacgiaovien.DanhSachGiaoVien)
		api.PUT("/giaovien/:id", thaotacgiaovien.CapNhatGiaoVien)
		api.DELETE("/giaovien/:id", thaotacgiaovien.XoaGiaoVien)
		api.GET("/export/giaovien", thaotacgiaovien.XuatDanhSachGiaoVien)
		api.POST("/import/giaovien", thaotacgiaovien.NhapDanhSachGiaoVien)
		//Lớp học
		api.POST("/lophoc", thaotaclophoc.TaoLopHoc)
		api.GET("/lophoc", thaotaclophoc.DanhSachLopHoc)
		api.PUT("/lophoc/:id", thaotaclophoc.CapNhatLopHoc)
		api.DELETE("/lophoc/:id", thaotaclophoc.XoaLopHoc)
		api.GET("/export/lophoc", thaotaclophoc.XuatDanhSachLopHoc)
		api.POST("/import/lophoc", thaotaclophoc.NhapDanhSachLopHoc)
		//Môn học
		api.POST("/monhoc", thaotacmonhoc.TaoMonHoc)
		api.GET("/monhoc", thaotacmonhoc.DanhSachMonHoc)
		api.PUT("/monhoc/:id", thaotacmonhoc.CapNhatMonHoc)
		api.DELETE("/monhoc/:id", thaotacmonhoc.XoaMonHoc)
		api.GET("/export/monhoc", thaotacmonhoc.XuatDanhSachMonHoc)
		api.POST("/import/monhoc", thaotacmonhoc.NhapDanhSachMonHoc)
		//Phân công
		api.POST("/phancong", thaotacphancong.TaoPhanCong)
		api.GET("/phancong", thaotacphancong.DanhSachPhanCong)
		api.PUT("/phancong/:id", thaotacphancong.CapNhatPhanCong)
		api.DELETE("/phancong/:id", thaotacphancong.XoaPhanCong)
		api.GET("/export/phancong", thaotacphancong.XuatDanhSachPhanCong)
		api.POST("/import/phancong", thaotacphancong.NhapDanhSachPhanCong)
	api.GET("/phancong/maphancong/:lop/:mon", thaotacphancong.LayMaPhanCongTheoLop_Mon)

	}

	r.Run(":8080")
}
