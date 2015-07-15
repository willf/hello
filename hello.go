package main

import "fmt"

func main() {
	greetings := map[string]string{
		"Chinese":    "你好世界",
		"Dutch":      "Hello wereld",
		"English":    "Hello world",
		"French":     "Bonjour monde",
		"German":     "Hallo Welt",
		"Greek":      "γειά σου κόσμος",
		"Italian":    "Ciao mondo",
		"Japanese":   "こんにちは世界",
		"Korean":     "여보세요 세계",
		"Portuguese": "Olá mundo",
		"Russian":    "Здравствулте мир",
		"Spanish":    "Hola mundo"}
	for key, value := range greetings {
		fmt.Println("In", key, "they say:", value)
	}
}
