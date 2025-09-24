package handlers

import (
	"fmt"
	"net/http"
	"os"

	"artsite/internal/mediaserver"
	"artsite/internal/templates"
	"artsite/internal/utils"
)

func UploadVideo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stdout, "Video Upload Handler Reached\n")
		r.ParseMultipartForm(30 << 20)

		file, _, err := r.FormFile("hero_video")
		if err != nil {
			utils.RespondError(w, r, http.StatusBadRequest, "Error retrieving file from request")
			fmt.Fprintf(os.Stderr, "Err: %v\n", err)
			return
		}

		upload_err := mediaserver.UploadVideo(file)
		if upload_err != nil {
			utils.RespondError(w, r, http.StatusInternalServerError, "Error uploading file to the server")
			fmt.Fprintf(os.Stderr, "Err: %v\n", upload_err)
			return
		}
	}
}

func UploadImage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stdout, "Upload Handler Reached")
		r.ParseMultipartForm(20 << 20)

		file, _, err := r.FormFile("image")
		if err != nil {
			utils.RespondError(w, r, http.StatusBadRequest, "Error retrieving file from request")
			fmt.Fprintf(os.Stderr, "Err: %v\n", err)
			return
		}

		image_url, err := mediaserver.Upload(file)
		if err != nil {
			utils.RespondError(w, r, http.StatusInternalServerError, "Error uploading file to the server")
			fmt.Fprintf(os.Stderr, "Err: %v\n", err)
			return
		}

		templates.ImagePreview(image_url).Render(r.Context(), w)
	}
}
