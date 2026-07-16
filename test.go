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
	tuan,err:=db.SelectTuanHoc()
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println("Tuần mặc định: ", tuan)
	fmt.Println("Chi tiết của giáo viên")
	chitiettheolop,err := db.SelectAllChiTietTheoLop(tuan)
	fmt.Println("Sau khi lấy thông tin")
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
