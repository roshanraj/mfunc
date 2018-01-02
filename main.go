package main

import (
	"fmt"
	math "math"
)

func Logit(p float64) float64 {
	return math.Log(p / (1 - p))
}
func SoftMax(p []float64) []float64 {
	return math.Exp(1)
}
func main() {
	p := 0.8
	fmt.Println("logit with value ", Logit(p))
	val := []float64{1, 2, 3, 4}
	fmt.Println("softmax with value ", SoftMax(val))
}
