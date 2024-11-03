package ppdb

import (
	// "bytes"
	"context"
	// JOEntity "ppdb-be/internal/entity/ppdb"
	jaegerLog "ppdb-be/pkg/log"

	// "context"
	ppdbEntity "ppdb-be/internal/entity/ppdb"

	"github.com/opentracing/opentracing-go"
)

type IppdbSvc interface {
	//get
	LoginAdmin(ctx context.Context, emailAdmin string, password string) (string, error)
	GetKontakSekolah(ctx context.Context) ([]ppdbEntity.TableKontakSekolah, error)
	GetDataAdminSlim(ctx context.Context, searchInput string, page, length int) ([]ppdbEntity.TableKelolaDataAdmin, interface{}, error)
	GetRole(ctx context.Context) ([]ppdbEntity.TableRole, error)
	GetStatus(ctx context.Context) ([]ppdbEntity.TableStatus, error)

	GetGambarInfoDaftar(ctx context.Context, infoID string) ([]byte, error)
	GetInfoDaftar(ctx context.Context) ([]ppdbEntity.TableInfoDaftar, error)

	GetGambarBanner(ctx context.Context, bannerID string) ([]byte, error)
	GetBanner(ctx context.Context) ([]ppdbEntity.TableBanner, error)

	GetGambarFasilitas(ctx context.Context, fasilitasID string) ([]byte, error)
	GetFasilitasSlim(ctx context.Context, searchInput string, page, length int) ([]ppdbEntity.TableFasilitas, interface{}, error)
	GetFasilitas(ctx context.Context) ([]ppdbEntity.TableFasilitas, error)

	GetPhotoStaff(ctx context.Context, staffID string) ([]byte, error)
	GetProfileStaffSlim(ctx context.Context, searchInput string, page, length int) ([]ppdbEntity.TableStaff, interface{}, error)
	GetProfileStaffUtama(ctx context.Context) ([]ppdbEntity.TableStaff, error)

	GetImageEvent(ctx context.Context, eventID string) ([]byte, error)
	GetEventSlim(ctx context.Context, searchInput string, page, length int) ([]ppdbEntity.TableEvent, interface{}, error)
	GetEventDetail(ctx context.Context, eventID string) (ppdbEntity.TableEvent, error)
	GetEventUtama(ctx context.Context) ([]ppdbEntity.TableEvent, error)

	//Peserta Didik
	GetLoginCheck(ctx context.Context, login ppdbEntity.TablePesertaDidik) (ppdbEntity.TablePesertaDidik, error)

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
}

type (
	// Handler ...
	Handler struct {
		ppdbSvc IppdbSvc
		tracer  opentracing.Tracer
		logger  jaegerLog.Factory
	}
)

// New for bridging product handler initialization
func New(is IppdbSvc, tracer opentracing.Tracer, logger jaegerLog.Factory) *Handler {
	return &Handler{
		ppdbSvc: is,
		tracer:  tracer,
		logger:  logger,
	}
}
