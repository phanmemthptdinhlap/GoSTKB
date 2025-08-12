package item

type Item struct {
	Field []string `json:"field"`
}
type Items struct {
	Name string `json:"name"`
	TotalFields int `json:"total_fields"`
	Field1 string `json:"field1" name: "field1"
}