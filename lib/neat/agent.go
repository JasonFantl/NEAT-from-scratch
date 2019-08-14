package neat

type Agent struct {
	Network Network
	Fitness float64
}

func NewAgent(inputSize, outputSize int) Agent {
	return Agent{NewNetwork(inputSize, outputSize), 0.0}
}

func (a Agent) UpdateFitness(input, outputs []float64) {

	a.Network.FeedForward(input)
}
