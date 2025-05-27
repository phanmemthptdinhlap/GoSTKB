package main

import (
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	// Import package libs
)

func main() {
	// Phục vụ các tệp tĩnh
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// Phục vụ các tệp HTML
	http.HandleFunc("/", ShowHome)
	// Định nghĩa các handler
	// Khởi động server
	http.ListenAndServe(":8080", nil)
}
func ShowHome(w http.ResponseWriter, r *http.Request) {
	// Hiển thị trang chủ
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
}
func AddGiaovien(w http.ResponseWriter, r *http.Request) {
	// Hiển thị trang thêm giáo viên
	tmpl := template.Must(template.ParseFiles("templates/addgiaovien.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
}
