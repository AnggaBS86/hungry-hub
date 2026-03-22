package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/example/hungry-hub/internal/config"
	"github.com/example/hungry-hub/internal/db"
	"github.com/example/hungry-hub/internal/httpapi"
	"github.com/example/hungry-hub/internal/migrate"
	"github.com/example/hungry-hub/internal/seed"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	conn, err := db.OpenMySQL(cfg.DB)
	if err != nil {
		log.Fatalf("db: %v", err)
	}

	sqlDB, err := conn.DB()
	if err != nil {
		log.Fatalf("db: %v", err)
	}
	defer sqlDB.Close()

	if cfg.RunMigrations {
		if err := migrate.RunMySQLMigrations(sqlDB, cfg.MigrationsDir); err != nil {
			log.Fatalf("migrations: %v", err)
		}
	}

	if cfg.RunSeed {
		if err := seed.Run(conn); err != nil {
			log.Fatalf("seed: %v", err)
		}
	}

	e := httpapi.NewServer(conn)

	srv := &http.Server{
		Addr:              ":" + cfg.AppPort,
		Handler:           e,
		ReadHeaderTimeout: 10 * time.Second,
	}

	go func() {
		log.Printf("listening on :%s", cfg.AppPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("http: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = srv.Shutdown(ctx)
	if err != nil {
		panic("Error when shutdown : " + err.Error())
	}
}
