package main

type Manager struct {
	cache   []*Handle
	handles []*Handle
	systems []*System
}

func (m *Manager) update() {
	for _, s := range m.systems {
		for _, h := range m.handles {
			e := h.Entity()
			if e != nil {
				if e.HasSignature(s.signature) {
					runSys := *s.f
					runSys(e.Components)
				}
			}
		}
	}
}

func (m *Manager) NewEntity(comps ...Component) *Handle {
	//Check if we can resue any of our entities before making a new one
	end := len(m.cache)
	if end > 0 {
		ret := m.cache[end]
		m.cache = m.cache[:end-1]
		return ret.Instatiate(comps...)
	}
	return NewEntity(comps...)
}

func NewManager(h []*Handle, s []*System) Manager {
	return Manager{
		handles: h,
		systems: s,
	}
}
