package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"go-jwt-auth/config"
	"go-jwt-auth/model"
	"go-jwt-auth/utility"

	"golang.org/x/oauth2"
)

func GoogleLogin(w http.ResponseWriter, r *http.Request) {

	url := config.GoogleOAuthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	//fmt.Println("URL String is   ", url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func GoogleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	token, err := config.GoogleOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Println(err)
		fmt.Println("+++++++++++++++++++++++")
		http.Error(w, "Token exchange failed", http.StatusBadRequest)
		return
	}

	idToken := token.Extra("id_token").(string)
	payload, err := utility.VerifyGoogleIDToken(context.Background(), idToken)
	if err != nil {
		http.Error(w, "Invalid ID token", http.StatusUnauthorized)
		return
	}

	var user model.User
	row := config.DB.QueryRow("SELECT id, email FROM users WHERE google_id=?", payload.Subject)
	if err := row.Scan(&user.ID, &user.Email); err == sql.ErrNoRows {
		res, _ := config.DB.Exec("INSERT INTO users (google_id, email, name, picture, created_at, last_login_at) VALUES (?, ?, ?, ?, NOW(), NOW())",
			payload.Subject, payload.Claims["email"], payload.Claims["name"], payload.Claims["picture"])
		user.ID, _ = res.LastInsertId()
	} else {
		config.DB.Exec("UPDATE users SET last_login_at=NOW() WHERE id=?", user.ID)
	}

	accessToken, _ := utility.GenerateAccessToken(user.ID)
	refreshToken, error := utility.RandomString(64)

	if error != nil {

		http.Error(w, "failed to generate refresh token", http.StatusInternalServerError)
		return
	}

	config.DB.Exec("UPDATE users SET refresh_token=? WHERE id=?", refreshToken, user.ID)

	resp := map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}
	json.NewEncoder(w).Encode(resp)
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {

	type req struct {
		RefreshToken string `json:"refresh_token"`
	}
	var request req
	//var req struct{ RefreshToken string `refresh_token` }
	json.NewDecoder(r.Body).Decode(&request)

	var userID int64
	err := config.DB.QueryRow("SELECT id FROM users WHERE refresh_token=?", request.RefreshToken).Scan(&userID)
	if err != nil {
		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
		return
	}

	newAccessToken, _ := utility.GenerateAccessToken(userID)
	json.NewEncoder(w).Encode(map[string]string{"access_token": newAccessToken})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int64)
	config.DB.Exec("UPDATE users SET refresh_token=NULL WHERE id=?", userID)
	w.WriteHeader(http.StatusOK)
}
