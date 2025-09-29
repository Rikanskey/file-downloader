package v1

import (
	"file-downloader/internal/app"
	"github.com/go-chi/render"
	"net/http"
)

func marshalRequestId(w http.ResponseWriter, r *http.Request, id string) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Location", id)
	render.Respond(w, r, nil)
}

func marshallStatus(status app.TaskStatus) DownloadLinkStatusResponseStatus {
	switch status {
	case app.Wait:
		return DownloadLinkStatusResponseStatusWait
	case app.Ready:
		return DownloadLinkStatusResponseStatusReady
	default:
		return DownloadLinkStatusResponseStatusError
	}
}

func marshallDownloadStatuses(w http.ResponseWriter, r *http.Request, ds []app.DownloadStatus) {
	drr := DownloadRequestResponse{Responses: make([]DownloadLinkStatusResponse, 0)}
	for _, d := range ds {
		drr.Responses = append(drr.Responses, marshallDownloadStatus(d))
	}

	render.Respond(w, r, drr)
}

func marshallDownloadStatus(ds app.DownloadStatus) DownloadLinkStatusResponse {
	return DownloadLinkStatusResponse{Status: marshallStatus(ds.TaskStatus), Link: ds.Link}
}
