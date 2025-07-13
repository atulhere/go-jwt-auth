package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-jwt-auth/config"
	"go-jwt-auth/model"
	"net/http"
)

var db *sql.DB

func main() {

	//Load Configs
	config.LoadConfig()

	// Connect to Database
	db = config.ConnectDB()

	//Registered routes
	routes()
	http.ListenAndServe(":8080", nil)

}

func routes() {

	// Create Handler for Healthz rout
	http.HandleFunc("/healthz", Healthz)
	http.HandleFunc("/society", Society)
	http.HandleFunc("POST /login", Login)

	//Create Handler for get Society API

}

func Healthz(w http.ResponseWriter, r *http.Request) {

}

// API to get list of all socities
func Society(w http.ResponseWriter, r *http.Request) {

	// Check if the request method is GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Set the response header to application/json
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "Maa")
	var society model.Society

	row := db.QueryRow("SELECT id, name,address,pin,created_at FROM `society` WHERE id = 1;").Scan(&society.Id, &society.Name, &society.Address, &society.Pin, &society.Created)

	if row != nil {

		// if row == sql.ErrNoRows {
		// 	// No rows found, handle accordingly
		// 	return http.NotFound(w, r) // Return 404 if no rows found

		// } else {

		// 	return http.Error(w, r) // Return 500 for other errors
		// }

		json.NewEncoder(w).Encode(society) // Encode the society data as JSON

		// Return the society data as JSON
		// Simple query to check DB connection
	}

	fmt.Println(society)
	// If query is successful, write the response

}

func Login(w http.ResponseWriter, r *http.Request) {

	// Validate the request
	var user = model.User{Username: "test", Password: "303030303030"}

	json.NewDecoder(r.Body).Decode(&user)

	// Check for empty credentials
	if user.Username == "" || user.Password == "" {

		http.Error(w, "Invalid credentials", http.StatusUnauthorized)

		return

	}
	// Check the credentials against the database
	if user.Username != "test" || user.Password != "303030303030" {

		http.Error(w, "Invalid credentials", http.StatusUnauthorized)

		return
	}

	// Generate the JWT token
	token, err := GenerateAuthToken(user.Username)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	// Generate the refresh token
	refreshToken, err := GenerateRefreshToken(user.Username)
	if err != nil {
		http.Error(w, "Could not generate refresh token", http.StatusInternalServerError)
	}

	// Set the token in the response header
	json.NewEncoder(w).Encode(map[string]string{
		"token":         token,
		"refresh_token": refreshToken})

}
