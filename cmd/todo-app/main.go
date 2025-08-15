// cmd/todo-app/main.go

package main

import (
    "TodoGolangAPI/internal/config"
    "log/slog"
)

const (
    envLocal = "local"
    envDev   = "dev"
    envProd  = "prod"
)

func setupLogger(env string) *slog.Logger {
    var log *slog.Logger

    switch env {
    case envLocal:
        log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
    case envDev:
        log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
    case envProd:
        log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
    }

    return log
}

func main() {
    cfg := config.MustLoad()

    log := setupLogger(cfg.Env)
    log = log.With(slog.String("env", cfg.Env)) // к каждому сообщению будет добавляться поле с информацией о текущем окружении

    log.Info("initializing server", slog.String("address", cfg.Address)) // Помимо сообщения выведем параметр с адресом
    log.Debug("logger debug mode enabled")

//section
  
