package evolution

type Specie struct {
	Agents    []Agent
	BestAgent Agent
}

func NewSpecie(firstAgent Agent) Specie {
	agents := make([]Agent, 1)
	agents[0] = firstAgent
	newSpecie := Specie{Agents: agents, BestAgent: firstAgent}

	return newSpecie
}

func (s Specie) GetMascot() Agent {
	return s.Agents[0]
}

func (s *Specie) AddAgent(agent Agent) {
	s.Agents = append(s.Agents, agent)
}
