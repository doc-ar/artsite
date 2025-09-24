package handlers

import (
	"net/http"

	"artsite/internal/templates"
)

func GetAdminPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		adminpage := templates.AdminPage()
		layout := templates.Layout(adminpage)
		layout.Render(r.Context(), w)
	}
}
