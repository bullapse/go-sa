package tspsa


type Route struct {
	R []Node			// Routes
	D float64 			// distance
}

func (r *Route) SwapNodes(n1 int, n2 int) {
	t1 := r.R[n1]
	t2 := r.R[n2]
	r.R[n1] = t2
	r.R[n2] = t1
}

/*
 * Create a new route from a given TSPMAP's Nodes
 */
func NewRoute(r []Node) Route {
	return Route{r,0}
}

/*
 * Return the nubmer of cities in the route
 * return: (int)
 */
func (r *Route) Cities() int {
	return len(r.R)
}

/*
 * Calculate the total route distance for the traveling salesman and return the float value
 * return: Route Distance (float64)
 */
func (r *Route) CalcDistance() float64 {
	if r.D == 0 {
		for i := 0; i < r.Cities(); i++ {
			s := r.R[i]
			var e Node
			if i+1 < r.Cities() {
				e = r.R[i+1]
			} else {
				e = r.R[0]
			}
			r.D += s.GetEuDistance(&e)
		}
	}
	return r.D
}

func (r *Route) String() string {
	t := ""
	for i := 0; i < r.Cities(); i++ {
		t += r.R[i].String()
	}
	return t
}