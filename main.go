package main

import (
	"fmt"
	"net/http"
)

func encryptar(texto, url, key string) (string, error) {

	_, err := http.Get(url)
	if err != nil {
		fmt.Println("Error al conectar con el servidor", err)
		return "", err
	}
	return texto, nil
}
func desencryptar(codigo, url, key string) (string, error) {

	_, err := http.Get(url)
	if err != nil {
		fmt.Println("Error al conectar con el servidor", err)
		return "", err
	}
	return codigo, nil
}

func main() {
	url := "https://github.com/cheatsnake/classify"
	texto := "Hola, soy una persona"
	key := "lkm7/sjd__kf174d"

	encryptar(texto, url, key)
}
