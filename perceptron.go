package main

type perceptron struct {
	capas []capa
}

func (el perceptron) newPerceptron(capas int) perceptron {
	result := perceptron{
		capas: make([]capa, capas),
	}
	return result
}
