package main

import (
	"fmt"
	"math"
	"math/rand/v2"
)

type Vitter[T any] struct {
	Results []struct {
		Index uint64
		Val   T
	}

	sequence []uint64
	index    int
	counter  uint64
}

func (v *Vitter[T]) Update(val T) bool {
	if v.index >= len(v.sequence) {
		return false
	}
	v.counter++
	if v.sequence[v.index] == v.counter {
		v.index++
		v.Results = append(v.Results, struct {
			Index uint64
			Val   T
		}{Index: v.counter, Val: val})
		return true
	}
	return false
}

func NewVitter[T any](length int) *Vitter[T] {
	return &Vitter[T]{sequence: vitterSequence(length), Results: make([]struct {
		Index uint64
		Val   T
	}, 0, length)}
}

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
	v := NewVitter[int](50)
	for range int(1e4) {
		rng := rand.IntN(1e5)
		_ = v.Update(rng)
	}
	fmt.Println(v.Results)
}
