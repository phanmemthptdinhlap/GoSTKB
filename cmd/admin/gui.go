package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
)

type MenuItem struct {
	Name string
	URL  string
}

// renderTemplate là hàm để hiển thị template với dữ liệu
type TemplateData struct {
	Title string
	Menu  []MenuItem
	Data  interface{}
}

func renderTemplate(w http.ResponseWriter, r *http.Request, templateName string, Data TemplateData) {
	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/menu.html",
		"templates/"+templateName+".html")
	if err != nil {
		http.Error(w, "Lỗi khi tải template", http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "base", Data)
	if err != nil {
		http.Error(w, "Lỗi khi hiển thị template", http.StatusInternalServerError)
		return
	}
	fmt.Printf("Đã hiển thị template %s với tiêu đề %s\n", templateName, Data.Title)
}

func Gui(db *sql.DB) {
	// Tạo một HTTP server
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", Home())
	http.HandleFunc("/teachers", listgiaovien(db))
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
func Home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tempdata := TemplateData{
			Title: "Trang chủ",
			Menu: []MenuItem{
				{Name: "Trang chủ", URL: "/"},
				{Name: "Danh sách giáo viên", URL: "/teachers"},
				{Name: "Thêm giáo viên", URL: "/add-teacher"},
			},
		}
		renderTemplate(w, r, "home", tempdata)
	}
}

func listgiaovien(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, _ := GetTeachers(db)
		if rows == nil {
			http.Error(w, "không lấy được danh sách giáo viên", http.StatusInternalServerError)
			return
		}
		tempdata := TemplateData{
			Title: "Danh sách giáo viên",
			Menu: []MenuItem{
				{Name: "Trang chủ", URL: "/"},
				{Name: "Danh sách giáo viên", URL: "/teachers"},
				{Name: "Thêm giáo viên", URL: "/add-teacher"},
			},
			Data: rows,
		}
		renderTemplate(w, r, "teachers", tempdata)

	}
}
