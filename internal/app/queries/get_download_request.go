package queries

import (
	"context"
	"file-downloader/internal/app"
)

type GetDownloadRequestModel interface {
	ReadTask(ctx context.Context, requestId string) ([]app.DownloadStatus, error)
}

type GetDownloadRequestHandler struct {
	getModel GetDownloadRequestModel
}

func NewGetDownloadRequestHandler(getModel GetDownloadRequestModel) *GetDownloadRequestHandler {
	return &GetDownloadRequestHandler{getModel: getModel}
}

func (h GetDownloadRequestHandler) Handle(ctx context.Context, query app.GetDownloadRequest) ([]app.DownloadStatus, error) {
	ds, err := h.getModel.ReadTask(ctx, query.DownloadRequestId)

	return ds, err
}
