package server

import (
	"file-downloader/internal/app"
	"file-downloader/internal/app/commands"
	"file-downloader/internal/app/queries"
	"file-downloader/internal/config"
	"file-downloader/internal/port/http"
	"file-downloader/internal/repository"
	"file-downloader/internal/server"
	"log"
)

func Start(configDir string) {
	cfg := newConfig(configDir)
	application := newApplication(cfg)
	startServer(&cfg.Http, application)
}

func newConfig(configDir string) *config.ServerConfig {
	cfg, err := config.NewServerConfig(configDir)
	if err != nil {
		log.Panicln(err)
	}

	return cfg
}

func newApplication(cfg *config.ServerConfig) app.Application {
	return app.Application{
		Commands: app.Commands{
			CreateDownloadRequest: commands.NewCreateDownloadTaskHandler(repository.NewTaskServerRep(cfg.RequestLogger)),
		},
		Queries: app.Queries{
			GetDownloadRequest: queries.NewGetDownloadRequestHandler(repository.NewTaskReader(cfg.DownloadTasksConfig)),
		},
	}
}

func startServer(cfg *config.HttpConfig, app app.Application) {
	log.Printf("Starting HTTP server on address: %s\n", cfg.Port)
	httpServer := server.New(cfg, http.NewHandler(app))

	err := httpServer.Run()

	log.Println(err)
}
