package config

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func ConnectoDb () *sqlx.DB {
	
	dsn := os.Getenv("DSN")

	db, err := sqlx.Connect("mysql",dsn)
	
	if err != nil{
		log.Fatal("Error connecting to db", err.Error())
	}

	log.Println("✅ Connected to Railway MySQL successfully!")
  return db
}
