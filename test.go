package main

import (
	"fmt"
	"encoding/json"
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
	chitiet:=[]ChiTiet{
		{PhanCongId: 1, Tuan: 1, Sotiet: 5},
		{PhanCongId: 2, Tuan: 1, Sotiet: 5},


	}

	sl,err := db.InsertChiTiet(chitiet)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Printf("Đã thêm %d Giáo viên vào CSDL!\n",sl)
	ct,err := db.SelectAllChiTiet()
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println("Chi tiết của giáo viên")
	fmt.Println(ct)
	chitiettheolop,err := db.SelectAllChiTietTheoLop(1)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println("Danh sách các Chi tiết theo lớp")
	for _,ct:=range chitiettheolop{
		vi,err:=json.MarshalIndent(ct,"","  ")
		if err!=nil{
			fmt.Println(err)
			return
		}
		fmt.Println(string(vi))
	}
}
