package network

import (
	"fmt"
)

type Key struct {
	In, Out int
}

type Network struct {
	Nodes                 map[int]Node
	Connections           map[Key]Connection
	NodeInnovation        *InnovationGenerator
	ConnectionInnovation  *InnovationGenerator
	InputSize, OutputSize int
}

func NewNetwork(inSize, outSize int) Network {
	g := Network{
		Nodes:                make(map[int]Node),
		Connections:          make(map[Key]Connection),
		NodeInnovation:       NewInnovationGenerator(),
		ConnectionInnovation: NewInnovationGenerator(),
		InputSize:            inSize,
		OutputSize:           outSize}
	for i := 0; i < inSize; i++ {
		g.AddNode(SENSOR)
	}
	for i := 0; i < outSize; i++ {
		g.AddNode(OUTPUT)
	}
	return g
}

func (g *Network) AddConnection(inNode int, outNode int, weight float64) Connection {
	//Check that the connection can exist
	if _, exist := g.Connections[Key{inNode, outNode}]; g.Nodes[outNode].State == SENSOR ||
		exist {
		return Connection{}
	}

	newConnection := NewConnection(g.Nodes[inNode].ID, g.Nodes[outNode].ID, weight, g.ConnectionInnovation.get())
	g.Connections[newConnection.GetKey()] = newConnection

	return newConnection
}

func (g *Network) AddNode(t nodeType) Node {
	newNode := NewNode(g.NodeInnovation.get(), t)
	g.Nodes[newNode.ID] = newNode

	return newNode
}

func (g Network) GetFormated() string {
	formatedString := ""

	for key, value := range g.Nodes {
		formatedString += fmt.Sprintf("%d,", key)
		formatedString += fmt.Sprintf("%d|", value.State-1)
	}

	formatedString = formatedString[:len(formatedString)-1] + ":"

	for _, value := range g.Connections {
		isEnabled := 0
		if value.Enabled {
			isEnabled = 1
		}
		formatedString += fmt.Sprintf("%d,%d,%v,%v,%d|", value.In, value.Out, isEnabled, value.Weight, value.Innovation)
	}

	formatedString = formatedString[:len(formatedString)-1]

	return formatedString
}

func (g Network) ToString() string {
	s := ""
	for _, node := range g.Nodes {
		s += fmt.Sprintf("node: ID %d, type %d\n", node.ID, node.State)
	}
	for _, conn := range g.Connections {
		s += fmt.Sprintf("connection: %d to %d, enabled=%t, innovation %d\n", conn.In, conn.Out, conn.Enabled, conn.Innovation)
	}
	return s
}

func (g Network) Copy() Network {
	copiedNodes := make(map[int]Node)
	copiedConnections := make(map[Key]Connection)
	for key, val := range g.Nodes {
		copiedNodes[key] = val
	}
	for key, val := range g.Connections {
		copiedConnections[key] = val
	}

	return Network{
		Nodes:                copiedNodes,
		Connections:          copiedConnections,
		NodeInnovation:       g.NodeInnovation,
		ConnectionInnovation: g.ConnectionInnovation,
		InputSize:            g.InputSize,
		OutputSize:           g.OutputSize}
}
