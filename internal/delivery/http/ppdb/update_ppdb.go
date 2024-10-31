package ppdb

import (
	// "encoding/json"
	// "io/ioutil"
	"context"
	"encoding/json"
	httpHelper "ppdb-be/internal/delivery/http"
	ppdbEntity "ppdb-be/internal/entity/ppdb"

	// JOEntity "ppdb-be/internal/entity/ppdb"
	"ppdb-be/pkg/response"
	// "encoding/json"
	// "io/ioutil"
	"log"
	"net/http"

	// "strconv"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"go.uber.org/zap"
)

func (h *Handler) UpdatePpdb(w http.ResponseWriter, r *http.Request) {
	var (
		result   interface{}
		metadata interface{}
		err      error
		resp     response.Response
		types    string
	)
	defer resp.RenderJSON(w, r)

	spanCtx, _ := h.tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	span := h.tracer.StartSpan("Getppdb", ext.RPCServerOption(spanCtx))
	defer span.Finish()

	ctx := r.Context()
	ctx = opentracing.ContextWithSpan(ctx, span)
	h.logger.For(ctx).Info("HTTP request received", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	// Your code here
	types = r.FormValue("type")
	switch types {
	// case "":

	}

	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		h.logger.For(ctx).Error("HTTP request error", zap.String("method", r.Method), zap.Stringer("url", r.URL), zap.Error(err))
		return
	}

	resp.Data = result
	resp.Metadata = metadata
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))

}

func (h *Handler) UpdateBanner(w http.ResponseWriter, r *http.Request) {
	// Parsing bannerID dari URL atau request parameter
	bannerID := r.URL.Query().Get("bannerID")
	if bannerID == "" {
		http.Error(w, "Banner ID is required", http.StatusBadRequest)
		return
	}

	// Decode JSON dari request body ke struct TableBanner
	var banner ppdbEntity.TableBanner
	if err := json.NewDecoder(r.Body).Decode(&banner); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Panggil service UpdateBanner
	result, err := h.ppdbSvc.UpdateBanner(context.Background(), banner, bannerID)
	if err != nil {
		http.Error(w, "Failed to update banner", http.StatusInternalServerError)
		return
	}

	// Response berhasil
	response := map[string]string{
		"message": result,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
