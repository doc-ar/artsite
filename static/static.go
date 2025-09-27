package static

import (
	"embed"
	"net/http"
)

//go:embed css/* script/* logo.svg
var files embed.FS

func StaticHandler() http.Handler {
	return http.StripPrefix("/static/", http.FileServer(http.FS(files)))
}
