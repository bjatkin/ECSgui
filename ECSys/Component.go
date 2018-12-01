package main

type Component interface {
	Type() uint64
}

type CType struct {
	T uint64
}

func (c *CType) Type() uint64 {
	return c.T
}

func FindComponents(comps []Component, Types ...uint64) []Component {
	ret := make([]Component, len(Types))
	for _, comp := range comps {
		for i, t := range Types {
			if comp.Type() == t {
				ret[i] = comp
			}
		}
	}
	return ret
}
