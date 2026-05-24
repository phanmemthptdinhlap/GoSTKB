package main

import (
    "fmt"
    "net/http"
)

type WebPage struct {
    Title   		string
		mux     		*http.ServeMux
		SetPageHome		func()
		SetPageLogin		func()
		SetPageClass		func()
		SetPageTeacher	func()
		SetPageSubject	func()
		SetPageAssignment	func()
		SetPageAdmin		func()
		SetPageAbout		func()
}

func (p *WebPage) init(mux *http.ServeMux) {
    p.mux = mux
		if p.SetPageHome != nil {
			p.SetPageHome()
		}
		if p.SetPageLogin != nil {
			p.SetPageLogin()
		}
		if p.SetPageClass != nil {
			p.SetPageClass()
		}
		if p.SetPageTeacher != nil {
			p.SetPageTeacher()
		}
		if p.SetPageSubject != nil {
			p.SetPageSubject()
		}
		if p.SetPageAssignment != nil {
			p.SetPageAssignment()
		}
		if p.SetPageAdmin != nil {
			p.SetPageAdmin()
		}
		if p.SetPageAbout != nil {
			p.SetPageAbout()
		}
}

// Khai báo biến toàn cục
var page WebPage

func main() {
    mux := http.NewServeMux()
		page.init(mux)
    fmt.Println("Server đang chạy tại http://localhost:8080")
    http.ListenAndServe(":8080", mux)
}
