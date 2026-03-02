package main

import "fmt"

func triangulo() {

	var lado1, lado2 int

	fmt.Println("Ingrese el lado 1:")
	fmt.Scanln(&lado1)
	fmt.Printf("El lado 1 es: %d\n", lado1)

	fmt.Println("Ingrese el lado 2:")
	fmt.Scanln(&lado2)
	fmt.Printf("El lado 2 es: %d\n", lado2)

	fmt.Printf("La hipotenusa es: %.2f\n", hipotenusa(lado1, lado2))
}

func hipotenusa(lado1, lado2 int) float64 {
	return float64(lado1*lado1+lado2*lado2) / 2
}
