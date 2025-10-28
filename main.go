package main

import (
	"fmt"
	"go-jwt-auth/config"
	"go-jwt-auth/handler"
	"go-jwt-auth/middleware"
	"net/http"
	"time"
)

func main() {

	fmt.Println("Starting Server on :8080 — Build:", time.Now())

	// Initialize configuration
	config.Init()
	// Registered routes
	routes()

	// Start Server
	http.ListenAndServe(":8080", nil)

}

func routes() {

	// Create Handler for Healthz rout
	http.HandleFunc("/healthz", Healthz)
	http.HandleFunc("/login", handler.GoogleLogin)
	http.HandleFunc("/callback", handler.GoogleCallback)
	http.HandleFunc("/token-refresh", handler.RefreshToken)

	// 🔐 Protected route
	http.Handle("/logout", middleware.AuthMiddleware(http.HandlerFunc(handler.Logout)))
	http.Handle("/profile", middleware.AuthMiddleware(http.HandlerFunc(handler.ProfileHandler)))

}

func Healthz(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
