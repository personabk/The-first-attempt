package server

import (
	"log"
	"net/http"
	"time" // Обязательно для таймаутов!

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// Server инкапсулирует http.Server и логгер.
type Server struct {
	Logger *log.Logger
	Server *http.Server
}

// тест
// NewServer создает и настраивает новый экземпляр сервера.
func NewServer(logger *log.Logger) *Server {
	// 1. Создаем http-роутер
	mux := http.NewServeMux()

	// 2. Регистрируем хендлеры
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	// 3. Создаем экземпляр http.Server с настройками из ТЗ
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger: logger,
		Server: srv,
	}
}
