package http

import (
	"net/http"

	"github.com/kenkoii/WordCosmos/datastore"
)

type Handler struct {
	Service datastore.AppReviewInfoService
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// handle request
}
