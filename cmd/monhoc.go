package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	. "GoSTKB/libsql"
)

// Bổ sung trường action để mapping với giao diện Vue


func (p *WebPage) SetPageMonHoc() {
	p.mux.HandleFunc("/monhoc", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/monhoc.html", "templates/base.html")
		if err != nil {
			fmt.Println("Lỗi parse template: ", err)
			http.Error(w, "Lỗi parse template: "+err.Error(), http.StatusInternalServerError)
			return
		}

		data := struct{ Title string }{Title: "Môn học - Website của tôi"}

		err = tmpl.ExecuteTemplate(w, "base", data)
		if err != nil {
			fmt.Println("Lỗi render: ", err)
			http.Error(w, "Lỗi render: "+err.Error(), http.StatusInternalServerError)
		}
	})

	// API Lấy danh sách
	p.mux.HandleFunc("GET /api/monhoc", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		monhoc,err := db.SelectAllMonHoc()
		if err != nil {
			fmt.Println("Lỗi lấy danh sách: ", err)
			http.Error(w, "Lỗi lấy danh sách: "+err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(monhoc)
	})

	// === API MỚI: ĐỒNG BỘ DỮ LIỆU HÀNG LOẠT ===
	p.mux.HandleFunc("POST /api/monhoc/sync", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var danhSachDongBo []	MonHoc
		
		err := json.NewDecoder(r.Body).Decode(&danhSachDongBo)		
		if err != nil {
			http.Error(w, "Lỗi decode: "+err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Printf("Nhận được %d bản ghi cần đồng bộ\n", len(danhSachDongBo))
		var Insert []MonHoc
		var Update []MonHoc
		var Delete []int	

		// Lấy danh sách các giao vien đã có trong DB

		// Phân loại và xử lý từng hành động
		for _, mon := range danhSachDongBo {
			switch mon.Action {
			case "thêm":
				Insert = append(Insert, mon)
			case "sửa":
				Update = append(Update, mon)
			case "xóa":
				Delete = append(Delete, mon.ID)
			}	
		}
		db.InsertMonHoc(Insert)
		db.EditMonHoc(Update)
		db.DeleteMonHoc(Delete)	
		// Trả về thành công
		json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Đồng bộ hoàn tất"})
	})

	// Các API cũ giữ nguyên phục vụ cho CRUD lẻ (nếu cần)
	p.mux.HandleFunc("POST /api/monhoc", func(w http.ResponseWriter, r *http.Request) { /* ... */ })
	p.mux.HandleFunc("PUT /api/monhoc", func(w http.ResponseWriter, r *http.Request) { /* ... */ })
	p.mux.HandleFunc("DELETE /api/monhoc/{ma}", func(w http.ResponseWriter, r *http.Request) { /* ... */ })	
}
