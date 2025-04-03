package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type ConnMySQL struct {
	DB  *sql.DB
	Err string
}

func GetDBPool() *ConnMySQL {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbSchema := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPass, dbHost, dbSchema)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return &ConnMySQL{Err: fmt.Sprintf("error al abrir la base de datos: %v", err)}
	}

	db.SetMaxOpenConns(10)
	if err := db.Ping(); err != nil {
		db.Close()
		return &ConnMySQL{Err: fmt.Sprintf("error al verificar la conexión a la base de datos: %v", err)}
	}

	return &ConnMySQL{DB: db}
}
