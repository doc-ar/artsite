package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

	"artsite/internal/handlers"
	"artsite/internal/mediaserver"
	"artsite/internal/utils"
	"artsite/static"
)

func main() {
	ctx := context.Background()
	router := http.NewServeMux()

	println(os.Getenv("DATABASE_URL"))

	dbpool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	utils.CreateAdmin(dbpool, ctx)

	// File and Media handlers
	router.Handle("GET /static/", static.StaticHandler())
	router.Handle("GET /media/", mediaserver.MediaHandler())

	// Page Handlers
	router.HandleFunc("GET /{$}", handlers.GetHomePage())
	router.HandleFunc("GET /gallery", handlers.GetGalleryPage())
	router.HandleFunc("GET /series", handlers.GetSeriesPage())
	router.HandleFunc("GET /series/{name}", handlers.GetSeriesDetailsPage(ctx, dbpool))
	router.HandleFunc("GET /admin", utils.AuthMiddleware(handlers.GetAdminPage()))

	// CRUD Handler
	router.HandleFunc("GET /art/", handlers.GetArtPortfolio(dbpool))
	router.HandleFunc("GET /portfolio/", handlers.GetPortfolio(dbpool))
	router.HandleFunc("GET /serieslist/", handlers.GetSeries(dbpool))
	router.HandleFunc("GET /series/{name}/list", handlers.GetSeriesDetails(dbpool))
	router.HandleFunc("GET /introvideo", handlers.GetIntroVideo())

	// Authenticated CRUD handlers
	router.HandleFunc("GET /admin/art/", utils.AuthMiddleware(handlers.GetArtAdmin(dbpool)))
	router.HandleFunc("GET /admin/art/form/{id}", utils.AuthMiddleware(handlers.GetArtForm(dbpool)))
	router.HandleFunc("GET /admin/series/", utils.AuthMiddleware(handlers.GetSeriesAdmin(dbpool)))
	router.HandleFunc("GET /admin/series/form/{id}", utils.AuthMiddleware(handlers.GetSeriesForm(dbpool)))
	router.HandleFunc("GET /admin/seriesselector", utils.AuthMiddleware(handlers.GetSeriesSelector(dbpool)))
	router.HandleFunc("GET /admin/portfolio/", utils.AuthMiddleware(handlers.GetPortfolioAdmin(dbpool)))
	router.HandleFunc("POST /admin/upload", utils.AuthMiddleware(handlers.UploadImage()))
	router.HandleFunc("POST /admin/uploadvideo", utils.AuthMiddleware(handlers.UploadVideo()))
	router.HandleFunc("POST /admin/art", utils.AuthMiddleware(handlers.PostArt(dbpool)))
	router.HandleFunc("POST /admin/series", utils.AuthMiddleware(handlers.PostSeries(dbpool)))
	router.HandleFunc("POST /admin/portfolio/{id}", utils.AuthMiddleware(handlers.PostPortfolio(dbpool)))
	router.HandleFunc("PUT /admin", utils.AuthMiddleware(handlers.UpdateAdmin(dbpool)))
	router.HandleFunc("PUT /admin/art", utils.AuthMiddleware(handlers.UpdateArt(dbpool)))
	router.HandleFunc("PUT /admin/series", utils.AuthMiddleware(handlers.UpdateSeries(dbpool)))
	router.HandleFunc("DELETE /admin/art/{id}", utils.AuthMiddleware(handlers.DeleteArt(dbpool)))
	router.HandleFunc("DELETE /admin/series/{id}", utils.AuthMiddleware(handlers.DeleteSeries(dbpool)))
	router.HandleFunc("DELETE /admin/portfolio/{id}", utils.AuthMiddleware(handlers.DeletePortfolio(dbpool)))

	// Login Handlers
	router.HandleFunc("POST /admin/login", handlers.Login(dbpool))

	log.Fatal(http.ListenAndServe("0.0.0.0:8000", router))
}
