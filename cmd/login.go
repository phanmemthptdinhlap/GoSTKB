package main

import (
	"html/template"
	"net/http"
)
func init() {
	page.SetPageClass=func() {
		page.mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
			tmpl, err := template.ParseFiles("templates/login.html","templates/base.html")
			if err != nil {
				panic(err)
				http.Error(w, "Lỗi tải tệp login.html", http.StatusInternalServerError)
				return
			}
			data := struct {
				Title string
			}{
				Title: "Đăng nhập",
			}
			err = tmpl.Execute(w, data)
			if err != nil {
				panic(err)
				http.Error(w, "Lỗi thực hiện tạo trang", http.StatusInternalServerError)
				return
			}
		})
	}
}
