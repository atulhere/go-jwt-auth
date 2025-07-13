package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

func ConnectDB() *sql.DB {

	fmt.Println("Connecting to Database...")
	databaseConfig := getDatabaseConfig()

	fmt.Println("Database Config:", databaseConfig.Database)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		databaseConfig.User,
		databaseConfig.Password,
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.Database,
	)

	db, err := sql.Open("mysql", dsn)

	fmt.Println("Opening DB...", err)
	db.Ping()
	if err != nil {

		log.Fatalf("Error opening DB: %v", err)
	}

	// Now check if we can actually connect
	if err := db.Ping(); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	return db

}
