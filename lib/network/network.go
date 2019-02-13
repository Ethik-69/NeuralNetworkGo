package network

import (
	"math"
	"math/rand"
)

type Network struct {
	numInput    int
	numHidden   int
	numOutput   int
	layer       [][]float64
	arrayInput  []float64
	arrayHidden []float64
	ArrayOutput []float64

	weightsInput  [][]float64
	weightsOutput [][]float64

	changeInput  [][]float64
	changeOutput [][]float64
}

func (n *Network) sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

func (n *Network) dsigmoid(y float64) float64 {
	return 1.0 - math.Pow(y, 2.0)
}

func (n *Network) Update(sensor [2]interface{}) {
	for i := 0; i < n.numInput-1; i++ {
		n.arrayInput[i] = float64(sensor[i].(float32))
	}

	for j := 0; j < n.numHidden; j++ {
		sum := 0.0
		for i := 0; i < n.numInput; i++ {
			sum += n.arrayInput[i] * n.weightsInput[i][j]
		}
		n.arrayHidden[j] = n.sigmoid(sum)
	}

	for k := 0; k < n.numOutput; k++ {
		sum := 0.0
		for j := 0; j < n.numHidden; j++ {
			sum += n.arrayHidden[j] * n.weightsOutput[j][k]
		}
		n.ArrayOutput[k] = n.sigmoid(sum)
	}
}

func (n *Network) BackPropagation(targets []float64) float64 {
	N := 0.5
	M := 0.1

	outputDeltas := make([]float64, n.numOutput)
	// Calculate error terms for output
	for k := 0; k < n.numOutput; k++ {
		error := targets[k] - n.ArrayOutput[k]
		outputDeltas[k] = n.dsigmoid(n.ArrayOutput[k]) * error
	}

	hiddenDeltas := make([]float64, n.numHidden)
	// Calculate error terms for hidden
	for j := 0; j < n.numHidden; j++ {
		error := 0.0
		for k := 0; k < n.numOutput; k++ {
			error += outputDeltas[k] * n.weightsOutput[j][k]
			hiddenDeltas[j] = n.dsigmoid(n.arrayHidden[j]) * error
		}
	}

	// Update output weights
	for j := 0; j < n.numHidden; j++ {
		n.changeOutput[j] = make([]float64, n.numOutput)
		for k := 0; k < n.numOutput; k++ {
			change := outputDeltas[k] * n.arrayHidden[j]
			n.weightsOutput[j][k] = n.weightsOutput[j][k] + N*change + M*n.changeOutput[j][k]
			n.changeOutput[j][k] = change
		}
	}

	// Update input weights
	for i := 0; i < n.numInput; i++ {
		n.changeInput[i] = make([]float64, n.numHidden)
		for j := 0; j < n.numHidden; j++ {
			change := hiddenDeltas[j] * n.arrayInput[i]
			n.weightsInput[i][j] = n.weightsInput[i][j] + N*change + M*n.changeInput[i][j]
			n.changeInput[i][j] = change
		}
	}

	// Calculate error
	error := 0.0
	for k := 0; k < len(targets); k++ {
		error += 0.5 * math.Pow((targets[k]-n.ArrayOutput[k]), 2)
	}
	return error
}

func (n *Network) generateRandomWeights() {
	min := -100.0
	max := 100.0
	for i := 0; i < n.numInput; i++ {
		n.weightsInput[i] = make([]float64, n.numHidden)
		for j := 0; j < n.numHidden; j++ {
			n.weightsInput[i][j] = min + rand.Float64()*(max-(min))
		}
	}

	for j := 0; j < n.numHidden; j++ {
		n.weightsOutput[j] = make([]float64, n.numOutput)
		for k := 0; k < n.numOutput; k++ {
			n.weightsOutput[j][k] = min + rand.Float64()*(max-(min))
		}
	}

	// Add another layer for biais
	for index := range n.weightsInput[len(n.weightsInput)-1] {
		n.weightsInput[len(n.weightsInput)-1][index] = -1.0
	}
}

func (n *Network) generateHiddenLayer() {

}

func (n *Network) Init(numInput, numHidden, numOutput int) {
	n.numInput = numInput + 1
	n.numHidden = numHidden
	n.numOutput = numOutput

	n.arrayInput = make([]float64, n.numInput)
	n.arrayHidden = make([]float64, n.numHidden)
	n.ArrayOutput = make([]float64, n.numOutput)

	n.weightsInput = make([][]float64, n.numInput)
	n.weightsOutput = make([][]float64, n.numHidden)

	n.changeInput = make([][]float64, n.numInput)
	n.changeOutput = make([][]float64, n.numHidden)

	n.generateRandomWeights()
}
