package webui

import (
	"ECSgui/ECSys"
)

type SVG struct {
	ecs.CType                        //We need this so that SVG adheres to the component interface
	Tag       string                 //Change this to a byte array to make make sure it's on the stack
	Hide      bool                   //Use this so every tag can be added and removed easily
	Attr      map[string]interface{} //So ugly XP is there a better way we could achieve something similar?
	Children  []*SVG                 //Maybe stick a limit on this to make sure as much is on the stack as posible
}

type State struct {
	Name        string
	Transitions []*Transition
}

type Transition struct {
	Name         string
	Event        string
	Target       *State
	Modification []*Mod
}

type Mod struct {
	Target *SVG
	Hide   bool
	Attr   map[string]interface{} //Again I kinda hate this. It dosen't feel very Go ish
}

type UIComponent struct {
	ecs.CType   //Make sure the UIComponent adheres to the component interface
	Initial     *SVG
	States      []State
	Transitions []Transition
}

//If we do it like this we will have very big files that will have to be updated to
//keep up with the spec. It's much less flexible. However it will be much clearer and more
//explicit to work with. I like this however It will also make this component extreamly heavy
//As memory will have to be reserved for every possible attribute.

// type SVG struct {
// 	Width   float64
// 	Height  float64
// 	ViewBox [4]float64

// 	Transform transform
// 	Children  []*Tag
// }

// type Tag struct {
// 	Id   string
// 	Type string

// 	Fill   color
// 	Stroke color

// 	D path

// 	StrokeWidth float64
// 	CX          float64
// 	CY          float64
// 	R           float64

// 	Children []*Tag
// }
