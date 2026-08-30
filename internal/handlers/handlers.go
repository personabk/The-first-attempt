package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// IndexHandler отдаёт статический файл index.html (страницу с формой).
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("index.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Не удалось найти index.html: %v", err), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка чтения файла: %v", err), http.StatusInternalServerError)
	}
}

// UploadHandler обрабатывает загрузку файла, конвертирует данные и показывает результат.
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается. Используйте POST.", http.StatusMethodNotAllowed)
		return
	}

	// Лимит загрузки 32 МБ
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, fmt.Sprintf("Ошибка парсинга формы: %v", err), http.StatusBadRequest)
		return
	}

	// Получаем файл из формы. Поле в HTML должно называться "file"
	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, fmt.Sprintf("Файл не найден в форме: %v", err), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Читаем содержимое файла
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка чтения файла: %v", err), http.StatusInternalServerError)
		return
	}

	// Очищаем строку от пробелов по краям
	inputText := strings.TrimSpace(string(data))

	if inputText == "" {
		http.Error(w, "Файл пуст", http.StatusBadRequest)
		return
	}

	// Вызываем конвертер из пакета service
	result, err := service.Convert(inputText)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка конвертации: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, result)

}

// truncate обрезает строку до указанной длины, чтобы красиво показать превью.
func truncate(s string, limit int) string {
	if len(s) <= limit {
		return s
	}
	return s[:limit] + "..."
}
