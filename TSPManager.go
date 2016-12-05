package tspsa

import (
	"github.com/bullapse/tspsa/sa"
	"math/rand"
)

/*
 * Create a random map for the TSP problem
 * Modifies the map in the subsequent given TSPMAP struct
 * x: Max x cord
 * y: Max y cord
 * p: Number of points
 */
func CreateRandomMAP(x int, y int, p int) []sa.Node{
	var r []sa.Node
	for i := 0; i < p; i++ {
		r = append(r, sa.NewBlankNode(rand.Intn(x), rand.Intn(y)))
	}
	return r
}
