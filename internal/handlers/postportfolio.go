package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	db "artsite/internal/db/queries"
	"artsite/internal/utils"
)

func PostPortfolio(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stderr, "POST portfolio endpoint reached \n")
		queries := db.New(pool)

		// Get form data
		id := r.PathValue("id")

		// Convert ID into uuid
		id_uuid, err := uuid.Parse(id)
		if err != nil {
			utils.RespondError(w, r, http.StatusInternalServerError, "Error parsing id")
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}

		// Add new art to db
		portfolio, err := queries.PostPortfolio(r.Context(), id_uuid)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error Creating: %v\n", err)
			utils.RespondError(w, r, http.StatusInternalServerError, "Failed to add art to portfolio")
			return
		}

		// Send response depending on accept header
		if utils.AcceptsJSON(r) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"art_id":   portfolio.ArtID,
				"added_at": portfolio.AddedAt.Time.String(),
			})
		} else {
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusCreated)
			// // In production, use templates. Here's a simple example:
			// fmt.Fprintf(w, "<h1>Admin Updated</h1><p>ID: %s</p><p>Title: %s</p><p>Description: %s</p><",
			// 	updatedadmin.ID.String(), updatedadmin.Username, updatedadmin.CreatedAt.Time.Format(time.RFC1123))
		}
	}
}
