package main

import "fmt"

type Match struct {
	Index   int
	Matched string
}

// Config represents the configuration of the DNA sequence (substring) that is to be matched.
// It is assumed that each string in the config is 3 characters long and each character is in the set {A, C, G, T}.
type Config struct {
	Preffix        map[string]struct{}
	InterfixLength int
	Suffix         map[string]struct{}
}

func DefaultConfig() *Config {
	return &Config{
		Preffix: map[string]struct{}{
			"ATG": {},
		},
		InterfixLength: 30,
		Suffix: map[string]struct{}{
			"TAA": {},
			"TAG": {},
			"TGA": {},
		},
	}

}

func findSubstrings(config *Config, str string) []Match {
	if config == nil {
		config = DefaultConfig()
	}
	var matches []Match

	fin := config.Suffix
	start := config.Preffix
	interfixLength := config.InterfixLength

	// State machine
	const (
		findPrefix = iota
		move
		findSuffix
	)

	state := findPrefix
	moveCounter := 0
	matchStart := 0
	i := 0

	for i+3 <= len(str) {
		subStr := str[i : i+3]
		// fmt.Println(state, moveCounter)
		switch state {
		case findPrefix:
			if _, ok := start[subStr]; ok {
				state = move
				matchStart = i
				// Reset moveCounter. Since find Prefix matches XYZ, we need to subtract (YZ)'s length from the counter.
				moveCounter = -2
			}
			i++
		case move:
			switch {
			}
			if moveCounter == interfixLength {
				state = findSuffix
				break
			}
			_, ok1 := start[subStr]
			_, ok2 := fin[subStr]
			if ok1 || ok2 {
				state = findPrefix
				break
			}

			i++
			moveCounter++
		case findSuffix:
			if _, ok := fin[subStr]; ok {
				matches = append(matches, Match{matchStart, str[matchStart : i+3]})
				state = findPrefix
				break
			} else if _, ok := start[subStr]; ok {
				state = findPrefix
				break
			}
			i++
		default:
			panic("unknown state")
		}
	}
	return matches
}

func main() {
	fmt.Println("Example 1")
	sampleString := "TAACGYTAATGCCCCCCTAGATGCCCCCCCTTTTTTTTTCAAAAAAACGGGGGGGGTGAAAAAAAA"
	matches := findSubstrings(nil, sampleString)

	for _, match := range matches {
		fmt.Printf("Matched: %s at index %d\n", match.Matched, match.Index)
	}

	// Examples ~Joanna Kulig
	fmt.Println("Example 2")
	sampleString = "ATGCACGTCCAACAAACATCAAAACAAAAAAAATAACTTTGATAATGCACGGTCCACAAACTCAAGGCAACAAAAAACTGA"
	matches = findSubstrings(nil, sampleString)
	for _, match := range matches {
		fmt.Printf("Matched: %s at index %d\n", match.Matched, match.Index)
	}

	fmt.Println("Example 3")
	sampleString = "ATGCCAAAAAAAAATGCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCTAA"
	matches = findSubstrings(nil, sampleString)
	for _, match := range matches {
		fmt.Printf("Matched: %s at index %d\n", match.Matched, match.Index)
	}
}
