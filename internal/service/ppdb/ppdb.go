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
	LoginAdmin(ctx context.Context, emailAdmin string, password string) (string, error)
	GetKontakSekolah(ctx context.Context) ([]ppdbEntity.TableKontakSekolah, error)
	GetDataAdmin(ctx context.Context, searchInput string, offset, limit int) ([]ppdbEntity.TableKelolaDataAdmin, error)
	GetDataAdminPagination(ctx context.Context, searchInput string) (int, error)

	GetRole(ctx context.Context) ([]ppdbEntity.TableRole, error)

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

	//insert
	InsertDataAdmin(ctx context.Context, admin ppdbEntity.TableAdmin) (string, error)
	InsertInfoDaftar(ctx context.Context, infoDaftar ppdbEntity.TableInfoDaftar) (string, error)

	InsertBanner(ctx context.Context, banner ppdbEntity.TableBanner) (string, error)
	InsertFasilitas(ctx context.Context, fasilitas ppdbEntity.TableFasilitas) (string, error)

	//delete
	DeleteAdmin(ctx context.Context, adminID string) (string, error)
	DeleteBanner(ctx context.Context, bannerID string) (string, error) 
	DeleteFasilitas(ctx context.Context, fasilitasID string) (string, error) 
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
