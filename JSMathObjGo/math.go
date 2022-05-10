package math

import (
	"math"
	"math/rand"
	"time"
)

// Generate a random number between 0 and 1 (float64)
func Random() float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()
}

// Generate a random number between X and Y (int)
func RandomBetween(x int, y int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(y-x) + x
}

// Round a number to the nearest integer
func Round(f float64) int {
	return int(f + 0.5)
}

// Get the absolute value of a number
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Get the power of a number
func Pow(x float64, y float64) float64 {
	return math.Pow(float64(x), float64(y))
}

// Get the square root of a number
func Sqrt(x float64) float64 {
	return math.Sqrt(float64(x))
}

// Get the logarithm of a number
func Log(x float64) float64 {
	return math.Log(float64(x))
}

// Get the ceiling of a number
func Ceil(x float64) float64 {
	return math.Ceil(float64(x))
}
