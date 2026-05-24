package main

import (
	"html/template"
	"net/http"
)
func init() {
	page.SetPageClass=func() {
		page.mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
			tmpl, err := template.ParseFiles("templates/about.html","templates/base.html")
			if err != nil {
				panic(err)
				http.Error(w, "Lỗi tải tệp about.html", http.StatusInternalServerError)
				return
			}
			data := struct {
				Title string
			}{
				Title: "Giới thiệu",
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
