package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	. "GoSTKB/libsql"
)

// Bổ sung trường action để mapping với giao diện Vue


func (p *WebPage) SetPageLopHoc() {
	p.mux.HandleFunc("/lophoc", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/lophoc.html", "templates/base.html")
		if err != nil {
			fmt.Println("Lỗi parse template: ", err)
			http.Error(w, "Lỗi parse template: "+err.Error(), http.StatusInternalServerError)
			return
		}

		data := struct{ Title string }{Title: "Lớp học - Website của tôi"}

		err = tmpl.ExecuteTemplate(w, "base", data)
		if err != nil {
			fmt.Println("Lỗi render: ", err)
			http.Error(w, "Lỗi render: "+err.Error(), http.StatusInternalServerError)
		}
	})

	// API Lấy danh sách
	p.mux.HandleFunc("GET /api/lophoc", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		lops,err := db.SelectAllLopHoc()
		if err != nil {
			fmt.Println("Lỗi lấy danh sách: ", err)
			http.Error(w, "Lỗi lấy danh sách: "+err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(lops)
	})

	// === API MỚI: ĐỒNG BỘ DỮ LIỆU HÀNG LOẠT ===
	p.mux.HandleFunc("POST /api/lophoc/sync", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var danhSachDongBo []LopHoc
		
		err := json.NewDecoder(r.Body).Decode(&danhSachDongBo)
		if err != nil {
			http.Error(w, "Lỗi decode: "+err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Printf("Nhận được %d bản ghi cần đồng bộ\n", len(danhSachDongBo))

		// Phân loại và xử lý từng hành động
		for _, lop := range danhSachDongBo {
			switch lop.Action {
			case "thêm":
				db.InsertLopHoc(lop)
			case "sửa":
				db.EditLopHoc(lop)
			case "xóa":
				db.DeleteLopHoc(lop.ID)
			}
		}

		// Trả về thành công
		json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Đồng bộ hoàn tất"})
	})

	// Các API cũ giữ nguyên phục vụ cho CRUD lẻ (nếu cần)
	p.mux.HandleFunc("POST /api/lophoc", func(w http.ResponseWriter, r *http.Request) { /* ... */ })
	p.mux.HandleFunc("PUT /api/lophoc", func(w http.ResponseWriter, r *http.Request) { /* ... */ })
	p.mux.HandleFunc("DELETE /api/lophoc/{ma}", func(w http.ResponseWriter, r *http.Request) { /* ... */ })
}
