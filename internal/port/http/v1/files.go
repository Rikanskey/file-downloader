package v1

import (
	"file-downloader/internal/app"
	"github.com/go-chi/render"
	"net/http"
)

func (h handler) CreateRequest(w http.ResponseWriter, r *http.Request) {
	q, ok := unmarshalDownloadRequest(w, r)
	if !ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
		render.Respond(w, r, nil)
		return
	}

	id, err := h.app.Commands.CreateDownloadRequest.Handle(r.Context(), q)
	if err == nil {
		marshalRequestId(w, r, id)
	}
}

func (h handler) GetFiles(w http.ResponseWriter, r *http.Request, requestId string) {
	ds, err := h.app.Queries.GetDownloadRequest.Handle(r.Context(), app.GetDownloadRequest{DownloadRequestId: requestId})
	if len(ds) == 0 {
		w.WriteHeader(http.StatusNotFound)
		render.Respond(w, r, nil)
	} else if err == nil {
		marshallDownloadStatuses(w, r, ds)
	}
}
