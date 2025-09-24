package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	db "artsite/internal/db/queries"
	"artsite/internal/utils"
)

func UpdateAdmin(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queries := db.New(pool)

		// Get form data
		username := r.FormValue("username")
		old_password := r.FormValue("old_password")
		new_password := r.FormValue("new_password")

		// Check inputs
		if username == "" || old_password == "" || new_password == "" {
			utils.RespondError(w, r, http.StatusBadRequest, "Username and password are required")
			return
		}

		// Get admin data from username
		admin, err := queries.GetAdmin(r.Context(), username)
		if err != nil {
			utils.RespondError(w, r, http.StatusInternalServerError, "Failed getting admin")
			return
		}

		// Verify if onld password is correct
		if !utils.VerifyPassword(old_password, admin.PasswordHash) {
			utils.RespondError(w, r, http.StatusForbidden, "Invalid Password Entered")
			return
		}

		// Hash the new password
		new_hash, err := utils.HashPassword(new_password)
		if err != nil {
			utils.RespondError(w, r, http.StatusInternalServerError, "Unable to hash new password")
			return
		}

		// Update the admin info
		updatedadminparams := db.UpdateAdminParams{
			ID:           admin.ID,
			PasswordHash: new_hash,
			Username:     username,
		}
		updatedadmin, err := queries.UpdateAdmin(r.Context(), updatedadminparams)
		if err != nil {
			log.Printf("database error: %v", err)
			utils.RespondError(w, r, http.StatusInternalServerError, "Failed to update admin")
			return
		}

		// Send response depending on accept header
		if utils.AcceptsJSON(r) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"id":         updatedadmin.ID.String(),
				"username":   updatedadmin.Username,
				"created_at": updatedadmin.CreatedAt,
			})
		} else {
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusCreated)
			// In production, use templates. Here's a simple example:
			fmt.Fprintf(w, "<h1>Admin Updated</h1><p>ID: %s</p><p>Username: %s</p><p>Created: %s</p>",
				updatedadmin.ID.String(), updatedadmin.Username, updatedadmin.CreatedAt.Time.Format(time.RFC1123))
		}
	}
}
