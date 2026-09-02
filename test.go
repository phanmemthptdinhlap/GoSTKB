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
	phancong, err := db.SelectAllPhanCongMonHoc()
	if err != nil {
		fmt.Println(err)
		return
	}
	phancongMap := make(map[string][]string)
	for _, l := range lophoc {
		phancongMap[l.TenLop] = make([]string, 0)
	}
	for _, pc := range phancong {
		phancongMap[pc.Lop] = append(phancongMap[pc.Lop], pc.Mon)
	}

	fmt.Println(lops)
	fmt.Println(mons)
	fmt.Println(phancongMap)

}
