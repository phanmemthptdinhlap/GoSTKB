package SQL

type TABLEVIEW struct {
	id       string
	theme_id string
	table    *TABLE
}

func NewTableView(table *TABLE) *TABLEVIEW {
	return &TABLEVIEW{
		table: table,
	}
}
func (tv *TABLEVIEW) ToHTML(isedit bool) string {
	html := "<table border='1' id='" + tv.id + "' class='" + tv.theme_id + "'>\n"
	html += "<tr>"
	for _, col := range tv.table.columns {
		html += "<th>" + col + "</th>"
	}
	html += "</tr>\n"
	for _, row := range tv.table.rows {
		html += "<tr>"
		for _, col := range tv.table.columns {
			value := row[col]
			if value == nil {
				value = "NULL"
			}
			html += "<td>" + value.(string) + "</td>"
			if isedit {
				html += `<td>
				<button class="edit-button" data-id="${row.id}">Sửa thông tin</button>
				<button class="delete-button" data-id="${row.id}">Xóa</button>
				</td>`
			}
		}
		html += "</tr>\n"
	}
	html += "</table>"
	if isedit {
		html += `<button class="add-btn">Thêm dòng mới</button>`
	}
	return html
}
