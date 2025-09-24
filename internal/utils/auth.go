package utils

import (
	"net/http"

	"artsite/internal/templates"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jwtCookie, err := r.Cookie("token")
		if err != nil {
			// Token missing: render login page
			loginpage := templates.LoginPage()
			layout := templates.Layout(loginpage)
			layout.Render(r.Context(), w)
			return
		}

		_, err = VerifyJWT(jwtCookie.Value)
		if err != nil {
			// Token invalid or expired: render login page
			loginpage := templates.LoginPage()
			layout := templates.Layout(loginpage)
			layout.Render(r.Context(), w)
			return
		}

		// Token valid: proceed to handler
		next.ServeHTTP(w, r)
	}
}
