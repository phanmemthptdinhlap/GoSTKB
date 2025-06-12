package SQL

import (
	"database/sql"
	"strings"
)

type TABLE struct {
	name    string
	columns []string
	rows    []map[string]interface{}
}

func NewTable(name string, columns []string) *TABLE {
	return &TABLE{
		name:    name,
		columns: columns,
		rows:    make([]map[string]interface{}, 0),
	}
}
func (t *TABLE) SyncDataBase(db *sql.DB) {
	for _, row := range t.rows {
		if row["id"] == nil {
			// Thêm mới
			db.Exec("INSERT INTO "+t.name+" ("+strings.Join(t.columns, ", ")+") VALUES ("+placeholders(len(t.columns))+")", rowValues(row, t.columns)...)
		} else {
			// Cập nhật
			updates := make([]string, 0)
			for _, col := range t.columns {
				if col != "id" {
					updates = append(updates, col+" = ?")
				}
			}
			query := "UPDATE " + t.name + " SET " + strings.Join(updates, ", ") + " WHERE id = ?"
			values := rowValues(row, t.columns)
			values = append(values, row["id"])
			_, err := db.Exec(query, values...)
			if err != nil {
				panic("Cập nhật dữ liệu thất bại: " + err.Error())
			}
		}
	}
	rows, err := db.Query("SELECT * FROM " + t.name)
	if err != nil {
		panic("Lấy dữ liệu từ CSDL thất bại: " + err.Error())
	}
	defer rows.Close()
	t.rows = make([]map[string]interface{}, 0)
	for rows.Next() {
		row := make(map[string]interface{})
		columns := make([]interface{}, len(t.columns))
		for i := range columns {
			columns[i] = new(interface{})
		}
		if err := rows.Scan(columns...); err != nil {
			panic("Lấy dữ liệu từ CSDL thất bại: " + err.Error())
		}
		for i, col := range t.columns {
			row[col] = *(columns[i].(*interface{}))
		}
		t.rows = append(t.rows, row)
	}
}

func placeholders(n int) string {
	if n <= 0 {
		return ""
	}
	s := make([]string, n)
	for i := range s {
		s[i] = "?"
	}
	return strings.Join(s, ", ")
}

func rowValues(row map[string]interface{}, s []string) []interface{} {
	values := make([]interface{}, len(s))
	for i, col := range s {
		values[i] = row[col]
	}
	return values
}
func (t *TABLE) GetNameAndColumns() (string, []string) {
	return t.name, t.columns
}
func (t *TABLE) AddRow(row map[string]interface{}) {
	if len(row) != len(t.columns) {
		panic("Số lượng cột không khớp với bảng")
	}
	t.rows = append(t.rows, row)
}
func (t *TABLE) GetRows() []map[string]interface{} {
	return t.rows
}
func (t *TABLE) GetName() string {
	return t.name
}
func (t *TABLE) GetColumns() []string {
	return t.columns
}
func (t *TABLE) GetRowCount() int {
	return len(t.rows)
}
func (t *TABLE) GetColumnCount() int {
	return len(t.columns)
}
func (t *TABLE) GetRow(index int) map[string]interface{} {
	if index < 0 || index >= len(t.rows) {
		panic("Chỉ mục hàng không hợp lệ")
	}
	return t.rows[index]
}
func (t *TABLE) GetColumn(index int) string {
	if index < 0 || index >= len(t.columns) {
		panic("Chỉ mục cột không hợp lệ")
	}
	return t.columns[index]
}
func (t *TABLE) GetValue(rowIndex int, columnIndex int) interface{} {
	if rowIndex < 0 || rowIndex >= len(t.rows) {
		panic("Chỉ mục hàng không hợp lệ")
	}
	if columnIndex < 0 || columnIndex >= len(t.columns) {
		panic("Chỉ mục cột không hợp lệ")
	}
	return t.rows[rowIndex][t.columns[columnIndex]]
}
func (t *TABLE) SetValue(rowIndex int, columnIndex int, value interface{}) {
	if rowIndex < 0 || rowIndex >= len(t.rows) {
		panic("Chỉ mục hàng không hợp lệ")
	}
	if columnIndex < 0 || columnIndex >= len(t.columns) {
		panic("Chỉ mục cột không hợp lệ")
	}
	t.rows[rowIndex][t.columns[columnIndex]] = value
}
func (t *TABLE) DeleteRow(index int) {
	if index < 0 || index >= len(t.rows) {
		panic("Chỉ mục hàng không hợp lệ")
	}
	t.rows = append(t.rows[:index], t.rows[index+1:]...)
}
func (t *TABLE) DeleteColumn(index int) {
	if index < 0 || index >= len(t.columns) {
		panic("Chỉ mục cột không hợp lệ")
	}
	columnName := t.columns[index]
	t.columns = append(t.columns[:index], t.columns[index+1:]...)
	for i := range t.rows {
		delete(t.rows[i], columnName)
	}
}
func (t *TABLE) RenameColumn(index int, newName string) {
	if index < 0 || index >= len(t.columns) {
		panic("Chỉ mục cột không hợp lệ")
	}
	oldName := t.columns[index]
	t.columns[index] = newName
	for i := range t.rows {
		t.rows[i][newName] = t.rows[i][oldName]
		delete(t.rows[i], oldName)
	}
}
func (t *TABLE) Clear() {
	t.rows = make([]map[string]interface{}, 0)
}
func (t *TABLE) GetColumnIndex(name string) int {
	for i, col := range t.columns {
		if col == name {
			return i
		}
	}
	return -1 // Trả về -1 nếu không tìm thấy cột
}
func (t *TABLE) GetRowIndex(row map[string]interface{}) int {
	for i, r := range t.rows {
		if len(r) != len(row) {
			continue // Bỏ qua nếu số lượng cột không khớp
		}
		match := true
		for key, value := range row {
			if r[key] != value {
				match = false
				break
			}
		}
		if match {
			return i // Trả về chỉ mục hàng nếu tìm thấy
		}
	}
	return -1 // Trả về -1 nếu không tìm thấy hàng
}
func (t *TABLE) GetColumnValues(columnName string) []interface{} {
	index := t.GetColumnIndex(columnName)
	if index == -1 {
		panic("Cột không tồn tại")
	}
	values := make([]interface{}, len(t.rows))
	for i, row := range t.rows {
		values[i] = row[t.columns[index]]
	}
	return values
}
func (t *TABLE) GetRowValues(rowIndex int) map[string]interface{} {
	if rowIndex < 0 || rowIndex >= len(t.rows) {
		panic("Chỉ mục hàng không hợp lệ")
	}
	return t.rows[rowIndex]
}
func (t *TABLE) GetColumnNames() []string {
	return t.columns
}
func (t *TABLE) GetRowNames() []map[string]interface{} {
	rowNames := make([]map[string]interface{}, len(t.rows))
	for i, row := range t.rows {
		rowNames[i] = make(map[string]interface{})
		for _, col := range t.columns {
			rowNames[i][col] = row[col]
		}
	}
	return rowNames
}
func (t *TABLE) GetTableInfo() map[string]interface{} {
	info := make(map[string]interface{})
	info["name"] = t.name
	info["columns"] = t.columns
	info["row_count"] = len(t.rows)
	info["rows"] = t.rows
	return info
}
func (t *TABLE) Clone() *TABLE {
	clone := NewTable(t.name, t.columns)
	for _, row := range t.rows {
		newRow := make(map[string]interface{})
		for key, value := range row {
			newRow[key] = value
		}
		clone.AddRow(newRow)
	}
	return clone
}
func (t *TABLE) Merge(other *TABLE) {
	if t.name != other.name {
		panic("Không thể hợp nhất hai bảng với tên khác nhau")
	}
	if len(t.columns) != len(other.columns) {
		panic("Số lượng cột không khớp giữa hai bảng")
	}
	for _, col := range other.columns {
		if t.GetColumnIndex(col) == -1 {
			panic("Cột " + col + " không tồn tại trong bảng gốc")
		}
	}
	for _, row := range other.rows {
		t.AddRow(row)
	}
}
func (t *TABLE) Filter(predicate func(map[string]interface{}) bool) *TABLE {
	filtered := NewTable(t.name, t.columns)
	for _, row := range t.rows {
		if predicate(row) {
			newRow := make(map[string]interface{})
			for key, value := range row {
				newRow[key] = value
			}
			filtered.AddRow(newRow)
		}
	}
	return filtered
}
func (t *TABLE) SortByColumn(columnName string, ascending bool) {
	index := t.GetColumnIndex(columnName)
	if index == -1 {
		panic("Cột không tồn tại")
	}
	sortFunc := func(i, j int) bool {
		if ascending {
			return t.rows[i][t.columns[index]].(string) < t.rows[j][t.columns[index]].(string)
		}
		return t.rows[i][t.columns[index]].(string) > t.rows[j][t.columns[index]].(string)
	}
	for i := 0; i < len(t.rows)-1; i++ {
		for j := i + 1; j < len(t.rows); j++ {
			if sortFunc(i, j) {
				t.rows[i], t.rows[j] = t.rows[j], t.rows[i]
			}
		}
	}
}
func (t *TABLE) GroupBy(columnName string) map[interface{}][]map[string]interface{} {
	index := t.GetColumnIndex(columnName)
	if index == -1 {
		panic("Cột không tồn tại")
	}
	grouped := make(map[interface{}][]map[string]interface{})
	for _, row := range t.rows {
		key := row[t.columns[index]]
		if _, exists := grouped[key]; !exists {
			grouped[key] = make([]map[string]interface{}, 0)
		}
		newRow := make(map[string]interface{})
		for key, value := range row {
			newRow[key] = value
		}
		grouped[key] = append(grouped[key], newRow)
	}
	return grouped
}
func (t *TABLE) Aggregate(columnName string, aggFunc func([]interface{}) interface{}) map[interface{}]interface{} {
	index := t.GetColumnIndex(columnName)
	if index == -1 {
		panic("Cột không tồn tại")
	}
	aggregated := make(map[interface{}]interface{})
	for _, row := range t.rows {
		key := row[t.columns[index]]
		if _, exists := aggregated[key]; !exists {
			aggregated[key] = make([]interface{}, 0)
		}
		aggregated[key] = append(aggregated[key].([]interface{}), row[t.columns[index]])
	}
	for key, values := range aggregated {
		aggregated[key] = aggFunc(values.([]interface{}))
	}
	return aggregated
}
