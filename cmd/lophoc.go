package main

import (
		"fmt"
		"encoding/json"
    "html/template" // Nhớ import thư viện này, KHÔNG dùng text/template
    "net/http"
)
type LopHoc struct {
	Ma string `json:"ma"`
	Ten string `json:"ten"`
	Gvcn string `json:"gvcn"`
}
var lop_test = []LopHoc{
	{Ma: "1", Ten: "10A1", Gvcn: "Mai"},
	{Ma: "2", Ten: "12A1", Gvcn: "DuyenA"},
	{Ma: "3", Ten: "12A2", Gvcn: "DuyenB"},
}
func getLopHoc() []LopHoc {
	return lop_test
}
func addLopHoc(lop LopHoc) []LopHoc {
	lop_test = append(lop_test, lop)
	return lop_test
}
func deleteLopHoc(ma string) []LopHoc {
	for i, v := range lop_test {
		if v.Ma == ma {
			lop_test = append(lop_test[:i], lop_test[i+1:]...)
		}
	}
	fmt.Println(lop_test)
	return lop_test
}
func updateLopHoc(lop LopHoc) []LopHoc {
	for i, v := range lop_test {
		if v.Ma == lop.Ma {
			lop_test[i] = lop
		}
	}
	return lop_test
}
func (p *WebPage) SetPageLopHoc() {
		p.mux.HandleFunc("/lophoc", func(w http.ResponseWriter, r *http.Request) {
			// 1. Đọc cả file base và file home
			// Lưu ý: Đường dẫn tính từ nơi bạn chạy lệnh "go run"
			tmpl, err := template.ParseFiles("templates/lophoc.html", "templates/base.html")
			if err != nil {
				panic(err)
				http.Error(w, "Lỗi tải giao diện: "+err.Error(), http.StatusInternalServerError)
				return
			}

			// 2. Chuẩn bị dữ liệu muốn truyền ra HTML
			data := struct {
				Title string
			}{
				Title: "Lớp học - Website của tôi",
			}

			// 3. Thực thi template có tên là "base" và truyền data vào
			err = tmpl.ExecuteTemplate(w, "base", data)
			if err != nil {
				panic(err)
				http.Error(w, "Lỗi render: "+err.Error(), http.StatusInternalServerError)
			}
		})
		p.mux.HandleFunc("GET /api/lophoc", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(getLopHoc())
		})
		p.mux.HandleFunc("POST /api/lophoc", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			var lop LopHoc
			err:= json.NewDecoder(r.Body).Decode(&lop)
			if err != nil {
				fmt.Println("Lối", err)
				http.Error(w, "Lỗi decode: "+err.Error(), http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(addLopHoc(lop))
		})
		p.mux.HandleFunc("DELETE /api/lophoc/{ma}", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			ma:=r.PathValue("ma")
			json.NewEncoder(w).Encode(deleteLopHoc(ma))
		})
}
