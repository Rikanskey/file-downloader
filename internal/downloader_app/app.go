package downloader_app

type Application struct {
	Commands Commands
	Queries  Queries
}

type (
	Commands struct {
		CreateDownloadFileCommand createDownloadFile
	}
	createDownloadFile interface {
		Handle(cmd CreateDownloadFileCommand) error
	}
)

type (
	Queries struct{}
)
