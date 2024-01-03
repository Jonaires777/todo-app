package config

import (
	"fmt"
	"os"
	"todo-app/models"
	"todo-app/repositories"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBconfig struct {
	User     string
	Password string
	DBname   string
}

func LoadEnv(path string) error {
	return godotenv.Load(path)
}

func CreateDBConnection(config DBconfig) (*gorm.DB, error) {
	dsn := fmt.Sprint("%s:%s@/%s?parseTime=True", config.User, config.Password, config.DBname)

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return conn, nil
}

func MigrateDB (conn *gorm.DB) error {
	if os.Getenv("MIGRATE") == "false" {
		return nil
	}

	return conn.Migrator().AutoMigrate(
		&models.Users{},
		&models.Todos{},
	)
}

func SetupDB(envPath string) (*repositories.Repository, error) {
	if dbEnvError := LoadEnv(envPath); dbEnvError != nil {
		return nil, dbEnvError
	}

	config := DBconfig{
		User: os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		DBname: os.Getenv("MYSQL_DATABASE"),
	}

	conn, err := CreateDBConnection(config)

	if err != nil {
		return nil, err
	}

	if migrateError := MigrateDB(conn); migrateError != nil{
		return nil, migrateError
	}

	return repositories.SetupRepository(conn), nil
}

func CloseDB (repository *repositories.Repository) error {
	sqlDB, _ := repository.DB.DB()

	return sqlDB.Close()
}

func PingDB (repository *repositories.Repository) error {
	sqlDB, _ := repository.DB.DB()

	return sqlDB.Ping()
}
