package utils

import "unicode"

// Функция для обработки имени
func CapitalizeFirstLetter(s string) string {
	if s == "" || s == "-" {
		return s
	}

	// Преобразуем первый символ в верхний регистр
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])

	// Оставляем остальные символы без изменений
	return string(runes)
}
