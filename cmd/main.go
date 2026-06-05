package main

import (
    "fmt"
    "net/http"
)

type WebPage struct {
    Title   		string
		mux     		*http.ServeMux
}
func (p *WebPage) SetStaticFile() {
		ps:=http.FileServer(http.Dir("./static"))
    p.mux.Handle("/static/", http.StripPrefix("/static/", ps))
}

func (p *WebPage) init(mux *http.ServeMux) {
    p.mux = mux
		p.SetPageTrangChu()
		p.SetPageLopHoc()
		p.SetPageGiaoVien()
		p.SetPageMonHoc()
		p.SetPagePhanCong()
		p.SetPageThongTin()
		p.SetStaticFile()
	}

// Khai báo biến toàn cục
var page WebPage

func main() {
    mux := http.NewServeMux()
		page.init(mux)
    fmt.Println("Server đang chạy tại http://localhost:8080")
    http.ListenAndServe(":8080", mux)
}
