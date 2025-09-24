package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func AcceptsJSON(r *http.Request) bool {
	accept := r.Header.Get("Accept")
	return strings.Contains(accept, "application/json")
}

func RespondError(w http.ResponseWriter, r *http.Request, code int, message string) {
	if AcceptsJSON(r) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(map[string]string{"error": message})
	} else {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(code)
		fmt.Fprintf(w, "<h1>Error %d</h1><p>%s</p>", code, message)
	}
}
