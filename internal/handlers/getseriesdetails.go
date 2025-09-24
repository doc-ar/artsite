package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

	db "artsite/internal/db/queries"
	"artsite/internal/templates"
	"artsite/internal/utils"
)

func GetSeriesDetails(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stdout, "Get series details endpoint reached\n")
		queries := db.New(pool)
		series_name := r.PathValue("name")

		artlist, err := queries.ListSeriesDetails(r.Context(), series_name)
		if err != nil {
			utils.RespondError(w, r, http.StatusInternalServerError, "Could not retrieve art from db")
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}

		templates.SeriesDetailsList(artlist).Render(r.Context(), w)

		fmt.Fprintf(os.Stdout, "Get series details endpoint exited\n\n")
	}
}
