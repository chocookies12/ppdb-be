package http

import (
	"errors"
	"log"
	"net/http"

	"ppdb-be/pkg/response"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
)

// Handler will initialize mux router and register handler
func (s *Server) Handler() *mux.Router {
	r := mux.NewRouter()
	// Jika tidak ditemukan, jangan diubah.
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	// Health Check
	r.HandleFunc("", defaultHandler).Methods("GET")
	r.HandleFunc("/", defaultHandler).Methods("GET")

	// Tambahan Prefix di depan API endpoint
	router := r.PathPrefix("/ppdb").Subrouter()

	router.HandleFunc("", defaultHandler).Methods("GET")
	router.HandleFunc("/", defaultHandler).Methods("GET")

	sub := router.PathPrefix("/v1").Subrouter()

	// Routes
	ppdb := sub.PathPrefix("/data").Subrouter()
	ppdb.HandleFunc("", s.Ppdb.GetPpdb).Methods("GET")
	ppdb.HandleFunc("", s.Ppdb.InsertPpdb).Methods("POST")
	ppdb.HandleFunc("", s.Ppdb.UpdatePpdb).Methods("PUT")
	ppdb.HandleFunc("", s.Ppdb.DeletePpdb).Methods("DELETE")

	ppdb.HandleFunc("/loginadmin", s.Ppdb.GetLoginAdmin).Methods("POST")

	ppdb.HandleFunc("/getkontaksekolah", s.Ppdb.GetKontakSekolah).Methods("GET")
	ppdb.HandleFunc("/getrole", s.Ppdb.GetRole).Methods("GET")
	ppdb.HandleFunc("/getagama", s.Ppdb.GetAgama).Methods("GET")
	ppdb.HandleFunc("/getstatus", s.Ppdb.GetStatus).Methods("GET")
	ppdb.HandleFunc("/getjurusan", s.Ppdb.GetJurusan).Methods("GET")

	//DataAdmin
	ppdb.HandleFunc("/getdataadmin", s.Ppdb.GetDataAdminSlim).Methods("GET")
	ppdb.HandleFunc("/insertdataadmin", s.Ppdb.InsertDataAdmin).Methods("POST")
	ppdb.HandleFunc("/deletedataadmin", s.Ppdb.DeleteAdmin).Methods("DELETE")
	//Info Penddaftaran
	ppdb.HandleFunc("/insertinfopendaftaran", s.Ppdb.InsertInfoDaftar).Methods("POST")
	ppdb.HandleFunc("/getgambarinfodaftar", s.Ppdb.GetGambarInfoDaftar).Methods("GET")
	ppdb.HandleFunc("/getinfodaftar", s.Ppdb.GetInfoDaftar).Methods("GET")
	ppdb.HandleFunc("/updateinfodaftar", s.Ppdb.UpdateInfoDaftar).Methods("PUT")

	//Banner
	ppdb.HandleFunc("/insertbanner", s.Ppdb.InsertBanner).Methods("POST")
	ppdb.HandleFunc("/getgambarbanner", s.Ppdb.GetGambarBanner).Methods("GET")
	ppdb.HandleFunc("/getbanner", s.Ppdb.GetBanner).Methods("GET")
	ppdb.HandleFunc("/deletebanner", s.Ppdb.DeleteBanner).Methods("DELETE")
	ppdb.HandleFunc("/updatebanner", s.Ppdb.UpdateBanner).Methods("PUT")

	//Fasilitas
	ppdb.HandleFunc("/insertfasilitas", s.Ppdb.InsertFasilitas).Methods("POST")
	ppdb.HandleFunc("/getgambarfasilitas", s.Ppdb.GetGambarFasilitas).Methods("GET")
	ppdb.HandleFunc("/getfasilitasslim", s.Ppdb.GetFasilitasSlim).Methods("GET")
	ppdb.HandleFunc("/getfasilitas", s.Ppdb.GetFasilitas).Methods("GET")
	ppdb.HandleFunc("/deletefasilitas", s.Ppdb.DeleteFasilitas).Methods("DELETE")
	ppdb.HandleFunc("/updatefasilitas", s.Ppdb.UpdateFasilitas).Methods("PUT")

	//Profile Staff
	ppdb.HandleFunc("/insertprofilestaff", s.Ppdb.InsertProfileStaff).Methods("POST")
	ppdb.HandleFunc("/getphotostaff", s.Ppdb.GetPhotoStaff).Methods("GET")
	ppdb.HandleFunc("/getprofilstaffslim", s.Ppdb.GetProfileStaffSlim).Methods("GET") // slim untuk pagination
	ppdb.HandleFunc("/deleteprofilstaff", s.Ppdb.DeleteProfileStaff).Methods("DELETE")
	ppdb.HandleFunc("/getprofilstaff", s.Ppdb.GetProfilStaffUtama).Methods("GET")
	ppdb.HandleFunc("/updateprofilstaff", s.Ppdb.UpdateProfileStaff).Methods("PUT")

	//Event Sekolah
	ppdb.HandleFunc("/insertevent", s.Ppdb.InsertEvent).Methods("POST")
	ppdb.HandleFunc("/getimageevent", s.Ppdb.GetImageEvent).Methods("GET")
	ppdb.HandleFunc("/geteventslim", s.Ppdb.GetEventSlim).Methods("GET")
	ppdb.HandleFunc("/geteventdetail", s.Ppdb.GetEventDetail).Methods("GET")
	ppdb.HandleFunc("/geteventutama", s.Ppdb.GetEventUtama).Methods("GET")
	ppdb.HandleFunc("/deleteevent", s.Ppdb.DeleteEvent).Methods("DELETE")
	ppdb.HandleFunc("/updateevent", s.Ppdb.UpdateEvent).Methods("PUT")

	//Peserta Didik
	ppdb.HandleFunc("/registeraccount", s.Ppdb.InsertPesertaDidik).Methods("POST")
	ppdb.HandleFunc("/login", s.Ppdb.GetLoginCheck).Methods("POST")

	ppdb.HandleFunc("/getpembayaranformulirdetail", s.Ppdb.GetPembayaranFormulirDetail).Methods("GET")
	ppdb.HandleFunc("/getformulirdetail", s.Ppdb.GetFormulirDetail).Methods("GET")
	ppdb.HandleFunc("/getberkasdetail", s.Ppdb.GetBerkasDetail).Methods("GET")
	ppdb.HandleFunc("/getjadwaltestdetail", s.Ppdb.GetJadwalTestDetail).Methods("GET")
	ppdb.HandleFunc("/getkartutest", s.Ppdb.GetGeneratedKartuTest).Methods("GET")
	ppdb.HandleFunc("/getformulir", s.Ppdb.GetGeneratedFormulir).Methods("GET")

	ppdb.HandleFunc("/updatepembayaranformulir", s.Ppdb.UpdatePembayaranFormulir).Methods("PUT")
	ppdb.HandleFunc("/updateformulir", s.Ppdb.UpdateFormulir).Methods("PUT")
	ppdb.HandleFunc("/updateberkas", s.Ppdb.UpdateBerkas).Methods("PUT")
	ppdb.HandleFunc("/updatejadwaltest", s.Ppdb.UpdateJadwalTest).Methods("PUT")
	ppdb.HandleFunc("/updatestatuspembayaranformulir", s.Ppdb.UpdateStatusPembayaranFormulir).Methods("PUT")


	//Jadwal Test Admin
	ppdb.HandleFunc("/getjadwaltestslim", s.Ppdb.GetJadwalTestSlim).Methods("GET")

	//Pembayaran Formulir Admin
	ppdb.HandleFunc("/getpembayaranformulirslim", s.Ppdb.GetPembayaranFormulirSlim).Methods("GET")

	//Formulir Admin 
	ppdb.HandleFunc("/getformulirslim", s.Ppdb.GetFormulirSlim).Methods("GET")

	//Dashboard
	ppdb.HandleFunc("/getcountdataweb", s.Ppdb.GetCountDataWeb).Methods("GET")
	ppdb.HandleFunc("/getcountdatappdb", s.Ppdb.GetCountDataPpdb).Methods("GET")



	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	return r
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Example Service API"))
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	var (
		resp   *response.Response
		err    error
		errRes response.Error
	)
	resp = &response.Response{}
	defer resp.RenderJSON(w, r)

	err = errors.New("404 Not Found")

	if err != nil {
		// Error response handling
		errRes = response.Error{
			Code:   404,
			Msg:    "404 Not Found",
			Status: true,
		}

		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		resp.StatusCode = 404
		resp.Error = errRes
		return
	}
}
