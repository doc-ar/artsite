package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	db "artsite/internal/db/queries"
	"artsite/internal/utils"
)

func UpdateSeries(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queries := db.New(pool)

		// Get form data
		id := r.FormValue("id")
		image_url := r.FormValue("image_url")
		name := r.FormValue("name")
		description := r.FormValue("description")

		// Convert series uuid from string
		series_uuid, err := uuid.Parse(id)
		if err != nil {
			utils.RespondError(w, r, http.StatusInternalServerError, "Error parsing art id into uuid")
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}

		// Make params for update art query
		seriesparams := db.UpdateSeriesParams{
			ID:          series_uuid,
			Name:        name,
			Description: description,
			CoverImg:    image_url,
		}

		// Execute the update art query
		updatedseries, err := queries.UpdateSeries(r.Context(), seriesparams)
		if err != nil {
			log.Printf("database error: %v", err)
			utils.RespondError(w, r, http.StatusInternalServerError, "Failed to update art")
			return
		}

		// Send response based on accept header
		if utils.AcceptsJSON(r) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"name":        updatedseries.Name,
				"description": updatedseries.Description,
				"cover_img":   updatedseries.CoverImg,
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
