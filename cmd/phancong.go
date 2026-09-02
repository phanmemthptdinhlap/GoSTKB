package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	. "GoSTKB/libsql"
)

// Bổ sung trường action để mapping với giao diện Vue

func phancongMonHoc() interface{} {
	lophoc, err := db.SelectAllLopHoc()
	if err != nil {
		fmt.Println("Lỗi lấy danh sách: ", err)
		return nil
	}
	monhoc, err := db.SelectAllMonHoc()
	if err != nil {
		fmt.Println("Lỗi lấy danh sách: ", err)
		return nil
	}
	mons:=make([]string,0)
	for _,m := range monhoc{
		mons=append(mons,m.TenMon)
	}
	phancong, err := db.SelectAllPhanCongMonHoc()
	if err != nil {
		fmt.Println("Lỗi lấy danh sách: ", err)
		return nil
	}
	phancongMap := make(map[string][]string)
	for _, l := range lophoc {
		phancongMap[l.TenLop] = make([]string, 0)
	}
	for _, pc := range phancong {
		phancongMap[pc.Lop] = append(phancongMap[pc.Lop], pc.Mon)
	}
	return map[string] interface{}{
		"mon": mons,
		"phancong": phancongMap,
	}
}

func (p *WebPage) SetPagePhanCong() {
	p.mux.HandleFunc("/pc_giaovien", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/phancong_giaovien.html", "templates/base.html")
		if err != nil {
			fmt.Println("Lỗi parse template: ", err)
			http.Error(w, "Lỗi parse template: "+err.Error(), http.StatusInternalServerError)
			return
		}

		data := struct{ Title string }{Title: "Phân công giáo viên - Website của tôi"}

		err = tmpl.ExecuteTemplate(w, "base", data)
		if err != nil {
			fmt.Println("Lỗi render: ", err)
			http.Error(w, "Lỗi render: "+err.Error(), http.StatusInternalServerError)
		}
	})
	p.mux.HandleFunc("/pc_monhoc", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/phancong_monhoc.html", "templates/base.html")
		if err != nil {
			fmt.Println("Lỗi parse template: ", err)
			http.Error(w, "Lỗi parse template: "+err.Error(), http.StatusInternalServerError)
			return
		}

		data := struct{ Title string }{Title: "Phân công môn học - Website của tôi"}
		err = tmpl.ExecuteTemplate(w, "base", data)
		if err != nil {
			fmt.Println("Lỗi render: ", err)
			http.Error(w, "Lỗi render: "+err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// API Lấy danh sách
	p.mux.HandleFunc("GET /api/phancong/mon", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		data:=phancongMonHoc()
		json.NewEncoder(w).Encode(data)
	})
	p.mux.HandleFunc("GET /api/phancong/giaovien", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode("Hello World")
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
			var Delete []int	
	
			// Lấy danh sách các giao vien đã có trong DB
	
			// Phân loại và xử lý từng hành động
				
			db.InsertPhanCong(Insert)
			db.DeletePhanCong(Delete)
	
			// Trả về thành công
			json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Đồng bộ hoàn tất"})
		})
	
		// Các API cũ giữ nguyên phục vụ cho CRUD lẻ (nếu cần)
		p.mux.HandleFunc("POST /api/phancong", func(w http.ResponseWriter, r *http.Request) { /* ... */ })
		p.mux.HandleFunc("PUT /api/phancong", func(w http.ResponseWriter, r *http.Request) { /* ... */ })
		p.mux.HandleFunc("DELETE /api/phancong/{ma}", func(w http.ResponseWriter, r *http.Request) { /* ... */ })
	}



