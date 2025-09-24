package handlers

import (
	"net/http"

	"artsite/internal/templates"
)

func GetGalleryPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contents := templates.GalleryPage()
		layout := templates.Layout(contents)
		layout.Render(r.Context(), w)
	}
}
