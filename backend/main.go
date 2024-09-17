package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var db *sql.DB

func connectDB() {
	var err error
	dbHost := "host=" + os.Getenv("DB_HOST")
	dbUser := "user=" + os.Getenv("DB_USER")
	dbPassword := "password=" + os.Getenv("DB_PASSWORD")

	// Fetch database credentials from environment variables or hardcode for development
	connStr := dbHost + " " + dbUser + " " + dbPassword + " " + "dbname=mydatabase sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Verify the connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database:", err)
	}
	fmt.Println("Successfully connected to the database!")
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Query the database
	rows, err := db.Query("SELECT id, name FROM mytable")
	if err != nil {
		http.Error(w, "Database query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Write the query results to the response
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			http.Error(w, "Failed to scan result", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "ID: %d, Name: %s\n", id, name)
	}
}

func main() {
	// Connect to the database
	connectDB()

	// Close the database connection when the app exits
	defer db.Close()

	http.HandleFunc("/", handler)
	fmt.Println("Server starting on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}
