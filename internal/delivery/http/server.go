package http

import (
	"net/http"

	"ppdb-be/pkg/grace"

	"github.com/rs/cors"
)

// Handler ...
type Handler interface {
	GetPpdb(w http.ResponseWriter, r *http.Request)
	InsertPpdb(w http.ResponseWriter, r *http.Request)
	DeletePpdb(w http.ResponseWriter, r *http.Request)
	UpdatePpdb(w http.ResponseWriter, r *http.Request)

	//website ppdb
	GetKontakSekolah(w http.ResponseWriter, r *http.Request)
	GetRole(w http.ResponseWriter, r *http.Request)
	//kelola Admin 
	GetDataAdminSlim(w http.ResponseWriter, r *http.Request)
	InsertDataAdmin(w http.ResponseWriter, r *http.Request)
	DeleteAdmin(w http.ResponseWriter, r *http.Request) 	
	//Info Pendaftaran 
	InsertInfoDaftar(w http.ResponseWriter, r *http.Request)
	GetGambarInfoDaftar(w http.ResponseWriter, r *http.Request)
	GetInfoDaftar(w http.ResponseWriter, r *http.Request)

	//Banner
	InsertBanner(w http.ResponseWriter, r *http.Request)
	GetGambarBanner(w http.ResponseWriter, r *http.Request)
	GetBanner(w http.ResponseWriter, r *http.Request) 
	
	//Fasilitas 
	InsertFasilitas(w http.ResponseWriter, r *http.Request)
}


// Server ...
type Server struct {
	Ppdb Handler
}

// Serve is serving HTTP gracefully on port x ...
func (s *Server) Serve(port string) error {
	handler := cors.AllowAll().Handler(s.Handler())
	return grace.Serve(port, handler)
}
