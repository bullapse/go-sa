package sa

import (
	"github.com/bullapse/tspsa"
	"math/rand"
)

type SA struct {
	t int		// Initial Temperature
	c float64	// Cooling Rate
	r tspsa.Route
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
	l := sa.r.Cities()
	// best route so far
	var b []Node
	// current
	cs := n
	for sa.t > 1 {
		// temp Route
		t := cs
		fp := rand.Intn(l)
		sp := rand.Intn(l)
		sa.r.SwapNodes(fp, sp)

	}


}



