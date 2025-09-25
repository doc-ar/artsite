package handlers

import (
	"net/http"

	"artsite/internal/mediaserver"
	"artsite/internal/templates"
)

func GetIntroVideo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		templates.IntroVideo(mediaserver.GetIntroVideoPath()).Render(r.Context(), w)
	}
}
