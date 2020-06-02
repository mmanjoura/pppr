package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mmanjoura/pppr/report"
	"github.com/mmanjoura/pppr/report/serializer/json"

	"github.com/pkg/errors"
)

// ReportHandler ...
type ReportHandler interface {
	Generate(http.ResponseWriter, *http.Request)
}

type handler struct {
	reportService report.Service
}

// NewHandler ...
func NewHandler(reportService report.Service) ReportHandler {
	return &handler{reportService: reportService}
}

func setupResponse(w http.ResponseWriter, contentType string, body []byte, statusCode int) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil {
		log.Println(err)
	}
}

func (h *handler) serializer(contentType string) report.Serializer {

	return &json.Report{}
}

func (h *handler) Generate(w http.ResponseWriter, r *http.Request) {
	rpt := report.Report{}
	contentType := r.Header.Get("Content-Type")

	// acquirerId := chi.URLParam(r, "acquirerid")
	// startDate := chi.URLParam(r, "startdate")
	// endDate := chi.URLParam(r, "enddate")
	reportType := chi.URLParam(r, "reporttype")

	// rpt.AcquirerID = acquirerId
	// rpt.StartDate = startDate
	// rpt.EndDate = endDate
	rpt.ReportType = reportType

	// We need to add a middleweare to generate report
	_, _, err := h.reportService.GetPayments("", "reports")
	if err != nil {
		if errors.Cause(err) == report.ErrTransactionNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseBody, err := h.serializer(contentType).EncodeGenerate(rpt)
	setupResponse(w, contentType, responseBody, http.StatusOK)
}
