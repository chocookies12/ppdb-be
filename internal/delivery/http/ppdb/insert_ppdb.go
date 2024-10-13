package ppdb

import (
	// "bytes"
	// "encoding/json"
	// "io/ioutil"
	"encoding/json"
	"io"
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

func (h *Handler) InsertInfoDaftar(w http.ResponseWriter, r *http.Request) {
	var (
		infoDaftar ppdbEntity.TableInfoDaftar
		resp       response.Response
	)

	// Parse multipart form with maximum file size of 10MB
	err := r.ParseMultipartForm(10 << 20) // Maksimum ukuran file 10MB
	if err != nil {
		http.Error(w, "Error memproses form-data", http.StatusBadRequest)
		return
	}

	// Mengambil file dari form-data
	file, _, err := r.FormFile("poster_daftar")
	if err != nil {
		http.Error(w, "Error mengambil file dari form-data", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Membaca isi file ke dalam byte array
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error membaca isi file", http.StatusInternalServerError)
		return
	}

	// Membaca data JSON lainnya dari form-data
	infoDaftar.PosterDaftar = []byte(fileBytes)
	// infoDaftar.LinkPosterDaftar = r.FormValue("link_poster_daftar")
	infoDaftar.AwalTahunAjar = r.FormValue("awal_tahun_ajar")
	infoDaftar.AkhirTahunAjar = r.FormValue("akhir_tahun_ajar")

	// Memanggil service untuk memasukkan data infoDaftar
	result, err := h.ppdbSvc.InsertInfoDaftar(r.Context(), infoDaftar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp.Data = result
	resp.Message = "Data Info Daftar berhasil dimasukkan"

	// Logging informasi request
	ctx := r.Context()
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	// Mengembalikan response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) InsertBanner(w http.ResponseWriter, r *http.Request) {
	var (
		banner ppdbEntity.TableBanner
		resp       response.Response
	)

	// Parse multipart form with maximum file size of 10MB
	err := r.ParseMultipartForm(10 << 20) // Maksimum ukuran file 10MB
	if err != nil {
		http.Error(w, "Error memproses form-data", http.StatusBadRequest)
		return
	}

	// Mengambil file dari form-data
	file, _, err := r.FormFile("banner_image")
	if err != nil {
		http.Error(w, "Error mengambil file dari form-data", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Membaca isi file ke dalam byte array
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error membaca isi file", http.StatusInternalServerError)
		return
	}

	// Membaca data JSON lainnya dari form-data
	banner.BannerName = r.FormValue("banner_name")
	banner.BannerImage = []byte(fileBytes)


	result, err := h.ppdbSvc.InsertBanner(r.Context(), banner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp.Data = result
	resp.Message = "Data Info Banner berhasil dimasukkan"

	// Logging informasi request
	ctx := r.Context()
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	// Mengembalikan response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
