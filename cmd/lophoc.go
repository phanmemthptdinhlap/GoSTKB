package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

// Bổ sung trường action để mapping với giao diện Vue
type LopHoc struct {
	Ma     string `json:"ma"`
	Ten    string `json:"ten"`
	Gvcn   string `json:"gvcn"`
	Action string `json:"action,omitempty"` 
}

var lop_test = []LopHoc{
	{Ma: "1", Ten: "10A1", Gvcn: "Mai"},
	{Ma: "2", Ten: "12A1", Gvcn: "DuyenA"},
	{Ma: "3", Ten: "12A2", Gvcn: "DuyenB"},
}

func getLopHoc() []LopHoc {
	return lop_test
}

func addLopHoc(lop LopHoc) {
	// Lọc bỏ trường action trước khi lưu vào DB (bằng cách gán rỗng)
	lop.Action = " "
	lop_test = append(lop_test, lop)
	fmt.Println("Thêm lớp học", lop_test)
}

func deleteLopHoc(ma string) {
	for i, v := range lop_test {
		if v.Ma == ma {
			lop_test = append(lop_test[:i], lop_test[i+1:]...)
			break // Thêm break để dừng vòng lặp sau khi xóa thành công, tránh lỗi out of bounds
		}
	}
	fmt.Println("Sau khi xóa:", lop_test)
}

func updateLopHoc(lop LopHoc) {
	fmt.Println("Cập nhật lớp học", lop)
	for i, v := range lop_test {
		if v.Ma == lop.Ma {
			fmt.Println("Cập nhật thành công mã", lop.Ma)
			lop.Action = "" // Xóa action trước khi lưu
			lop_test[i] = lop
			break
		}
	}
	fmt.Println("Cập nhật lớp học", lop_test)
}

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
		json.NewEncoder(w).Encode(getLopHoc())
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
				addLopHoc(lop)
			case "sửa":
				updateLopHoc(lop)
			case "xóa":
				deleteLopHoc(lop.Ma)
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
