package webui

import (
	"ECSgui/ECSys"
	"time"
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
	manager := ecs.NewManager(
		[]*ecs.Handle{
			MoveButton,
		},
		[]*ecs.System{
			DrawSVG(),
			HandleClick(),
		})

	manager.Update()
	moveBttn.Children[0].Attr["fill"] = "blue"
	time.Sleep(1000 * time.Millisecond)
	manager.Update()
}
