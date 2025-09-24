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

func GetPortfolioAdmin(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stdout, "Get Admin Portfolio List endpoint reached\n")
		queries := db.New(pool)

		portfolio, err := queries.ListPortfolio(r.Context())
		if err != nil {
			utils.RespondError(w, r, http.StatusInternalServerError, "Could not retrieve art from db")
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}

		templates.AdminPortfolio(portfolio).Render(r.Context(), w)

		fmt.Fprintf(os.Stdout, "Get Admin Portfolio endpoint exited\n\n")
	}
}
