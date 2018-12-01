package components

//--------------------------------
//This the the runtime implementation
//--------------------------------

type svg struct {
	Tag      string
	Hide     bool
	Attr     map[string]interface{}
	Children []*svg
}

type state struct {
	Name        string
	Transitinos []*transition
}

type transition struct {
	Name   string
	Event  string
	Target *state
	Mod    []*mod
}

type mod struct {
	target *svg
	Hide   bool
	Attr   map[string]string
}

type UIComponent struct {
	Initial     string
	Disp        svg
	States      []state
	Transitions []transition
}

//---------------------------------
//This is the compile time implementation
//---------------------------------

func NewButtonComponent() UIComponent {
	ret := UIComponent{
		Initial: "unpressed",
		Disp: svg{
			Tag:  "svg",
			Attr: map[string]interface{}{"width": 500, "height": 500},
			Children: []*svg{
				&svg{
					Tag: "rect",
					Attr: map[string]interface{}{"x": 50, "y": 50, "rx": 20, "ry": 20, "width": 400, "height": 150,
						"style": "fill:red;stroke:black;stroke-width:1;opacity:0.5"},
				},
			},
		},
	}

	trans := []transition{
		transition{
			Name:  "hover-over",
			Event: "on-hover-start",
			Mod: []*mod{
				&mod{
					target: ret.Disp.Children[0],
					Attr:   map[string]string{"style": "fill:red;stroke:black;stroke-width:1;"},
				},
			},
		},
	}

	state := []state{
		state{
			Name: "unpressed",
		},
		state{
			Name: "hovered",
		},
	}

	trans[0].Target = &state[1]

	state[0].Transitinos = []*transition{
		&trans[0],
	}

	ret.States = state
	ret.Transitions = trans

	return ret
}
