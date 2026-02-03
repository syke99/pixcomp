package pixcomp

type ycgco struct {
	y  uint8
	cg int16
	co int16
	a  uint8
}

func (yc *ycgco) RGBA() (uint8, uint8, uint8, uint8) {
	y32 := int32(yc.y)
	cg32 := int32(yc.cg)
	co32 := int32(yc.co)

	t := y32 - (cg32 >> 1)
	g := cg32 + t
	b := t - (co32 >> 1)
	r := b + co32

	return clamp8(r), clamp8(g), clamp8(b), yc.a
}

func (yc *ycgco) YCgCoA() (y uint8, cg, co int16, alpha uint8) {
	return yc.y, yc.cg, yc.co, yc.a
}
