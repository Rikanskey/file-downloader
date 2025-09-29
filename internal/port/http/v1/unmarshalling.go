package v1

import (
	"file-downloader/internal/app"
	"github.com/go-chi/render"
	"net/http"
)

func decode(w http.ResponseWriter, r *http.Request, v interface{}) bool {
	if err := render.Decode(r, v); err != nil {
		render.Respond(w, r, err)
		return false
	}
	return true
}

func unmarshalDownloadRequest(w http.ResponseWriter, r *http.Request) (q app.CreateDownloadTaskQuery, ok bool) {
	var createDownloadRequest CreateDownloadRequest
	if ok := decode(w, r, &createDownloadRequest); !ok {
		return app.CreateDownloadTaskQuery{}, ok
	}
	return app.CreateDownloadTaskQuery{
		Links: createDownloadRequest.Links,
	}, true
}
