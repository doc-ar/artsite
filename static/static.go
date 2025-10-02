package static

import (
	"embed"
	"net/http"
)

//go:embed css/* script/* favicon.png
var files embed.FS

func StaticHandler() http.Handler {
	return http.StripPrefix("/static/", http.FileServer(http.FS(files)))
}
