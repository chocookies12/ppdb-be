package ppdb

import (
	"encoding/json"
	"fmt"
	"io"
	httpHelper "ppdb-be/internal/delivery/http"
	"strings"
	"time"

	ppdbEntity "ppdb-be/internal/entity/ppdb"
	"ppdb-be/pkg/response"

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
	var (
		err  error
		resp response.Response
	)

	// Parse multipart form with a maximum file size of 10MB
	if err := r.ParseMultipartForm(10 << 20); err != nil { // Maksimum ukuran file 10MB
		http.Error(w, "Error memproses form-data", http.StatusBadRequest)
		return
	}

	// Retrieve the file from form-data
	file, _, err := r.FormFile("banner_image")
	if err != nil {
		http.Error(w, "Error mengambil file dari form-data", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Read the file contents into a byte array
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error membaca isi file", http.StatusInternalServerError)
		return
	}

	// Extract other JSON data from form-data
	bannerID := r.FormValue("bannerID")
	TableBanner := ppdbEntity.TableBanner{
		BannerID:    bannerID,
		BannerName:  r.FormValue("banner_name"),
		BannerImage: fileBytes,
	}

	// Update data in the database via the UpdateDestinasi service
	result, err := h.ppdbSvc.UpdateBanner(r.Context(), TableBanner, bannerID)
	if err != nil {
		http.Error(w, "Gagal memperbarui data banner", http.StatusInternalServerError)
		log.Printf("[ERROR] %s %s - %s\n", r.Method, r.URL, err.Error())
		return
	}

	resp.Data = result

	ctx := r.Context()
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	// Response with success
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
	w.Write([]byte("Banner berhasil diperbarui"))
}

func (h *Handler) UpdateInfoDaftar(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		resp response.Response
	)

	// Parse multipart form with a maximum file size of 10MB
	if err := r.ParseMultipartForm(10 << 20); err != nil { // Maksimum ukuran file 10MB
		http.Error(w, "Error memproses form-data", http.StatusBadRequest)
		return
	}

	// Retrieve the file from form-data
	file, _, err := r.FormFile("poster_daftar")
	if err != nil {
		http.Error(w, "Error mengambil file dari form-data", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Read the file contents into a byte array
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error membaca isi file", http.StatusInternalServerError)
		return
	}

	// Extract other JSON data from form-data
	infoID := r.FormValue("infoID")
	TableInfoDaftar := ppdbEntity.TableInfoDaftar{
		InfoID:         infoID,
		PosterDaftar:   fileBytes,
		AwalTahunAjar:  r.FormValue("awal_tahun_ajar"),
		AkhirTahunAjar: r.FormValue("akhir_tahun_ajar"),
	}

	// Update data in the database via the UpdateDestinasi service
	result, err := h.ppdbSvc.UpdateInfoDaftar(r.Context(), TableInfoDaftar, infoID)
	if err != nil {
		http.Error(w, "Gagal memperbarui data informasi daftar", http.StatusInternalServerError)
		log.Printf("[ERROR] %s %s - %s\n", r.Method, r.URL, err.Error())
		return
	}

	resp.Data = result

	ctx := r.Context()
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	// Response with success
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
	w.Write([]byte("Informasi Pendaftaran berhasil diperbarui"))
}

func (h *Handler) UpdateFasilitas(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		resp response.Response
	)

	// Parse multipart form with a maximum file size of 10MB
	if err := r.ParseMultipartForm(10 << 20); err != nil { // Maksimum ukuran file 10MB
		http.Error(w, "Error memproses form-data", http.StatusBadRequest)
		return
	}

	// Retrieve the file from form-data
	file, _, err := r.FormFile("fasilitas_image")
	if err != nil {
		http.Error(w, "Error mengambil file dari form-data", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Read the file contents into a byte array
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error membaca isi file", http.StatusInternalServerError)
		return
	}

	// Extract other JSON data from form-data
	fasilitasID := r.FormValue("fasilitasID")
	TableFasilitas := ppdbEntity.TableFasilitas{
		FasilitasID:    fasilitasID,
		FasilitasName:  r.FormValue("fasilitas_name"),
		FasilitasImage: fileBytes,
	}

	// Update data in the database via the UpdateDestinasi service
	result, err := h.ppdbSvc.UpdateFasilitas(r.Context(), TableFasilitas, fasilitasID)
	if err != nil {
		http.Error(w, "Gagal memperbarui data fasilitas", http.StatusInternalServerError)
		log.Printf("[ERROR] %s %s - %s\n", r.Method, r.URL, err.Error())
		return
	}

	resp.Data = result

	ctx := r.Context()
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	// Response with success
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
	w.Write([]byte("Fasilitas berhasil diperbarui"))
}

func (h *Handler) UpdateProfileStaff(w http.ResponseWriter, r *http.Request) {
	var (
		staff ppdbEntity.TableStaff
		resp  response.Response
		err   error
	)

	fmt.Println("masuk1")

	// Parse multipart form dengan ukuran maksimum 10MB
	err = r.ParseMultipartForm(10 << 20) // Maksimum ukuran file 10MB
	if err != nil {
		fmt.Println("Error memproses bagian dari form-data:", err)
		http.Error(w, "Error memproses form-data", http.StatusBadRequest)
		return
	}

	// Mengambil file gambar dari form-data
	file, _, err := r.FormFile("staff_photo")
	if err != nil {
		fmt.Println("Error mengambil file dari form-data:", err)
		http.Error(w, "Error mengambil file dari form-data", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Membaca isi file ke dalam byte array
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error membaca isi file ke dalam byte array:", err)
		http.Error(w, "Error membaca isi file", http.StatusInternalServerError)
		return
	}

	// Parse tanggal lahir
	staffTglLahir := r.FormValue("staff_tgl_lahir")
	var dateOnly *time.Time

	if staffTglLahir != "" {
		if strings.TrimSpace(staffTglLahir) != "" { // Memeriksa apakah string tidak kosong
			parsedDate, err := time.Parse("2006-01-02", staffTglLahir)
			// Jika parsing dengan format "2006-01-02" gagal, coba format RFC3339
			parsedDate, err = time.Parse(time.RFC3339, staffTglLahir)
			if err != nil {
				log.Printf("Error parsing date: %s, Error: %v", staffTglLahir, err)
				http.Error(w, "Error memproses tanggal awal", http.StatusBadRequest)
				return
			}
			dateOnly = &parsedDate
		}
	}
	fmt.Println("tetsinggg3", staffTglLahir)

	// Membaca tempat lahir dari form-data
	staffTmptLahir := r.FormValue("staff_tmpt_lahir")
	var tempatLahir *string
	if staffTmptLahir != "" {
		tempatLahir = &staffTmptLahir
	}
	fmt.Println("tetsinggg4", staffTglLahir)

	// Membaca data JSON yang lain dari form-data
	staffID := r.FormValue("staffID")
	staff = ppdbEntity.TableStaff{
		StaffID:        staffID,
		StaffPhoto:     fileBytes,
		StaffName:      r.FormValue("staff_name"),
		StaffGender:    r.FormValue("staff_gender"),
		StaffPosition:  r.FormValue("staff_position"),
		StaffTglLahir:  dateOnly,    // Mengatur nilai pointer
		StaffTmptLahir: tempatLahir, // Mengatur nilai pointer
	}
	fmt.Println("tetsinggg5", staffTglLahir)

	// Memperbarui data ke dalam database melalui layanan UpdateProfileStaff
	result, err := h.ppdbSvc.UpdateProfileStaff(r.Context(), staff, staffID)
	if err != nil {
		resp.SetError(err, http.StatusInternalServerError)
		resp.StatusCode = 500
		log.Printf("[ERROR] %s %s - %s\n", r.Method, r.URL, err.Error())
		json.NewEncoder(w).Encode(resp)
		return
	}
	fmt.Println("tetsinggg6", staffTglLahir)

	// Mengembalikan response
	resp.Data = result
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
	w.Write([]byte("Data Staff berhasil diperbarui"))
}

func (h *Handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	var (
		event ppdbEntity.TableEvent
		resp  response.Response
		err   error
	)

	// Parse multipart form dengan ukuran maksimum 10MB
	err = r.ParseMultipartForm(10 << 20) // Maksimum ukuran file 10MB
	if err != nil {
		fmt.Println("Error memproses bagian dari form-data:", err)
		http.Error(w, "Error memproses form-data", http.StatusBadRequest)
		return
	}

	// Mengambil file gambar dari form-data
	file, _, err := r.FormFile("event_image")
	if err != nil {
		fmt.Println("Error mengambil file dari form-data:", err)
		http.Error(w, "Error mengambil file dari form-data", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Membaca isi file ke dalam byte array
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error membaca isi file ke dalam byte array:", err)
		http.Error(w, "Error membaca isi file", http.StatusInternalServerError)
		return
	}

	eventStartDate := r.FormValue("event_start_date")

	if eventStartDate != "" {
		if strings.TrimSpace(eventStartDate) != "" { // Memeriksa apakah string tidak kosong
			parsedDate, err := time.Parse("2006-01-02", eventStartDate)
			if err != nil {
				// Jika parsing dengan format "2006-01-02" gagal, coba format RFC3339
				parsedDate, err = time.Parse(time.RFC3339, eventStartDate)
				if err != nil {
					log.Printf("Error parsing date: %s, Error: %v", eventStartDate, err)
					http.Error(w, "Error memproses tanggal awal", http.StatusBadRequest)
					return
				}
			}
			event.EventStartDate = parsedDate
		}
	}

	eventEndDate := r.FormValue("event_end_date")
	var dateOnly *time.Time

	if eventEndDate != "" {
		if strings.TrimSpace(eventEndDate) != "" { // Memeriksa apakah string tidak kosong
			parsedDate, err := time.Parse("2006-01-02", eventEndDate)
			if err != nil {
				// Jika parsing dengan format "2006-01-02" gagal, coba format RFC3339
				parsedDate, err = time.Parse(time.RFC3339, eventEndDate)
				if err != nil {
					log.Printf("Error parsing date: %s, Error: %v", eventEndDate, err)
					http.Error(w, "Error memproses tanggal akhir", http.StatusBadRequest)
					return
				}
			}
			dateOnly = &parsedDate
		}
	}

	// Membaca data JSON yang lain dari form-data
	eventID := r.FormValue("eventID")
	event = ppdbEntity.TableEvent{
		EventID:        eventID,
		EventHeader:    r.FormValue("event_header"),
		EventStartDate: event.EventStartDate,
		EventEndDate:   dateOnly,
		EventDesc:      r.FormValue("event_desc"),
		EventImage:     fileBytes,
	}

	result, err := h.ppdbSvc.UpdateEvent(r.Context(), event, eventID)
	if err != nil {
		resp.SetError(err, http.StatusInternalServerError)
		resp.StatusCode = 500
		log.Printf("[ERROR] %s %s - %s\n", r.Method, r.URL, err.Error())
		json.NewEncoder(w).Encode(resp)
		return
	}

	// Mengembalikan response
	resp.Data = result
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
	w.Write([]byte("Data Staff berhasil diperbarui"))
}

func (h *Handler) UpdatePembayaranFormulir(w http.ResponseWriter, r *http.Request) {
	var (
		pembayaranformulir ppdbEntity.TablePembayaranFormulir
		resp               response.Response
		fileData           []byte
	)

	pembayaranformulir.PembayaranID = r.FormValue("pembayaran_id")
	pembayaranformulir.StatusID = r.FormValue("status_id")

	file, _, err := r.FormFile("bukti_pembayaran")
	if err != nil && file != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR] Unable to parse form file")
		return
	}

	if file != nil {
		fileData, err = io.ReadAll(file)
		if err != nil {
			resp = httpHelper.ParseErrorCode(err.Error())
			log.Printf("[ERROR] Unable to read file")
			return
		}
		defer file.Close()
	} else {
		fileData = nil
	}

	pembayaranformulir.BuktiPembayaran = fileData

	// Panggil service untuk memasukkan data pembayaranformulir
	result, err := h.ppdbSvc.UpdatePembayaranFormulir(r.Context(), pembayaranformulir)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp.Data = result
	resp.Message = "Pembayaran formulir data updated successfully" // Menyusun pesan respons

	// Mengambil konteks dari request
	ctx := r.Context()
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	// Mengembalikan response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) UpdateFormulir(w http.ResponseWriter, r *http.Request) {
	var (
		formulir ppdbEntity.TableDataFormulir
		resp     response.Response
	)

	// Decode JSON dari body request
	err := json.NewDecoder(r.Body).Decode(&formulir)
	if err != nil {
		fmt.Println("err", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Panggil service untuk memasukkan data formulir
	result, err := h.ppdbSvc.UpdateFormulir(r.Context(), formulir)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp.Data = result
	resp.Message = "Formulir data updated successfully" // Menyusun pesan respons

	// Mengambil konteks dari request
	ctx := r.Context()
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	// Mengembalikan response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) UpdateBerkas(w http.ResponseWriter, r *http.Request) {
	var (
		berkas ppdbEntity.TableBerkas
		resp   response.Response
	)

	berkas.BerkasID = r.FormValue("berkas_id")
	berkas.StatusID = r.FormValue("status_id")

	fileakta, _, err := r.FormFile("akta_lahir")
	if err != nil && fileakta != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR] Unable to parse form fileakta")
		return
	}

	if fileakta != nil {
		berkas.AktalLahir, err = io.ReadAll(fileakta)
		if err != nil {
			resp = httpHelper.ParseErrorCode(err.Error())
			log.Printf("[ERROR] Unable to read file")
			return
		}
		defer fileakta.Close()
	} else {
		berkas.AktalLahir = nil
	}

	filephoto, _, err := r.FormFile("pas_photo")
	if err != nil && filephoto != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR] Unable to parse form filephoto")
		return
	}

	if filephoto != nil {
		berkas.PasPhoto, err = io.ReadAll(filephoto)
		if err != nil {
			resp = httpHelper.ParseErrorCode(err.Error())
			log.Printf("[ERROR] Unable to read filephoto")
			return
		}
		defer filephoto.Close()
	} else {
		berkas.PasPhoto = nil
	}

	filerapor, _, err := r.FormFile("rapor")
	if err != nil && filerapor != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR] Unable to parse form filerapor")
		return
	}

	if filerapor != nil {
		berkas.Rapor, err = io.ReadAll(filerapor)
		if err != nil {
			resp = httpHelper.ParseErrorCode(err.Error())
			log.Printf("[ERROR] Unable to read filerapor")
			return
		}
		defer filerapor.Close()
	} else {
		berkas.Rapor = nil
	}

	// Panggil service untuk memasukkan data berkas
	result, err := h.ppdbSvc.UpdateBerkas(r.Context(), berkas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp.Data = result
	resp.Message = "Berkas data updated successfully" // Menyusun pesan respons

	// Mengambil konteks dari request
	ctx := r.Context()
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	// Mengembalikan response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) UpdateJadwalTest(w http.ResponseWriter, r *http.Request) {
	var (
		jadwaltest ppdbEntity.TableJadwalTest
		resp       response.Response
	)

	// Decode JSON dari body request
	err := json.NewDecoder(r.Body).Decode(&jadwaltest)
	if err != nil {
		fmt.Println("err", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Panggil service untuk memasukkan data jadwaltest
	result, err := h.ppdbSvc.UpdateJadwalTest(r.Context(), jadwaltest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp.Data = result
	resp.Message = "Jadwal test data updated successfully" // Menyusun pesan respons

	// Mengambil konteks dari request
	ctx := r.Context()
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	// Mengembalikan response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
