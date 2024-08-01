package boot

import (

	// "ppdb-be/internal/data/auth"
	// "ppdb-be/pkg/httpclient"
	"log"
	"net/http"
	"ppdb-be/pkg/tracing"

	"ppdb-be/internal/config"
	jaegerLog "ppdb-be/pkg/log"

	"github.com/fsnotify/fsnotify"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	ppdbData "ppdb-be/internal/data/ppdb"
	ppdbServer "ppdb-be/internal/delivery/http"
	ppdbHandler "ppdb-be/internal/delivery/http/ppdb"
	ppdbService "ppdb-be/internal/service/ppdb"
	// coreData "ppdb-be/internal/data/core"
)

// HTTP will load configuration, do dependency injection and then start the HTTP server
func HTTP() error {
	err := config.Init()
	if err != nil {
		log.Fatalf("[CONFIG] Failed to initialize config: %v", err)
	}
	cfg := config.Get()
	// Open MySQL DB Connection
	db, err := openDatabases(cfg)
	if err != nil {
		log.Fatalf("[DB] Failed to initialize database connection: %v", err)
	}

	// //
	// docs.SwaggerInfo.Host = cfg.Swagger.Host
	// docs.SwaggerInfo.Schemes = cfg.Swagger.Schemes

	// Set logger used for jaeger
	logger, _ := zap.NewDevelopment(
		zap.AddStacktrace(zapcore.FatalLevel),
		zap.AddCallerSkip(1),
	)
	zapLogger := logger.With(zap.String("service", "ppdb"))
	zlogger := jaegerLog.NewFactory(zapLogger)

	// Set tracer for service
	tracer, closer := tracing.Init("ppdb", zlogger)
	defer closer.Close()

	// httpc := httpclient.NewClient(tracer)
	// ad := auth.New(nil, cfg.API.Auth)

	// cod := coreData.New(nil, cfg.API.Core)

	// Diganti dengan domain yang anda buat
	sd := ppdbData.New(db, tracer, zlogger)
	// ss := joborderService.New(sd, ad, nil, tracer, zlogger)
	ss := ppdbService.New(sd, tracer, zlogger)
	sh := ppdbHandler.New(ss, tracer, zlogger)

	config.PrepareWatchPath()
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		err := config.Init()
		if err != nil {
			log.Printf("[VIPER] Error get config file, %v", err)
		}
		cfg := config.Get()
		masterNew, err := openDatabases(cfg)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize database connection: %v", err)
		} else {
			*db = *masterNew
			sd.InitStmt()
		}

	})
	s := ppdbServer.Server{
		Ppdb: sh,
	}

	if err := s.Serve(cfg.Server.Port); err != http.ErrServerClosed {
		return err
	}

	return nil
}

func openDatabases(cfg *config.Config) (master *sqlx.DB, err error) {
	master, err = openConnectionPool("mysql", cfg.Database.Master)
	if err != nil {
		return master, err
	}

	return master, err
}

func openConnectionPool(driver string, connString string) (db *sqlx.DB, err error) {
	db, err = sqlx.Open(driver, connString)
	if err != nil {
		return db, err
	}

	err = db.Ping()
	if err != nil {
		return db, err
	}

	return db, err
}
