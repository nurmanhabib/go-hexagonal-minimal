package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"

	"github.com/pressly/goose/v3"
)

func main() {
	if len(os.Args) < 2 {
		os.Args = append(os.Args, "up")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	db, errDB := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName))
	if errDB != nil {
		log.Fatal(errDB)
	}
	defer db.Close()

	if err := goose.SetDialect("mysql"); err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[1]
	if err := goose.RunContext(context.Background(), cmd, db, "migrations"); err != nil {
		log.Fatal(err)
	}

	log.Println("Migration completed successfully")
}
