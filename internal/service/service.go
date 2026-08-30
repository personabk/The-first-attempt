package service

import (
	"errors"
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Convert определяет тип строки (текст или Морзе) и конвертирует её в противоположный вид.
func Convert(input string) (string, error) {
	// 1. Базовая проверка на пустоту
	if input == "" {
		return "", errors.New("входная строка пуста")
	}

	// 2. Определяем, что нам дали: текст или Морзе
	if isMorseString(input) {
		// Если это Морзе -> конвертируем в текст
		return morse.ToText(input), nil
	} else {
		// Если это текст -> конвертируем в Морзе
		return morse.ToMorse(input), nil
	}
}

// isMorseString возвращает true, если строка похожа на код Морзе.
// Логика простая: если в строке есть точки и тире, но почти нет букв — это Морзе.
func isMorseString(s string) bool {
	hasDotOrDash := strings.Contains(s, ".") || strings.Contains(s, "-")

	// Считаем количество букв в строке
	letterCount := 0
	for _, r := range s {
		if unicode.IsLetter(r) {
			letterCount++
		}
	}

	// Эвристика:
	// Если есть символы Морзе (точка/тире) И букв очень мало (или нет вообще),
	// то считаем, что это код Морзе.
	// Почему letterCount < 2? Чтобы фраза "Hi" (2 буквы) не считалась Морзе,
	// а вот ".-" (0 букв) считалась.
	return hasDotOrDash && letterCount < 2
}
