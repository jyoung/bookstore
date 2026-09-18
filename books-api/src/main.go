package main

import (
	"books-api/config"
	"books-api/database"
	"books-api/server"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {

	cfg := config.Load()

	db, err := database.NewDatabase(cfg.DB)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	defer db.Close()
	log.Println("connected to database")

	router := server.NewRouter(db)

	srv := server.NewServer(cfg.Port, router)

	if err := srv.Run(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
