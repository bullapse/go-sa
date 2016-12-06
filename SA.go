package tspsa

import (
	"math/rand"
	"math"
	"fmt"
)

/*
 * Wrapper to call SA with random points
 * it: Itterations
 * t: temp
 * tmin: Min Temp
 * c: coolrate
 * x: Number of x coordinates
 * y: Number of y coordinates
 * p: Number of Points
 */
func RunSARandom(it int, t float64, tmin float64, c float64, x int, y int, p int) {
	RunSA(it, t, tmin, c, CreateRandomMAP(x, y, p))
}

/* TODO
 * Run Simmulated Annealing Algorithm
 */
func RunSAFromFile(s string) {

}

/*
 * Run the SA algorithm with the given itterations, temp, min temp, coolrate, and a splice of Node objects
 *
 * i:    Itterations
 * t: 	 temp
 * tmin: temp min
 * c: 	 coolrate
 */
func RunSA(it int, t float64, tmin float64, c float64, n []Node) Route {
	// Set the Routes for the current SA given the list of nodes
	cur  := NewRoute(n)
	cur.CalcDistance() // old cost
	best := cur
	next := cur
	l := cur.Nodes() // Length
	for t > tmin {
		for i := 0; i < it; i++ {
			next.SwapNodes(rand.Intn(l), rand.Intn(l))
			next.CalcDistance()
			if acceptProb(cur.D, next.D, t) > float64(rand.Intn(l)) {
				cur = next
			}
			if (cur.D < best.D) {
				best = cur
			}
		}
		t = t * c
	}
	fmt.Println(best.String())
	fmt.Printf("Distance: %.6f\n", best.D)
	return best
}

/*
 * acceptance Probability Calculation
 * This should return 1.0 if the new energy is greater than the current energy
 * Else it will return the acceptance Probability
 */
func acceptProb(e float64, ne float64, t float64) float64 {
	if ne > e {
		return 1
	} else {
		// acceptance probability
		return math.Exp(-1 * (e - ne) / t)
	}
}





