package neat

import "math/rand"

type Species []Agent

type Population struct {
	Agents    []Species
	BestAgent Agent
}

func NewPopulation(popSize, networkInputSize, networkOutputSize int) Population {
	newSpecies := make([]Species, popSize)
	for _, val := range newSpecies {
		val = make([]Agent, 1)
		val[0] = NewAgent(networkInputSize, networkOutputSize)
	}
	newPopulation := Population{Agents: newSpecies}

	return newPopulation
}

func (p *Population) UpdateFitness(input, outputs []float64) {
	for _, species := range p.Agents {
		for _, agent := range species {
			agent.UpdateFitness(input, outputs)
		}
	}
}

func Crossover(n1, n2 Agent) Network {

	equalFitness := n1.Fitness == n2.Fitness
	mostFitNetwork := n1.Network
	lessFitNetwork := n2.Network
	if n2.Fitness > n1.Fitness {
		mostFitNetwork = n2.Network
		lessFitNetwork = n1.Network
	}

	mostFitNetwork = mostFitNetwork.Copy()
	//add all nodes
	for _, node := range lessFitNetwork.Nodes {
		mostFitNetwork.Nodes[node.ID] = node
	}

	for _, c := range lessFitNetwork.Connections {
		if _, exists := mostFitNetwork.Connections[c.getKey()]; exists || equalFitness {
			if rand.Intn(2) == 0 {
				mostFitNetwork.Connections[c.getKey()] = c
			}
		}
	}

	return mostFitNetwork
}

func Distance(a1, a2 Agent) float64 {

	g1 := a1.Network
	g2 := a2.Network

	C1 := 0.5
	C2 := 0.2

	g1Length := len(g1.Connections)
	g2Length := len(g2.Connections)

	numberOfGenes := 1
	// if g1Length > 20 || g2Length > 20 {
	if g1Length > g2Length {
		numberOfGenes = g1Length
	} else {
		numberOfGenes = g2Length
	}
	// }

	numberOfMatching := 0
	weightDifference := 0.0
	for key, val1 := range g1.Connections {
		if val2, exists := g2.Connections[key]; exists {
			numberOfMatching++
			weightDifference += Abs(val1.Weight - val2.Weight)
		}
	}
	weightDifference /= float64(numberOfMatching)

	numberNotMatching := g1Length - numberOfMatching + g2Length - numberOfMatching

	theta := C1*float64(numberNotMatching)/float64(numberOfGenes) + C2*weightDifference

	return theta

}
