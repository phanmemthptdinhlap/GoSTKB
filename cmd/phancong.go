package main

import (
    "html/template" // Nhớ import thư viện này, KHÔNG dùng text/template
    "net/http"
)
func (p *WebPage) SetPagePhanCong() {
		p.mux.HandleFunc("/assignment", func(w http.ResponseWriter, r *http.Request) {
			// 1. Đọc cả file base và file home
			// Lưu ý: Đường dẫn tính từ nơi bạn chạy lệnh "go run"
			tmpl, err := template.ParseFiles("templates/base.html", "templates/assignment.html")
			if err != nil {
				panic(err)
				http.Error(w, "Lỗi tải giao diện: "+err.Error(), http.StatusInternalServerError)
				return
			}

			// 2. Chuẩn bị dữ liệu muốn truyền ra HTML
			data := struct {
				Title string
			}{
				Title: "Phân phối - Website của tôi",
			}

			// 3. Thực thi template có tên là "base" và truyền data vào
			err = tmpl.ExecuteTemplate(w, "base", data)
			if err != nil {
				panic(err)
				http.Error(w, "Lỗi render: "+err.Error(), http.StatusInternalServerError)
			}
		})
	}
