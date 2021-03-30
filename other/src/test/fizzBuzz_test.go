package test

import (
	"testing"
)

func fizzBuzzCount(start, end int) map[string]int {
	resultHash := map[string]int{}
	fizzCount, buzzCount, fizzBuzzCount := 0, 0, 0
	for ; start < end; start++ {
		if start%3 == 0 && start%5 == 0 {
			fizzBuzzCount++
		} else if start%3 == 0 {
			fizzCount++
		} else if start%5 == 0 {
			buzzCount++
		}
	}
	return resultHash
}

func TestFizzBuzzCount(t *testing.T) {
	start, end := 1, 10
	fizzResult := fizzBuzzCount(start, end)["fizz"] != 3
	buzzResult := fizzBuzzCount(start, end)["buzz"] != 2
	fizzBuzzResult := fizzBuzzCount(start, end)["fizzBuzz"] != 0
	asertMap := map[string]int{"fizz": 3, "buzz": 2, "fizzBuzz": 0}

	if fizzResult && buzzResult && fizzBuzzResult {
		t.Errorf("fizzBuzzCount() = %v, want %v", fizzBuzzCount(1, 10), asertMap)
	}

}
