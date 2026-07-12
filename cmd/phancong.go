package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	. "GoSTKB/libsql"
)

// Bổ sung trường action để mapping với giao diện Vue

func (p *WebPage) SetPagePhanCong() {
	p.mux.HandleFunc("/phancong", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/phancong.html", "templates/base.html")
		if err != nil {
			fmt.Println("Lỗi parse template: ", err)
			http.Error(w, "Lỗi parse template: "+err.Error(), http.StatusInternalServerError)
			return
		}

		data := struct{ Title string }{Title: "Phân phối - Website của tôi"}

		err = tmpl.ExecuteTemplate(w, "base", data)
		if err != nil {
			fmt.Println("Lỗi render: ", err)
			http.Error(w, "Lỗi render: "+err.Error(), http.StatusInternalServerError)
		}
	})

	// API Lấy danh sách
	p.mux.HandleFunc("GET /api/phancong", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		phancong,err := db.SelectAllPhanCong()
		if err != nil {
			fmt.Println("Lỗi lấy danh sách: ", err)
			http.Error(w, "Lỗi lấy danh sách: "+err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println("Lấy danh sách phancong: ", phancong)
		json.NewEncoder(w).Encode(phancong)
	})
		
		// API MỚI: ĐỒNG BỘ DỮ LIỆU HÀNG LOẠT
	p.mux.HandleFunc("POST /api/phancong/sync", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var danhSachDongBo []	PhanCong	
		err := json.NewDecoder(r.Body).Decode(&danhSachDongBo)
			if err != nil {
				http.Error(w, "Lỗi decode: "+err.Error(), http.StatusBadRequest)
				return
			}
	
			fmt.Printf("Nhận được %d bản ghi cần đồng bộ\n", len(danhSachDongBo))
			fmt.Println("Nhận được phancong: ", danhSachDongBo)
			var Insert []PhanCong
			var Update []PhanCong
			var Delete []int	
	
			// Lấy danh sách các giao vien đã có trong DB
	
			// Phân loại và xử lý từng hành động
			for _, phancong := range danhSachDongBo {
				switch phancong.Action {
				case "thêm":
					Insert = append(Insert, phancong)
				case "sửa":
					Update = append(Update, phancong)
				case "xóa":
					Delete = append(Delete, phancong.ID)
				}
			}
	
			db.InsertPhanCong(Insert)
			db.EditPhanCong(Update)
			db.DeletePhanCong(Delete)
	
			// Trả về thành công
			json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Đồng bộ hoàn tất"})
		})
	
		// Các API cũ giữ nguyên phục vụ cho CRUD lẻ (nếu cần)
		p.mux.HandleFunc("POST /api/phancong", func(w http.ResponseWriter, r *http.Request) { /* ... */ })
		p.mux.HandleFunc("PUT /api/phancong", func(w http.ResponseWriter, r *http.Request) { /* ... */ })
		p.mux.HandleFunc("DELETE /api/phancong/{ma}", func(w http.ResponseWriter, r *http.Request) { /* ... */ })
	}



