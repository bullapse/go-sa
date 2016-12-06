package tspsa

import "testing"

func TestRandom(t *testing.T) {
	RunSARandom(100, 1, 0.00001, 0.4, 200, 200, 50, true)
}