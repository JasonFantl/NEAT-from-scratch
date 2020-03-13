package evolution

import (
	network "NEAT/lib/neat/network"
	"errors"
	math "math"
)

type Agent struct {
	Network network.Network
	Fitness float64
}

func NewAgent(inputSize, outputSize int) Agent {
	return Agent{network.NewNetwork(inputSize, outputSize), 0.0}
}

func BuildAgent(inNetwork network.Network) Agent {
	return Agent{inNetwork, 0}
}

func (a Agent) UpdateFitness(input, outputs []float64) {

	results := a.Network.FeedForward(input)
	vecDiff, _ := vectorSub(results, outputs)
	a.Fitness = vectorNormalSqr(vecDiff)
}

func vectorNormalSqr(vector []float64) float64 {
	sum := 0.0
	for _, val := range vector {
		sum += math.Pow(val, 2.0)
	}
	return sum
}

func vectorSub(v1, v2 []float64) ([]float64, error) {
	if len(v1) != len(v2) {
		return make([]float64, 0), errors.New("vectors not of same length")
	}

	v3 := make([]float64, len(v1))
	for i := range v1 {
		v3[i] = v1[i] - v2[i]
	}

	return v3, nil
}

func (a *Agent) Mutate() {
	a.Network.Mutate()
}
