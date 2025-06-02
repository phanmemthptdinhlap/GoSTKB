package main

import (
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Kết nối đến cơ sở dữ liệu SQLite
	db, _ := ConnectSTKB()
	Gui(db)
	// Phục vụ các tệp tĩnh
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// Phục vụ các tệp HTML
	//http.HandleFunc("/", ShowHome)
	// Định nghĩa các handler
	// Khởi động server
	//http.ListenAndServe(":8080", nil)
}
