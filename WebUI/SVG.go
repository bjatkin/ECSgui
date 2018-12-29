package webui

import (
	"ECSgui/ECSys"

	"github.com/dennwc/dom"
)

type SVG struct {
	ecs.CType                        //We need this so that SVG adheres to the component interface
	Tag       string                 //Change this to a byte array to make make sure it's on the stack
	Hide      bool                   //Use this so every tag can be added and removed easily
	Attr      map[string]interface{} //So ugly XP is there a better way we could achieve something similar?
	Children  []*SVG                 //Maybe stick a limit on this to make sure as much is on the stack as posible
	Element   *dom.Element
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
