package main

import (
	"GoSTKB/cmd/admin"
	"GoSTKB/cmd/user"
	"GoSTKB/libs/database"
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
	http.HandleFunc("/", user.Show_home())
	//admin
	http.HandleFunc("/admin", admin.RequireAuth(admin.ListUser(db)))
	http.HandleFunc("/admin/user", admin.RequireAuth(admin.ListUser(db)))
	http.HandleFunc("/admin/user/add", admin.RequireAuth(admin.CreateUser(db)))
	http.HandleFunc("/admin/user/edit", admin.RequireAuth(admin.EditUser(db)))
	http.HandleFunc("/admin/user/delete", admin.RequireAuth(admin.DeleteUser(db)))
	http.HandleFunc("/admin/login", admin.Do_admin(db))
	http.HandleFunc("/admin/logout", admin.Logout())
	// Khởi động server
	http.ListenAndServe(":8080", nil)
}
