package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Sesuaikan dengan database kamu
	host := "localhost"
	user := "postgres"
	password := "12345"
	dbname := "todo_db"
	port := "5432"

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal konek database: ", err)
		os.Exit(1)
	}

	DB = database
	fmt.Println("Database connected ✔")
}
