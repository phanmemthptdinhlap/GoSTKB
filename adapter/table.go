package adapter

type Table struct {
	Name string        `json:"table_name"`
	Head []string      `json:"table_head"`
	Row  []interface{} `json:"table_row"`
}
type Teacher struct {
	ID        string `json:"teacher_id"`
	FullName  string `json:"full_name"`
	ShortName string `json:"short_name"`
}
type Class struct {
	ID   string `json:"class_id"`
	Name string `json:"class_name"`
}
