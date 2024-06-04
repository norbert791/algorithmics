package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strings"

	"github.com/twmb/murmur3"
)

func main() {
	b, err := os.ReadFile("book.txt")
	if err != nil {
		panic(err)
	}
	text := string(b)

	seeds := []uint64{10, 20, 100, 1000, 10000}
	text = strings.ReplaceAll(text, ",", "")
	text = strings.ReplaceAll(text, ".", "")
	text = strings.ReplaceAll(text, "’", "")
	text = strings.ReplaceAll(text, "“", "")
	text = strings.ReplaceAll(text, "”", "")
	text = strings.ReplaceAll(text, "!", "")
	text = strings.ReplaceAll(text, "?", "")
	text = strings.ReplaceAll(text, ";", "")
	text = strings.ReplaceAll(text, ":", "")
	text = strings.ReplaceAll(text, "(", "")
	text = strings.ReplaceAll(text, ")", "")
	text = strings.ReplaceAll(text, ".", "")

	splitted := strings.Split(text, " ")

	// Multiply and Add (MAD) method.
	mad := func(x uint64) uint64 {
		return x % 21
	}

	for _, seed := range seeds {
		histogram := make([]uint64, 21)
		h := murmur3.SeedNew64(seed)
		hash := func(s string) uint64 {
			_, err := h.Write([]byte(s))
			if err != nil {
				panic(err)
			}
			v := mad(h.Sum64())
			return v
		}
		d := make(map[string]struct{})
		for _, elem := range splitted {
			d[elem] = struct{}{}
		}
		for word := range d {
			histogram[hash(word)]++
		}
		fmt.Println("seed: ", seed, ",hist", histogram)
	}

	s1 := "avocado"
	s2 := "banana"

	var colCounter int

	for range 1000 {
		seed := rand.Int64()
		h := murmur3.SeedNew64(uint64(seed))
		_, err := h.Write([]byte(s1))
		if err != nil {
			panic(err)
		}
		v1 := h.Sum64()
		v1 = mad(v1)
		h.Reset()
		_, err = h.Write([]byte(s2))
		if err != nil {
			panic(err)
		}
		v2 := h.Sum64()
		v2 = mad(v2)
		if v1 == v2 {
			fmt.Println("Collision found with seed:", seed)
			colCounter++
		}
	}
	fmt.Println("Number of collisions found:", colCounter)
}
