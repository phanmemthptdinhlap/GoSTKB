package main 

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"GoSTKB/models" // Adjust the import path as necessary
)
func main() {
	// Initialize Gin router
	router := gin.Default()

	// Connect to the database
	db, err := gorm.Open(sqlite.Open("stkb.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Migrate the GiaoVien model
	if err := db.AutoMigrate(&models.GiaoVien{}); err != nil {
		panic("failed to migrate database")
	}
}