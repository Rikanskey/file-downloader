package downloader

import (
	"file-downloader/internal/config"
	"file-downloader/internal/downloader_app"
	"file-downloader/internal/downloader_app/download_commands"
	"file-downloader/internal/repository"
	"log"
)

func StartDownloader(configDir string) {
	cfg := newConfig(configDir)
	app := newApplication(cfg)
	da := downloader_app.NewDownloadApplication(cfg.TaskWriter.Path, 15, repository.NewTaskServerRep(cfg.TaskWriter),
		repository.NewTaskReader(cfg.DownloadTasks), app)

	err := da.Start()
	if err != nil {
		log.Panicln(err)
	}
}

func newConfig(configDir string) *config.DownloaderConfig {
	cfg, err := config.NewDownloaderConfig(configDir)

	if err != nil {
		log.Panicln(err.Error())
	}

	return cfg
}

func newApplication(cfg *config.DownloaderConfig) downloader_app.Application {
	return downloader_app.Application{Commands: downloader_app.Commands{
		CreateDownloadFileCommand: download_commands.NewCreateDownloadFileHandler(repository.NewUrlDownloader(cfg.FilesDirectory.Path))},
	}
}
