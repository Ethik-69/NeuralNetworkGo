package network

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Network struct {
	numInput    int
	numHidden   int
	numOutput   int
	layer       [][]float64
	ArrayOutput []float64

	weights [][]float64

	changeInput  [][]float64
	changeOutput [][]float64
}

func (n *Network) sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

func (n *Network) dsigmoid(y float64) float64 {
	return 1.0 - math.Pow(y, 2.0)
}

/*
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
*/

func (n *Network) generateRandomWeights(layerNumber int) {
	min := -100.0
	max := 100.0

	// For each layer
	for i := 0; i < layerNumber-1; i++ {
		fmt.Print(n.layer[i])
		// For each node in the current layer
		for j := 0; j < len(n.layer[i])-1; j++ {
			n.weights[j] = make([]float64, len(n.layer[i]))
			// For each node in the next layer
			for k := 0; k < len(n.layer[i+1]); k++ {
				n.weights[j][k] = min + rand.Float64()*(max-(min))
			}
		}
	}
	fmt.Print("Weights: ")
	fmt.Print(n.weights)

	// Add another layer for biais
	//for i := 0; i < len(n.weights[0][len(n.weights[0])-1]); i++ {
	//	n.weights[0][len(n.weights[0])-1][i] = -1.0
	//}
}

func (n *Network) generateHiddenLayer(layerNumber int) {
	NodeNumberMin := 3
	NodeNumberMax := 10
	for i := 1; i < layerNumber-1; i++ {
		n.layer[i] = make([]float64, rand.Intn(NodeNumberMax-NodeNumberMin)+NodeNumberMin)
	}
}

func (n *Network) Init(numInput, numHidden, numOutput int) {
	n.numInput = numInput + 1
	n.numHidden = numHidden
	n.numOutput = numOutput

	layerNumberMin := 3
	layerNumberMax := 13
	layerNumber := rand.Intn(layerNumberMax-layerNumberMin) + layerNumberMin
	n.layer = make([][]float64, layerNumber)

	n.layer[0] = make([]float64, n.numInput)
	n.layer[len(n.layer)-1] = make([]float64, n.numOutput)

	n.generateHiddenLayer(layerNumber)
	fmt.Print("Layer: ")
	//fmt.Print(n.layer)

	n.ArrayOutput = make([]float64, n.numOutput)

	n.changeInput = make([][]float64, n.numInput)
	n.changeOutput = make([][]float64, n.numHidden)

	n.weights = make([][]float64, layerNumber-1)
	n.generateRandomWeights(layerNumber)
	time.Sleep(2000 * time.Millisecond)
}
