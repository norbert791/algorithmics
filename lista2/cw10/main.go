package main

import (
	"fmt"
	"math"
	"math/rand/v2"
)

type MonteCarloConfig struct {
	// UpperBound is the upper bound of given function's image
	UpperBound float64
	// LowerBound is the lower bound of given function's image
	LowerBound float64
	// Points is the number of points to generate
	Points uint
	// Function is the function to integrate
	Function func(float64) float64
	// Precision is 'height' of the area around the function's value
	Precision float64
}

func MonteCarloIntegral(config MonteCarloConfig) func(start, end float64) float64 {
	return func(start, end float64) float64 {
		var posHits uint
		var negHits uint
		var posTotal uint
		var negTotal uint
		for range config.Points {
			// Generate random point
			x := start + (end-start)*rand.Float64()
			val := config.Function(x)
			if val >= 0 {
				posTotal++
				y := config.UpperBound * rand.Float64()
				if y <= val+config.Precision {
					posHits++
				}
			} else {
				negTotal++
				y := config.LowerBound * rand.Float64()
				if y >= val-config.Precision {
					negHits++
				}
			}
		}

		var negArea, posArea float64
		if negTotal > 0 {
			// Note: config.LowerBound < 0 if this branch is taken
			negArea = float64(negHits) / float64(negTotal) * (end - start) * config.LowerBound
		}
		if posTotal > 0 {
			// Note: config.UpperBound > 0 if this branch is taken
			posArea = float64(posHits) / float64(posTotal) * (end - start) * config.UpperBound
		}

		return posArea + negArea
	}
}

func main() {
	// Compute integral of sin(x) from 0 to pi
	integral := MonteCarloIntegral(MonteCarloConfig{
		UpperBound: 1,
		LowerBound: -1,
		Points:     1000,
		Function:   math.Sin,
		Precision:  0.005,
	})
	fmt.Println(integral(0, math.Pi))
}
