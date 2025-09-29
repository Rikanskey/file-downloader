package download_commands

import (
	"file-downloader/internal/downloader_app"
	"file-downloader/internal/repository"
)

type Downloader interface {
	Download(url string) error
}

type CreateDownloadFileHandler struct {
	downloader Downloader
}

func NewCreateDownloadFileHandler(downloader *repository.UrlDownloader) CreateDownloadFileHandler {
	return CreateDownloadFileHandler{
		downloader: downloader,
	}
}

func (h CreateDownloadFileHandler) Handle(cmd downloader_app.CreateDownloadFileCommand) error {
	err := h.downloader.Download(cmd.Url)
	return err
}
