package handlers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	db "artsite/internal/db/queries"
	"artsite/internal/utils"
)

func Login(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queries := db.New(pool)

		// Get admin data from db
		username := r.FormValue("username")
		admin, err := queries.GetAdmin(r.Context(), username)
		if err != nil {
			utils.RespondError(w, r, http.StatusBadRequest, "User does not exist")
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			return
		}

		// Get and verify password
		password := r.FormValue("password")
		if !utils.VerifyPassword(password, admin.PasswordHash) {
			utils.RespondError(w, r, http.StatusForbidden, "Incorrect Password")
			fmt.Fprintf(os.Stderr, "Error: Incorrect Password\n")
			return
		}

		// Generate jwt token
		jwt, err := utils.CreateJWT(admin.Username)
		if err != nil {
			utils.RespondError(w, r, http.StatusInternalServerError, "Error generating JWT")
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			return
		}

		// Set secure jwt cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    jwt,
			Path:     "/",
			HttpOnly: true,
			// Secure:   true,
			SameSite: http.SameSiteStrictMode,
			Expires:  time.Now().Add(20 * time.Minute),
		})
	}
}
