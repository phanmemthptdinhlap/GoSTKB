package VIEW

import (
	"GoSTKB/SQL"
	"bytes"
	"fmt"
	"text/template"
)

type FORMVIEW struct {
	id       string
	theme_id string
	title    string
	table    *SQL.TABLE
}

func NewFormView(table *SQL.TABLE, ptitle *string) *FORMVIEW {
	id := "id_" + table.GetName()
	theme_id := "theme_default"
	title := table.GetName()
	if ptitle != nil {
		title = *ptitle
	}
	return &FORMVIEW{
		id:       id,
		theme_id: theme_id,
		title:    title,
		table:    table,
	}
}

type FormViewData struct {
	ID      string
	ThemeID string
	Title   string
	Name    string
	Columns []string
	Row     map[string]interface{}
	IsEdit  bool
}

func (fv *FORMVIEW) ToHTML(isedit bool) string {
	const tpl = `
<p class='title'>{{.Title}}</p>
<form id='{{.ID}}' class='{{.ThemeID}}'>
<table border='0'>
	{{range .Columns}}
    <tr>
		<td><label for="{{.}}">{{.}}</label></td>
		<td><input type="text" name="{{.}}" id="{{.}}" required>
    </tr>
	{{end}}
</table>
{{if .IsEdit}}
	<button class="add-btn">Thêm mới</button>
{{else}}
	<button class="edit-btn">Cập nhật</button>
{{end}}
<button class="can-btn">Hủy</button>`
	data := FormViewData{
		ID:      fv.id,
		ThemeID: fv.theme_id,
		Title:   fv.title,
		Name:    fv.table.GetName(),
		Columns: fv.table.GetColumns(),
		Row:     fv.table.GetRow(1),
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
