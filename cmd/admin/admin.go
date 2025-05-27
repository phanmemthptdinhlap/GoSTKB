package admin

import (
	"GoSTKB/libs/database"
	"GoSTKB/libs/myauth"
	"database/sql"
	"html/template"
	"net/http"
)

func IsAuthenticated(r *http.Request) bool {
	session, _ := myauth.Store.Get(r, "session-name")
	auth, ok := session.Values["authenticated"].(bool)
	role := session.Values["role"].(string)
	return ok && auth && role == "admin"
}

func Do_admin(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// Lấy dữ liệu từ form
			user := r.FormValue("user")
			password := r.FormValue("password")

			// Kiểm tra tài khoản admin
			var hashedPassword string
			err := db.QueryRow("SELECT password FROM admin WHERE user = ?", user).Scan(&hashedPassword)
			if err != nil {
				http.Error(w, "Invalid username or password", http.StatusUnauthorized)
				return
			}

			// Kiểm tra mật khẩu
			if !myauth.CheckPasswordHash(password, hashedPassword) {
				// Nếu mật khẩu không đúng
				http.Error(w, "Invalid username or password", http.StatusUnauthorized)
				return
			}

			// Lưu thông tin đăng nhập vào session
			session, _ := myauth.Store.Get(r, "session-name")
			session.Values["authenticated"] = true
			session.Values["user"] = user
			session.Values["role"] = "admin" // Lưu vai trò người dùng
			session.Save(r, w)
			// Chuyển hướng đến trang dashboard
			// Nếu đăng nhập thành công, chuyển hướng đến trang dashboard
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
			return
		}

		// Hiển thị form đăng nhập
		tmpl := template.Must(template.ParseFiles("templates/admin/login.html"))
		tmpl.Execute(w, nil)
	}
}
func Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Xử lý đăng xuất
		session, _ := myauth.Store.Get(r, "session-name")
		session.Values["authenticated"] = false
		session.Values["user"] = nil
		session.Values["role"] = nil // Xóa vai trò người dùng
		session.Save(r, w)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !IsAuthenticated(r) {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
			return
		}

		next(w, r) // Hiển thị danh sách người dùng
	}
}

func ListUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		rows, err := db.Query("SELECT id, name, email, password FROM users")
		if err != nil {
			http.Error(w, "Lỗi truy vấn", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		users := []User{}
		for rows.Next() {
			var u User
			if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password); err != nil {
				http.Error(w, "Lỗi đọc dữ liệu", http.StatusInternalServerError)
				return
			}
			users = append(users, u)
		}

		tmpl, err := template.ParseFiles("templates/admin/user/list.html")
		if err != nil {
			http.Error(w, "Lỗi template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, users)
	}
}

// Handler: Tạo người dùng mới
func CreateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			tmpl, _ := template.ParseFiles("templates/admin/user/create.html")
			tmpl.Execute(w, nil)
			return
		}
		if r.Method == http.MethodPost {
			name := r.FormValue("name")
			email := r.FormValue("email")
			password := r.FormValue("password")
			// Kiểm tra xem email đã tồn tại chưa
			err := database.CreateUser(db, name, email, password)
			if err != nil {
				http.Error(w, "Lỗi thêm người dùng", http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/admin/user", http.StatusSeeOther)
		}
	}
}

// Handler: Sửa thông tin người dùng
func EditUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if r.Method == http.MethodGet {
			row := db.QueryRow("SELECT id, name, email, password FROM users WHERE id = ?", id)
			var u User
			if err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password); err != nil {
				http.Error(w, "Lỗi truy vấn", http.StatusInternalServerError)
				return
			}
			tmpl, _ := template.ParseFiles("templates/admin/user/edit.html")
			tmpl.Execute(w, u)
			return
		}
		if r.Method == http.MethodPost {
			name := r.FormValue("name")
			email := r.FormValue("email")
			password := r.FormValue("password")
			_, err := db.Exec("UPDATE users SET name = ?, email = ?, password WHERE id = ?", name, email, password, id)
			if err != nil {
				http.Error(w, "Lỗi cập nhật", http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/admin/user", http.StatusSeeOther)
		}
	}
}

// Handler: Xóa người dùng
func DeleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		_, err := db.Exec("DELETE FROM users WHERE id = ?", id)
		if err != nil {
			http.Error(w, "Lỗi xóa", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/user", http.StatusSeeOther)
	}
}
