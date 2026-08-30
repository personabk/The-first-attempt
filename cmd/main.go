package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final-tpl/internal/server"
)

func main() {
	// 1. Создаем логгер.
	// log.New создает логгер, который пишет в os.Stderr.
	// "INFO:" - префикс сообщения.
	// log.LstdFlags - стандартные флаги (дата, время, файл, строка).
	logger := log.New(os.Stderr, "INFO:", log.LstdFlags)

	// 2. Создаем сервер, передавая ему наш логгер.
	srv := server.NewServer(logger)

	// 3. Запускаем сервер.
	// ListenAndServe блокирует выполнение программы, пока сервер работает.
	// Если возникает ошибка (например, порт 8080 уже занят), функция вернет ошибку.
	logger.Println("Server started on :8080")

	if err := srv.Server.ListenAndServe(); err != nil {
		// 4. Если ошибка при запуске — логируем её как FATAL и завершаем программу.
		// log.Fatal делает то же самое, что log.Println + os.Exit(1)
		logger.Fatalf("Failed to start server: %v", err)
	}
}
