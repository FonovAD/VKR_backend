package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"time"
	"vkr/internal/logger"
	"vkr/internal/presenter/http/handler"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/file"
	_ "github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"

	"vkr/internal/infrastructure/db/pg" // ← путь к твоему pg-пакету
	"vkr/internal/interactor"
	"vkr/internal/presenter/http/router"
)

// Config — корневая структура конфигурации приложения
type Config struct {
	Postgres pg.PGConfig `yaml:"postgres"`
	Port     string      `yaml:"port" required:"true"`
}

func main() {
	rootPath := parseRootPath()
	cfg, err := loadConfig(rootPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	conn, err := pg.NewPGConnection(cfg.Postgres)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("failed to close db connection: %v", err)
		}
	}()

	e := echo.New()
	e.HideBanner = true

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			err := next(c)
			duration := time.Since(start)
			log.Printf("HTTP %s %s %d %v", c.Request().Method, c.Path(), c.Response().Status, duration)
			return err
		}
	})

	handler.AppMiddleware(e)

	appLogger := logger.NewLogger()
	appLogger.LogInfo("main", nil, "application starting")

	i := interactor.NewInteractor(conn, appLogger)
	h := i.NewAppHandler()

	router.NewRouter(e, h)

	addr := ":" + cfg.Port
	log.Printf("server starting on %s", addr)

	go func() {
		if err := e.Start(addr); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server shutdown unexpectedly: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Fatalf("failed to shutdown server: %v", err)
	}
}

func parseRootPath() string {
	var rootPath string
	flag.StringVar(&rootPath, "rootPath", ".", "root folder of the project")
	flag.Parse()
	return rootPath
}

func loadConfig(rootPath string) (Config, error) {
	path := fmt.Sprintf("%s/config/default.yml", rootPath)

	loader := confita.NewLoader(file.NewBackend(path))
	ctx := context.Background()

	var cfg Config
	if err := loader.Load(ctx, &cfg); err != nil {
		return Config{}, fmt.Errorf("failed to load config from %s: %w", path, err)
	}

	return cfg, nil
}
