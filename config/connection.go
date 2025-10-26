package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

func ConnectDB() *sql.DB {

	fmt.Println("Connecting to Database...")
	var databaseConfig = GetDatabaseConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		databaseConfig.User,
		databaseConfig.Password,
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.Database,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening DB: %v", err)
	}
	defer db.Close() // Ensure the database connection is closed when done

	fmt.Println("Opening DB...")
	fmt.Println(dsn)

	if err := db.Ping(); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	return db

}
