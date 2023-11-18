package config

import (
	"database/sql"
	"fmt"
	"todo-app/models"
	"todo-app/repositories"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() {
	var err error
	dsn := "root:root@/todoapp?parseTime=True"
	repositories.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	fmt.Println("Connection Opened to database")
	repositories.DBConn.AutoMigrate(&models.Todos{})
	repositories.DBConn.AutoMigrate(&models.Users{})
	fmt.Println("Database Migrated")
}

func CloseDB() error {
	var sqlDB *sql.DB

	return sqlDB.Close()
}