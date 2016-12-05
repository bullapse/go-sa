package main

import (
	"github.com/bullapse/tspsa/sanode"
	"math/rand"
)

type TSPMAP struct {
	m []sanode.Node
}

/*
 * Create a random map for the TSP problem
 * x: Max x cord
 * y: Max y cord
 * p: Number of points
 */
func (m *TSPMAP) createRandomMAP(x int, y int, p int) {
	for i := 0; i < p; i++ {
		append(m.m, sanode.NewBlankNode(rand.Intn(x), rand.Intn(y)))
	}
}

/* TODO
 * Create a map from a CSV file
 */
func (m *TSPMAP) createMAPFromFile(n string) {

}

func main() {

}

