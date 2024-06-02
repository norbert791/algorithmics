package main

import (
	"fmt"
	"math"
	"math/rand/v2"
)

func vitterSequence(length int) []uint64 {
	if length <= 0 {
		return nil
	}
	sequence := make([]uint64, length)
	sequence[0] = 1
	for i := 1; i < length; i++ {
		prev := sequence[i-1]
		rng := rand.Float64()
		sequence[i] = prev + uint64(math.Ceil(rng*float64(prev)/(1-rng)))
	}
	return sequence

}

func main() {

	sequence := vitterSequence(50)
	fmt.Println(sequence)
}
