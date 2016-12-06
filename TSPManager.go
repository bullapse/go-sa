/*
 *http://www.theprojectspot.com/tutorial-post/simulated-annealing-algorithm-for-beginners/6
 */
package tspsa

import (
	"math/rand"
	"fmt"
)

/*
 * Create a random map for the TSP problem
 * Modifies the map in the subsequent given TSPMAP struct
 * x: Max x cord
 * y: Max y cord
 * p: Number of points
 */
func CreateRandomMAP(x int, y int, p int) []Node{
	var r []Node
	for i := 0; i < p; i++ {
		r = append(r, NewBlankNode(rand.Intn(x), rand.Intn(y)))
		fmt.Printf("X: %i, Y: %")
	}
	return r
}