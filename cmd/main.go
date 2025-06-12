package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"task-runner-api/internal/config"
	handler "task-runner-api/internal/http"
	"task-runner-api/internal/task"
	"task-runner-api/pkg/logger"

	"github.com/gorilla/mux"
)

func main() {
	logger.ZapLoggerInit()
	ctx := context.Background()
	cfg := config.MustInit()

	taskManager := task.NewManager()
	h := handler.NewHandler(taskManager)

	r := mux.NewRouter()
	h.RegisterRoutes(r)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: r,
	}

	go func() {
		logger.Info(fmt.Sprintf("Starting HTTP server on :%d", cfg.Port))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal(fmt.Sprintf("Failed to start server: %v", err))
		}
	}()

	awaitStop(ctx, srv)
}

func awaitStop(ctx context.Context, srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	sig := <-quit
	logger.Info(fmt.Sprintf("Shutting down server... signal: %v", sig))

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error(fmt.Sprintf("HTTP server Shutdown: %v", err))
	}
}
