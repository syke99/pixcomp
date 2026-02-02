package pixcomp

type rgba struct {
	r, g, b, a uint8
}

func (r *rgba) RGBA() (uint8, uint8, uint8, uint8) {
	return r.r, r.g, r.b, r.a
}

func (r *rgba) YCgCoA() (uint8, int16, int16, uint8) {
	rr := int16(r.r)
	g := int16(r.g)
	b := int16(r.b)

	co := rr - b
	t := b + co/2
	cg := g + t
	y := uint8(t + cg/2)

	return y, cg, co, r.a
}
