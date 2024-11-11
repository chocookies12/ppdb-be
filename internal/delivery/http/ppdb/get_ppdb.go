package ppdb

import (
	// "internal/itoa"

	"encoding/json"
	"log"
	"net/http"
	httpHelper "ppdb-be/internal/delivery/http"
	"ppdb-be/pkg/response"
	"strconv"
	"strings"

	// "strconv"
	ppdbEntity "ppdb-be/internal/entity/ppdb"
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

	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	role, err := h.ppdbSvc.GetRole(ctx)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		h.logger.For(ctx).Error("HTTP request error", zap.String("method", r.Method), zap.Stringer("url", r.URL), zap.Error(err))
		return
	}

	resp.Data = role
	resp.Metadata = nil

	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))
}

func (h *Handler) GetAgama(w http.ResponseWriter, r *http.Request) {

	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	agama, err := h.ppdbSvc.GetAgama(ctx)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		h.logger.For(ctx).Error("HTTP request error", zap.String("method", r.Method), zap.Stringer("url", r.URL), zap.Error(err))
		return
	}

	resp.Data = agama
	resp.Metadata = nil

	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))
}

func (h *Handler) GetJurusan(w http.ResponseWriter, r *http.Request) {

	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	jurusan, err := h.ppdbSvc.GetJurusan(ctx)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		h.logger.For(ctx).Error("HTTP request error", zap.String("method", r.Method), zap.Stringer("url", r.URL), zap.Error(err))
		return
	}

	resp.Data = jurusan
	resp.Metadata = nil

	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))
}


func (h *Handler) GetStatus(w http.ResponseWriter, r *http.Request) {

	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	status, err := h.ppdbSvc.GetStatus(ctx)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		h.logger.For(ctx).Error("HTTP request error", zap.String("method", r.Method), zap.Stringer("url", r.URL), zap.Error(err))
		return
	}

	resp.Data = status
	resp.Metadata = nil

	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))
}


func (h *Handler) GetGambarInfoDaftar(w http.ResponseWriter, r *http.Request) {
	infoID := r.URL.Query().Get("infoID")
	if infoID == "" {
		http.Error(w, "infoID is required", http.StatusBadRequest)
		return
	}

	poster, err := h.ppdbSvc.GetGambarInfoDaftar(r.Context(), infoID)
	if err != nil {
		http.Error(w, "Failed to get poster image", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.WriteHeader(http.StatusOK)
	w.Write(poster)
}
func (h *Handler) GetInfoDaftar(w http.ResponseWriter, r *http.Request) {
	// Membuat response default
	resp := response.Response{}
	defer resp.RenderJSON(w, r) // Pastikan response selalu dikembalikan dalam format JSON

	// Memperoleh context dari request
	ctx := r.Context()

	// Memanggil service untuk mendapatkan data Info Daftar
	infoDaftar, err := h.ppdbSvc.GetInfoDaftar(ctx)
	if err != nil {
		// Jika terjadi error, gunakan ParseErrorCode untuk memparsing error
		resp = httpHelper.ParseErrorCode(err.Error())
		h.logger.For(ctx).Error("HTTP request error", zap.String("method", r.Method), zap.Stringer("url", r.URL), zap.Error(err))
		return
	}

	// Mengisi field data dan metadata dalam response
	resp.Data = infoDaftar // Set data dengan info daftar yang didapat
	resp.Metadata = nil    // Jika ada metadata tambahan, bisa diset di sini

	// Logging informasi request yang berhasil
	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))
}

func (h *Handler) GetGambarBanner(w http.ResponseWriter, r *http.Request) {
	bannerID := r.URL.Query().Get("bannerID")
	if bannerID == "" {
		http.Error(w, "bannerID is required", http.StatusBadRequest)
		return
	}

	poster, err := h.ppdbSvc.GetGambarBanner(r.Context(), bannerID)
	if err != nil {
		http.Error(w, "Failed to get poster image", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.WriteHeader(http.StatusOK)
	w.Write(poster)
}

func (h *Handler) GetBanner(w http.ResponseWriter, r *http.Request) {
	// Membuat response default
	resp := response.Response{}
	defer resp.RenderJSON(w, r) // Pastikan response selalu dikembalikan dalam format JSON

	// Memperoleh context dari request
	ctx := r.Context()

	// Memanggil service untuk mendapatkan data banner
	banners, err := h.ppdbSvc.GetBanner(ctx)
	if err != nil {
		// Jika terjadi error, gunakan ParseErrorCode untuk memparsing error
		resp = httpHelper.ParseErrorCode(err.Error())
		h.logger.For(ctx).Error("HTTP request error", zap.String("method", r.Method), zap.Stringer("url", r.URL), zap.Error(err))
		return
	}

	// Mengisi field data dan metadata dalam response
	resp.Data = banners // Set data dengan banner yang didapat
	resp.Metadata = nil // Jika Anda memiliki metadata (misal: pagination), bisa diset di sini

	// Logging informasi request yang berhasil
	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))
}

func (h *Handler) GetGambarFasilitas(w http.ResponseWriter, r *http.Request) {
	fasilitasID := r.URL.Query().Get("fasilitasID")
	if fasilitasID == "" {
		http.Error(w, "fasilitasID is required", http.StatusBadRequest)
		return
	}

	poster, err := h.ppdbSvc.GetGambarFasilitas(r.Context(), fasilitasID)
	if err != nil {
		http.Error(w, "Failed to get poster image", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.WriteHeader(http.StatusOK)
	w.Write(poster)
}

func (h *Handler) GetFasilitasSlim(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	searchInput := r.FormValue("searchInput")
	page, _ := strconv.Atoi(r.FormValue("page"))
	length, _ := strconv.Atoi(r.FormValue("length"))

	ctx := r.Context()

	// Get fasilitas data with pagination
	fasilitas, metadata, err := h.ppdbSvc.GetFasilitasSlim(ctx, searchInput, page, length)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		h.logger.For(ctx).Error("HTTP request error", zap.String("method", r.Method), zap.Stringer("url", r.URL), zap.Error(err))
		return
	}

	// Prepare response data
	resp.Data = fasilitas
	resp.Metadata = metadata

	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))
}

func (h *Handler) GetFasilitas(w http.ResponseWriter, r *http.Request) {
	// Membuat response default
	resp := response.Response{}
	defer resp.RenderJSON(w, r) // Pastikan response selalu dikembalikan dalam format JSON

	// Memperoleh context dari request
	ctx := r.Context()

	// Memanggil service untuk mendapatkan data banner
	fasilitas, err := h.ppdbSvc.GetFasilitas(ctx)
	if err != nil {
		// Jika terjadi error, gunakan ParseErrorCode untuk memparsing error
		resp = httpHelper.ParseErrorCode(err.Error())
		h.logger.For(ctx).Error("HTTP request error", zap.String("method", r.Method), zap.Stringer("url", r.URL), zap.Error(err))
		return
	}

	// Mengisi field data dan metadata dalam response
	resp.Data = fasilitas // Set data dengan banner yang didapat
	resp.Metadata = nil   // Jika Anda memiliki metadata (misal: pagination), bisa diset di sini

	// Logging informasi request yang berhasil
	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))
}

func (h *Handler) GetPhotoStaff(w http.ResponseWriter, r *http.Request) {
	staffID := r.URL.Query().Get("staffID")
	if staffID == "" {
		http.Error(w, "staffID is required", http.StatusBadRequest)
		return
	}

	poster, err := h.ppdbSvc.GetPhotoStaff(r.Context(), staffID)
	if err != nil {
		http.Error(w, "Failed to get poster image", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.WriteHeader(http.StatusOK)
	w.Write(poster)
}

func (h *Handler) GetProfileStaffSlim(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	searchInput := r.FormValue("searchInput")
	page, _ := strconv.Atoi(r.FormValue("page"))
	length, _ := strconv.Atoi(r.FormValue("length"))

	ctx := r.Context()

	// Get staff data with pagination
	staff, metadata, err := h.ppdbSvc.GetProfileStaffSlim(ctx, searchInput, page, length)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		h.logger.For(ctx).Error("HTTP request error", zap.String("method", r.Method), zap.Stringer("url", r.URL), zap.Error(err))
		return
	}

	// Prepare response data
	resp.Data = staff
	resp.Metadata = metadata

	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))
}

func (h *Handler) GetProfilStaffUtama(w http.ResponseWriter, r *http.Request) {
	// Membuat response default
	resp := response.Response{}
	defer resp.RenderJSON(w, r) // Pastikan response selalu dikembalikan dalam format JSON

	// Memperoleh context dari request
	ctx := r.Context()

	// Memanggil service untuk mendapatkan data banner
	staff, err := h.ppdbSvc.GetProfileStaffUtama(ctx)
	if err != nil {
		// Jika terjadi error, gunakan ParseErrorCode untuk memparsing error
		resp = httpHelper.ParseErrorCode(err.Error())
		h.logger.For(ctx).Error("HTTP request error", zap.String("method", r.Method), zap.Stringer("url", r.URL), zap.Error(err))
		return
	}

	resp.Data = staff
	resp.Metadata = nil

	// Logging informasi request yang berhasil
	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))
}

func (h *Handler) GetImageEvent(w http.ResponseWriter, r *http.Request) {
	eventID := r.URL.Query().Get("eventID")
	if eventID == "" {
		http.Error(w, "eventID is required", http.StatusBadRequest)
		return
	}

	poster, err := h.ppdbSvc.GetImageEvent(r.Context(), eventID)
	if err != nil {
		http.Error(w, "Failed to get poster image", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.WriteHeader(http.StatusOK)
	w.Write(poster)
}

func (h *Handler) GetEventSlim(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	searchInput := r.FormValue("searchInput")
	page, _ := strconv.Atoi(r.FormValue("page"))
	length, _ := strconv.Atoi(r.FormValue("length"))

	ctx := r.Context()

	// Get staff data with pagination
	event, metadata, err := h.ppdbSvc.GetEventSlim(ctx, searchInput, page, length)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		h.logger.For(ctx).Error("HTTP request error", zap.String("method", r.Method), zap.Stringer("url", r.URL), zap.Error(err))
		return
	}

	// Prepare response data
	resp.Data = event
	resp.Metadata = metadata

	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))
}

func (h *Handler) GetEventDetail(w http.ResponseWriter, r *http.Request) {

	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	eventID := r.URL.Query().Get("eventID")
	if eventID == "" {
		resp = httpHelper.ParseErrorCode("eventID is required")
		h.logger.For(ctx).Error("HTTP request error - missing eventID", zap.String("method", r.Method), zap.Stringer("url", r.URL))
		return
	}

	event, err := h.ppdbSvc.GetEventDetail(ctx, eventID)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		h.logger.For(ctx).Error("HTTP request error", zap.String("method", r.Method), zap.Stringer("url", r.URL), zap.Error(err))
		return
	}

	resp.Data = event
	resp.Metadata = nil

	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))
}

func (h *Handler) GetEventUtama(w http.ResponseWriter, r *http.Request) {
	// Membuat response default
	resp := response.Response{}
	defer resp.RenderJSON(w, r) // Pastikan response selalu dikembalikan dalam format JSON

	// Memperoleh context dari request
	ctx := r.Context()

	// Memanggil service untuk mendapatkan data event
	events, err := h.ppdbSvc.GetEventUtama(ctx)
	if err != nil {
		// Jika terjadi error, gunakan ParseErrorCode untuk memparsing error
		resp = httpHelper.ParseErrorCode(err.Error())
		h.logger.For(ctx).Error("HTTP request error", zap.String("method", r.Method), zap.Stringer("url", r.URL), zap.Error(err))
		return
	}

	resp.Data = events
	resp.Metadata = nil

	// Logging informasi request yang berhasil
	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))
}

func (h *Handler) GetLoginCheck(w http.ResponseWriter, r *http.Request) {
	var (
		login ppdbEntity.TablePesertaDidik
		resp  response.Response
	)

	// Decode JSON dari body request
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Panggil service untuk memasukkan data pesertadidik
	result, err := h.ppdbSvc.GetLoginCheck(r.Context(), login)
	if err != nil {
		if strings.Contains(err.Error(), "bcrypt") {
			http.Error(w, "Invalid email or password", http.StatusInternalServerError)
		} else if strings.Contains(err.Error(), "no rows") {
			http.Error(w, "Email not found", http.StatusInternalServerError)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	resp.Data = result
	resp.Message = "Login berhasil"

	// Mengambil konteks dari request
	ctx := r.Context()
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	// Mengembalikan response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) GetPembayaranFormulirDetail(w http.ResponseWriter, r *http.Request) {
	var (
		resp  response.Response
	)

	// Panggil service untuk memasukkan data pesertadidik
	result, err := h.ppdbSvc.GetPembayaranFormulirDetail(r.Context(), r.FormValue("idpesertadidik"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp.Data = result
	resp.Message = "Berhasil mendapatkan pembayaran formulir"

	// Mengambil konteks dari request
	ctx := r.Context()
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	// Mengembalikan response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) GetFormulirDetail(w http.ResponseWriter, r *http.Request) {
	var (
		resp  response.Response
	)

	// Panggil service untuk memasukkan data pesertadidik
	result, err := h.ppdbSvc.GetFormulirDetail(r.Context(), r.FormValue("idpesertadidik"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp.Data = result
	resp.Message = "Berhasil mendapatkan formulir"

	// Mengambil konteks dari request
	ctx := r.Context()
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	// Mengembalikan response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}


func (h *Handler) GetBerkasDetail(w http.ResponseWriter, r *http.Request) {
	var (
		resp  response.Response
	)

	// Panggil service untuk memasukkan data pesertadidik
	result, err := h.ppdbSvc.GetBerkasDetail(r.Context(), r.FormValue("idpesertadidik"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp.Data = result
	resp.Message = "Berhasil mendapatkan berkas"

	// Mengambil konteks dari request
	ctx := r.Context()
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	// Mengembalikan response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}


func (h *Handler) GetJadwalTestDetail(w http.ResponseWriter, r *http.Request) {
	var (
		resp  response.Response
	)

	// Panggil service untuk memasukkan data pesertadidik
	result, err := h.ppdbSvc.GetJadwalTestDetail(r.Context(), r.FormValue("idpesertadidik"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp.Data = result
	resp.Message = "Berhasil mendapatkan jadwal test"

	// Mengambil konteks dari request
	ctx := r.Context()
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	// Mengembalikan response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) GetGeneratedKartuTest(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	ctx := r.Context()

	result, err := h.ppdbSvc.GetGeneratedKartuTest(ctx, r.FormValue("idpesertadidik"))
	if err != nil {
		defer resp.RenderJSON(w, r)
		
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "inline; filename=Kartu_Test.pdf")
	w.Write(result)

	resp.Data = result
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
}

func (h *Handler) GetGeneratedFormulir(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	ctx := r.Context()

	result, err := h.ppdbSvc.GetGeneratedFormulir(ctx)
	if err != nil {
		defer resp.RenderJSON(w, r)
		
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "inline; filename=Formulir.pdf")
	w.Write(result)

	resp.Data = result
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
}