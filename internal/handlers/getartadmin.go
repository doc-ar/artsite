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

func GetArtAdmin(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stdout, "Get Admin Art List endpoint reached\n")
		queries := db.New(pool)

		artlist, err := queries.ListArt(r.Context())
		if err != nil {
			utils.RespondError(w, r, http.StatusInternalServerError, "Could not retrieve art from db")
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}

		templates.AdminArtList(artlist).Render(r.Context(), w)

		fmt.Fprintf(os.Stdout, "Get Art endpoint exited\n\n")
	}
}
