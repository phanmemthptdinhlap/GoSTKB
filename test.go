package main

import (
	"fmt"
	. "GoSTKB/libsql"
)

func main() {
	fmt.Println("Hello World!")
	db, err := ConnectSTKB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	mh:=[]MonHoc{
		{TenMon: "Lý", LoaiMon: "Tự chọn"},
		{TenMon: "Hóa", LoaiMon: "Tự chọn"},
		{TenMon: "Sinh", LoaiMon: "Tự chọn"},
		{TenMon: "Tin", LoaiMon: "Tự chọn"},
		{TenMon: "Văn", LoaiMon: "Bắt buộc"},
		{TenMon: "Sử", LoaiMon: "Bắt buộc"},
		{TenMon: "Địa", LoaiMon: "Tự chọn"},
		{TenMon: "Ngoại ngữ", LoaiMon: "Tự chọn"},
		{TenMon: "GDKTPL", LoaiMon: "Tự chọn"},
		{TenMon: "QPAN", LoaiMon: "Bắt buộc"},
		{TenMon: "GDTC", LoaiMon: "Tự chọn"},
	}
	sl,err := db.InsertMonHoc(mh)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Printf("Đã thêm %d môn học vào CSDL!\n",sl)
	mh,err = db.SelectAllMonHoc()
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println("Danh sách các môn học")
	for _,v:=range mh{
		fmt.Println(v)
	}
}
