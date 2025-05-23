package admin

import (
	"GoSTKB/libs/myauth"
	"database/sql"
	"html/template"
	"net/http"
)

func Check_login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Xử lý đăng nhập
		if r.Method == http.MethodPost {
			// ...kiểm tra tài khoản...
			session, _ := myauth.Store.Get(r, "session-name")
			session.Values["authenticated"] = true
			session.Save(r, w)
			http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
			return
		}
		// Hiển thị form đăng nhập
		tmpl := template.Must(template.ParseFiles("templates/admin/login.html"))
		tmpl.Execute(w, nil)
	}

}

func Do_login(db *sql.DB) http.HandlerFunc {
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
			session.Save(r, w)
			// Chuyển hướng đến trang dashboard
			// Nếu đăng nhập thành công, chuyển hướng đến trang dashboard
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
