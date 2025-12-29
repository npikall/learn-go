package clockface

import (
	"math"
	"time"
)

// A Point represents a 2D Cartesian coordinate
type Point struct {
	X float64
	Y float64
}

// Scale the unit vector of a point
func (p *Point) Scale(scale float64) {
	p.X *= scale
	p.Y *= scale
}

// Flip the y-axis
func (p *Point) FlipY() {
	p.Y = -p.Y
}

// Moves a Point by the given vector/point
func (p *Point) Offset(o Point) {
	p.X += o.X
	p.Y += o.Y
}

const secondHandLength = 90
const clockCenterX = 150
const clockCenterY = 150

// SecondHand is the unit vector of the second hand of an analogue clock at time `t`
// represented as a Point.
func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	p.Scale(secondHandLength)
	p.FlipY()
	p.Offset(Point{X: clockCenterX, Y: clockCenterY})
	return p
}

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / float64(t.Second())))
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	X := math.Sin(angle)
	Y := math.Cos(angle)
	return Point{X: X, Y: Y}
}
