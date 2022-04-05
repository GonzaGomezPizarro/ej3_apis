package main

import (
	"fmt"
	"net/http"
)

func encryptar(texto, url string) (string, error) {

	_, err := http.Get(url)
	if err != nil {
		fmt.Println("Error al conectar con el servidor", err)
		return "", err
	}
	return texto, nil
}

func main() {
	url := "https://github.com/cheatsnake/classify"
	texto := "Hola, soy una persona"
	encryptar(texto, url)
}
