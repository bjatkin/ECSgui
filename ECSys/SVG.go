package main

type SVG struct {
	Width   float64
	Height  float64
	ViewBox [4]float64

	Transform transform
	Children  []*Tag
}

type Tag struct {
	Id   string
	Type string

	Fill   color
	Stroke color

	D path

	StrokeWidth float64
	CX          float64
	CY          float64
	R           float64

	Children []*Tag
}
