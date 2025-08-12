package view

type BaseView struct {
	ID          string `json:"id"`
	FullName    string `json:"full_name"`
	DisplayName string `json:"DisplayName"`
}

func GetView() []BaseView {
	view := []BaseView{
		{"1", "ma_giao_vien", "Mã Giáo Viên"},
		{"2", "ten_giao_vien", "Tên Giáo Viên"},
		{"3", "ten_TKB", "Tên trên TKB"},
	}
	return view
}
