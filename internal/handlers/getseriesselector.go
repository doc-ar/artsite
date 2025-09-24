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

func GetSeriesSelector(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stdout, "Get series selector endpoint reached\n")
		queries := db.New(pool)

		serieslist, err := queries.ListSeries(r.Context())
		if err != nil {
			utils.RespondError(w, r, http.StatusInternalServerError, "Could not retrieve series from db")
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}

		templates.SeriesSelector(serieslist).Render(r.Context(), w)

		fmt.Fprintf(os.Stdout, "Get series selector endpoint exited\n\n")
	}
}
