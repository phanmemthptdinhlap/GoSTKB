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
	db.InsertMonHoc(MonHoc{TenMon: "Toán", LoaiMon: "Tự nhiên"})
	db.InsertMonHoc(MonHoc{TenMon: "Văn", LoaiMon: "Xã hội"})
	mh,err := db.SelectAllMonHoc()
	if err!=nil{
		fmt.Println(err)
		return
	}
	for _,v:=range mh{
		fmt.Println(v)
	}
}
