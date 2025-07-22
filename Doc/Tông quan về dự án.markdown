Tôi sẽ hướng dẫn bạn tạo một web server bằng Go sử dụng **GORM** (ORM cho SQLite) và **Gin** (web framework) để quản lý thời khóa biểu theo yêu cầu của bạn. Dự án sẽ bao gồm các bảng: **Giáo viên**, **Lớp học**, **Phân công giảng dạy**, **Môn học**, **Tiết học theo tuần**, cùng với các file HTML, CSS, JS và sử dụng template. Tôi sẽ giải thích chi tiết từng bước, bao gồm các lệnh và cách sử dụng thư viện.

---

### **1. Thiết lập môi trường phát triển**
Trước tiên, bạn cần cài đặt Go và thiết lập môi trường làm việc.

#### **1.1. Cài đặt Go**
- Tải và cài đặt Go từ [trang chính thức](https://golang.org/dl/) (phiên bản mới nhất tính đến 07/2025 là Go 1.22 hoặc cao hơn).
- Kiểm tra cài đặt:
  ```bash
  go version
  ```

#### **1.2. Tạo thư mục dự án**
```bash
mkdir timetable-manager
cd timetable-manager
go mod init timetable-manager
```

#### **1.3. Cài đặt các thư viện cần thiết**
Cài đặt **Gin**, **GORM**, và driver SQLite:
```bash
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
```

- **Giải thích**:
  - `github.com/gin-gonic/gin`: Framework web nhẹ, nhanh, dùng để tạo API và xử lý HTTP request.
  - `gorm.io/gorm`: ORM giúp tương tác với cơ sở dữ liệu (SQLite) dễ dàng.
  - `gorm.io/driver/sqlite`: Driver để kết nối GORM với SQLite.

#### **1.4. Cấu trúc thư mục**
Tạo cấu trúc thư mục như sau:
```
timetable-manager/
├── main.go
├── models/
│   ├── teacher.go
│   ├── class.go
│   ├── subject.go
│   ├── assignment.go
│   ├── schedule.go
├── handlers/
│   ├── teacher_handler.go
│   ├── class_handler.go
│   ├── subject_handler.go
│   ├── assignment_handler.go
│   ├── schedule_handler.go
├── templates/
│   ├── index.html
│   ├── teacher.html
│   ├── class.html
│   ├── subject.html
│   ├── assignment.html
│   ├── schedule.html
├── static/
│   ├── css/
│   │   └── style.css
│   ├── js/
│   │   └── main.js
├── go.mod
├── go.sum
├── timetable.db
```

---

### **2. Thiết lập cơ sở dữ liệu với GORM**
Tạo các model trong thư mục `models/` tương ứng với các bảng.

#### **2.1. Model Giáo viên (`models/teacher.go`)**
```go
package models

import "gorm.io/gorm"

type Teacher struct {
    gorm.Model
    TeacherID   string `gorm:"unique;not null" json:"teacher_id"`
    FullName    string `json:"full_name"`
    DisplayName string `json:"display_name"`
}
```

- **Giải thích**:
  - `gorm.Model`: Cung cấp các trường mặc định như `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` cho quản lý bản ghi.
  - `TeacherID`: Mã giáo viên, duy nhất, không null.
  - `json:"..."`: Tag để định dạng dữ liệu khi trả về JSON.

#### **2.2. Model Lớp học (`models/class.go`)**
```go
package models

import "gorm.io/gorm"

type Class struct {
    gorm.Model
    ClassID string `gorm:"unique;not null" json:"class_id"`
    ClassName string `json:"class_name"`
    Grade     string `json:"grade"`
}
```

#### **2.3. Model Môn học (`models/subject.go`)**
```go
package models

import "gorm.io/gorm"

type Subject struct {
    gorm.Model
    SubjectID   string `gorm:"unique;not null" json:"subject_id"`
    SubjectName string `json:"subject_name"`
}
```

#### **2.4. Model Phân công giảng dạy (`models/assignment.go`)**
```go
package models

import "gorm.io/gorm"

type Assignment struct {
    gorm.Model
    AssignmentID string `gorm:"unique;not null" json:"assignment_id"`
    TeacherID    string `json:"teacher_id"`
    SubjectID    string `json:"subject_id"`
    ClassID      string `json:"class_id"`
}
```

#### **2.5. Model Tiết học theo tuần (`models/schedule.go`)**
```go
package models

import "gorm.io/gorm"

type Schedule struct {
    gorm.Model
    Week         int    `json:"week"`
    AssignmentID string `json:"assignment_id"`
    TotalLessons int    `json:"total_lessons"`
}
```

#### **2.6. Kết nối và khởi tạo cơ sở dữ liệu**
Tạo file `main.go` để thiết lập kết nối SQLite và khởi tạo bảng:
```go
package main

import (
    "timetable-manager/models"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func main() {
    // Kết nối SQLite
    db, err := gorm.Open(sqlite.Open("timetable.db"), &gorm.Config{})
    if err != nil {
        panic("Không thể kết nối cơ sở dữ liệu")
    }

    // Tự động tạo bảng
    db.AutoMigrate(&models.Teacher{}, &models.Class{}, &models.Subject{}, &models.Assignment{}, &models.Schedule{})

    // Khởi tạo Gin
    r := gin.Default()

    // Load template HTML
    r.LoadHTMLGlob("templates/*")

    // Cấu hình static file
    r.Static("/static", "./static")

    // Định nghĩa route cho trang chủ
    r.GET("/", func(c *gin.Context) {
        c.HTML(200, "index.html", nil)
    })

    // Khởi động server
    r.Run(":8080")
}
```

- **Giải thích**:
  - `gorm.Open(sqlite.Open("timetable.db"), &gorm.Config{})`: Kết nối đến SQLite, tạo file `timetable.db` nếu chưa tồn tại.
  - `db.AutoMigrate`: Tự động tạo các bảng dựa trên model.
  - `r.LoadHTMLGlob`: Load tất cả file HTML trong thư mục `templates`.
  - `r.Static`: Phục vụ file tĩnh (CSS, JS) từ thư mục `static`.
  - `r.Run(":8080")`: Khởi động server ở cổng 8080.

---

### **3. Xử lý API với Gin**
Tạo các handler trong thư mục `handlers/` để xử lý CRUD (Create, Read, Update, Delete) cho từng bảng.

#### **3.1. Handler Giáo viên (`handlers/teacher_handler.go`)**
```go
package handlers

import (
    "timetable-manager/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
)

// Dependency Injection: Truyền db vào handler
type TeacherHandler struct {
    DB *gorm.DB
}

// Tạo giáo viên
func (h *TeacherHandler) CreateTeacher(c *gin.Context) {
    var teacher models.Teacher
    if err := c.ShouldBindJSON(&teacher); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    h.DB.Create(&teacher)
    c.JSON(http.StatusOK, teacher)
}

// Lấy danh sách giáo viên
func (h *TeacherHandler) GetTeachers(c *gin.Context) {
    var teachers []models.Teacher
    h.DB.Find(&teachers)
    c.JSON(http.StatusOK, teachers)
}

// Cập nhật giáo viên
func (h *TeacherHandler) UpdateTeacher(c *gin.Context) {
    id := c.Param("id")
    var teacher models.Teacher
    if err := h.DB.First(&teacher, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy giáo viên"})
        return
    }
    if err := c.ShouldBindJSON(&teacher); err != thế

System: Bạn đang xây dựng một dự án web server bằng Go sử dụng **GORM** và **Gin** để quản lý thời khóa biểu. Dưới đây là hướng dẫn chi tiết tiếp tục từ phần trước, hoàn thiện các handler, giao diện frontend, và cung cấp code đầy đủ cho CSS, JS, và template HTML. Tôi sẽ giải thích rõ ràng từng lệnh và cách sử dụng các thư viện.

---

### **3. Xử lý API với Gin (tiếp tục)**

Tiếp tục hoàn thiện handler cho bảng **Giáo viên** và thêm các handler cho các bảng còn lại (**Lớp học**, **Môn học**, **Phân công giảng dạy**, **Tiết học theo tuần**).

#### **3.1. Hoàn thiện Handler Giáo viên (`handlers/teacher_handler.go`)**
```go
package handlers

import (
    "timetable-manager/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
)

// TeacherHandler chứa DB để truy cập cơ sở dữ liệu
type TeacherHandler struct {
    DB *gorm.DB
}

// Tạo giáo viên
func (h47 *TeacherHandler) CreateTeacher(c *gin.Context) {
    var teacher models.Teacher
    if err := c.ShouldBindJSON(&teacher); err != nil {
        c.JSON(http Debtors.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    h.DB.Create(&teacher)
    c.JSON(http.StatusOK, teacher)
}

// Lấy danh sách giáo viên
func (h *TeacherHandler) GetTeachers(c *gin.Context) {
    var teachers []models.Teacher
    h.DB.Find(&teachers)
    c.JSON(http.StatusOK, teachers)
}

// Cập nhật giáo viên
func (h *TeacherHandler) UpdateTeacher(c *gin.Context) {
    id := c.Param("id")
    var teacher models.Teacher
    if err := h.DB.First(&teacher, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy giáo viên"})
        return
    }
    if err := c.ShouldBindJSON(&teacher); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    h.DB.Save(&teacher)
    c.JSON(http.StatusOK, teacher)
}

// Xóa giáo viên
func (h *TeacherHandler) DeleteTeacher(c *gin.Context) {
    id := c.Param("id")
    var teacher models.Teacher
    if err := h.DB.First(&teacher, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy giáo viên"})
        return
    }
    h.DB.Delete(&teacher)
    c.JSON(http.StatusOK, gin.H{"message": "Xóa giáo viên thành công"})
}
```

- **Giải thích**:
  - `c.ShouldBindJSON(&teacher)`: Ràng buộc dữ liệu JSON từ request vào struct `Teacher`.
  - `h.DB.Create/Save/Delete`: Các phương thức GORM để tạo, cập nhật, xóa bản ghi.
  - `c.JSON`: Trả về phản hồi JSON với mã trạng thái HTTP (200, 400, 404).

#### **3.2. Handler Lớp học (`handlers/class_handler.go`)**
```go
package handlers

import (
    "timetable-manager/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
)

type ClassHandler struct {
    DB *gorm.DB
}

func (h *ClassHandler) CreateClass(c *gin.Context) {
    var class models.Class
    if err := c.ShouldBindJSON(&class); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    h.DB.Create(&class)
    c.JSON(http.StatusOK, class)
}

func (h *ClassHandler) GetClasses(c *gin.Context) {
    var classes []models.Class
    h.DB.Find(&classes)
    c.JSON(http.StatusOK, classes)
}

func (h *ClassHandler) UpdateClass(c *gin.Context) {
    id := c.Param("id")
    var class models.Class
    if err := h.DB.First(&class, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy lớp học"})
        return
    }
    if err := c.ShouldBindJSON(&class); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    h.DB.Save(&class)
    c.JSON(http.StatusOK, class)
}

func (h *ClassHandler) DeleteClass(c *gin.Context) {
    id := c.Param("id")
    var class models.Class
    if err := h.DB.First(&class, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy lớp học"})
        return
    }
    h.DB.Delete(&class)
    c.JSON(http.StatusOK, gin.H{"message": "Xóa lớp học thành công"})
}
```

#### **3.3. Handler Môn học, Phân công giảng dạy, Tiết học theo tuần**
Tương tự, tạo các file `subject_handler.go`, `assignment_handler.go`, `schedule_handler.go` với cấu trúc giống `teacher_handler.go` và `class_handler.go`, thay đổi model tương ứng (`Subject`, `Assignment`, `Schedule`). Tôi sẽ cung cấp một ví dụ cho `subject_handler.go`, các file khác bạn có thể tự triển khai tương tự:

```go
package handlers

import (
    "timetable-manager/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
)

type SubjectHandler struct {
    DB *gorm.DB
}

func (h *SubjectHandler) CreateSubject(c *gin.Context) {
    var subject models.Subject
    if err := c.ShouldBindJSON(&subject); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    h.DB.Create(&subject)
    c.JSON(http.StatusOK, subject)
}

func (h *SubjectHandler) GetSubjects(c *gin.Context) {
    var subjects []models.Subject
    h.DB.Find(&subjects)
    c.JSON(http.StatusOK, subjects)
}

func (h *SubjectHandler) UpdateSubject(c *gin.Context) {
    id := c.Param("id")
    var subject models.Subject
    if err := h.DB.First(&subject, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy môn học"})
        return
    }
    if err := c.ShouldBindJSON(&subject); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    h.DB.Save(&subject)
    c.JSON(http.StatusOK, subject)
}

func (h *SubjectHandler) DeleteSubject(c *gin.Context) {
    id := c.Param("id")
    var subject models.Subject
    if err := h.DB.First(&subject, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy môn học"})
        return
    }
    h.DB.Delete(&subject)
    c.JSON(http.StatusOK, gin.H{"message": "Xóa môn học thành công"})
}
```

#### **3.4. Cập nhật `main.go` để thêm route**
Cập nhật file `main.go` để thêm các route API và trang HTML:

```go
package main

import (
    "timetable-manager/handlers"
    "timetable-manager/models"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func main() {
    // Kết nối SQLite
    db, err := gorm.Open(sqlite.Open("timetable.db"), &gorm.Config{})
    if err != nil {
        panic("Không thể kết nối cơ sở dữ liệu")
    }

    // Tự động tạo bảng
    db.AutoMigrate(&models.Teacher{}, &models.Class{}, &models.Subject{}, &models.Assignment{}, &models.Schedule{})

    // Khởi tạo Gin
    r := gin.Default()

    // Load template HTML
    r.LoadHTMLGlob("templates/*")

    // Cấu hình static file
    r.Static("/static", "./static")

    // Khởi tạo handler
    teacherHandler := &handlers.TeacherHandler{DB: db}
    classHandler := &handlers.ClassHandler{DB: db}
    subjectHandler := &handlers.SubjectHandler{DB: db}
    assignmentHandler := &handlers.AssignmentHandler{DB: db}
    scheduleHandler := &handlers.ScheduleHandler{DB: db}

    // Route trang HTML
    r.GET("/", func(c *gin.Context) {
        c.HTML(200, "index.html", nil)
    })
    r.GET("/teachers", func(c *gin.Context) {
        c.HTML(200, "teacher.html", nil)
    })
    r.GET("/classes", func(c *gin.Context) {
        c.HTML(200, "class.html", nil)
    })
    r.GET("/subjects", func(c *gin.Context) {
        c.HTML(200, "subject.html", nil)
    })
    r.GET("/assignments", func(c *gin.Context) {
        c.HTML(200, "assignment.html", nil)
    })
    r.GET("/schedules", func(c *gin.Context) {
        c.HTML(200, "schedule.html", nil)
    })

    // Route API
    api := r.Group("/api")
    {
        // Giáo viên
        api.POST("/teachers", teacherHandler.CreateTeacher)
        api.GET("/teachers", teacherHandler.GetTeachers)
        api.PUT("/teachers/:id", teacherHandler.UpdateTeacher)
        api.DELETE("/teachers/:id", teacherHandler.DeleteTeacher)

        // Lớp học
        api.POST("/classes", classHandler.CreateClass)
        api.GET("/classes", classHandler.GetClasses)
        api.PUT("/classes/:id", classHandler.UpdateClass)
        api.DELETE("/classes/:id", classHandler.DeleteClass)

        // Môn học
        api.POST("/subjects", subjectHandler.CreateSubject)
        api.GET("/subjects", subjectHandler.GetSubjects)
        api.PUT("/subjects/:id", subjectHandler.UpdateSubject)
        api.DELETE("/subjects/:id", subjectHandler.DeleteSubject)

        // Phân công giảng dạy
        api.POST("/assignments", assignmentHandler.CreateAssignment)
        api.GET("/assignments", assignmentHandler.GetAssignments)
        api.PUT("/assignments/:id", assignmentHandler.UpdateAssignment)
        api.DELETE("/assignments/:id", assignmentHandler.DeleteAssignment)

        // Tiết học theo tuần
        api.POST("/schedules", scheduleHandler.CreateSchedule)
        api.GET("/schedules", scheduleHandler.GetSchedules)
        api.PUT("/schedules/:id", scheduleHandler.UpdateSchedule)
        api.DELETE("/schedules/:id", scheduleHandler.DeleteSchedule)
    }

    // Khởi động server
    r.Run(":8080")
}
```

- **Giải thích**:
  - `r.Group("/api")`: Nhóm các route API dưới prefix `/api`.
  - Mỗi handler được gắn với các phương thức HTTP (`POST`, `GET`, `PUT`, `DELETE`) để xử lý CRUD.

---

### **4. Giao diện Frontend (HTML, CSS, JS)**

#### **4.1. Template HTML (`templates/index.html`)**
Tạo trang chủ với menu điều hướng:
```html
<!DOCTYPE html>
<html lang="vi">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Quản lý thời khóa biểu</title>
    <link rel="stylesheet" href="/static/css/style.css">
</head>
<body>
    <header>
        <h1>Quản lý thời khóa biểu</h1>
        <nav>
            <a href="/teachers">Giáo viên</a>
            <a href="/classes">Lớp học</a>
            <a href="/subjects">Môn học</a>
            <a href="/assignments">Phân công giảng dạy</a>
            <a href="/schedules">Tiết học theo tuần</a>
        </nav>
    </header>
    <main>
        <h2>Chào mừng đến với hệ thống quản lý thời khóa biểu</h2>
    </main>
    <script src="/static/js/main.js"></script>
</body>
</html>
```

#### **4.2. Template Giáo viên (`templates/teacher.html`)**
Tạo form và bảng để quản lý giáo viên:
```html
<!DOCTYPE html>
<html lang="vi">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Quản lý giáo viên</title>
    <link rel="stylesheet" href="/static/css/style.css">
</head>
<body>
    <header>
        <h1>Quản lý giáo viên</h1>
        <nav>
            <a href="/">Trang chủ</a>
        </nav>
    </header>
    <main>
        <h2>Thêm giáo viên</h2>
        <form id="teacherForm">
            <input type="text" id="teacher_id" placeholder="Mã giáo viên" required>
            <input type="text" id="full_name" placeholder="Họ tên" required>
            <input type="text" id="display_name" placeholder="Họ tên trên TKB" required>
            <button type="submit">Thêm</button>
        </form>
        <h2>Danh sách giáo viên</h2>
        <table>
            <thead>
                <tr>
                    <th>Mã giáo viên</th>
                    <th>Họ tên</th>
                    <th>Họ tên trên TKB</th>
                    <th>Hành động</th>
                </tr>
            </thead>
            <tbody id="teacherTable"></tbody>
        </table>
    </main>
    <script src="/static/js/main.js"></script>
</body>
</html>
```

- Tương tự, tạo các file `class.html`, `subject.html`, `assignment.html`, `schedule.html` với form và bảng phù hợp với từng bảng dữ liệu.

#### **4.3. CSS (`static/css/style.css`)**
```css
body {
    font-family: Arial, sans-serif;
    margin: 0;
    padding: 0;
    background-color: #f4f4f4;
}

header {
    background-color: #333;
    color: white;
    padding: 10px;
    text-align: center;
}

nav a {
    color: white;
    margin: 0 15px;
    text-decoration: none;
}

nav a:hover {
    text-decoration: underline;
}

main {
    padding: 20px;
    max-width: 1200px;
    margin: auto;
}

form {
    margin-bottom: 20px;
}

form input {
    padding: 8px;
    margin: 5px;
    width: 200px;
}

form button {
    padding: 8px 16px;
    background-color: #28a745;
    color: white;
    border: none;
    cursor: pointer;
}

form button:hover {
    background-color: #218838;
}

table {
    width: 100%;
    border-collapse: collapse;
    margin-top: 20px;
}

table, th, td {
    border: 1px solid #ddd;
}

th, td {
    padding: 10px;
    text-align: left;
}

th {
    background-color: #f2f2f2;
}

button.delete {
    background-color: #dc3545;
}

button.delete:hover {
    background-color: #c82333;
}
```

- **Giải thích**:
  - CSS cơ bản, sử dụng màu sắc tương phản tốt cho cả theme sáng và tối.
  - Form và bảng được thiết kế đơn giản, dễ sử dụng.

#### **4.4. JavaScript (`static/js/main.js`)**
```javascript
// Gọi API để lấy danh sách giáo viên
function loadTeachers() {
    fetch('/api/teachers')
        .then(response => response.json())
        .then(data => {
            const table = document.getElementById('teacherTable');
            table.innerHTML = '';
            data.forEach(teacher => {
                const row = `<tr>
                    <td>${teacher.teacher_id}</td>
                    <td>${teacher.full_name}</td>
                    <td>${teacher.display_name}</td>
                    <td>
                        <button onclick="editTeacher(${teacher.id})">Sửa</button>
                        <button class="delete" onclick="deleteTeacher(${teacher.id})">Xóa</button>
                    </td>
                </tr>`;
                table.innerHTML += row;
            });
        });
}

// Thêm giáo viên
document.getElementById('teacherForm').addEventListener('submit', function(e) {
    e.preventDefault();
    const teacher = {
        teacher_id: document.getElementById('teacher_id').value,
        full_name: document.getElementById('full_name').value,
        display_name: document.getElementById('display_name').value
    };
    fetch('/api/teachers', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(teacher)
    }).then(() => {
        loadTeachers();
        document.getElementById('teacherForm').reset();
    });
});

// Xóa giáo viên
function deleteTeacher(id) {
    fetch(`/api/teachers/${id}`, { method: 'DELETE' })
        .then(() => loadTeachers());
}

// Sửa giáo viên (tạm thời hiển thị form nhập lại)
function editTeacher(id) {
    alert('Chức năng sửa đang được phát triển! ID: ' + id);
}

// Load danh sách khi trang được tải
document.addEventListener('DOMContentLoaded', loadTeachers);
```

- **Giải thích**:
  - `fetch`: Gửi request HTTP tới API để lấy, thêm, xóa giáo viên.
  - `loadTeachers`: Cập nhật bảng danh sách giáo viên.
  - Chức năng sửa (`editTeacher`) hiện chỉ là placeholder, bạn có thể mở rộng bằng cách thêm form sửa.

- Tương tự, viết JS cho các bảng khác (`classes`, `subjects`, `assignments`, `schedules`) bằng cách thay đổi endpoint API và dữ liệu.

---

### **5. Chạy và kiểm tra ứng dụng**
1. Chạy server:
   ```bash
   go run main.go
   ```
2. Mở trình duyệt, truy cập `http://localhost:8080`.
3. Kiểm tra các trang `/teachers`, `/classes`, v.v., và thử thêm/xóa dữ liệu.

---

### **6. Lưu ý và mở rộng**
- **Xử lý lỗi**: Thêm kiểm tra dữ liệu đầu vào (validation) trong handler.
- **Bảo mật**: Thêm xác thực (authentication) và phân quyền (authorization) nếu cần.
- **Tối ưu giao diện**: Sử dụng framework như Bootstrap để giao diện đẹp hơn.
- **Tính năng nâng cao**: Thêm chức năng tìm kiếm, lọc, hoặc hiển thị thời khóa biểu dạng lưới.

Nếu bạn cần thêm chi tiết về bất kỳ phần nào (ví dụ: handler cho các bảng khác, JS chi tiết, hoặc tính năng nâng cao), hãy cho tôi biết!