package mediaserver

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/HugoSmits86/nativewebp"
)

func MediaHandler() http.Handler {
	fs := http.FileServer(http.Dir("./media"))
	return http.StripPrefix("/media/", fs)
}

func UploadVideo(file multipart.File) error {
	defer file.Close()
	err := os.MkdirAll("media", 0755)
	if err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	os.Remove("media/hero_video.webm")
	introfile, err := os.Create("media/hero_video.webm")
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
	err := os.MkdirAll("media", 0755)
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
