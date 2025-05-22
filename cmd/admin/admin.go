package admin

import (
	"GoSTKB/libs/myauth"
	"database/sql"
	"html/template"
	"net/http"
)

func Check_login(db *sql.DB) http.HandlerFunc {
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

			// Đăng nhập thành công
			http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
			return
		}

		// Hiển thị form đăng nhập
		tmpl := template.Must(template.ParseFiles("templates/admin/login.html"))
		tmpl.Execute(w, nil)
	}
}
func Dashboard(w http.ResponseWriter, r *http.Request) {
	// Kiểm tra xem người dùng đã đăng nhập chưa
	if r.Method != http.MethodGet {
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}

	// Hiển thị trang dashboard
	tmpl := template.Must(template.ParseFiles("templates/admin/dashboard.html"))
	tmpl.Execute(w, nil)
}
