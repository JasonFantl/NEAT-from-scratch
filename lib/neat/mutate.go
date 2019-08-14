package neat

import (
	"math/rand"
)

func (n *Network) Mutate() {
	switch rnd := rand.Intn(4); rnd {
	case 0:
		n.MutateConnection()
	case 1:
		n.MutateAddNode()
	default:
		n.MutateAddConnection()
	}
}

func (g *Network) SplitConnection(connectionKey Key) {

	newNode := g.AddNode(HIDDEN)

	g.AddConnection(g.Connections[connectionKey].In, newNode.ID, 1.0)
	g.AddConnection(newNode.ID, g.Connections[connectionKey].Out, g.Connections[connectionKey].Weight)

	disabledConnection := g.Connections[connectionKey]
	disabledConnection.Disable()
	g.Connections[connectionKey] = disabledConnection

}

//MutateAddConnection - randomly create and add a new connection to connections array
func (g *Network) MutateAddConnection() {

	node1, _ := randomNodeKey(g.Nodes)
	node2, _ := randomNodeKey(g.Nodes)

	weight := rand.Float64()*2 - 1

	g.AddConnection(node1, node2, weight)
}

func (g *Network) MutateAddNode() {

	connectionKey, exists := randomConnectionKey(g.Connections)

	if !exists {
		return
	}

	g.SplitConnection(connectionKey)

}

func (g *Network) MutateConnection() {

	connectionKey, exists := randomConnectionKey(g.Connections)

	if !exists {
		return
	}

	newWeightConnection := g.Connections[connectionKey]
	newWeightConnection.Weight = rand.Float64()*2 - 1
	g.Connections[connectionKey] = newWeightConnection
}

func randomNodeKey(m map[int]Node) (key int, exists bool) {
	if len(m) == 0 {
		return 0, false
	}
	r := rand.Intn(len(m))
	for k := range m {
		if r == 0 {
			return k, true
		}
		r--
	}
	return 0, false
}

func randomConnectionKey(m map[Key]Connection) (key Key, exists bool) {
	if len(m) == 0 {
		return Key{}, false
	}
	r := rand.Intn(len(m))
	for k := range m {
		if r == 0 {
			return k, true
		}
		r--
	}
	return Key{}, false
}
