package ppdb

import (
	// "bytes"
	// "encoding/json"
	// "io/ioutil"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

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
		resp   response.Response
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

func (h *Handler) InsertFasilitas(w http.ResponseWriter, r *http.Request) {
	var (
		fasilitas ppdbEntity.TableFasilitas
		resp      response.Response
	)

	// Parse multipart form with maximum file size of 10MB
	err := r.ParseMultipartForm(10 << 20) // Maksimum ukuran file 10MB
	if err != nil {
		http.Error(w, "Error memproses form-data", http.StatusBadRequest)
		return
	}

	// Mengambil file dari form-data
	file, _, err := r.FormFile("fasilitas_image")
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
	fasilitas.FasilitasName = r.FormValue("fasilitas_name")
	fasilitas.FasilitasImage = []byte(fileBytes)

	result, err := h.ppdbSvc.InsertFasilitas(r.Context(), fasilitas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp.Data = result
	resp.Message = "Data Fasilitas Sekolah berhasil dimasukkan"

	// Logging informasi request
	ctx := r.Context()
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	// Mengembalikan response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) InsertProfileStaff(w http.ResponseWriter, r *http.Request) {
	var (
		staff ppdbEntity.TableStaff
		resp  response.Response
	)

	// Parse multipart form dengan ukuran maksimum 10MB
	err := r.ParseMultipartForm(10 << 20) // Maksimum ukuran file 10MB
	if err != nil {
		http.Error(w, "Error memproses form-data", http.StatusBadRequest)
		return
	}

	// Mengambil file dari form-data
	file, _, err := r.FormFile("staff_photo")
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
	staff.StaffPhoto = fileBytes
	staff.StaffName = r.FormValue("staff_name")
	staff.StaffGender = r.FormValue("staff_gender")
	staff.StaffPosition = r.FormValue("staff_position")

	// Mengambil tempat lahir
	staffTmptLahir := r.FormValue("staff_tmpt_lahir")
	if staffTmptLahir != "" {
		staff.StaffTmptLahir = &staffTmptLahir // Menggunakan pointer untuk menyimpan tempat lahir
	} else {
		staff.StaffTmptLahir = nil // Atur ke nil jika tidak ada tempat lahir yang diberikan
	}

	// Parse tanggal lahir dengan format RFC1123
	staffTglLahir := r.FormValue("staff_tgl_lahir")
	log.Printf("Tanggal lahir yang diterima: %s", staffTglLahir) // Logging nilai yang diterima
	if staffTglLahir != "" {
		// Parsing sesuai format waktu RFC1123
		parsedDate, err := time.Parse(time.RFC1123, staffTglLahir)
		if err != nil {
			log.Printf("Error parsing date: %s, Error: %v", staffTglLahir, err) // Logging tambahan
			http.Error(w, "Error memproses tanggal lahir", http.StatusBadRequest)
			return
		}

		// Menyimpan hanya tanggal tanpa waktu
		dateOnly := time.Date(parsedDate.Year(), parsedDate.Month(), parsedDate.Day(), 0, 0, 0, 0, parsedDate.Location())
		staff.StaffTglLahir = &dateOnly // Menggunakan pointer untuk menyimpan tanggal
	} else {
		staff.StaffTglLahir = nil // Atur ke nil jika tidak ada tanggal yang diberikan
	}

	// Memanggil service untuk memasukkan data staff
	result, err := h.ppdbSvc.InsertProfileStaff(r.Context(), staff)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp.Data = result
	resp.Message = "Data Staff berhasil dimasukkan"

	// Logging informasi request
	ctx := r.Context()
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	// Mengembalikan response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) InsertEvent(w http.ResponseWriter, r *http.Request) {
	var (
		event ppdbEntity.TableEvent
		resp  response.Response
	)

	// Parse multipart form dengan ukuran maksimum 10MB
	err := r.ParseMultipartForm(10 << 20) // Maksimum ukuran file 10MB
	if err != nil {
		http.Error(w, "Error memproses form-data", http.StatusBadRequest)
		return
	}

	// Mengambil file dari form-data
	file, _, err := r.FormFile("event_image")
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

	// Menyimpan gambar ke dalam struct event
	event.EventImage = fileBytes
	event.EventHeader = r.FormValue("event_header")

	// Mengambil dan memproses tanggal mulai event
	eventStartDate := r.FormValue("event_start_date")
	if eventStartDate != "" {
		parsedStartDate, err := time.Parse("2006-01-02", eventStartDate)
		if err != nil {
			log.Printf("Error parsing start date: %s, Error: %v", eventStartDate, err)
			http.Error(w, "Error memproses tanggal mulai event", http.StatusBadRequest)
			return
		}
		event.EventStartDate = parsedStartDate
	}

	// Mengambil dan memproses tanggal akhir event
	eventEndDate := r.FormValue("event_end_date")
	if eventEndDate != "" {
		parsedEndDate, err := time.Parse("2006-01-02", eventEndDate)
		if err != nil {
			log.Printf("Error parsing end date: %s, Error: %v", eventEndDate, err)
			http.Error(w, "Error memproses tanggal akhir event", http.StatusBadRequest)
			return
		}
		event.EventEndDate = &parsedEndDate
	} else {
		event.EventEndDate = nil
	}

	// Menyimpan deskripsi event
	event.EventDesc = r.FormValue("event_desc")

	// Memanggil service untuk memasukkan data event
	result, err := h.ppdbSvc.InsertEvent(r.Context(), event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp.Data = result
	resp.Message = "Data Event berhasil dimasukkan"

	// Logging informasi request
	ctx := r.Context()
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	// Mengembalikan response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) InsertPesertaDidik(w http.ResponseWriter, r *http.Request) {
	var (
		pesertadidik ppdbEntity.TablePesertaDidik
		resp         response.Response
	)

	// Decode JSON dari body request
	err := json.NewDecoder(r.Body).Decode(&pesertadidik)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Panggil service untuk memasukkan data pesertadidik
	result, err := h.ppdbSvc.InsertPesertaDidik(r.Context(), pesertadidik)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp.Data = result
	resp.Message = "Peserta didik data inserted successfully" // Menyusun pesan respons

	// Mengambil konteks dari request
	ctx := r.Context()
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	// Mengembalikan response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
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
