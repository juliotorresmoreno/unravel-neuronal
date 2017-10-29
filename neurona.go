package main

import (
	"math/rand"
	"time"
)

type neurona struct {
	pesos  []float32
	umbral float32
	sesgo  float32
}

func (el neurona) salida(entrada ...float32) float32 {
	return el.sigmoide(el.neurona(entrada...))
}

func (el neurona) neurona(entrada ...float32) float32 {
	var resultado = el.umbral
	for i, v := range entrada {
		resultado += v * el.pesos[i]
	}
	return resultado
}

func (el neurona) sigmoide(d float32) float32 {
	if d > 0 {
		return 1
	}
	return 0
}

func newNeurona(length int, sesgo float32) neurona {
	var ramdom = rand.New(rand.NewSource(time.Now().UnixNano()))
	var _neurona = neurona{}
	_neurona.pesos = make([]float32, length)
	for index := 0; index < length; index++ {
		_neurona.pesos[index] = ramdom.Float32() - ramdom.Float32()
	}
	_neurona.umbral = ramdom.Float32() - ramdom.Float32()
	_neurona.sesgo = sesgo
	return _neurona
}

func (el *neurona) aprender(esperado float32, entrada ...float32) bool {
	result := el.salida(entrada...)
	if result != esperado {
		err := el.sesgo * (esperado - result)
		pesos := make([]float32, len(el.pesos))
		for index := 0; index < len(entrada); index++ {
			pesos[index] = el.pesos[index] + err*entrada[index]
		}
		el.pesos = pesos
		el.umbral = el.umbral + err
		return false
	}
	return true
}
