package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"

	httpAdapter "hexagonal-minimal/internal/adapter/http"
	mysqlAdapter "hexagonal-minimal/internal/adapter/mysql"
	"hexagonal-minimal/internal/domain/user"
)

func main() {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	db, errDB := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName))
	if errDB != nil {
		panic(errDB)
	}

	userRepo := mysqlAdapter.NewUserRepository(db)
	userService := user.NewService(userRepo)
	handler := httpAdapter.NewHandler(userService)

	http.HandleFunc("/users", handler.Create)
	http.HandleFunc("/users/get", handler.Get)
	http.HandleFunc("/users/delete", handler.Delete)

	http.ListenAndServe(":8080", nil)
}
