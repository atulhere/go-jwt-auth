package handler

import (
	"encoding/json"
	"net/http"
)

// UserProfile — structure for response data
type UserProfile struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name,omitempty"`
}

// ProfileHandler — returns profile info if authorized
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Extract user info from context (set by AuthMiddleware)
	userId, ok := r.Context().Value("user_id").(int64)
	if !ok || userId < 0 {
		http.Error(w, "Unauthorized: no valid user session", http.StatusUnauthorized)
		return
	}

	// You could also fetch user data from DB here.
	profile := UserProfile{
		ID:    userId,
		Email: "atulsolanki30@gmail.com",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}
