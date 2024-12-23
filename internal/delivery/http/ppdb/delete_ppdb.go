package ppdb

import (
	"log"
	"net/http"
	httpHelper "ppdb-be/internal/delivery/http"
	"ppdb-be/pkg/response"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"go.uber.org/zap"
)

func (h *Handler) DeletePpdb(w http.ResponseWriter, r *http.Request) {
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

	types = r.FormValue("type")
	switch types {

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

func (h *Handler) DeleteAdmin(w http.ResponseWriter, r *http.Request) {
	var (
		resp    response.Response
		adminID = r.URL.Query().Get("adminID") // Get adminID from query
	)
	defer resp.RenderJSON(w, r)

	ctx := r.Context()
	result, err := h.ppdbSvc.DeleteAdmin(ctx, adminID) // Call service to delete admin
	if err != nil {
		log.Printf("[ERROR] %v\n", err)
		resp = httpHelper.ParseErrorCode(err.Error())
		return
	}

	resp.Data = result
}

func (h *Handler) DeleteBanner(w http.ResponseWriter, r *http.Request) {
	var (
		resp    response.Response
		bannerID = r.URL.Query().Get("bannerID")
	)
	defer resp.RenderJSON(w, r)

	ctx := r.Context()
	result, err := h.ppdbSvc.DeleteBanner(ctx, bannerID) 
	if err != nil {
		log.Printf("[ERROR] %v\n", err)
		resp = httpHelper.ParseErrorCode(err.Error())
		return
	}

	resp.Data = result
}

func (h *Handler) DeleteFasilitas(w http.ResponseWriter, r *http.Request) {
	var (
		resp    response.Response
		fasilitasID = r.URL.Query().Get("fasilitasID")
	)
	defer resp.RenderJSON(w, r)

	ctx := r.Context()
	result, err := h.ppdbSvc.DeleteFasilitas(ctx, fasilitasID) 
	if err != nil {
		log.Printf("[ERROR] %v\n", err)
		resp = httpHelper.ParseErrorCode(err.Error())
		return
	}

	resp.Data = result
}

func (h *Handler) DeleteProfileStaff(w http.ResponseWriter, r *http.Request) {
	var (
		resp    response.Response
		staffID = r.URL.Query().Get("staffID")
	)
	defer resp.RenderJSON(w, r)

	ctx := r.Context()
	result, err := h.ppdbSvc.DeleteProfileStaff(ctx, staffID) 
	if err != nil {
		log.Printf("[ERROR] %v\n", err)
		resp = httpHelper.ParseErrorCode(err.Error())
		return
	}

	resp.Data = result
}

func (h *Handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	var (
		resp    response.Response
		eventID = r.URL.Query().Get("eventID")
	)
	defer resp.RenderJSON(w, r)

	ctx := r.Context()
	result, err := h.ppdbSvc.DeleteEvent(ctx, eventID) 
	if err != nil {
		log.Printf("[ERROR] %v\n", err)
		resp = httpHelper.ParseErrorCode(err.Error())
		return
	}

	resp.Data = result
}
