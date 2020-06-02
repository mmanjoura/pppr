package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/mmanjoura/pppr/acquirer"
	"github.com/mmanjoura/pppr/acquirer/serializer/json"

	"github.com/pkg/errors"
)

// AcquirerHandler ...
type AcquirerHandler interface {
	GetAcquirerPayments(http.ResponseWriter, *http.Request)
	GetMerchantPayments(http.ResponseWriter, *http.Request)
	GetReports(http.ResponseWriter, *http.Request)
	GetTransactions(http.ResponseWriter, *http.Request)
	GetExchangeRates(http.ResponseWriter, *http.Request)
	GetLogMessages(http.ResponseWriter, *http.Request)
	GetProcessStates(http.ResponseWriter, *http.Request)
	PutProcessState(http.ResponseWriter, *http.Request)
}

type handler struct {
	acquirerService acquirer.Service
}

// NewHandler ...
func NewHandler(acquirerService acquirer.Service) AcquirerHandler {
	return &handler{acquirerService: acquirerService}
}

func setupResponse(w http.ResponseWriter, contentType string, body []byte, statusCode int) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil {
		log.Println(err)
	}
}

func (h *handler) serializer(contentType string) acquirer.Serializer {

	return &json.Acquirer{}
}

func (h *handler) GetAcquirerPayments(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	ai := chi.URLParam(r, "acquirerid")
	aPayments, err := h.acquirerService.GetAcquirerPayments(ai)
	if err != nil {
		if errors.Cause(err) == acquirer.ErrTransactionNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseBody, err := h.serializer(contentType).EncodeGetAcquirerPayments(aPayments)
	setupResponse(w, contentType, responseBody, http.StatusOK)

}

func (h *handler) GetMerchantPayments(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	mid := chi.URLParam(r, "mid")
	mPayments, err := h.acquirerService.GetMerchantPayments(mid)
	if err != nil {
		if errors.Cause(err) == acquirer.ErrTransactionNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseBody, err := h.serializer(contentType).EncodeGetMerchantPayments(mPayments)
	setupResponse(w, contentType, responseBody, http.StatusOK)

}

func (h *handler) GetReports(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	ai := chi.URLParam(r, "acquirerid")
	reports, err := h.acquirerService.GetReports(ai)
	if err != nil {
		if errors.Cause(err) == acquirer.ErrReportNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseBody, err := h.serializer(contentType).EncodeGetReports(reports)
	setupResponse(w, contentType, responseBody, http.StatusOK)

}

func (h *handler) GetTransactions(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	ai := chi.URLParam(r, "acquirerid")
	transactions, err := h.acquirerService.GetTransactions(ai)
	if err != nil {
		if errors.Cause(err) == acquirer.ErrTransactionNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseBody, err := h.serializer(contentType).EncodeGetTransactions(transactions)
	setupResponse(w, contentType, responseBody, http.StatusOK)

}

func (h *handler) GetExchangeRates(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	date := chi.URLParam(r, "date")
	exchangeRates, err := h.acquirerService.GetExchangeRates(date)
	if err != nil {
		if errors.Cause(err) == acquirer.ErrExchangeRateNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseBody, err := h.serializer(contentType).EncodeGetExchangeRates(exchangeRates)
	setupResponse(w, contentType, responseBody, http.StatusOK)

}

func (h *handler) GetLogMessages(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	date := chi.URLParam(r, "date")
	logMessages, err := h.acquirerService.GetLogMessages(date)
	if err != nil {
		if errors.Cause(err) == acquirer.ErrLogMessageNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseBody, err := h.serializer(contentType).EncodeGetLogMessages(logMessages)
	setupResponse(w, contentType, responseBody, http.StatusOK)

}

func (h *handler) GetProcessStates(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	approved, err := strconv.ParseBool(chi.URLParam(r, "approved"))
	if err != nil {
		if errors.Cause(err) == acquirer.ErrProcessStateNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	processStates, err := h.acquirerService.GetProcessStates(approved)
	if err != nil {
		if errors.Cause(err) == acquirer.ErrProcessStateNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseBody, err := h.serializer(contentType).EncodeGetProcessStates(processStates)
	setupResponse(w, contentType, responseBody, http.StatusOK)

}

func (h *handler) PutProcessState(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	processid := chi.URLParam(r, "processid")

	err := h.acquirerService.PutProcessState(processid, acquirer.ProcessState{})
	if err != nil {
		if errors.Cause(err) == acquirer.ErrProcessStateNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	// TO DO, to be looked at !!!
	_ = h.serializer(contentType).EncodePutProcessState(acquirer.ProcessState{})
	setupResponse(w, contentType, nil, http.StatusOK)

}
