package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	db "artsite/internal/db/queries"
	"artsite/internal/utils"
)

func UpdateArt(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queries := db.New(pool)

		// Get form data
		id := r.FormValue("id")
		image_url := r.FormValue("image_url")
		title := r.FormValue("title")
		description := r.FormValue("description")
		width := r.FormValue("width")
		height := r.FormValue("height")
		series_id := r.FormValue("series_id")
		category := r.FormValue("category")
		medium := r.FormValue("medium")
		price := r.FormValue("price")

		// Convert Art uuid from string
		id_uuid, err := uuid.Parse(id)
		if err != nil {
			utils.RespondError(w, r, http.StatusInternalServerError, "Error parsing art id into uuid")
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}

		// Convert Series UUID from string
		series_uuid, err := uuid.Parse(series_id)
		if err != nil {
			utils.RespondError(w, r, http.StatusInternalServerError, "Error parsing series id into uuid")
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}

		// Convert integer values
		width_int, err := strconv.Atoi(width)
		if err != nil {
			utils.RespondError(w, r, http.StatusInternalServerError, "Error converting width to int")
			fmt.Fprintf(os.Stderr, "Err: %v\n", err)
			return
		}
		height_int, err := strconv.Atoi(height)
		if err != nil {
			utils.RespondError(w, r, http.StatusInternalServerError, "Error converting height to int")
			fmt.Fprintf(os.Stderr, "Err: %v\n", err)
			return
		}
		price_int, err := strconv.Atoi(price)
		if err != nil {
			utils.RespondError(w, r, http.StatusInternalServerError, "Error converting price to int")
			fmt.Fprintf(os.Stderr, "Err: %v\n", err)
			return
		}

		// Make params for update art query
		artparams := db.UpdateArtParams{
			ID:          id_uuid,
			Title:       title,
			Description: description,
			SeriesID:    series_uuid,
			Category:    category,
			Medium:      medium,
			ImageUrl:    image_url,
			Width:       int32(width_int),
			Height:      int32(height_int),
			Price:       int32(price_int),
		}

		// Execute the update art query
		updatedart, err := queries.UpdateArt(r.Context(), artparams)
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
				"title":       updatedart.Title,
				"description": updatedart.Description,
				"series_id":   updatedart.SeriesID,
				"category":    updatedart.Category,
				"medium":      updatedart.Medium,
				"image_url":   updatedart.ImageUrl,
				"width":       updatedart.Width,
				"height":      updatedart.Height,
				"price":       updatedart.Price,
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
