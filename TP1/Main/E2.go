package main

import "fmt"

var Temperatura float32
var Humedad int32
var Presion float32

func main() {
	Temperatura = 35.5
	Humedad = 80
	Presion = 101.325
	fmt.Println("La temperatura es de ", Temperatura, " grados")
	fmt.Println("La humedad es de ", Humedad, " %")
	fmt.Println("La presión es de ", Presion, " hectoPascales")
}
