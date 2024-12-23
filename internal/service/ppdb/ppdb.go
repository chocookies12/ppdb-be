package ppdb

import (
	"ppdb-be/internal/entity"
	// "ppdb-be/internal/entity/auth"
	"context"
	"errors"
	jaegerLog "ppdb-be/pkg/log"

	// "time"

	ppdbEntity "ppdb-be/internal/entity/ppdb"

	"github.com/opentracing/opentracing-go"
)

type Data interface {
	//get
	// LoginAdmin(ctx context.Context, emailAdmin string, password string) (string, error)
	GetLoginAdmin(ctx context.Context, login ppdbEntity.TableAdmin) (ppdbEntity.TableAdmin, error)
	GetKontakSekolah(ctx context.Context) ([]ppdbEntity.TableKontakSekolah, error)
	GetDataAdmin(ctx context.Context, searchInput string, offset, limit int) ([]ppdbEntity.TableKelolaDataAdmin, error)
	GetDataAdminPagination(ctx context.Context, searchInput string) (int, error)

	GetRole(ctx context.Context) ([]ppdbEntity.TableRole, error)
	GetAgama(ctx context.Context) ([]ppdbEntity.TableAgama, error)
	GetStatus(ctx context.Context) ([]ppdbEntity.TableStatus, error)
	GetJurusan(ctx context.Context) ([]ppdbEntity.TableJurusan, error)

	//info Daftar
	GetGambarInfoDaftar(ctx context.Context, infoID string) ([]byte, error)
	GetInfoDaftar(ctx context.Context) ([]ppdbEntity.TableInfoDaftar, error)

	//banner
	GetGambarBanner(ctx context.Context, bannerID string) ([]byte, error)
	GetBanner(ctx context.Context) ([]ppdbEntity.TableBanner, error)

	//fasilitas
	GetGambarFasilitas(ctx context.Context, fasilitasID string) ([]byte, error)
	GetFasilitas(ctx context.Context, searchInput string, offset, limit int) ([]ppdbEntity.TableFasilitas, error)
	GetFasilitasPagination(ctx context.Context, searchInput string) (int, error)
	GetFasilitasUtama(ctx context.Context) ([]ppdbEntity.TableFasilitas, error)

	//Profile Staff
	GetPhotoStaff(ctx context.Context, staffID string) ([]byte, error)
	GetProfileStaff(ctx context.Context, searchInput string, offset, limit int) ([]ppdbEntity.TableStaff, error)
	GetProfileStaffPagination(ctx context.Context, searchInput string) (int, error)
	GetProfileStaffUtama(ctx context.Context) ([]ppdbEntity.TableStaff, error)

	//Event
	GetImageEvent(ctx context.Context, eventID string) ([]byte, error)
	GetEvent(ctx context.Context, searchInput string, offset, limit int) ([]ppdbEntity.TableEvent, error)
	GetEventPagination(ctx context.Context, searchInput string) (int, error)
	GetEventDetail(ctx context.Context, eventID string) (ppdbEntity.TableEvent, error)
	GetEventUtama(ctx context.Context) ([]ppdbEntity.TableEvent, error)

	//Peserta Didik
	GetLoginCheck(ctx context.Context, login ppdbEntity.TablePesertaDidik) (ppdbEntity.TablePesertaDidik, error)

	GetPesertaDidikDetail(ctx context.Context, idpesertadidik string) (ppdbEntity.TablePesertaDidik, error)

	GetPembayaranFormulirDetail(ctx context.Context, idpesertadidik string) (ppdbEntity.TablePembayaranFormulir, error)

	GetFormulirDetail(ctx context.Context, idpesertadidik string) (ppdbEntity.TableDataFormulir, error)

	GetBerkasDetail(ctx context.Context, idpesertadidik string) (ppdbEntity.TableBerkas, error)

	GetJadwalTestDetail(ctx context.Context, idpesertadidik string) (ppdbEntity.TableJadwalTest, error)

	//JadwalTest Admin
	GetJadwalTestAll(ctx context.Context, searchInput string, offset, limit int) ([]ppdbEntity.TableJadwalTest, error)
	GetJadwalTestPagination(ctx context.Context, searchInput string) (int, error)

	//Pembayaran Formulir Admin
	GetPembayaranFormulirAll(ctx context.Context, searchInput string, offset, limit int) ([]ppdbEntity.TablePembayaranFormulir, error)
	GetPembayaranFormulirPagination(ctx context.Context, searchInput string) (int, error)

	//Data Siswa Admin
	GetFormulirAll(ctx context.Context, searchInput string, offset, limit int) ([]ppdbEntity.TableDataFormulir, error)
	GetFormulirPagination(ctx context.Context, searchInput string) (int, error)

	//PesertaName
	GetPesertaName(ctx context.Context, idpembayaran string) (ppdbEntity.TablePesertaDidik, error)

	//Dashboard
	GetCountDataWeb(ctx context.Context) (ppdbEntity.CountDataWeb, error)
	GetCountPesertaDidik(ctx context.Context, tahun int) (ppdbEntity.CountPesertaDidik, error)
	GetCountBuktiPembayaran(ctx context.Context, tahun int) (ppdbEntity.CountBuktiPembayaran, error)
	GetCountFormulir(ctx context.Context, tahun int) (ppdbEntity.CountFormulir, error)

	//insert
	InsertDataAdmin(ctx context.Context, admin ppdbEntity.TableAdmin) (string, error)
	InsertInfoDaftar(ctx context.Context, infoDaftar ppdbEntity.TableInfoDaftar) (string, error)

	InsertBanner(ctx context.Context, banner ppdbEntity.TableBanner) (string, error)
	InsertFasilitas(ctx context.Context, fasilitas ppdbEntity.TableFasilitas) (string, error)
	InsertProfileStaff(ctx context.Context, staff ppdbEntity.TableStaff) (string, error)
	InsertEvent(ctx context.Context, event ppdbEntity.TableEvent) (string, error)

	InsertPesertaDidik(ctx context.Context, pesertadidik ppdbEntity.TablePesertaDidik) (string, error)

	//delete
	DeleteAdmin(ctx context.Context, adminID string) (string, error)
	DeleteBanner(ctx context.Context, bannerID string) (string, error)
	DeleteFasilitas(ctx context.Context, fasilitasID string) (string, error)
	DeleteProfileStaff(ctx context.Context, staffID string) (string, error)
	DeleteEvent(ctx context.Context, eventID string) (string, error)

	//update
	UpdateBanner(ctx context.Context, banner ppdbEntity.TableBanner, bannerID string) (string, error)
	UpdateInfoDaftar(ctx context.Context, infoDaftar ppdbEntity.TableInfoDaftar, infoID string) (string, error)
	UpdateFasilitas(ctx context.Context, fasilitas ppdbEntity.TableFasilitas, fasilitasID string) (string, error)
	UpdateProfileStaff(ctx context.Context, staff ppdbEntity.TableStaff, staffID string) (string, error)
	UpdateEvent(ctx context.Context, event ppdbEntity.TableEvent, eventID string) (string, error)
	UpdatePembayaranFormulir(ctx context.Context, pembayaranformulir ppdbEntity.TablePembayaranFormulir) (string, error)
	UpdateFormulir(ctx context.Context, formulir ppdbEntity.TableDataFormulir) (string, error)
	UpdateKontakPeserta(ctx context.Context, kontakpeserta ppdbEntity.TableDataFormulir) (string, error)
	UpdateOrtu(ctx context.Context, ortu ppdbEntity.TableDataFormulir) (string, error)
	UpdateBerkas(ctx context.Context, berkas ppdbEntity.TableBerkas) (string, error)
	UpdateJadwalTest(ctx context.Context, jadwalTest ppdbEntity.TableJadwalTest) (string, error)
	UpdateStatusPembayaranFormulir(ctx context.Context, pembayaranformulir ppdbEntity.TablePembayaranFormulir) (string, error)
}

type Service struct {
	ppdb   Data
	tracer opentracing.Tracer
	logger jaegerLog.Factory
}

// New ...
// Tambahkan parameter sesuai banyak data layer yang dibutuhkan
func New(ppdbData Data, tracer opentracing.Tracer, logger jaegerLog.Factory) Service {
	// Assign variable dari parameter ke object
	return Service{
		ppdb:   ppdbData,
		tracer: tracer,
		logger: logger,
	}
}

func (s Service) checkPermission(ctx context.Context, _permissions ...string) error {
	claims := ctx.Value(entity.ContextKey("claims"))
	if claims != nil {
		actions := claims.(entity.ContextValue).Get("permissions").(map[string]interface{})
		for _, action := range actions {
			permissions := action.([]interface{})
			for _, permission := range permissions {
				for _, _permission := range _permissions {
					if permission.(string) == _permission {
						return nil
					}
				}
			}
		}
	}
	return errors.New("401 unauthorized")
}
