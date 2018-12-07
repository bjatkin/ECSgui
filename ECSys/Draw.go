package main

import "fmt"

func exampleGUI() {
	moveBttn := SVG{
		CType: CType{1},
		Tag:   "g",
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

	MoveButton := NewEntity(&moveBttn)
	DrawSVG := NewSystem([]uint16{1}, func(c []Component) {
		comp := FindComponents(c, 1)[0]
		//Draw the svg here
		svg := comp.(*SVG)
		fmt.Printf("%#v\n\n", svg)

		fmt.Print("<" + svg.Tag)
		for k, v := range svg.Attr {
			fmt.Printf(" %s=\"%d\"", k, v)
		}
		fmt.Print(">")
		fmt.Println("\n\n")
	})

	manager := NewManager([]*Handle{MoveButton}, []*System{&DrawSVG})

	manager.update()
}
