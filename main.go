package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type RESPONSEJSON struct {
	Result string `json:"result"`
}

func encryptar(texto, key string) RESPONSEJSON {

	response, err := http.Post("https://classify-web.herokuapp.com/api/encrypt", "application/json", strings.NewReader(
		`{
			"data":"`+texto+`",
			"key": "`+key+`"
		}`,
	))
	if err != nil {
		fmt.Println("Error al conectar con el servidor", err)
		return RESPONSEJSON{}
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error al leer la respuesta del servidor", err)
		return RESPONSEJSON{}
	}

	var codigo RESPONSEJSON
	err = json.Unmarshal(bytes, &codigo)
	if err != nil {
		fmt.Println("Error al deserializar la respuesta del servidor", err)
		return RESPONSEJSON{}
	}

	return codigo
}
func desencryptar(codigo, key string) RESPONSEJSON {

	response, err := http.Post("https://classify-web.herokuapp.com/api/decrypt", "application/json", strings.NewReader(
		`{
			"data":"`+codigo+`",
			"key": "`+key+`"
		}`,
	))
	if err != nil {
		fmt.Println("Error al conectar con el servidor", err)
		return RESPONSEJSON{}
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error al leer la respuesta del servidor", err)
		return RESPONSEJSON{}
	}

	var texto RESPONSEJSON
	err = json.Unmarshal(bytes, &texto)
	if err != nil {
		fmt.Println("Error al deserializar la respuesta del servidor", err)
		return RESPONSEJSON{}
	}

	return texto
}

func main() {
	texto := "Hola, soy una persona"
	key := "lkm7/sjd__kf174d"

	fmt.Println("\n -> El texto a encriptar es:", texto)
	fmt.Println(" -> La clave es:", key)

	textoEncryptado := encryptar(texto, key)

	fmt.Println("\n -> El texto encriptado es:", textoEncryptado)
	fmt.Println("\n -> El texto desencriptado es:", desencryptar(textoEncryptado.Result, key))

}
