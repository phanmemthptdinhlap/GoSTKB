package main

import (
	"GOSTKB/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	_ "gorm.io/driver/sqlite" // Import the SQLite driver
)
func main() {
	db, err := gorm.Open(sqlite.Open("timetable.db"), &gorm.Config{})
    if err != nil {
        panic("Không thể kết nối cơ sở dữ liệu")
    }
}