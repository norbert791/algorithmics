package main

import "fmt"

const startSeq = "ATG"

// suffixState has complexity O(1) and returns the next state of the suffix automaton.
func suffixState(r rune, state string) string {
	switch r {
	case 'T':
		if state == "" {
			return "T"
		}
	case 'A':
		if state == "T" {
			return "TA"
		}
		if state == "TG" {
			return "0"
		}
		if state == "TA" {
			return "0"
		}
	case 'G':
		if state == "T" {
			return "TG"
		}
		if state == "TA" {
			return "0"
		}
	default:
		return ""
	}

	return ""
}

// interfixState has complexity O(1) and returns the next state of the interfix automaton.
func interfixState(r rune, state string) string {
	switch {
	case r == 'A' && state == "":
		return "A"
	case r == 'T' && state == "A":
		return "AT"
	case r == 'G' && state == "AT":
		return "0"
	}

	return ""
}

type Match struct {
	Index   int
	Matched string
}

// Config represents the configuration of the DNA sequence (substring) that is to be matched.
// It is assumed that each string in the config is 3 characters long and each character is in the set {A, C, G, T}.
type Config struct {
	Preffix        map[string]struct{}
	InterfixLength int
	InterfixBanned map[string]struct{}
	Suffix         map[string]struct{}
}

func DefaultConfig() *Config {
	return &Config{
		Preffix: map[string]struct{}{
			"ATG": {},
		},
		InterfixLength: 30,
		InterfixBanned: map[string]struct{}{
			"ATG": {},
			"TAA": {}, "TAG": {}, "TGA": {},
		},
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

	banned := config.InterfixBanned
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
		fmt.Println(state, moveCounter)
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
			if _, ok := banned[subStr]; ok {
				state = findPrefix
				break
			}
			if moveCounter == interfixLength {
				state = findSuffix
				break
			}
			i++
			moveCounter++
		case findSuffix:
			if _, ok := banned[subStr]; ok {
				state = findPrefix
			}
			if _, ok := fin[subStr]; ok {
				matches = append(matches, Match{matchStart, str[matchStart : i+3]})
			}
			i++
		default:
			panic("unknown state")
		}
	}
	return matches
}

func main() {
	sampleString := "TAACGYTAATGCCCCCCTAGATGCCCCCCCTTTTTTTTTCAAAAAAACGGGGGGGGTGAAAAAAAA"
	matches := findSubstrings(nil, sampleString)

	for _, match := range matches {
		fmt.Printf("Matched: %s at index %d\n", match.Matched, match.Index)
	}
}
