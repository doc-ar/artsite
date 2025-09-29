package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

	db "artsite/internal/db/queries"
	"artsite/internal/mediaserver"
	"artsite/internal/templates"
	"artsite/internal/utils"
)

func UploadVideo(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stdout, "Video Upload Handler Reached\n")
		r.ParseMultipartForm(30 << 20)

		file, _, err := r.FormFile("hero_video")
		if err != nil {
			utils.RespondError(w, r, http.StatusBadRequest, "Error retrieving file from request")
			fmt.Fprintf(os.Stderr, "Err: %v\n", err)
		}

		video_url, err := mediaserver.UploadVideoCLD(file)
		if err != nil || video_url == "" {
			utils.RespondError(w, r, http.StatusInternalServerError, "Error uploading file to the server")
			fmt.Fprintf(os.Stderr, "Err: %v\n", err)
		}

		queries := db.New(pool)
		updated_video, err := queries.SetVideo(r.Context(), video_url)

		fmt.Fprintf(os.Stdout, "Uploaded Video: %s\n", updated_video.Url)
		fmt.Fprintf(os.Stdout, "Video Upload Handler Exited\n\n")
	}
}

func UploadImage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stdout, "Upload Image Handler Reached\n")
		r.ParseMultipartForm(20 << 20)

		file, _, err := r.FormFile("image_input")
		if err != nil {
			utils.RespondError(w, r, http.StatusBadRequest, "Error retrieving file from request")
			fmt.Fprintf(os.Stderr, "Err: %v\n", err)
			return
		}

		image_url, err := mediaserver.UploadImageCLD(file)
		if err != nil {
			utils.RespondError(w, r, http.StatusInternalServerError, "Error uploading file to the server")
			fmt.Fprintf(os.Stderr, "Err %d: %v\n", http.StatusInternalServerError, err)
			return
		}

		templates.ImagePreview(image_url).Render(r.Context(), w)
	}
}
