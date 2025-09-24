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

func DeleteArt(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stdout, "DELETE /art endpoint reached\n")
		queries := db.New(pool)

		id := r.PathValue("id")
		id_uuid, err := uuid.Parse(id)
		if err != nil {
			utils.RespondError(w, r, http.StatusInternalServerError, "Error in ID conversion")
			fmt.Fprintf(os.Stderr, "Err: %v\n", err)
			return
		}

		deletedart, err := queries.DeleteArt(r.Context(), id_uuid)
		if err != nil {
			utils.RespondError(w, r, http.StatusInternalServerError, "Error Deleting Art")
			fmt.Fprintf(os.Stderr, "Err: %v\n", err)
			return
		}

		fmt.Fprintf(os.Stdout, "Deleted Art: %v\n", deletedart)
		fmt.Fprintf(os.Stdout, "DELETE /art endpoint successfully exited\n\n")
	}
}
