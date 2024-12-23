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
	GetLoginAdmin(w http.ResponseWriter, r *http.Request)
	GetKontakSekolah(w http.ResponseWriter, r *http.Request)
	GetRole(w http.ResponseWriter, r *http.Request)
	GetStatus(w http.ResponseWriter, r *http.Request)
	GetAgama(w http.ResponseWriter, r *http.Request)
	GetJurusan(w http.ResponseWriter, r *http.Request)

	//kelola Admin
	GetDataAdminSlim(w http.ResponseWriter, r *http.Request)
	InsertDataAdmin(w http.ResponseWriter, r *http.Request)
	DeleteAdmin(w http.ResponseWriter, r *http.Request)

	//Info Pendaftaran
	InsertInfoDaftar(w http.ResponseWriter, r *http.Request)
	GetGambarInfoDaftar(w http.ResponseWriter, r *http.Request)
	GetInfoDaftar(w http.ResponseWriter, r *http.Request)
	UpdateInfoDaftar(w http.ResponseWriter, r *http.Request)

	//Banner
	InsertBanner(w http.ResponseWriter, r *http.Request)
	GetGambarBanner(w http.ResponseWriter, r *http.Request)
	GetBanner(w http.ResponseWriter, r *http.Request)
	DeleteBanner(w http.ResponseWriter, r *http.Request)
	UpdateBanner(w http.ResponseWriter, r *http.Request)

	//Fasilitas
	InsertFasilitas(w http.ResponseWriter, r *http.Request)
	GetGambarFasilitas(w http.ResponseWriter, r *http.Request)
	GetFasilitasSlim(w http.ResponseWriter, r *http.Request)
	GetFasilitas(w http.ResponseWriter, r *http.Request)
	DeleteFasilitas(w http.ResponseWriter, r *http.Request)
	UpdateFasilitas(w http.ResponseWriter, r *http.Request)

	//Profile Staff
	InsertProfileStaff(w http.ResponseWriter, r *http.Request)
	GetPhotoStaff(w http.ResponseWriter, r *http.Request)
	GetProfileStaffSlim(w http.ResponseWriter, r *http.Request)
	DeleteProfileStaff(w http.ResponseWriter, r *http.Request)
	GetProfilStaffUtama(w http.ResponseWriter, r *http.Request)
	UpdateProfileStaff(w http.ResponseWriter, r *http.Request)

	//Event Sekolah
	InsertEvent(w http.ResponseWriter, r *http.Request)
	GetImageEvent(w http.ResponseWriter, r *http.Request)
	GetEventSlim(w http.ResponseWriter, r *http.Request)
	GetEventDetail(w http.ResponseWriter, r *http.Request)
	GetEventUtama(w http.ResponseWriter, r *http.Request)
	DeleteEvent(w http.ResponseWriter, r *http.Request)
	UpdateEvent(w http.ResponseWriter, r *http.Request)

	//Peserta Didik
	InsertPesertaDidik(w http.ResponseWriter, r *http.Request)
	GetLoginCheck(w http.ResponseWriter, r *http.Request)

	GetPembayaranFormulirDetail(w http.ResponseWriter, r *http.Request)
	GetFormulirDetail(w http.ResponseWriter, r *http.Request)
	GetBerkasDetail(w http.ResponseWriter, r *http.Request)
	GetJadwalTestDetail(w http.ResponseWriter, r *http.Request)
	GetGeneratedKartuTest(w http.ResponseWriter, r *http.Request)
	GetGeneratedFormulir(w http.ResponseWriter, r *http.Request)

	UpdatePembayaranFormulir(w http.ResponseWriter, r *http.Request)
	UpdateFormulir(w http.ResponseWriter, r *http.Request)
	UpdateBerkas(w http.ResponseWriter, r *http.Request)
	UpdateJadwalTest(w http.ResponseWriter, r *http.Request)

	UpdateStatusPembayaranFormulir(w http.ResponseWriter, r *http.Request)

	//JadwalTest, pembayaran formulir, formulir di admin
	GetJadwalTestSlim(w http.ResponseWriter, r *http.Request)
	GetPembayaranFormulirSlim(w http.ResponseWriter, r *http.Request) 
	GetFormulirSlim(w http.ResponseWriter, r *http.Request)

	//Dashboard
	GetCountDataWeb(w http.ResponseWriter, r *http.Request)
	GetCountDataPpdb(w http.ResponseWriter, r *http.Request)
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
