package main

import (
	"fmt"
	"html/template"
	"net/http"
	"encoding/json"
)
type GiaoVien struct{
	Ma string `json:"ma"`
	HoTen string `json:"hoTen"`
	TenTKB string `json:"tenTKB"`
	MonDay string `json:"monDay"`
	LopDay string `json:"lopDay"`
	Action string `json:"action,omitempty"`
}
var dsgiaovien_test=[]GiaoVien{
	{Ma: "GV001", HoTen: "Lâm Kiên", TenTKB: "KienL", MonDay: "Tin học, Trai nghiem", LopDay: "12A1, 12A2"},
	{Ma: "GV002", HoTen: "Kiên", TenTKB: "Kien", MonDay: "Tin học", LopDay: "12C1, 12C2"},
}
func getGiaoVien() []GiaoVien{
	return dsgiaovien_test
}
func addGiaovien(gv GiaoVien) {
	gv.Action = " "
	for _,v:=range dsgiaovien_test{
		if v.Ma==gv.Ma{
			fmt.Println("Giao vien đã tồn tại")
			return
		}
	}
	dsgiaovien_test=append(dsgiaovien_test,gv)
}
func updateGiaovien(gv GiaoVien) {
	gv.Action = " "
	for i,v:=range dsgiaovien_test{
		if v.Ma==gv.Ma{
			dsgiaovien_test[i]=gv
			return
		} 
	}
	dsgiaovien_test=append(dsgiaovien_test,gv)
}

func deleteGiaovien(maGV string) {
		for i,v:=range dsgiaovien_test{
			if v.Ma==maGV{
				dsgiaovien_test=append(dsgiaovien_test[:i],dsgiaovien_test[i+1:]...)
				return 
			}
		}
		fmt.Println("Giao vien không tồn tại")
}
func (p *WebPage) SetPageGiaoVien() {
	p.mux.HandleFunc("/giaovien", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/giaovien.html", "templates/base.html")
		if err != nil {
			fmt.Println("Lỗi parse template: ", err)
			http.Error(w, "Lỗi parse template: "+err.Error(), http.StatusInternalServerError)
			return
		}

		data := struct{ Title string }{Title: "Giáo viên - Website của tôi"}

		err = tmpl.ExecuteTemplate(w, "base", data)
		if err != nil {
			fmt.Println("Lỗi render: ", err)
			http.Error(w, "Lỗi render: "+err.Error(), http.StatusInternalServerError)
		}
	})

	// API Lấy danh sách
	p.mux.HandleFunc("GET /api/giaovien", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(getGiaoVien())
	})

	// === API MỚI: ĐỒNG BỘ DỮ LIỆU HÀNG LOẠT ===
	p.mux.HandleFunc("POST /api/giaovien/sync", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var danhSachDongBo []GiaoVien
		
		err := json.NewDecoder(r.Body).Decode(&danhSachDongBo)
		if err != nil {
			http.Error(w, "Lỗi decode: "+err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Printf("Nhận được %d bản ghi cần đồng bộ\n", len(danhSachDongBo))

		// Phân loại và xử lý từng hành động
		for _, gv := range danhSachDongBo {
			switch gv.Action {
			case "thêm":
				addGiaovien(gv)
			case "sửa":
				updateGiaovien(gv)
			case "xóa":
				deleteGiaovien(gv.Ma)
			}
		}

		// Trả về thành công
		json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Đồng bộ hoàn tất"})
	})

	// Các API cũ giữ nguyên phục vụ cho CRUD lẻ (nếu cần)
	p.mux.HandleFunc("POST /api/giaovien", func(w http.ResponseWriter, r *http.Request) { /* ... */ })
	p.mux.HandleFunc("PUT /api/giaovien", func(w http.ResponseWriter, r *http.Request) { /* ... */ })
	p.mux.HandleFunc("DELETE /api/giaovien/{ma}", func(w http.ResponseWriter, r *http.Request) { /* ... */ })
}
