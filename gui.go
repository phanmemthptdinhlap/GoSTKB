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

func (p *Table_view) GetHtml(id_table, class_table string, edit bool) string {
	texthtml := `<table id=` + id_table + ` class="` + class_table + `">
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
		col = []string{"id", "teacher_id", "subject_id", "class_id", "day_of_week", "start_time", "end_time"}
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
		htmlTable := data.GetHtml("table_"+table_name, "table", true)
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
func ShowForm(db *sql.DB, table_name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Xử lý dữ liệu từ form nếu có
		if r.Method == http.MethodPost {
			// Lấy dữ liệu từ form và lưu vào cơ sở dữ liệu
			// Ví dụ: Lấy tên giáo viên từ form
			name := r.FormValue("name")
			// Thực hiện lưu dữ liệu vào cơ sở dữ liệu
			_, err := db.Exec("INSERT INTO "+table_name+" (name) VALUES (?)", name)
			if err != nil {
				http.Error(w, fmt.Sprintf("Lỗi khi lưu dữ liệu: %v", err), http.StatusInternalServerError)
				return
			}
			// Chuyển hướng về trang danh sách sau khi lưu thành công
			http.Redirect(w, r, "/"+table_name, http.StatusSeeOther)
			return
		}
		// Hiển thị form để thêm mới dữ liệu
		data := &WebGui{
			Title:  "Thêm mới " + table_name,
			Header: "Thêm mới " + table_name,
			Footer: "© 2023 Quản lý thời khóa biểu",
			Body: `<form method="post">
				<label for="name">Tên:</label>
				<input type="text" id="name" name="name" required>
				<button type="submit">Lưu</button>
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
	// Gui khởi tạo giao diện web cho quản lý thời khóa biểu
}

func Gui(db *sql.DB) {
	// Tạo một template mới
	// Tạo một HTTP handler để phục vụ trang web
	http.HandleFunc("/teachers", ShowTable(db, "teachers"))
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
