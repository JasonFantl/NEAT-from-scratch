package evolution

import (
	math "math"
	"math/rand"
)

type Population struct {
	Species   []Specie
	BestAgent Agent
}

func NewPopulation(popSize, networkInputSize, networkOutputSize int) Population {
	newAgents := make([]Agent, popSize)
	for i := range newAgents {
		newAgents[i] = NewAgent(networkInputSize, networkOutputSize)
		newAgents[i].Mutate()
	}

	newSpecies := Speciate(newAgents)

	newPopulation := Population{Species: newSpecies}

	return newPopulation
}

func (p *Population) UpdateFitness(input, outputs []float64) {
	for _, specie := range p.Species {
		for _, agent := range specie.Agents {
			agent.UpdateFitness(input, outputs)
		}
	}
}

func Crossover(n1, n2 Agent) Agent {

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

	// randomly add connections from less fit network
	for _, c := range lessFitNetwork.Connections {
		if _, exists := mostFitNetwork.Connections[c.GetKey()]; exists || equalFitness {
			if rand.Intn(2) == 0 {
				mostFitNetwork.Connections[c.GetKey()] = c
			}
		}
	}

	return Agent{mostFitNetwork, 0}
}

func Distance(a1, a2 Agent) float64 {

	g1 := a1.Network
	g2 := a2.Network

	C1 := 0.5
	C2 := 0.2

	g1Length := len(g1.Connections)
	g2Length := len(g2.Connections)

	numberOfGenes := 1

	if g1Length > g2Length {
		numberOfGenes = g1Length
	} else {
		numberOfGenes = g2Length
	}

	numberOfMatching := 0
	weightDifference := 0.0
	for key, val1 := range g1.Connections {
		if val2, exists := g2.Connections[key]; exists {
			numberOfMatching++
			weightDifference += math.Abs(val1.Weight - val2.Weight)
		}
	}
	weightDifference /= float64(numberOfMatching)

	numberNotMatching := g1Length - numberOfMatching + g2Length - numberOfMatching

	theta := C1*float64(numberNotMatching)/float64(numberOfGenes) + C2*weightDifference

	return theta

}

func Speciate(agents []Agent) []Specie {

	discriminationThresh := 0.5

	species := make([]Specie, 0)

	for _, agent := range agents {
		placed := false

		for _, specie := range species {
			if Distance(agent, specie.GetMascot()) < discriminationThresh {
				specie.AddAgent(agent)
				placed = true
				break
			}
		}

		if !placed {
			species = append(species, NewSpecie(agent))
		}
	}

	return species
}

func (p *Population) getNextGeneration() []Agent {

}

func (p *Population) Repopulate() {
	newAgents := p.getNextGeneration()
	p.Species = Speciate(newAgents)
}
