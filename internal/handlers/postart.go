package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	db "artsite/internal/db/queries"
	"artsite/internal/utils"
)

func PostArt(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stderr, "Endpoint for posting art is reached \n\n")
		queries := db.New(pool)

		// Get form data
		image_url := r.FormValue("image_url")
		title := r.FormValue("title")
		description := r.FormValue("description")
		width := r.FormValue("width")
		height := r.FormValue("height")
		series_id := r.FormValue("series_id")
		category := r.FormValue("category")
		medium := r.FormValue("medium")
		price := r.FormValue("price")

		// Convert width and height to integers
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

		// Convert Series ID into uuid
		series_uuid, err := uuid.Parse(series_id)
		if err != nil {
			utils.RespondError(w, r, http.StatusInternalServerError, "Error parsing id")
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}

		// Create params for db
		artparams := db.CreateArtParams{
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

		// Add new art to db
		newart, err := queries.CreateArt(r.Context(), artparams)
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
				"title":       newart.Title,
				"description": newart.Description,
				"series_id":   newart.SeriesID,
				"category":    newart.Category,
				"medium":      newart.Medium,
				"image_url":   newart.ImageUrl,
				"width":       newart.Width,
				"height":      newart.Height,
				"price":       newart.Price,
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
