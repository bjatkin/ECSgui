package main

type color struct {
	R   int32
	G   int32
	B   int32
	A   float32
	Hex string
}

type path struct {
	path []*pathPoint
}

type pathPoint struct {
	style string
	pos   float64
}

type transform struct {
	rotate    [3]float64
	translate [2]float64
	skewX     float64
	skewY     float64
	scale     [2]float64
}
