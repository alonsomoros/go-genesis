package main

import (
	"fmt"
	"math"
)

func triangulo() {

	var lado1, lado2 float64
	var hipotenusa, perimetro float64

	fmt.Println("Ingrese el lado 1:")
	fmt.Scanln(&lado1)
	fmt.Printf("El lado 1 es: %.2f\n", lado1)

	fmt.Println("Ingrese el lado 2:")
	fmt.Scanln(&lado2)
	fmt.Printf("El lado 2 es: %.2f\n", lado2)

	hipotenusa = calcular_hipotenusa(lado1, lado2)
	fmt.Printf("La hipotenusa es: %.2f\n", hipotenusa)

	fmt.Printf("El área del triángulo es: %.2f\n", calcular_area(lado1, lado2))

	perimetro = hipotenusa + float64(lado1) + float64(lado2)
	fmt.Printf("El perímetro del triángulo es: %.2f\n", perimetro)

}

func calcular_hipotenusa(lado1, lado2 float64) float64 {
	return math.Hypot(lado1, lado2)
}

func calcular_area(lado1, lado2 float64) float64 {
	return float64(lado1*lado2) / 2
}
