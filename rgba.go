package pixcomp

type rgba struct {
	r, g, b, a uint8
}

func (r *rgba) RGBA() (uint8, uint8, uint8, uint8) {
	return r.r, r.g, r.b, r.a
}

func (r *rgba) YCgCoA() (uint8, int16, int16, uint8) {
	R := int32(r.r)
	G := int32(r.g)
	B := int32(r.b)

	co32 := R - B
	t := B + (co32 >> 1)
	cg32 := G - t
	y32 := t + (cg32 >> 1)

	return clamp8(y32), int16(cg32), int16(co32), r.a
}
