package main

import (
	"GoSTKB/cmd/admin"
	"GoSTKB/libs/database"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	// Import package libs
)

func main() {
	// Kết nối SQLite
	db, _ := database.Connect()
	database.CreateTable_admin(db)
	database.CreateTable_users(db)
	// Định nghĩa các handler
	http.HandleFunc("/admin", admin.Check_login(db))
	http.HandleFunc("/admin/dashboard", func(w http.ResponseWriter, r *http.Request) {
		// Xử lý trang dashboard admin
		if r.Method != http.MethodGet {
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
			return
		}
		// Hiển thị trang dashboard
		tmpl := template.Must(template.ParseFiles("templates/admin/dashboard.html"))
		tmpl.Execute(w, nil)
	})
	// Khởi động server
	http.ListenAndServe(":8080", nil)
}
