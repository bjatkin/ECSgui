package gui

type SVGEntity struct {
	Signature []uint64
	GUI       SVGComponent
	Pos       PositionComponent
}

func NewSVGEntity(gui SVGComponent, x, y float32, layer int32) *SVGEntity {
	return &SVGEntity{
		Signature: GenerageSignature([]uint64{1, 2}),
		GUI:       gui,
		Pos:       PositionComponent{X: x, Y: y, Layer: layer},
	}
}
