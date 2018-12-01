package main

import (
	"fmt"
)

func main() {
	pos := Position{
		CType: CType{1},
		X:     10.0,
		Y:     10.0,
	}

	emoji := Emoji{
		CType: CType{2},
		GUI:   ";)",
		Name:  "Winky Face",
	}

	player := NewEntity(&pos, &emoji)
	other := NewEntity(&pos)

	echo := NewSystem([]uint16{1}, func(c []Component) {
		comps := FindComponents(c, 1)
		pos := comps[0].(*Position)

		fmt.Printf("X: %f, Y: %f\n\n", pos.X, pos.Y)
	})

	drawFace := NewSystem([]uint16{1, 2}, func(c []Component) {
		comps := FindComponents(c, 1, 2)
		pos, emoji := comps[0].(*Position), comps[1].(*Emoji)

		fmt.Printf("Hey Look a %s\n%s\n", emoji.Name, emoji.GUI)
		pos.X++
		pos.Y++
	})

	manager := NewManager([]*Handle{player, other}, []*System{&echo, &drawFace})
	manager.update()
	manager.update()
	manager.update()
	manager.update()
}

type Position struct {
	CType
	X float64
	Y float64
}

type Emoji struct {
	CType
	GUI  string
	Name string
}
