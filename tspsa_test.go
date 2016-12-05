package tspsa

import "testing"

func TestRandom(t *testing.T) {
	t1 := NewSA(10000, 0.003)
	t1.RunSARandom(200, 200, 50)
}