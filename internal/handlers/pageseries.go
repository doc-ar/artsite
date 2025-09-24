package handlers

import (
	"net/http"
	"artsite/internal/templates"
)

func GetSeriesPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contents := templates.SeriesPage()
		layout := templates.Layout(contents)
		layout.Render(r.Context(), w)
	}
}
