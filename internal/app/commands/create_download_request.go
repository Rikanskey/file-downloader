package commands

import (
	"context"
	"file-downloader/internal/app"
	"file-downloader/internal/repository"
	"github.com/google/uuid"
)

type TaskWriter interface {
	WriteTask(request app.Task)
}

type CreateDownloadTaskHandler struct {
	taskWriter TaskWriter
}

func NewCreateDownloadTaskHandler(requestLogger *repository.TaskServerRep) CreateDownloadTaskHandler {

	return CreateDownloadTaskHandler{
		taskWriter: requestLogger,
	}
}

func (h CreateDownloadTaskHandler) Handle(ctx context.Context, query app.CreateDownloadTaskQuery) (string, error) {
	id := uuid.New().String()
	h.taskWriter.WriteTask(app.Task{Id: id, Links: query.Links})
	return id, nil
}
