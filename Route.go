package main

import "github.com/bullapse/tspsa/sa"

type Route struct {
	r []sa.Node			// Routes
	d float64 			// distance
}

/*
 * Create a new route from a given TSPMAP's Nodes
 */
func NewRoute(r []sa.Node) Route {
	return Route{r,0}
}

/*
 * Return the nubmer of cities in the route
 * return: (int)
 */
func (r *Route) cities() int {
	return len(r.r)
}

/*
 * Calculate the total route distance for the traveling salesman and return the float value
 * return: Route Distance (float64)
 */
func (r *Route) calcDistance() float64 {
	if r.d == 0 {
		for i := 0; i < r.cities(); i++ {
			s := r.r[i]
			var e sa.Node
			if i+1 < r.cities() {
				e = r.r[i+1]
			} else {
				e = r.r[0]
			}
			r.d += s.GetEuDistance(&e)
		}
	}
	return r.d
}

func (r *Route) String() string {
	t := ""
	for i := 0; i < r.cities(); i++ {
		t += r.r[i].String()
	}
	return t
}