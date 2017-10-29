package main

import (
	"fmt"
)

func main() {
	var intentos int
	var _neurona = newNeurona(2, 0.1)
	var sw = true

	esperado := [2][2]float32{[2]float32{0, 0}, [2]float32{0, 1}}

	for sw {
		sw = false
		intentos++
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				if !_neurona.aprender(esperado[i][j], float32(i), float32(j)) {
					sw = true
				}
			}
		}
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println("entrada ", i, j, _neurona.salida(float32(i), float32(j)))
		}
	}

	fmt.Println("pesos", _neurona.pesos)
	fmt.Println("umbral", _neurona.umbral)
	fmt.Println("intentos", intentos)
}
