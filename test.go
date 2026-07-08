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
	gv:=[]GiaoVien{
		{TenNgan: "Kiên", HoTen:"Lương Lâm Kiên",MonChinhId: 4},
		{TenNgan: "Viên", HoTen:"Nguyễn Thị Viện", MonChinhId: 5},
	}
	sl,err := db.InsertGiaoVien(gv)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Printf("Đã thêm %d Giáo viên vào CSDL!\n",sl)
	gvs,err := db.SelectAllGiaoVien()
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println("Danh sách các Giáo viên")
	for _,g:=range gvs{
		fmt.Println(g)
	}
}
