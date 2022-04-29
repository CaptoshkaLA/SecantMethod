package main

import "math"

const e = 2.718

func Function(x float64) float64 {
	return math.Pow(x+1.2, 4)
}

func Derivative(x float64) float64 {
	return 4 * math.Pow(x+1.2, 3)
}
