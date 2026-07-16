package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Bổ sung trường action để mapping với giao diện Vue


func (p *WebPage) SetPageTest() {
	p.mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/test.html", "templates/base.html")
		if err != nil {
			fmt.Println("Lỗi parse template: ", err)
			http.Error(w, "Lỗi parse template: "+err.Error(), http.StatusInternalServerError)
			return
		}

		data := struct{ Title string }{Title: "Test - Website của tôi"}

		err = tmpl.ExecuteTemplate(w, "base", data)
		if err != nil {
			fmt.Println("Lỗi render: ", err)
			http.Error(w, "Lỗi render: "+err.Error(), http.StatusInternalServerError)
		}
	})

}
