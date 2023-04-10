package main

import (
	"log"

	"github.com/dimasyudhana/alterra-group-project-2/config"
)

func main() {
	// Database connection
	cfg := config.InitConfiguration()
	db, err := config.GetConnection(cfg)
	if err != nil {
		log.Fatalf("cannot connect to database: %v", err)
	}

	// Check database connection
	sqlDb, err := db.DB()
	if err != nil {
		log.Fatalf("cannot get sql.DB instance: %v", err)
	}
	err = sqlDb.Ping()
	if err != nil {
		log.Fatalf("cannot ping database: %v", err)
	}
	log.Println("Connected with database!")
}
