package main

import (
	"fmt"

	"github.com/GonzaGomezPizarro/ArquiSoft_GomezPizarro/archivos/ej3_apis/funciones/funciones"
)

func main() {
	texto := "Hola, soy una persona"
	key := "lkm7/sjd__kf174d"

	fmt.Println("\n -> El texto a encriptar es:", texto)
	fmt.Println(" -> La clave es:", key)

	textoEncryptado := funciones.encryptar(texto, key)

	fmt.Println("\n -> El texto encriptado es:", textoEncryptado)
	fmt.Println("\n -> El texto desencriptado es:", funciones.desencryptar(textoEncryptado.Result, key))

	var Largo string
	var simbolos string
	fmt.Println("\n -> Ingrese el largo de la clave:")
	fmt.Scanln(&Largo)
	fmt.Println(" -> Con simbolos? (s=1/n=0)")
	fmt.Scanln(&simbolos)

	fmt.Println(fmt.Sprintf("\n -> La clave generada es: %s\n", funciones.Keygen(Largo, simbolos)))

}
