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

func GetIntroVideo(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queries := db.New(pool)

		video, err := queries.GetURL(r.Context())
		if err != nil {
			utils.RespondError(w, r, http.StatusInternalServerError, "Could not retrieve video url from db")
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}

		templates.IntroVideo(video).Render(r.Context(), w)
	}
}
