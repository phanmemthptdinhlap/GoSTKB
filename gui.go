package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// Showtable Hiển thị bảng dữ liệu
type Table_view struct {
	Name    string
	Columns []string
	Rows    []map[string]interface{}
}

func (p *Table_view) GetHtml(name, class_table string, edit bool) string {
	texthtml := `<table id="table_` + name + `" data-table-name="` + name + `" class="` + class_table + `">
	<thead>
	<tr>`
	for _, col := range p.Columns {
		texthtml += `<th>` + col + `</th>`
	}
	if edit {
		texthtml += `<th>Hành động</th>` // Thêm cột Actions nếu cần
	}
	texthtml += `</tr>
	</thead>
	<tbody>`
	for _, row := range p.Rows {
		texthtml += `<tr>`
		for _, col := range p.Columns {
			texthtml += `<td>` + fmt.Sprintf("%v", row[col]) + `</td>`
		}
		if edit {
			texthtml += `<td>
				<button class="edit-btn" data-id="` + fmt.Sprintf("%v", row["id"]) + `">Sửa</button>
				<button class="delete-btn" data-id="` + fmt.Sprintf("%v", row["id"]) + `">Xóa</button>
			</td>` // Thêm nút Sửa và Xóa
			texthtml += `</tr>`
		}
	}
	texthtml += `</tbody>
	</table>
	<button class="add-btn">Thêm mới Giáo viên</button>`
	return texthtml
}

// ‌ WebGui là hàm chính để khởi tạo giao diện web
type WebGui struct {
	Title  string
	Header string
	Footer string
	Body   string
}

// Gui khởi tạo giao diện web cho quản lý thời khóa biểu
func (w *WebGui) Init(title, header, footer, body string) {
	w.Title = title
	w.Header = header
	w.Footer = footer
	w.Body = body
}
func (w *WebGui) TextHtml() string {
	return fmt.Sprintf(`
<!DOCTYPE html>	
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>%s</title>
	<link rel="stylesheet" href="/static/css/style.css">
	<script src="/static/js/script.js"></script>
</head>	
<body>
	<header>
		<div id="header">%s</div>
	</header>	
	<div id="content">
		%s
	</div>
	<footer>
		<div id="footer">%s</div>
	</footer>
</body>
</html>
`, w.Title, w.Header, w.Body, w.Footer)
}
func (w *WebGui) Template() *template.Template {
	tmpl, err := template.New("webgui").Parse(w.TextHtml())
	if err != nil {
		log.Fatalf("Lỗi khi tạo template: %v", err)
	}
	return tmpl
}

func getData(tableName string, db *sql.DB) (table []map[string]interface{}, col []string, title string, err error) {
	switch tableName {
	case "teachers":
		table, err = GetTeachers(db)
		col = []string{"id", "name", "tkb_name"}
		title = "Danh sách Giáo viên"
	case "subjects":
		table, err = GetSubjects(db)
		col = []string{"id", "name"}
		title = "Danh sách Môn học"
	case "classes":
		table, err = GetClasses(db)
		col = []string{"id", "name"}
		title = "Danh sách Lớp học"
	case "assignments":
		table, err = GetAssignments(db)
		col = []string{"id", "teacher_name", "subject_name", "class_name", "day_of_week", "start_time", "end_time"}
		title = "Danh sách Phân công"
	// Thêm các trường hợp khác nếu cần
	default:
		table = nil
		err = fmt.Errorf("bảng không hợp lệ: %s", tableName)
	}
	return table, col, title, err
}

func ShowTable(db *sql.DB, table_name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		table, col, title, err := getData(table_name, db)
		if err != nil {
			http.Error(w, fmt.Sprintf("Lỗi khi lấy dữ liệu: %v", err), http.StatusInternalServerError)
			return
		}
		data := &Table_view{
			Name:    title,
			Columns: col,
			Rows:    table,
		}
		htmlTable := data.GetHtml(table_name, "table", true)
		// Tạo dữ liệu để truyền vào template
		tmpl := (&WebGui{
			Title:  "Quản lý thời khóa biểu",
			Header: "Giáo viên",
			Footer: "© 2023 Quản lý thời khóa biểu",
			Body:   htmlTable,
		}).Template()
		err = tmpl.Execute(w, data)
		if err != nil {
			fmt.Println(data)
			fmt.Println("Lỗi khi thực thi template:", err)
			return
		}
	}
}

func AddGiaovien(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Xử lý thêm giáo viên mới vào cơ sở dữ liệu
		if r.Method == http.MethodPost {
			name := r.FormValue("name")
			tkbName := r.FormValue("tkb_name")
			// Thêm giáo viên vào cơ sở dữ liệu
			_, err := db.Exec("INSERT INTO teachers (name, tkb_name) VALUES (?, ?)", name, tkbName)
			if err != nil {
				http.Error(w, fmt.Sprintf("Lỗi khi thêm giáo viên: %v", err), http.StatusInternalServerError)
				return
			}
			// Chuyển hướng về trang danh sách giáo viên
			http.Redirect(w, r, "/teachers", http.StatusSeeOther)
		} else {
			// Hiển thị form thêm giáo viên
			data := &WebGui{
				Title:  "Thêm Giáo viên",
				Header: "Thêm Giáo viên",
				Footer: "© 2023 Quản lý thời khóa biểu",
				Body: `<form method="POST" action="/teachers/add">
					<label for="name">Tên Giáo viên:</label>
					<input type="text" id="name" name="name" required>
					<label for="tkb_name">Tên Thời khóa biểu:</label>
					<input type="text" id="tkb_name" name="tkb_name" required>
					<button type="submit">Thêm Giáo viên</button>
				</form>`,
			}
			// Tạo template từ dữ liệu
			tmpl := data.Template()
			// Thực thi template và gửi kết quả đến trình duyệt
			err := tmpl.Execute(w, data)
			if err != nil {
				http.Error(w, "Lỗi khi hiển thị trang", http.StatusInternalServerError)
				return
			}
		}
	}
	// Gui khởi tạo giao diện web cho quản lý thời khóa biểu
}
func EditGiaovien(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Xử lý sửa thông tin giáo viên
		if r.Method == http.MethodPost {
			id := r.FormValue("id")
			name := r.FormValue("name")
			tkbName := r.FormValue("tkb_name")
			// Cập nhật thông tin giáo viên trong cơ sở dữ liệu
			_, err := db.Exec("UPDATE teachers SET name = ?, tkb_name = ? WHERE id = ?", name, tkbName, id)
			if err != nil {
				http.Error(w, fmt.Sprintf("Lỗi khi cập nhật giáo viên: %v", err), http.StatusInternalServerError)
				return
			}
			// Chuyển hướng về trang danh sách giáo viên
			http.Redirect(w, r, "/teachers", http.StatusSeeOther)
		} else {
			// Hiển thị form sửa thông tin giáo viên
			data := &WebGui{
				Title:  "Sửa Giáo viên",
				Header: "Sửa Giáo viên",
				Footer: "© 2023 Quản lý thời khóa biểu",
				Body: `<form method="POST" action="/teachers/edit">
					<label for="id">ID Giáo viên:</label>
					<input type="text" id="id" name="id" required>
					<label for="name">Tên Giáo viên:</label>
					<input type="text" id="name" name="name" required>
					<label for="tkb_name">Tên Thời khóa biểu:</label>
					<input type="text" id="tkb_name" name="tkb_name" required>
					<button type="submit">Cập nhật Giáo viên</button>
				</form>`,
			}
			tmpl := data.Template()
			err := tmpl.Execute(w, data)
			if err != nil {
				http.Error(w, "Lỗi khi hiển thị trang", http.StatusInternalServerError)
				return
			}
		}
	}
}
func DelGiaovien(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Xử lý xóa giáo viên
		if r.Method == http.MethodPost {
			id := r.FormValue("id")
			// Xóa giáo viên khỏi cơ sở dữ liệu
			_, err := db.Exec("DELETE FROM teachers WHERE id = ?", id)
			if err != nil {
				http.Error(w, fmt.Sprintf("Lỗi khi xóa giáo viên: %v", err), http.StatusInternalServerError)
				return
			}
			// Chuyển hướng về trang danh sách giáo viên
			http.Redirect(w, r, "/teachers", http.StatusSeeOther)
		} else {
			http.Error(w, "Phương thức không hợp lệ", http.StatusMethodNotAllowed)
		}
	}
}

// Gui khởi tạo giao diện web cho quản lý thời khóa biểu

func Gui(db *sql.DB) {
	// Tạo một template mới
	// Tạo một HTTP handler để phục vụ trang web
	http.HandleFunc("/teachers", ShowTable(db, "teachers"))
	http.HandleFunc("/add/teachers", AddGiaovien(db))   // Thêm route cho thêm giáo viên
	http.HandleFunc("/edit/teachers", EditGiaovien(db)) // Thêm route cho sửa giáo viên
	http.HandleFunc("/del/teachers", DelGiaovien(db))   // Thêm route cho xóa giáo viên
	http.HandleFunc("/subjects", ShowTable(db, "subjects"))
	http.HandleFunc("/classes", ShowTable(db, "classes"))
	http.HandleFunc("/assignments", ShowTable(db, "assignments"))
	// Hiển thị trang chủ
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Tạo dữ liệu để truyền vào template
		data := &WebGui{
			Title:  "Quản lý thời khóa biểu",
			Header: "Trang chủ",
			Footer: "© 2023 Quản lý thời khóa biểu",
			Body: `<h1>Chào mừng đến với Quản lý thời khóa biểu</h1><p>Vui lòng chọn một mục từ menu bên dưới dây.</p>
			<nav>
				<ul>
					<li><a href="/teachers">Danh sách Giáo viên</a></li>
					<li><a href="/subjects">Danh sách Môn học</a></li>
					<li><a href="/classes">Danh sách Lớp học</a></li>
					<li><a href="/assignments">Danh sách Phân công</a></li>
				</ul>
			</nav>`,
		}
		// Tạo template từ dữ liệu
		tmpl := data.Template()
		// Thực thi template và gửi kết quả đến trình duyệt
		err := tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Lỗi khi hiển thị trang", http.StatusInternalServerError)
			return
		}
	})
	// Phục vụ các tệp tĩnh
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// Khởi động server
	log.Println("Đang chạy server trên cổng 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Lỗi khi khởi động server: %v", err)
	}
	log.Println("Server đã dừng.")
}

// Hàm main để khởi động ứng dụng
