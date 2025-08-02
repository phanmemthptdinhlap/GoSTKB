package main

import (
	"GoSTKB/handlers"
	"log"
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
	r.LoadHTMLGlob("templates/**/*.html")
	//Cấu hình file tĩnh
	r.Static("static", "./static")
	//khỏi tạo các thao tác
	thaotacgiaovien := &handlers.ThaoTac_GiaoVien{DB: db}
	thaotaclophoc := &handlers.ThaoTac_LopHoc{DB: db}

	//Điều phối truy vấn trang HTML
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	r.GET("/giaovien", func(c *gin.Context) {
		c.HTML(200, "giaovien.html", nil)
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
		api.GET("/lophoc", thaotacglophoc.DanhSachLopHoc)
		api.PUT("/lophoc/:id", thaotacglophoc.CapNhatLopHoc)
		api.DELETE("/lophoc/:id", thaotacglophoc.XoaLopHoc)
		api.GET("/export/lophoc", thaotacglophoc.XuatDanhSachLopHoc)
		api.POST("/import/lophoc", thaotacglophoc.NhapDanhSachLopHoc)
		//Học sinh

	}

	r.Run(":8080")
}
