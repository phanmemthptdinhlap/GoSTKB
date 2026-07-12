package main

import (
    "fmt"
    "net/http"
	. "GoSTKB/libsql"
)

type WebPage struct {
  Title   string 
	mux     *http.ServeMux
}
func (p *WebPage) SetStaticFile() {
    ps := http.FileServer(http.Dir("./static"))    
    // Dùng Handle thay vì HandleFunc
    p.mux.Handle("/static/", http.StripPrefix("/static/", ps))
}
// ... (Giữ nguyên các hàm SetStaticFile và init của WebPage) ...

var page WebPage
var db *SqlTKB // Biến toàn cục

// Khởi tạo các thông số tĩnh nhẹ nhàng

func main() {
	// Dùng dấu = (không dùng :=) để gán vào biến toàn cục db
	var err error
	db, err = ConnectSTKB()
	if err != nil {
		fmt.Println("Lỗi kết nối CSDL:", err)
		return
	}
	
	// Đặt defer ở hàm main để giữ CSDL sống suốt quá trình chạy server
	defer func() {
		if err := db.Close(); err != nil {
			fmt.Println("Lỗi đóng CSDL:", err)
		}
	}()
		
	// Chạy server
		page.Title = "Trang chủ của tôi"
    mux := http.NewServeMux()
	  page.mux = mux
		page.SetStaticFile()
		page.SetPageTrangChu()
		page.SetPageMonHoc()
		page.SetPageLopHoc()
		page.SetPageGiaoVien()
		page.SetPagePhanCong()
    fmt.Println("Server đang chạy tại http://localhost:8080")
    http.ListenAndServe(":8080", mux)
}
