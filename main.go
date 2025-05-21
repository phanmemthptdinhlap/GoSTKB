package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3" // Import SQLite driver
)

type User struct {
	ID    int
	Name  string
	Email string
}

func main() {
	// Kết nối SQLite
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Tạo bảng nếu chưa tồn tại
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        email TEXT
    )`)
	if err != nil {
		panic(err)
	}

	// Định nghĩa các handler
	http.HandleFunc("/", listUsers(db))
	http.HandleFunc("/create", createUser(db))
	http.HandleFunc("/edit", editUser(db))
	http.HandleFunc("/delete", deleteUser(db))

	// Khởi động server
	http.ListenAndServe(":8080", nil)
}

// Handler: Hiển thị danh sách người dùng
func listUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, name, email FROM users")
		if err != nil {
			http.Error(w, "Lỗi truy vấn", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		users := []User{}
		for rows.Next() {
			var u User
			if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
				http.Error(w, "Lỗi đọc dữ liệu", http.StatusInternalServerError)
				return
			}
			users = append(users, u)
		}

		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Lỗi template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, users)
	}
}

// Handler: Tạo người dùng mới
func createUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			tmpl, _ := template.ParseFiles("templates/create.html")
			tmpl.Execute(w, nil)
			return
		}
		if r.Method == http.MethodPost {
			name := r.FormValue("name")
			email := r.FormValue("email")
			_, err := db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", name, email)
			if err != nil {
				http.Error(w, "Lỗi thêm người dùng", http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

// Handler: Sửa thông tin người dùng
func editUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if r.Method == http.MethodGet {
			row := db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id)
			var u User
			if err := row.Scan(&u.ID, &u.Name, &u.Email); err != nil {
				http.Error(w, "Lỗi truy vấn", http.StatusInternalServerError)
				return
			}
			tmpl, _ := template.ParseFiles("templates/edit.html")
			tmpl.Execute(w, u)
			return
		}
		if r.Method == http.MethodPost {
			name := r.FormValue("name")
			email := r.FormValue("email")
			_, err := db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", name, email, id)
			if err != nil {
				http.Error(w, "Lỗi cập nhật", http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

// Handler: Xóa người dùng
func deleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		_, err := db.Exec("DELETE FROM users WHERE id = ?", id)
		if err != nil {
			http.Error(w, "Lỗi xóa", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
