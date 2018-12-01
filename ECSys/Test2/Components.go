package gui

type SVGComponent struct {
	RootTag *Tag
	Visible bool
}

type PositionComponent struct {
	X     float32
	Y     float32
	Layer int32
}
