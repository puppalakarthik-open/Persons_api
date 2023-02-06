package storage

import (
	"fmt"
	"os"
	"log"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

)
type Config struct {
	Host   string
	Port   string
	User   string
	DBname string
}
func NewConnection() (*gorm.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	var config = Config{
		Host:   os.Getenv("DB_HOST"),
		Port:   os.Getenv("DB_PORT"),
		User:   os.Getenv("DB_USER"),
		DBname: os.Getenv("DB_NAME"),
	}
	dsn := fmt.Sprintf("host=%s user=%s password='' dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.Host,config.User,config.DBname,config.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return db, err
	}
	return db, nil
}