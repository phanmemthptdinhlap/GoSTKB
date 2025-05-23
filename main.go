package main

import (
	"GoSTKB/cmd/admin"
	"GoSTKB/libs/database"
	"GoSTKB/libs/myauth"
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
	http.HandleFunc("/admin/login", admin.Do_login(db))
	http.HandleFunc("/admin/logout", func(w http.ResponseWriter, r *http.Request) {
		// Xử lý đăng xuất
		session, _ := myauth.Store.Get(r, "session-name")
		session.Values["authenticated"] = false
		session.Save(r, w)
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	})
	http.HandleFunc("/admin/dashboard", myauth.RequireAuth(func(w http.ResponseWriter, r *http.Request) {
		// Hiển thị trang dashboard
		tmpl := template.Must(template.ParseFiles("templates/admin/dashboard.html"))
		tmpl.Execute(w, nil)
	}))
	// Khởi động server
	http.ListenAndServe(":8080", nil)
}
