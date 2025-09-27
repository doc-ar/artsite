package mediaserver

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/HugoSmits86/nativewebp"
)

func MediaHandler() http.Handler {
	fs := http.FileServer(http.Dir("./media"))
	return http.StripPrefix("/media/", fs)
}

func GetIntroVideoPath() string {
	matches, err := filepath.Glob("media/intro_video_*.webp")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error finding video: %v\n", err)
		return ""
	}

	if len(matches) == 0 {
		fmt.Fprintf(os.Stderr, "Could not find any intro video\n")
		return ""
	}

	return matches[0]
}

func RemoveExistingVideos() {
	matches, err := filepath.Glob("media/intro_video_*.webp")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error finding video: %v\n", err)
	}

	if len(matches) == 0 {
		fmt.Fprintf(os.Stderr, "Could not find any intro video\n")
	}

	for _, path := range matches {
		os.Remove(path)
	}
}

func UploadVideo(file multipart.File) error {
	defer file.Close()
	err := os.MkdirAll("media", 0o755)
	if err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	video_url := fmt.Sprintf("media/intro_video_%d.webm", time.Now().Unix())

	RemoveExistingVideos()
	introfile, err := os.Create(video_url)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer introfile.Close()

	_, read_err := introfile.ReadFrom(file)
	if read_err != nil {
		return fmt.Errorf("error reading from uploaded file: %w", read_err)
	}

	return nil
}

func Upload(file multipart.File) (string, error) {
	err := os.MkdirAll("media", 0o755)
	if err != nil {
		return "", fmt.Errorf("error creating directory: %w", err)
	}

	filename := fmt.Sprintf("media/%d.webp", time.Now().Unix())
	webpFile, err := os.Create(filename)
	if err != nil {
		return "", fmt.Errorf("error creating file: %w", err)
	}
	defer webpFile.Close()

	img, err := fileToImage(file)
	if err != nil {
		return "", fmt.Errorf("error decoding image: %w", err)
	}

	err = nativewebp.Encode(webpFile, img, nil)
	if err != nil {
		return "", fmt.Errorf("error encoding to webp: %w", err)
	}

	return filename, nil
}

func fileToImage(file multipart.File) (image.Image, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}
