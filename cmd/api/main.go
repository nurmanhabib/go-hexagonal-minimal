package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	httpAdapter "hexagonal-minimal/internal/adapter/http"
	mongoAdapter "hexagonal-minimal/internal/adapter/mongodb"
	mysqlAdapter "hexagonal-minimal/internal/adapter/mysql"
	"hexagonal-minimal/internal/domain/user"
)

func main() {
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8080"
	}

	var userRepo user.Repository

	selectedAdapter := "mysql" // or mongo
	userRepo = resolveUserRepository(selectedAdapter)

	userService := user.NewService(userRepo)
	handler := httpAdapter.NewHandler(userService)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /users", handler.Get)
	mux.HandleFunc("POST /users", handler.Create)
	mux.HandleFunc("DELETE /users", handler.Delete)

	addr := fmt.Sprintf(":%s", appPort)

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	log.Printf("HTTP server running on port %s\n", addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func resolveUserRepository(adapter string) user.Repository {
	switch adapter {
	case "mysql":
		db, err := createMySQLConnection()
		if err != nil {
			log.Fatal(err)
		}

		return mysqlAdapter.NewUserRepository(db)

	case "mongo":
		db, err := createMongoDatabase()
		if err != nil {
			log.Fatal(err)
		}

		return mongoAdapter.NewUserRepository(db)

	default:
		log.Fatalf("unknown adapter: %s", adapter)
		return nil
	}
}

func createMySQLConnection() (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	return sql.Open("mysql", dsn)
}

func createMongoDatabase() (*mongo.Database, error) {
	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB_NAME")

	if mongoURI == "" || dbName == "" {
		return nil, fmt.Errorf("MONGO_URI or MONGO_DB_NAME is not set")
	}

	client, err := mongo.Connect(
		options.Client().ApplyURI(mongoURI),
	)
	if err != nil {
		return nil, err
	}

	return client.Database(dbName), nil
}
