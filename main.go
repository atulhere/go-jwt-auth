package main

import (
	"database/sql"
	"go-jwt-auth/config"
	"go-jwt-auth/handler"
	"go-jwt-auth/middleware"
	"net/http"
)

var db *sql.DB

func main() {

	//Load Configs
	config.LoadConfig()

	return
	// Connect to Database
	db = config.ConnectDB()

	//Registered routes
	routes()
	http.ListenAndServe(":8080", nil)

}

func routes() {

	// Create Handler for Healthz rout
	http.HandleFunc("/healthz", Healthz)
	http.HandleFunc("/auth/google/login", handler.GoogleLogin)
	http.HandleFunc("/auth/google/callback", handler.GoogleCallback)
	http.HandleFunc("/auth/refresh", handler.RefreshToken)
	http.Handle("/logout", middleware.AuthMiddleware(http.HandlerFunc(handler.Logout)))

}

func Healthz(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
