package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/FelipePn10/URLs-Shortener-API/api"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Failed to execute command", "error", err)
		return
	}
	slog.Info("All systems offline")
}

func run() error {
	db := make(map[string]string)

	handler := api.NewHandler(db)

	s := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  120 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      handler,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
