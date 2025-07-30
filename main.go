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
	/*
		// Kết nối CSDL
		db, err := gorm.Open(sqlite.Open("stkb.db"), &gorm.Config{})
		if err != nil {
			panic("Không kết nối được với cơ sở dữ liệu")
		}

		// Khởi tạo cấu trúc CSDL
		if err := db.AutoMigrate(&models.GiaoVien{},
			models.LopHoc{},
			models.MonHoc{},
			models.PhanCong{},
			models.TietDay{}); err != nil {
			panic("Lỗi khởi tạo CSDL")
		}*/
	//Khởi tạo trình quản lý truy vấn web
	r := gin.Default()
	//Tải template
	r.LoadHTMLGlob("templates/*")
	//Cấu hình file tĩnh
	r.Static("static", "./static")
	//khỏi tạo các thao tác
	thaotacgiaovien := &handlers.ThaoTac_GiaoVien{DB: db}

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
		//Lớp học

	}

	r.Run(":8080")
}
