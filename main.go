package main

import (
	"GoSTKB/SQL"
	"GoSTKB/VIEW"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Kết nối đến cơ sở dữ liệu SQLite
	db, _ := SQL.ConnectSTKB()
	if db == nil {
		panic("Không thể kết nối đến cơ sở dữ liệu")
	}
	// Tạo bảng nếu chưa tồn tại
	VIEW.Gui(db)
}
