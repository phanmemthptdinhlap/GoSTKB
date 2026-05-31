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
}
var dsgiaovien_test=[]GiaoVien{
	{Ma: "GV001", HoTen: "Lâm Kiên", TenTKB: "KienL", MonDay: "Tin học, Trai nghiem", LopDay: "12A1, 12A2"},
	{Ma: "GV002", HoTen: "Kiên", TenTKB: "Kien", MonDay: "Tin học", LopDay: "12C1, 12C2"},
}
func getGiaoVien() []GiaoVien{
	return dsgiaovien_test
}
func addGiaovien(gv GiaoVien) []GiaoVien{
	for _,v:=range dsgiaovien_test{
		if v.Ma==gv.Ma{
			fmt.Println("Giao vien đã tồn tại")
			return dsgiaovien_test
		}
	}
	dsgiaovien_test=append(dsgiaovien_test,gv)
	return dsgiaovien_test
}
func updateGiaovien(gv GiaoVien) []GiaoVien{
	for i,v:=range dsgiaovien_test{
		if v.Ma==gv.Ma{
			dsgiaovien_test[i].HoTen=gv.HoTen
			dsgiaovien_test[i].TenTKB=gv.TenTKB
			return dsgiaovien_test
		}
	}
	dsgiaovien_test=append(dsgiaovien_test,gv)
	return dsgiaovien_test
}

func deleteGiaovien(maGV string) []GiaoVien{
		for i,v:=range dsgiaovien_test{
			if v.Ma==maGV{
				dsgiaovien_test=append(dsgiaovien_test[:i],dsgiaovien_test[i+1:]...)
				return dsgiaovien_test
			}
		}
		fmt.Println("Giao vien không tồn tại")
		return dsgiaovien_test
}
func (page *WebPage) SetPageGiaoVien(){
		page.mux.HandleFunc("/giaovien",func(w http.ResponseWriter, r *http.Request){
			temp,err:=template.ParseFiles("templates/giaovien.html","templates/base.html")
			if err!=nil{
				fmt.Println(err)
				http.Error(w,"Internal Server Error",http.StatusInternalServerError)
				return
			}
			data:=struct{
				Title string
			}{
				Title:"Giáo viên - GoSTKB",
			}	
			err=temp.ExecuteTemplate(w,"base",data)
			if err!=nil{
				fmt.Println(err)
				http.Error(w,"Lỗi trang web",http.StatusInternalServerError)
				return
			}
		})
		page.mux.HandleFunc("GET /api/giaovien",func(w http.ResponseWriter, r *http.Request){
			w.Header().Set("Content-Type","application/json")
			json.NewEncoder(w).Encode(getGiaoVien())
		})
		page.mux.HandleFunc("POST /api/giaovien",func(w http.ResponseWriter, r *http.Request){
			var gv GiaoVien
			if err:=json.NewDecoder(r.Body).Decode(&gv);err!=nil{
				fmt.Println("Lỗi",err.Error())
				http.Error(w,"Lỗi gửi đi",http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type","application/json")
			json.NewEncoder(w).Encode(addGiaovien(gv))
		})
		page.mux.HandleFunc("PUT /api/giaovien",func(w http.ResponseWriter, r *http.Request){
			var gv GiaoVien
			if err:=json.NewDecoder(r.Body).Decode(&gv);err!=nil{
				http.Error(w,"Lỗi gửi đi",http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type","application/json")
			json.NewEncoder(w).Encode(updateGiaovien(gv))
		})
		page.mux.HandleFunc("DELETE /api/giaovien/{ma}",func(w http.ResponseWriter, r *http.Request){
			ma:=r.PathValue("ma")
			w.Header().Set("Content-Type","application/json")
			json.NewEncoder(w).Encode(deleteGiaovien(ma))
		})
	}
