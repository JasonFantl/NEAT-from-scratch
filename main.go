package main

import (
	evolution "NEAT/lib/neat/evolution"
	neat "NEAT/lib/neat/network"
	"fmt"
)

var filepath string = "C:/Users/jason/Documents/Coding/Processing/Created sketches/displayNEAT/networks.txt"

func main() {
	generateCrazyMutated()
}

func generateBasicNet() {
	net := neat.NewNetwork(3, 2)

	net.AddConnection(1, 5, 0.5)
	net.AddConnection(5, 5, 0.5)
	net.SplitConnection(neat.Key{1, 5})

	inputs := []float64{10, 1, -100}
	neat.WriteToFile(filepath, net.GetFormated())
	fmt.Println(net.GetFormated())
	fmt.Printf("\nEval of net with input as %v. eval(net) = %v", inputs, net.FeedForward(inputs))
}

func generateCrazyMutated() {
	net := neat.NewNetwork(3, 2)

	for i := 0; i < 20; i++ {
		net.Mutate()
	}

	inputs := []float64{100, 5, -10}
	neat.WriteToFile(filepath, net.GetFormated())
	fmt.Println(net.GetFormated())
	fmt.Printf("\nEval of net with input as %v. eval(net) = %v", inputs, net.FeedForward(inputs))
}

func generatePapersExample() {
	net1 := neat.NewNetwork(3, 1)
	net1.AddConnection(1, 4, 1)
	net1.AddConnection(2, 4, 1)
	net1.AddConnection(3, 4, 1)
	net1.SplitConnection(neat.Key{2, 4})

	net2 := net1.Copy()

	net2.SplitConnection(neat.Key{5, 4})

	net1.AddConnection(1, 5, 1)
	net2.AddConnection(3, 5, 1)
	net2.AddConnection(1, 6, 1)

	for i := 0; i < 10; i++ {
		net1.MutateConnection()
	}
	net3 := evolution.Crossover(evolution.Agent{net1, 1}, evolution.Agent{net2, 1})

	neat.WriteToFile(filepath, net1.GetFormated())
	neat.AddToFile(filepath, net2.GetFormated())
	neat.AddToFile(filepath, net3.GetFormated())

}
