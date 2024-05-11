package helper

import "math"

func AlmostEqual(a, b, e float64) bool {
    return math.Abs(a - b) <= e
}