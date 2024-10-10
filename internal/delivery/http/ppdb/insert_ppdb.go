package ppdb

import (
	// "bytes"
	// "encoding/json"
	// "io/ioutil"
	"encoding/json"
	"log"
	"net/http"

	// JOEntity "ppdb-be/internal/entity/ppdb"
	ppdbEntity "ppdb-be/internal/entity/ppdb"
	"ppdb-be/pkg/response"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"go.uber.org/zap"
)

func (h *Handler) InsertPpdb(w http.ResponseWriter, r *http.Request) {
	var (
		result   interface{}
		metadata interface{}
		err      error

		resp  response.Response
		types string

		// InsertPack	JOEntity.InsertUnit
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
	//ini ketika submit masuk ke halaman admin
	case "loginadmin":
		result, err = h.ppdbSvc.LoginAdmin(ctx, r.FormValue("emailAdmin"), r.FormValue("password"))
	}

	if err != nil {
		resp.SetError(err, http.StatusInternalServerError)
		resp.StatusCode = 500
		log.Printf("[ERROR] %s %s - %s\n", r.Method, r.URL, err.Error())
		return
	}

	resp.Data = result
	resp.Metadata = metadata
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))

}

func (h *Handler) InsertDataAdmin(w http.ResponseWriter, r *http.Request) {
	var (
		admin ppdbEntity.TableAdmin
		resp  response.Response
	)

	// Decode JSON dari body request
	err := json.NewDecoder(r.Body).Decode(&admin)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Panggil service untuk memasukkan data admin
	result, err := h.ppdbSvc.InsertDataAdmin(r.Context(), admin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp.Data = result
	resp.Message = "Admin data inserted successfully" // Menyusun pesan respons

	// Mengambil konteks dari request
	ctx := r.Context()
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	// Mengembalikan response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

