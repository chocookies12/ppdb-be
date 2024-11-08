package ppdb

import (
	"ppdb-be/internal/entity"
	// "ppdb-be/internal/entity/auth"
	"context"
	"errors"
	jaegerLog "ppdb-be/pkg/log"

	// "time"

	"github.com/opentracing/opentracing-go"
)

type Data interface {
	//get
	LoginAdmin(ctx context.Context, admin_id string, admin_password string) (string, error)
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
