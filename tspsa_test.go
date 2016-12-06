package tspsa

import "testing"

func TestRandom(t *testing.T) {
	RunSARandom(1000, 1, 0.00001, 0.9, 200, 200, 50)
}