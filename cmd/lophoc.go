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
		lophoc,err := db.SelectAllLopHoc()
		if err != nil {
			fmt.Println("Lỗi lấy danh sách: ", err)
			http.Error(w, "Lỗi lấy danh sách: "+err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(lophoc)
	})

	// === API MỚI: ĐỒNG BỘ DỮ LIỆU HÀNG LOẠT ===
	p.mux.HandleFunc("POST /api/lophoc/sync", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var danhSachDongBo []	LopHoc
		
		err := json.NewDecoder(r.Body).Decode(&danhSachDongBo)
		if err != nil {
			http.Error(w, "Lỗi decode: "+err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Printf("Nhận được %d bản ghi cần đồng bộ\n", len(danhSachDongBo))
		var Insert []LopHoc
		var Update []LopHoc
		var Delete []int	

		// Lấy danh sách các giao vien đã có trong DB

		// Phân loại và xử lý từng hành động
		for _, lophoc := range danhSachDongBo {
			switch lophoc.Action {
			case "thêm":
				Insert = append(Insert, lophoc)
			case "sửa":
				Update = append(Update, lophoc)
			case "xóa":
				Delete = append(Delete, lophoc.ID)
			}
		}
		db.InsertLopHoc(Insert)
		db.EditLopHoc(Update)
		db.DeleteLopHoc(Delete)

		// Trả về thành công
		json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Đồng bộ hoàn tất"})
	})

	// Các API cũ giữ nguyên phục vụ cho CRUD lẻ (nếu cần)
	p.mux.HandleFunc("POST /api/lophoc", func(w http.ResponseWriter, r *http.Request) { /* ... */ })
	p.mux.HandleFunc("PUT /api/lophoc", func(w http.ResponseWriter, r *http.Request) { /* ... */ })
	p.mux.HandleFunc("DELETE /api/lophoc/{ma}", func(w http.ResponseWriter, r *http.Request) { /* ... */ })
}
