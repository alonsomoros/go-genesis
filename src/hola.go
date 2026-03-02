package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hola mundo")

	// Tradicional
	var nombreAlon, apellidoAlon string = "Alon", "Klk"
	var edadAlon int = 21
	fmt.Println("Hello", nombreAlon, apellidoAlon, "with age", edadAlon)

	// Grupal
	var (
		nombreAlan   string = "Alan"
		apellidoAlan string = "Klk"
		edadAlan     int    = 21
	)
	fmt.Println("Hello", nombreAlan, apellidoAlan, "with age", edadAlan)

	// Dinámico
	var nombreAlen, apellidoAlen, edadAlen = "Alen", "Klk", 21
	fmt.Println("Hello", nombreAlen, apellidoAlen, "with age", edadAlen)

	// Dinámico v2
	nombreAlin, apellidoAlin, edadAlin := "Alin", "Klk", 21
	fmt.Println("Hello", nombreAlin, apellidoAlin, "with age", edadAlin)
}
