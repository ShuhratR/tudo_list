package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/todo_list_shuhrat/config"
	"github.com/todo_list_shuhrat/internal/handlers"
	"github.com/todo_list_shuhrat/pkg/logger"
	"github.com/todo_list_shuhrat/pkg/shutdown"
)

func main() {
	// Загружаем конфигурацию
	cfg := config.LoadConfig()

	// Инициализируем логгер
	logger.InitLogger()

	// Создаем роутер
	r := mux.NewRouter()

	// Добавляем маршруты
	r.HandleFunc("/ping", handlers.Ping).Methods("GET")

	// Добавляем еще маршруты для авторизации и задач

	// Запускаем сервер
	server := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: r,
	}

	// Graceful Shutdown
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %s", err)
		}
	}()
	shutdown.GracefulShutdown(server)
}
