package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

	db "artsite/internal/db/queries"
	"artsite/internal/utils"
)

func PostSeries(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stdout, "Post series endpointed reached \n")
		queries := db.New(pool)

		// Get form data
		image_url := r.FormValue("image_url")
		name := r.FormValue("name")
		description := r.FormValue("description")

		// Create params for db
		seriesparams := db.PostSeriesParams{
			Name:        name,
			Description: description,
			CoverImg:    image_url,
		}

		// Add new art to db
		newseries, err := queries.PostSeries(r.Context(), seriesparams)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error Creating: %v\n", err)
			utils.RespondError(w, r, http.StatusInternalServerError, "Failed to update admin")
			return
		}

		// Send response depending on accept header
		if utils.AcceptsJSON(r) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"name":        newseries.Name,
				"description": newseries.Description,
				"cover_img":   newseries.CoverImg,
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
