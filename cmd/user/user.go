package user

import (
	"GoSTKB/libs/myauth"
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
)

func Check_user(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Xử lý đăng nhập
		if r.Method == http.MethodPost {
			// ...kiểm tra tài khoản...
			session, _ := myauth.Store.Get(r, "session-name")
			session.Values["authenticated"] = true
			session.Values["role"] = "user" // Lưu vai trò người dùng
			session.Save(r, w)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		// Hiển thị form đăng nhập
		tmpl := template.Must(template.ParseFiles("templates/login.html"))
		tmpl.Execute(w, nil)
	}

}

func Do_user(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// Lấy dữ liệu từ form
			email := r.FormValue("email")
			password := r.FormValue("password")

			// Kiểm tra tài khoản admin
			var hashedPassword string
			var name string
			err := db.QueryRow("SELECT name, password FROM user WHERE email = ?", email).Scan(&name, &hashedPassword)
			if err != nil {
				http.Error(w, "Email không tồn tại", http.StatusUnauthorized)
				return
			}

			// Kiểm tra mật khẩu
			if !myauth.CheckPasswordHash(password, hashedPassword) {
				// Nếu mật khẩu không đúng
				http.Error(w, "Mã đăng nhập không đúng", http.StatusUnauthorized)
				return
			}

			// Lưu thông tin đăng nhập vào session
			session, _ := myauth.Store.Get(r, "session-name")
			session.Values["authenticated"] = true
			session.Values["user"] = email
			session.Values["name"] = name   // Lưu tên người dùng
			session.Values["role"] = "user" // Lưu vai trò người dùng
			session.Save(r, w)
			// Chuyển hướng đến trang dashboard
			// Nếu đăng nhập thành công, chuyển hướng đến trang dashboard
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		// Hiển thị form đăng nhập
		tmpl := template.Must(template.ParseFiles("templates/login.html"))
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

func Show_home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Kiểm tra xem người dùng đã đăng nhập chưa
		loginStatus := myauth.IsAuthenticated(r)
		type Page struct {
			Msg string
		}
		var page Page
		if !loginStatus {
			page.Msg = "Bạn chưa đăng nhập"
		} else {
			session, _ := myauth.Store.Get(r, "session-name")
			page.Msg = "Chào " + session.Values["user"].(string)
		}
		tmpl := template.Must(template.ParseFiles("templates/home.html"))
		err := tmpl.Execute(w, page)
		if err != nil {
			fmt.Println("Error executing template:", err)
			return
		}
	}
}
