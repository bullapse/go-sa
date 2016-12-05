package tspsa

import (
	"math/rand"
	"math"
	"fmt"
)

type SA struct {
	i int 				// Itterations
	maxt float64			// Max Temp
	mint float64			// Min Temp
	t float64			// Initial Temperature
	c float64			// Cooling Rate
}

/*
 * Create a new SA. (Not needed, but easier to use)
 * t: Initial Temperature
 * c: Cooling Rate
 */
func NewSA(t float64, c float64) SA {
	return SA{t,c}
}

/*
 * Wrapper to call SA with random points
 * x: Number of x coordinates
 * y: Number of y coordinates
 * p: Number of Points
 */
func (sa *SA) RunSARandom(x int, y int, p int) {
	sa.RunSA(CreateRandomMAP(x, y, p))
}

/* TODO
 * Run Simmulated Annealing Algorithm
 */
func (sa *SA) RunSAFromFile(s string) {

}

func (sa *SA) RunSA(n []Node) []Node {
	// Set the Routes for the current SA given the list of nodes
	best := NewRoute(n)
	cur := NewRoute(n)
	next := NewRoute(n)
	// get the length
	l := cur.Cities()
	for sa.t > 1 {
		fp := rand.Intn(l)
		sp := rand.Intn(l)
		next.SwapNodes(fp, sp)
		// get the current energy
		ce := cur.CalcDistance()
		// get the next energy
		ne := next.CalcDistance()
		// Should we accept next? TODO not sure if this is correct. Look at other algorithms
		if acceptProb(ce, ne, sa.t) > float64(rand.Intn(l)) {
			cur = next
		}
		if (cur.D < best.D) {
			best = cur
		}
		sa.t *= 1-sa.c
	}
	fmt.Println(best.String())
	fmt.Printf("Distance: %.6f\n", best.D)
	return best.R
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





