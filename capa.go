package main

type capa struct {
	neuronas []neurona
}

func (el capa) newCapa(neuronas int, neuronaCapa []int) capa {
	result := capa{
		neuronas: make([]neurona, neuronas),
	}
	for i := 0; i < neuronas; i++ {
		result.neuronas[i] = newNeurona(neuronaCapa[i], 0.1)
	}
	return result
}
