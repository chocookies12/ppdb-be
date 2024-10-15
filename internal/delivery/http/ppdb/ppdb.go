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

	GetGambarInfoDaftar(ctx context.Context, infoID string) ([]byte, error)
	GetInfoDaftar(ctx context.Context) ([]ppdbEntity.TableInfoDaftar, error)

	GetGambarBanner(ctx context.Context, bannerID string) ([]byte, error) 
	GetBanner(ctx context.Context) ([]ppdbEntity.TableBanner, error) 

	//insert
	InsertDataAdmin(ctx context.Context, admin ppdbEntity.TableAdmin) (string, error)
	InsertInfoDaftar(ctx context.Context, infoDaftar ppdbEntity.TableInfoDaftar) (string, error)
	InsertBanner(ctx context.Context, banner ppdbEntity.TableBanner) (string, error)
	InsertFasilitas(ctx context.Context, fasilitas ppdbEntity.TableFasilitas) (string, error) 

	//delete
	DeleteAdmin(ctx context.Context, adminID string) (string, error)
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
