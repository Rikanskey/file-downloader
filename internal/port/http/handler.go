package http

import (
	"file-downloader/internal/app"
	v1 "file-downloader/internal/port/http/v1"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
)

func NewHandler(app app.Application) http.Handler {
	apiRouter := chi.NewRouter()
	addMiddlewares(apiRouter)

	rootRouter := chi.NewRouter()
	rootRouter.Mount("/v1", v1.NewHandler(app, apiRouter))

	return rootRouter
}

func addMiddlewares(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	addCorsMiddleware(router)
}

func addCorsMiddleware(router *chi.Mux) {
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Accept", "Content-Type"},
		MaxAge:         300,
	})

	router.Use(corsMiddleware.Handler)
}
