package main

import (
	"fmt"
)

// const spanish = "Spanish"
// const french = "French"
// const frenchHelloPrefix = "Bonjour, "
// const spanishHelloPrefix = "Hola, "
// const englishHelloPrefix = "Hello, "

const separator = ", "

// I'd do a mapping (key, value) of (language, helloPrefix)
var helloLanguageMap = map[string]string{
	"French":  "Bonjour",
	"Spanish": "Hola",
	"English": "Hello",
	"":        "Hello",
}

// I'd pass the separator
func Hello(name, language string) string {

	if name == "" {
		name = "World"
	}
	return helloLanguageMap[language] + separator + name
	// return greetingPrefix(language) + name
}

// func greetingPrefix(language string) (prefix string) {
// 	switch language {
// 	case french:
// 		prefix = frenchHelloPrefix
// 	case spanish:
// 		prefix = spanishHelloPrefix
// 	default:
// 		prefix = englishHelloPrefix
// 	}
// 	return
// }

func main() {
	fmt.Println(Hello("world", ""))
}
