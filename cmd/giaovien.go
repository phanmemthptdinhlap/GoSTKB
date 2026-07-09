package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	. "GoSTKB/libsql"
)

// Bổ sung trường action để mapping với giao diện Vue


func (p *WebPage) SetPageGiaoVien() {
	p.mux.HandleFunc("/giaovien", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/giaovien.html", "templates/base.html")
		if err != nil {
			fmt.Println("Lỗi parse template: ", err)
			http.Error(w, "Lỗi parse template: "+err.Error(), http.StatusInternalServerError)
			return
		}

		data := struct{ Title string }{Title: "Giao viên - Website của tôi"}

		err = tmpl.ExecuteTemplate(w, "base", data)
		if err != nil {
			fmt.Println("Lỗi render: ", err)
			http.Error(w, "Lỗi render: "+err.Error(), http.StatusInternalServerError)
		}
	})

	// API Lấy danh sách
	p.mux.HandleFunc("GET /api/giaovien", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		giaovien,err := db.SelectAllGiaoVien()
		if err != nil {
			fmt.Println("Lỗi lấy danh sách: ", err)
			http.Error(w, "Lỗi lấy danh sách: "+err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(giaovien)
	})

	// === API MỚI: ĐỒNG BỘ DỮ LIỆU HÀNG LOẠT ===
	p.mux.HandleFunc("POST /api/giaovien/sync", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var danhSachDongBo []	GiaoVien
		
		err := json.NewDecoder(r.Body).Decode(&danhSachDongBo)
		if err != nil {
			http.Error(w, "Lỗi decode: "+err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Printf("Nhận được %d bản ghi cần đồng bộ\n", len(danhSachDongBo))
		var Insert []GiaoVien
		var Update []GiaoVien

		// Lấy danh sách các giao vien đã có trong DB

		// Phân loại và xử lý từng hành động
		for _, giaoVien := range danhSachDongBo {
			switch giaoVien.Action {
			case "thêm":
				Insert = append(Insert, giaoVien)
			case "sửa":
				Update = append(Update, giaoVien)
			case "xóa":
				db.DeleteGiaoVien(giaoVien.ID)
			}
		}
		db.InsertGiaoVien(Insert)
		db.UpdateGiaoVien(Update)

		// Trả về thành công
		json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Đồng bộ hoàn tất"})
	})

	// Các API cũ giữ nguyên phục vụ cho CRUD lẻ (nếu cần)
	p.mux.HandleFunc("POST /api/giaovien", func(w http.ResponseWriter, r *http.Request) { /* ... */ })
	p.mux.HandleFunc("PUT /api/giaovien", func(w http.ResponseWriter, r *http.Request) { /* ... */ })
	p.mux.HandleFunc("DELETE /api/giaovien/{ma}", func(w http.ResponseWriter, r *http.Request) { /* ... */ })
}
