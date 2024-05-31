package main

import "testing"

func SumTest(t *testing.T) {
	total := sum(10, 10)

	if total != 20 {
		t.Errorf("result incorrect")
	}
}
