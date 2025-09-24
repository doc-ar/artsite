package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	db "artsite/internal/db/queries"
	"artsite/internal/templates"
	"artsite/internal/utils"
)

func GetSeriesForm(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stdout, "Get Series Edit Form endpoint reached\n")
		queries := db.New(pool)

		id := r.PathValue("id")
		id_uuid, err := uuid.Parse(id)
		if err != nil {
			utils.RespondError(w, r, http.StatusInternalServerError, "Error parsing id")
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}

		seriesItem, err := queries.GetSeries(r.Context(), id_uuid)
		if err != nil {
			utils.RespondError(w, r, http.StatusInternalServerError, "Could not retrieve art from db")
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}

		templates.EditSeriesForm(seriesItem).Render(r.Context(), w)

		fmt.Fprintf(os.Stdout, "Get Series Edit Form endpoint exited\n\n")
	}
}

func GetArtForm(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stdout, "Get Art Edit Form endpoint reached\n")
		queries := db.New(pool)

		id := r.PathValue("id")
		id_uuid, err := uuid.Parse(id)
		if err != nil {
			utils.RespondError(w, r, http.StatusInternalServerError, "Error parsing id")
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}

		artItem, err := queries.GetArt(r.Context(), id_uuid)
		if err != nil {
			utils.RespondError(w, r, http.StatusInternalServerError, "Could not retrieve art from db")
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}

		templates.EditArtForm(artItem).Render(r.Context(), w)

		fmt.Fprintf(os.Stdout, "Get Art Edit Form endpoint exited\n\n")
	}
}
