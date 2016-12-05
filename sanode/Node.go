package sanode

import "math"

type Node struct {
	x int		// x cord
	y int 		// y cord
}

/*
 *  Calculate the Euclidean distance between the two nodes
 *  return the distance (float64)
 */
func (i *Node) getEuDistance(j *Node) float64 {
	return math.Sqrt(math.Pow(float64(i.x - j.x), 2) + math.Pow(float64(i.y - j.y), 2))
}

/*
 * String function for Node
 */
func (i *Node) String() string {
	return "NODE: {\n\tX: " +  string(i.x) + "\n\tY: " + string(i.y) + "\n}"
}

/*
 * Create a new blank node
 */
func NewBlankNode(x int, y int) Node {
	return Node{x,y}
}

