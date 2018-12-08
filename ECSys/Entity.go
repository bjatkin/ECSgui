package ecs

import (
	"math"
)

//Counters for Entity and Component ids
var (
	nextEntityID uint32
	nextHandleID uint32
)

//TODO Figure out a way too keep these from getting garbage collected
//TODO Should we declare a size for the component slice so we get it declared on the stack?

//Entity is a lightweight container for components
type Entity struct {
	id         uint32
	phase      uint8
	Signature  [2]uint64
	Components []Component
}

//Release the Entity back to the pool of available entities
func (e *Entity) Release() {
	e.phase++
}

//Clean cleans components off the Entity
func (e *Entity) Clean() *Entity {
	e.Components = nil
	return e
}

//HasSignature checks if this entity matches the given signature
func (e *Entity) HasSignature(sig [2]uint64) bool {
	var mask uint64 = math.MaxUint64
	for i, key := range sig {
		check := key ^ mask //Invert the key so that when we OR it with the entities signature it will be MaxUint64
		if e.Signature[i]|check != mask {
			return false
		}
	}
	return true
}

//Handle is a safe accesor to an entity that allows us to keep released entities away from the GC
//by preventing us from having to constantly new up more structs in memory
type Handle struct {
	id     uint32
	entity *Entity
	phase  uint8
}

//Entity returns the entity the handle points to unless it's ben released already
func (h *Handle) Entity() *Entity {
	if h.phase != h.entity.phase {
		return nil
	}
	return h.entity
}

func (h *Handle) Instatiate(comps ...Component) *Handle {
	h.phase = h.entity.phase
	h.entity.Components = comps
	h.entity.Signature = GenerageSignatureFromComponents(comps...)
	return h
}

//NewEntity hands back a  handle to a new Entity
func NewEntity(components ...Component) *Handle {
	//Figure out the signature of this Entity

	nextEntityID++
	return NewHandle(
		&Entity{
			Signature:  GenerageSignatureFromComponents(components...),
			id:         nextEntityID,
			phase:      0,
			Components: components,
		})
}

//NewHandle makes a new handle that points to the given entity
func NewHandle(e *Entity) *Handle {
	nextHandleID++
	return &Handle{
		id:     nextHandleID,
		entity: e,
		phase:  e.phase,
	}
}
