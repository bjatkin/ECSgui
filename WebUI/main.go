package webui

import (
	"ECSgui/ECSys"

	"github.com/dennwc/dom"
)

func ExampleGUI() {
	moveBttn := SVG{
		CType: ecs.CType{1},
		Tag:   "svg",
		Attr: map[string]interface{}{
			"width":  100,
			"height": 100,
		},
		Children: []*SVG{
			&SVG{
				Tag: "rect",
				Attr: map[string]interface{}{
					"x":      0,
					"y":      0,
					"width":  50,
					"height": 50,
					"fill":   "red",
					"rx":     7,
					"ry":     7,
				},
			},
			&SVG{
				Tag: "path",
				Attr: map[string]interface{}{
					"d":    "M20 17 L25 7 L30 17 Z",
					"fill": "white",
				},
			},
			&SVG{
				Tag: "path",
				Attr: map[string]interface{}{
					"d":         "M20 17 L25 7 L30 17 Z",
					"fill":      "white",
					"transform": "rotate(90 25 25)",
				},
			},
			&SVG{
				Tag: "path",
				Attr: map[string]interface{}{
					"d":         "M20 17 L25 7 L30 17 Z",
					"fill":      "white",
					"transform": "rotate(180 25 25)",
				},
			},
			&SVG{
				Tag: "path",
				Attr: map[string]interface{}{
					"d":         "M20 17 L25 7 L30 17 Z",
					"fill":      "white",
					"transform": "rotate(270 25 25)",
				},
			},
		},
	}

	MoveButton := ecs.NewEntity(&moveBttn)
	DrawSVG := ecs.NewSystem([]uint16{1}, func(c []ecs.Component) {
		comp := ecs.FindComponents(c, 1)[0]
		//Draw the svg here
		svg := comp.(*SVG)

		var build func(*dom.Element, []*SVG)
		build = func(parent *dom.Element, children []*SVG) {
			for _, c := range children {
				tag := dom.Doc.CreateElementNS("http://www.w3.org/2000/svg", c.Tag)
				for k, v := range c.Attr {
					tag.SetAttribute(k, v)
				}
				build(tag, c.Children)
				parent.AppendChild(tag)
			}
		}

		main := dom.Doc.CreateElementNS("http://www.w3.org/2000/svg", svg.Tag)
		for k, v := range svg.Attr {
			main.SetAttribute(k, v)
		}

		build(main, svg.Children)
		dom.Body.AppendChild(main)
	})

	manager := ecs.NewManager([]*ecs.Handle{MoveButton}, []*ecs.System{&DrawSVG})

	manager.Update()
}
