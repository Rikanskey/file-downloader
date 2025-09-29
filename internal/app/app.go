package app

import "context"

type Application struct {
	Commands Commands
	Queries  Queries
}

type (
	Commands struct {
		CreateDownloadRequest createDownloadRequest
	}
	createDownloadRequest interface {
		Handle(ctx context.Context, q CreateDownloadTaskQuery) (string, error)
	}
)

type (
	Queries struct {
		GetDownloadRequest getDownloadRequest
	}
	getDownloadRequest interface {
		Handle(ctx context.Context, q GetDownloadRequest) ([]DownloadStatus, error)
	}
)
