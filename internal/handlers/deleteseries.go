package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	db "artsite/internal/db/queries"
	"artsite/internal/utils"
)

func DeleteSeries(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stderr, "DELETE series endpoint reached\n")
		queries := db.New(pool)

		id := r.PathValue("id")
		id_uuid, err := uuid.Parse(id)
		if err != nil {
			utils.RespondError(w, r, http.StatusInternalServerError, "Error in ID conversion")
			fmt.Fprintf(os.Stderr, "Err: %v\n", err)
			return
		}

		deletedseries, err := queries.DeleteSeries(r.Context(), id_uuid)
		if err != nil {
			utils.RespondError(w, r, http.StatusInternalServerError, "Error Deleting Art")
			fmt.Fprintf(os.Stderr, "Err: %v\n", err)
			return
		}

		fmt.Fprintf(os.Stdout, "Deleted Series: %v\n", deletedseries)
		fmt.Fprintf(os.Stdout, "DELETE series endpoint successfully exited\n\n")
	}
}
