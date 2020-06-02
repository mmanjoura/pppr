package api

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mmanjoura/pppr/payment"
	"github.com/mmanjoura/pppr/payment/serializer/json"

	"github.com/pkg/errors"
)

// TransactionHandler ...
type PaymentHandler interface {
	Post(http.ResponseWriter, *http.Request)
}

type handler struct {
	paymentService payment.Service
}

// NewHandler ...
func NewHandler(paymentService payment.Service) PaymentHandler {
	return &handler{paymentService: paymentService}
}

func setupResponse(w http.ResponseWriter, contentType string, body []byte, statusCode int) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil {
		log.Println(err)
	}
}

func (h *handler) serializer(contentType string) payment.Serializer {

	return &json.Payment{}
}

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	collection := chi.URLParam(r, "collection")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	runParams, err := h.serializer(contentType).Decode(requestBody)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	//Added temporary to fix the build "payment.PAcquierer{}"

	var pAcquier = payment.PAcquirer{}
	err = h.paymentService.Run(runParams, collection, &pAcquier)
	if err != nil {
		if errors.Cause(err) == payment.ErrTransactionInvalid {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseBody, err := h.serializer(contentType).Encode(runParams)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	setupResponse(w, contentType, responseBody, http.StatusCreated)
}
