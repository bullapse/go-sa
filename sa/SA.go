package sa

import (
	"github.com/bullapse/tspsa"
	"math/rand"
)

type SA struct {
	t int		// Initial Temperature
	c float64	// Cooling Rate
	b tspsa.Route   // best Route
	n tspsa.Route	// Next Current Route
}

/*
 * Create a new SA. (Not needed, but easier to use)
 * t: Initial Temperature
 * c: Cooling Rate
 */
func NewSA(t int, c float64) SA {
	return SA{t,c}
}

/*
 * Wrapper to call SA with random points
 * x: Number of x coordinates
 * y: Number of y coordinates
 * p: Number of Points
 */
func (sa *SA) RunSARandom(x int, y int, p int) {
	sa.RunSA(tspsa.CreateRandomMAP(x, y, p))
}

/* TODO
 * Run Simmulated Annealing Algorithm
 */
func (sa *SA) RunSAFromFile(s string) {

}

func (sa *SA) RunSA(n []Node) {
	// Set the Routes for the current SA given the list of nodes
	sa.b = n
	sa.c = n
	// get the length
	l := sa.n.Cities()
	// current
	cs := n
	for sa.t > 1 {
		// temp Route
		t := cs
		fp := rand.Intn(l)
		sp := rand.Intn(l)
		sa.n.SwapNodes(fp, sp)
		// get the current energy
		// get the next energy
	}


}



