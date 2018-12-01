package main

type System struct {
	signature [2]uint64 //Note this will only allow up to 128 systems to be in play at a time. is this too many? to few?
	f         *func([]Component)
}

func NewSystem(signature []uint16, callBack func([]Component)) System {
	systemSig := [2]uint64{0, 0}
	for _, sig := range signature {
		index := sig / 64
		systemSig[index] = 1 << (sig - 64*index)
	}

	return System{
		signature: systemSig,
		f:         &callBack,
	}
}
