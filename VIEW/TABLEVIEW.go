package VIEW

import (
	"GoSTKB/SQL"
	"bytes"
	"fmt"
	"text/template"
)

type TABLEVIEW struct {
	id       string
	theme_id string
	title    string
	table    *SQL.TABLE
}

func NewTableView(table *SQL.TABLE, ptitle *string) *TABLEVIEW {
	id := "id_" + table.GetName()
	theme_id := "theme_default"
	title := table.GetName()
	if ptitle != nil {
		title = *ptitle
	}
	return &TABLEVIEW{
		id:       id,
		theme_id: theme_id,
		title:    title,
		table:    table,
	}
}

type TableViewData struct {
	ID      string
	ThemeID string
	Name    string
	Columns []string
	Rows    []map[string]interface{}
	IsEdit  bool
}

func (tv *TABLEVIEW) ToHTML(isedit bool) string {
	const tpl = `
<p class='title'> Dữ liệu trong {{.Name}}</p>
<table border='1' id='{{.ID}}' class='{{.ThemeID}}'>
    <tr>
        {{range .Columns}}<th>{{.}}</th>{{end}}
    </tr>
    {{range .Rows}}
    <tr>
        {{range $.Columns}}
        <td>{{if .}}{{index $.Rows $.Index . }}{{else}}NULL{{end}}</td>
        {{end}}
        {{if $.IsEdit}}
        <td>
            <button class="edit-button" data-id="{{index . "id"}}">Sửa thông tin</button>
            <button class="delete-button" data-id="{{index . "id"}}">Xóa</button>
        </td>
        {{end}}
    </tr>
    {{end}}
</table>
{{if .IsEdit}}
<button class="add-btn">Thêm dòng mới</button>
{{end}}`

	data := TableViewData{
		ID:      tv.id,
		ThemeID: tv.theme_id,
		Name:    tv.table.GetName(),
		Columns: tv.table.GetColumns(),
		Rows:    tv.table.GetRows(),
		IsEdit:  isedit,
	}

	t, err := template.New("table").Parse(tpl)
	if err != nil {
		return fmt.Sprintf("%v", err) // Hoặc xử lý lỗi
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return fmt.Sprintf("%v", err) // Hoặc xử lý lỗi
	}
	return buf.String()
}
