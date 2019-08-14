package neat

type InnovationGenerator struct {
	value int
}

func (inno *InnovationGenerator) get() int {
	inno.value++
	return inno.value - 1
}

func NewInnovationGenerator() *InnovationGenerator {
	return &InnovationGenerator{1}
}
