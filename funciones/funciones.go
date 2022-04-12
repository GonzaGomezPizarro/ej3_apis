package funciones

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type ResponseJson struct {
	Result string `json:"result"`
}
type KeyJson struct {
	Key string `json:"key"`
}

func encryptar(texto, key string) ResponseJson {

	response, err := http.Post("https://classify-web.herokuapp.com/api/encrypt", "application/json", strings.NewReader(
		`{
			"data":"`+texto+`",
			"key": "`+key+`"
		}`,
	))
	if err != nil {
		fmt.Println("Error al conectar con el servidor", err)
		return ResponseJson{}
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error al leer la respuesta del servidor", err)
		return ResponseJson{}
	}

	var codigo ResponseJson
	err = json.Unmarshal(bytes, &codigo)
	if err != nil {
		fmt.Println("Error al deserializar la respuesta del servidor", err)
		return ResponseJson{}
	}

	return codigo
}
func desencryptar(codigo, key string) ResponseJson {

	response, err := http.Post("https://classify-web.herokuapp.com/api/decrypt", "application/json", strings.NewReader(
		`{
			"data":"`+codigo+`",
			"key": "`+key+`"
		}`,
	))
	if err != nil {
		fmt.Println("Error al conectar con el servidor", err)
		return ResponseJson{}
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error al leer la respuesta del servidor", err)
		return ResponseJson{}
	}

	var texto ResponseJson
	err = json.Unmarshal(bytes, &texto)
	if err != nil {
		fmt.Println("Error al deserializar la respuesta del servidor", err)
		return ResponseJson{}
	}

	return texto
}

func Keygen(Largo, simbolos string) KeyJson {
	response, err := http.Get("https://classify-web.herokuapp.com/api/keygen?length=" + Largo + "&symbols=" + simbolos)
	if err != nil {
		fmt.Println("Error al conectar con el servidor", err)
		return KeyJson{}
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error al leer la respuesta del servidor", err)
		return KeyJson{}
	}
	var key KeyJson
	err = json.Unmarshal(bytes, &key)
	if err != nil {
		fmt.Println("Error al deserializar la respuesta del servidor", err)
		return KeyJson{}
	}
	return key

}
