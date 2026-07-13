package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	. "GoSTKB/libsql"
)

// Bổ sung trường action để mapping với giao diện Vue

func (p *WebPage) SetPageChiTiet() {
	p.mux.HandleFunc("/chitiet", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/chitiet.html", "templates/base.html")
		if err != nil {
			fmt.Println("Lỗi parse template: ", err)
			http.Error(w, "Lỗi parse template: "+err.Error(), http.StatusInternalServerError)
			return
		}

		data := struct{ Title string }{Title: "Chi tiết - Website của tôi"}

		err = tmpl.ExecuteTemplate(w, "base", data)
		if err != nil {
			fmt.Println("Lỗi render: ", err)
			http.Error(w, "Lỗi render: "+err.Error(), http.StatusInternalServerError)
		}
	})

	// API Lấy danh sách
	p.mux.HandleFunc("GET /api/chitiet", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		chitiet,err := db.SelectAllChiTiet()
		if err != nil {
			fmt.Println("Lỗi lấy danh sách: ", err)
			http.Error(w, "Lỗi lấy danh sách: "+err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println("Lấy danh sách chitiet: ", chitiet)
		json.NewEncoder(w).Encode(chitiet)
	})
		
		// API MỚI: ĐỒNG BỘ DỮ LIỆU HÀNG LOẠT
	p.mux.HandleFunc("POST /api/chitiet/sync", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var danhSachDongBo []	ChiTiet
		err := json.NewDecoder(r.Body).Decode(&danhSachDongBo)
			if err != nil {
				http.Error(w, "Lỗi decode: "+err.Error(), http.StatusBadRequest)
				return
			}
	
			fmt.Printf("Nhận được %d bản ghi cần đồng bộ\n", len(danhSachDongBo))
			fmt.Println("Nhận được phancong: ", danhSachDongBo)
			var Insert []ChiTiet
			var Update []ChiTiet
			var Delete []int	
	
			// Lấy danh sách các giao vien đã có trong DB
	
			// Phân loại và xử lý từng hành động
			for _, ct := range danhSachDongBo {
				switch ct.Action {
				case "thêm":
					Insert = append(Insert, ct)
				case "sửa":
					Update = append(Update, ct)
				case "xóa":
					Delete = append(Delete, ct.ID)
				}
			}
	
			db.InsertChiTiet(Insert)
			db.EditChiTiet(Update)
			db.DeleteChiTiet(Delete)
	
			// Trả về thành công
			json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Đồng bộ hoàn tất"})
		})
	
		// Các API cũ giữ nguyên phục vụ cho CRUD lẻ (nếu cần)
		p.mux.HandleFunc("POST /api/chitiet", func(w http.ResponseWriter, r *http.Request) { /* ... */ })
		p.mux.HandleFunc("PUT /api/chitiet", func(w http.ResponseWriter, r *http.Request) { /* ... */ })
		p.mux.HandleFunc("DELETE /api/chitiet/{ma}", func(w http.ResponseWriter, r *http.Request) { /* ... */ })
	}



