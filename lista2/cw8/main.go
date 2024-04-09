package main

import (
	"fmt"
	"runtime"
	"sync"

	"cloudeng.io/algo/lcs"
)

func computeLCS(s1, s2 []byte) int {
	alg := lcs.NewDP(s1, s2)
	return len(alg.LCS())
}

// generateBitSlice generates a slice of bytes representing the binary little endian representation of a number.
func generateBitSlice(num int, length int) []byte {
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = byte(num & 1)
		num >>= 1
	}
	return result
}

func generateLCSPairs(length int) <-chan [2][]byte {
	ch := make(chan [2][]byte, 100)
	maxNum := (1 << length) - 1

	go func() {
		defer close(ch)
		for i := 0; i <= maxNum; i++ {
			for j := 0; j <= maxNum; j++ {
				s1 := generateBitSlice(i, length)
				s2 := generateBitSlice(j, length)
				ch <- [2][]byte{s1, s2}
			}
		}
	}()
	return ch
}

func lcsEXP(length int, numOfWorkers int) float64 {
	workerResults := make([]int, numOfWorkers)
	pairs := generateLCSPairs(length)
	var wg sync.WaitGroup

	for i := range numOfWorkers {
		id := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			for pair := range pairs {
				workerResults[id] += computeLCS(pair[0], pair[1])
			}
		}()
	}

	wg.Wait()

	result := 0
	for _, r := range workerResults {
		result += r
	}

	numOfPairs := 1 << (length * 2)

	return float64(result) / float64(numOfPairs)
}

func main() {
	fmt.Println(lcsEXP(5, runtime.NumCPU()))
}
