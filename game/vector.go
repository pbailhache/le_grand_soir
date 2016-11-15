package game

import "math"

type Vector struct {
	x int
	y int
}

func (v *Vector) Magnitude() float64 {
	float := float64(v.x ^ v.x + v.y ^ v.y)
	return math.Abs(math.Sqrt(float))
}
