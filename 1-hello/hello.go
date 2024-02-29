package main

import (
	"fmt"
)

const languagePortuguese = "portuguese"
const languageSpanish = "spanish"

const prefixPortuguese = "Olá, "
const prefixSpanish = "Hola, "
const prefixEnglish = "Hello, "

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case languagePortuguese:
		prefix = prefixPortuguese
	case languageSpanish:
		prefix = prefixSpanish
	default:
		prefix = prefixEnglish
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
