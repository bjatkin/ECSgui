package webui

import (
	"ECSgui/ECSys"

	"github.com/dennwc/dom"
)

var ()

func DrawSVG() *ecs.System {
	ret := ecs.NewSystem(
		[]uint16{1}, //The signature expected by this system
		func(c []ecs.Component) { //The actual functionality
			comp := ecs.FindComponents(c, 1)[0]
			svg := comp.(*SVG)

			var build func(*dom.Element, []*SVG)
			build = func(parent *dom.Element, children []*SVG) {
				for i, c := range children {
					tag := c.Element
					if tag == nil {
						tag = dom.Doc.CreateElementNS("http://www.w3.org/2000/svg", c.Tag)
					}
					for k, v := range c.Attr {
						tag.SetAttribute(k, v)
					}
					build(tag, c.Children)
					children[i].Element = tag
					parent.AppendChild(tag)
				}
			}

			main := svg.Element
			if main == nil {
				main = dom.Doc.CreateElementNS("http://www.w3.org/2000/svg", svg.Tag)
			}

			for k, v := range svg.Attr {
				main.SetAttribute(k, v)
			}

			build(main, svg.Children)
			svg.Element = main

			dom.Body.AppendChild(main)
		})
	return &ret
}

func HandleClick() *ecs.System {
	ret := ecs.NewSystem(
		[]uint16{1, 2},
		func(c []ecs.Component) {

		})
	return &ret
}
