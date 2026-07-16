package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
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
	// API Lấy thông tin tuần mặc định
	p.mux.HandleFunc("GET /api/chitiet/tuan", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		tuan, err := db.SelectTuanHoc()
		if err != nil {
			fmt.Println("Lỗi lấy tuần mặc định: ", err)
			http.Error(w, "Lỗi lấy tuần mặc định: "+err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println("Lấy tuần mặc định: ", tuan)
		json.NewEncoder(w).Encode(tuan)
	})

	// API Lấy danh sách
	p.mux.HandleFunc("GET /api/chitiet/{tuan}", func(w http.ResponseWriter, r *http.Request) {
		tuan, err := strconv.Atoi(r.PathValue("tuan"))
		if err != nil {
			fmt.Println("Lỗi lấy tuần mặc định: ", err)
			http.Error(w, "Lỗi lấy tuần mặc định: "+err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println("Danh sách chi tiet theo tuần: ", tuan)
		w.Header().Set("Content-Type", "application/json")
		chitiet,err := db.SelectAllChiTietTheoLop(tuan)
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
			fmt.Println("Nhận được phancong: ")
			for _,ct:=range danhSachDongBo{
				vi,err:=json.MarshalIndent(ct,"","  ")
				if err!=nil{
					fmt.Println(err)
					return
				}
				fmt.Println(string(vi))
			}
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



