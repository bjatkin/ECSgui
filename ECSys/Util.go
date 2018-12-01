package main

func GenerageSignature(signature ...uint64) [2]uint64 {
	ret := [2]uint64{}
	for _, sig := range signature {
		index := sig / 64
		ret[index] = 1 << (sig - 64*index)
	}

	return ret
}

func GenerageSignatureFromComponents(comps ...Component) [2]uint64 {
	ret := [2]uint64{}
	for _, c := range comps {
		index := c.Type() / 64
		ret[index] = 1 << (c.Type() - 64*index)
	}

	return ret
}
