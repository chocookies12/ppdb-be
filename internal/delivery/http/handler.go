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

	ppdb.HandleFunc("/getkontaksekolah", s.Ppdb.GetKontakSekolah).Methods("GET")
	ppdb.HandleFunc("/getrole", s.Ppdb.GetRole).Methods("GET")
	//DataAdmin
	ppdb.HandleFunc("/getdataadmin", s.Ppdb.GetDataAdminSlim).Methods("GET")
	ppdb.HandleFunc("/insertdataadmin", s.Ppdb.InsertDataAdmin).Methods("POST")
	ppdb.HandleFunc("/deletedataadmin", s.Ppdb.DeleteAdmin).Methods("DELETE")
	//Info Penddaftaran
	ppdb.HandleFunc("/insertinfopendaftaran", s.Ppdb.InsertInfoDaftar).Methods("POST")
	ppdb.HandleFunc("/getgambarinfodaftar", s.Ppdb.GetGambarInfoDaftar).Methods("GET")
	ppdb.HandleFunc("/getinfodaftar", s.Ppdb.GetInfoDaftar).Methods("GET")
	//Banner
	ppdb.HandleFunc("/insertbanner", s.Ppdb.InsertBanner).Methods("POST")
	ppdb.HandleFunc("/getgambarbanner", s.Ppdb.GetGambarBanner).Methods("GET")
	ppdb.HandleFunc("/getbanner", s.Ppdb.GetBanner).Methods("GET")

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
