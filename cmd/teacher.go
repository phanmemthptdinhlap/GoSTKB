package main

import (
	"fmt"
	"html/template"
	"net/http"
)
func (p *WebPage) SetPageTeacher(){
		page.mux.HandleFunc("/teacher",func(w http.ResponseWriter, r *http.Request){
			temp,err:=template.ParseFiles("templates/teacher.html","templates/base.html")
			if err!=nil{
				fmt.Println(err)
				http.Error(w,"Internal Server Error",http.StatusInternalServerError)
				return
			}
			data:=struct{
				Title string
			}{
				Title:"Teacher",
			}	
			err=temp.ExecuteTemplate(w,"base",data)
			if err!=nil{
				fmt.Println(err)
				http.Error(w,"Lỗi trang web",http.StatusInternalServerError)
				return
			}
		})
	}
