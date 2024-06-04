package main

import (
	"fmt"
	"math/rand/v2"
)

func randomBinarySequence(n int) []int {
	result := make([]int, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, rand.IntN(2))
	}
	return result
}

// randomBinarySequenceDerandomized version of randomBinarySequence
// that returns the same result for the same input
// Derandomization was by optimizing the probability success
// for each bit. Bellow is the optimized version of the function.
func randomBinarySequenceDerandomized(n int) []int {
	return make([]int, n)
}

func main() {

	seq1 := randomBinarySequence(15)
	seq2 := randomBinarySequenceDerandomized(15)

	fmt.Println(seq1)
	fmt.Println(seq2)
}
