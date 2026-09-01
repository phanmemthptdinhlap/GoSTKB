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
	lophoc, err := db.SelectAllLopHoc()
	if err != nil {
		fmt.Println(err)
		return
	}
	lops:=make([]string,0)
	for _,l := range lophoc{
		lops = append(lops,l.TenLop)
	}
	monhoc, err:=db.SelectAllMonHoc()
	if err !=nil {
		fmt.Println(err)
		return
	}
	mons:=make([]string,0)
	for _,m := range monhoc{
		mons = append(mons,m.TenMon)
	}
	phancong, err := db.SelectAllPhanCong()
	if err != nil {
		fmt.Println(err)
		return
	}
	phancongMap := make(map[int]map[int][2]int)
	for _, v := range phancong {
		if _, ok := phancongMap[v.LopId]; !ok {
			phancongMap[v.LopId] = make(map[int][2]int)
		}
		phancongMap[v.LopId][v.MonHocId] = [2]int{v.GiaoVienId, v.TongTiet}
	}

	fmt.Println(lops)
	fmt.Println(mons)
}
