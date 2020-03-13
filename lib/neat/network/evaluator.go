package network

import (
	"fmt"
	"sort"
)

type evaluatedConnection struct {
	In      *FullyConnectedNode //
	Out     *FullyConnectedNode
	Weight  float64
	Enabled bool
	visted  bool
}
type FullyConnectedNode struct {
	ID             int
	State          nodeType
	ConnectionsIn  map[Key]*evaluatedConnection
	ConnectionsOut map[Key]*evaluatedConnection
	visted         bool
	value          float64
}

//Does not currently cyclic graphs, it ignores the cycles
//But we want it to handle cyclic grpahs and even connections of nodes into themselves
func (n Network) FeedForward(input []float64) []float64 {
	output := make([]float64, n.OutputSize)
	if len(input) != n.InputSize {
		panic(fmt.Sprintf("Invalid input size, wanted %v, but got %v", n.InputSize, len(input)))
	}

	connectedNodes := make(map[int]*FullyConnectedNode)

	index := 0
	for key, val := range n.Nodes {
		inputVal := 0.0
		if val.State == SENSOR {
			inputVal = input[index]
			index++
		}
		connectedNodes[key] = &FullyConnectedNode{
			val.ID, val.State,
			make(map[Key]*evaluatedConnection),
			make(map[Key]*evaluatedConnection),
			false,
			inputVal}
	}

	for key, val := range n.Connections {
		connectedConnection := evaluatedConnection{
			connectedNodes[val.In],
			connectedNodes[val.Out],
			val.Weight,
			val.Enabled,
			false}
		connectedNodes[val.In].ConnectionsOut[key] = &connectedConnection
		connectedNodes[val.Out].ConnectionsIn[key] = &connectedConnection
	}

	index = 0
	for _, val := range connectedNodes {
		if val.State == OUTPUT {
			output[index] = evalRecurseBack(val)
			index++
		}
	}
	return output
}

func evalRecurseBack(node *FullyConnectedNode) float64 {
	if node.visted || node.State == SENSOR {
		return node.value
	}

	node.visted = true

	evaluatedVal := 0.0
	node.value = evaluatedVal

	sortedConnections := make([]*evaluatedConnection, len(node.ConnectionsIn))
	i := 0
	for _, val := range node.ConnectionsIn {
		sortedConnections[i] = val
		i++
	}

	sort.Slice(sortedConnections, func(i, j int) bool {
		if sortedConnections[i].In.ID > sortedConnections[j].In.ID {
			return true
		}
		if sortedConnections[i].In.ID < sortedConnections[j].In.ID {
			return false
		}
		//sorts by Innovation of in connection, then by out innovation
		return sortedConnections[i].Out.ID > sortedConnections[j].Out.ID
	})

	for _, val := range sortedConnections {
		if val.Enabled && !val.visted {
			evaluatedVal += val.Weight * evalRecurseBack(val.In)
		}
	}

	node.value = evaluatedVal

	return evaluatedVal

}

func sigmoidActivation(x float64) float64 {
	return 1.0 / (1.0 + exp1(-x))
}

// Speed up over math.Exp by using less precision
// https://codingforspeed.com/using-faster-exponential-approximation/
func exp1(x float64) float64 {
	x = 1.0 + x/256.0
	x *= x
	x *= x
	x *= x
	x *= x
	x *= x
	x *= x
	x *= x
	x *= x
	return x
}
