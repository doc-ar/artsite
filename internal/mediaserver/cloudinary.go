package mediaserver

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadVideoCLD(file multipart.File) (string, error) {
	ctx := context.Background()

	cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		return "", fmt.Errorf("error creating cloudinary instance: %w", err)
	}

	params := uploader.UploadParams{
		PublicID:     "intro_video",
		ResourceType: "video",
		Folder:       "intro",
		Overwrite:    api.Bool(true),
		Invalidate:   api.Bool(true),
	}

	result, err := cld.Upload.Upload(ctx, file, params)
	if err != nil {
		return "", fmt.Errorf("error uploading to cloudinary: %w", err)
	}

	return result.SecureURL, nil
}

func UploadImageCLD(file multipart.File) (string, error) {
	ctx := context.Background()

	cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		return "", fmt.Errorf("error creating cloudinary instance: %w", err)
	}

	params := uploader.UploadParams{
		PublicID:     strconv.FormatInt(time.Now().UnixNano(), 10),
		Folder:       "gallery",
		ResourceType: "image",
	}

	result, err := cld.Upload.Upload(ctx, file, params)
	if err != nil {
		return "", fmt.Errorf("error uploading to cloudinary: %w", err)
	}

	optimized_url := fmt.Sprintf("https://res.cloudinary.com/%s/image/upload/f_auto,q_auto/%s", cld.Config.Cloud.CloudName, result.PublicID)

	return optimized_url, nil
}

func DeleteImageCLD(imageurl string) error {
	ctx := context.Background()
	cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		return fmt.Errorf("error creating cloudinary instance: %w", err)
	}

	imageid := urlToId(imageurl)
	if imageid == "" {
		return fmt.Errorf("could not extract public_id from URL: %s", imageurl)
	}

	params := admin.DeleteAssetsParams{
		PublicIDs: []string{imageid},
		AssetType: api.Image,
	}

	_, deleteErr := cld.Admin.DeleteAssets(ctx, params)
	if deleteErr != nil {
		return fmt.Errorf("error deleting file from cloudinary: %w", deleteErr)
	}

	return nil
}

func urlToId(fileURL string) string {
	// Drop query string if present
	parts := strings.SplitN(fileURL, "?", 2)
	clean := parts[0]

	// Find the “/upload/” part — after this is (transformations +) public_id + extension
	idx := strings.Index(clean, "/upload/")
	if idx < 0 {
		return ""
	}
	tail := clean[idx+len("/upload/"):]

	// Sometimes there are multiple segments (e.g. transformations), so detect and drop them.
	// A typical tail may be like: "f_auto,q_auto/gallery/12345.webp"
	segs := strings.Split(tail, "/")
	// If the first segment contains commas or underscores (common in transformation strings), drop it
	if len(segs) > 1 && (strings.Contains(segs[0], ",") || strings.Contains(segs[0], "_")) {
		segs = segs[1:]
	}
	// Rejoin the rest as the public_id + extension
	pidWithExt := strings.Join(segs, "/")
	// Strip the file extension
	ext := filepath.Ext(pidWithExt)
	pid := strings.TrimSuffix(pidWithExt, ext)

	return pid
}
