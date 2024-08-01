package ppdb

import (
	// "bytes"
	"context"
	// JOEntity "ppdb-be/internal/entity/ppdb"
	jaegerLog "ppdb-be/pkg/log"

	// "context"

	"github.com/opentracing/opentracing-go"
)

type IppdbSvc interface {
	//get
	LoginAdmin(ctx context.Context, admin_id string, admin_password string) (string, error)
	
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
