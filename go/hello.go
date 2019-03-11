package main

import "fmt"

const englishPrefix = "Hello, "
const chinese = "chinese" 
const chinesePrefix = "你好, "
const spanish = "spanish"
const spanishPrefix = "Hola, "

func Hello(name,language string) string {
	if name == "" {
        name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case chinese:
		prefix = chinesePrefix
	case spanish:
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", "english"))
}
