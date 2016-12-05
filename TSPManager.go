package main

import (
	"github.com/bullapse/tspsa/sa"
	"math/rand"
)

type TSPMAP struct {
	m []sa.Node
}

/*
 * Create a random map for the TSP problem
 * Modifies the map in the subsequent given TSPMAP struct
 * x: Max x cord
 * y: Max y cord
 * p: Number of points
 */
func (m *TSPMAP) createRandomMAP(x int, y int, p int) {
	for i := 0; i < p; i++ {
		m.m = append(m.m, sa.NewBlankNode(rand.Intn(x), rand.Intn(y)))
	}
}

/*
 * Calculate the number of nodes in our map
 * return: (int)
 */
func (m *TSPMAP) numberOfNodes() int {
	return len(m.m)
}

/*
 * Return a node inside TSPMAP's map of Nodes at index i
 * i: index
 * return: (*Node)
 */
func (m *TSPMAP) getNode(i int) *sa.Node {
	return &m.m[i]
}


/* TODO
 * Create a map from a CSV file
 */
func (m *TSPMAP) createMAPFromFile(n string) {

}

func main() {

}

