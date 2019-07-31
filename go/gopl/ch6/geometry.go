package main

import "math"

type Point struct {
	X,Y float64
}

func Distance(p, q Point) float64 {
	return math.Hypot(p.X - q.X, p.Y - q.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X - q.X, p.Y - q.Y)
}
