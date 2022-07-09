package utils

import "strings"

func CamelCase(text string) string {

	words := strings.Split(text, " ")
	key := strings.ToLower(words[0])
	for _, word := range words[1:] {
		key += strings.Title(word)
	}

	return key
}
