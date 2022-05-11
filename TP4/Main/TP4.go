package main

import (
	"fmt"
)

func promedio(copia []int) float32 {
	var suma float32
	suma = 0.0
	for i := 0; i < len(copia); i++ {
		suma = suma + float32(copia[i])
	}
	return suma / float32(len(copia))
}

func main() {
	var slice1 []int
	fmt.Println("Ejercicio 1 TP Clase 4")
	for i := 0; i < 10; i++ {
		slice1 = append(slice1, i+1)
	}
	fmt.Println("El slice1 contiene: ", slice1)
	fmt.Println("Los elementos impares son: ")
	for i := 0; i < len(slice1); i++ {
		if slice1[i]%2 == 1 {
			fmt.Print(slice1[i], " ")
		}
	}
	fmt.Println("\n ")
	slice2 := slice1[3:7]
	fmt.Println("El slice2 contiene: ", slice2)
	copia := make([]int, len(slice1))
	copy(copia, slice1)
	fmt.Println("El promedio del slice1 es: ", promedio(copia))
}
