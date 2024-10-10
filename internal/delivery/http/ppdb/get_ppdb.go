package ppdb

import (
	// "internal/itoa"

	"log"
	"net/http"
	httpHelper "ppdb-be/internal/delivery/http"
	"ppdb-be/pkg/response"
	"strconv"

	// "strconv"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"go.uber.org/zap"
)

// @Router /v1/profiles [get]
func (h *Handler) GetPpdb(w http.ResponseWriter, r *http.Request) {
	var (
		result   interface{}
		metadata interface{}
		err      error
		resp     response.Response
		types    string
	)
	defer resp.RenderJSON(w, r)

	// ptid, _ := strconv.Atoi(r.FormValue("ptid"))
	// page, _ := strconv.Atoi(r.FormValue("page"))
	// limit, _ := strconv.Atoi(r.FormValue("limit"))

	spanCtx, _ := h.tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	span := h.tracer.StartSpan("GetPpdb", ext.RPCServerOption(spanCtx))
	defer span.Finish()

	ctx := r.Context()
	ctx = opentracing.ContextWithSpan(ctx, span)
	h.logger.For(ctx).Info("HTTP request received", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	// Your code here
	types = r.FormValue("type")
	switch types {
	case "getkontaksekolah":

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

func (h *Handler) GetKontakSekolah(w http.ResponseWriter, r *http.Request) {
	// Membuat response default
	resp := response.Response{}
	defer resp.RenderJSON(w, r) // Pastikan response selalu dikembalikan dalam format JSON

	// Memperoleh context dari request
	ctx := r.Context()

	// Memanggil service untuk mendapatkan data kontak sekolah
	kontakSekolah, err := h.ppdbSvc.GetKontakSekolah(ctx)
	if err != nil {
		// Jika terjadi error, gunakan ParseErrorCode untuk memparsing error
		resp = httpHelper.ParseErrorCode(err.Error())
		h.logger.For(ctx).Error("HTTP request error", zap.String("method", r.Method), zap.Stringer("url", r.URL), zap.Error(err))
		return
	}

	// Mengisi field data dan metadata dalam response
	resp.Data = kontakSekolah
	resp.Metadata = nil // Jika Anda memiliki metadata (misal: pagination), bisa diset di sini

	// Logging informasi request yang berhasil
	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))
}

func (h *Handler) GetDataAdminSlim(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	searchInput := r.FormValue("searchInput")
	page, _ := strconv.Atoi(r.FormValue("page"))
	length, _ := strconv.Atoi(r.FormValue("length"))

	ctx := r.Context()

	// Get admin data with pagination
	admins, metadata, err := h.ppdbSvc.GetDataAdminSlim(ctx, searchInput, page, length)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		h.logger.For(ctx).Error("HTTP request error", zap.String("method", r.Method), zap.Stringer("url", r.URL), zap.Error(err))
		return
	}

	// Prepare response data
	resp.Data = admins
	resp.Metadata = metadata

	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))
}

func (h *Handler) GetRole(w http.ResponseWriter, r *http.Request) {
	// Membuat response default
	resp := response.Response{}
	defer resp.RenderJSON(w, r) // Pastikan response selalu dikembalikan dalam format JSON

	// Memperoleh context dari request
	ctx := r.Context()

	// Memanggil service untuk mendapatkan data kontak sekolah
	role, err := h.ppdbSvc.GetRole(ctx)
	if err != nil {
		// Jika terjadi error, gunakan ParseErrorCode untuk memparsing error
		resp = httpHelper.ParseErrorCode(err.Error())
		h.logger.For(ctx).Error("HTTP request error", zap.String("method", r.Method), zap.Stringer("url", r.URL), zap.Error(err))
		return
	}

	// Mengisi field data dan metadata dalam response
	resp.Data = role
	resp.Metadata = nil // Jika Anda memiliki metadata (misal: pagination), bisa diset di sini

	// Logging informasi request yang berhasil
	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))
}
