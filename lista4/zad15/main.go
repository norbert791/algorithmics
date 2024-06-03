package main

import (
	"fmt"
	"math"
	"math/rand/v2"
)

func randomBinarySequence(n int) []int {
	result := make([]int, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, rand.IntN(2))
	}
	return result
}

func randomBinarySequenceDerandomized(n int) []int {
	resultProb := 1.0 / float64(n)
	result := make([]int, n)
	for k := range n {
		prob := 1.0 / math.Pow(2, float64(n-k-1))
		if prob < resultProb {
			result[k] = 0
		} else {
			result[k] = 1
		}
	}
	return result
}

func main() {

	seq1 := randomBinarySequence(15)
	seq2 := randomBinarySequenceDerandomized(15)

	fmt.Println(seq1)
	fmt.Println(seq2)
}
