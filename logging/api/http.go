package api

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mmanjoura/pppr/logging"
	"github.com/mmanjoura/pppr/logging/serializer/json"
	"github.com/pkg/errors"
)

// LoggingHandler ...
type LoggingHandler interface {
	Post(http.ResponseWriter, *http.Request)
}

type handler struct {
	loggingService logging.LoggingService
}

// NewHandler ...
func NewHandler(loggingService logging.LoggingService) LoggingHandler {
	return &handler{loggingService: loggingService}
}

func setupResponse(w http.ResponseWriter, contentType string, body []byte, statusCode int) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil {
		log.Println(err)
	}
}

func (h *handler) serializer(contentType string) logging.LoggingSerializer {

	return &json.Message{}
}

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	collection := chi.URLParam(r, "collection")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	trx, err := h.serializer(contentType).Decode(requestBody)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = h.loggingService.Save(trx, collection)
	if err != nil {
		if errors.Cause(err) == logging.ErrLoggingInvalid {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseBody, err := h.serializer(contentType).Encode(trx)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	setupResponse(w, contentType, responseBody, http.StatusCreated)
}
