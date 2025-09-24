package handlers

import (
	"net/http"

	"artsite/internal/templates"
)

func GetHomePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contents := templates.HomePage()
		layout := templates.Layout(contents)
		layout.Render(r.Context(), w)
	}
}
