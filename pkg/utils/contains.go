package utils

import "strings"

func StringContains(arr []string, substr string) bool {
	for _, v := range arr {
		if v == substr {
			return true
		}
	}

	return false
}

func YesContain(name string) string {
	if strings.Contains(name, "yes") {
		return "Хаст"
	}
	return "Нест"
}
