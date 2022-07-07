package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) http.Handler {
	h := &Handler{db: db}
	r := mux.NewRouter()
	r.HandleFunc("/", h.handleRoot).Methods("GET")
	r.HandleFunc("/user/{id}", h.handleUser).Methods("GET")
	return r
}

func main() {
	// Load environment variables from file.
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load environment variables: %v", err)
	}

	// Connect to PlanetScale database using DSN environment variable.
	db, err := gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatalf("failed to connect to planetscale: %v", err)
	}

	// Create an API handler which serves data from PlanetScale.
	handler := NewHandler(db)
	fmt.Println("first commit")
}
