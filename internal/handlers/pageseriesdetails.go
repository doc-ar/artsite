package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

	db "artsite/internal/db/queries"
	"artsite/internal/templates"
	"artsite/internal/utils"
)

func GetSeriesDetailsPage(ctx context.Context, pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := db.New(pool)
		series_name := r.PathValue("name")
		series, err := query.GetSeriesFromName(ctx, series_name)
		if err != nil {
			utils.RespondError(w, r, http.StatusInternalServerError, "Could not retrieve series details from db")
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}

		contents := templates.SeriesDetailsPage(series)
		layout := templates.Layout(contents)
		layout.Render(r.Context(), w)
	}
}
