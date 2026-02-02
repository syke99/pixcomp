package pixcomp

type ycgco struct {
	y  uint8
	cg int16
	co int16
	a  uint8
}

func (yc *ycgco) RGBA() (uint8, uint8, uint8, uint8) {
	y := int16(yc.y)

	t := y - yc.cg/2
	g := yc.cg + t
	b := t - yc.co/2
	r := b + yc.co

	return uint8(r), uint8(g), uint8(b), yc.a
}

func (yc *ycgco) YCgCoA() (y uint8, cg, co int16, alpha uint8) {
	return yc.y, yc.cg, yc.co, yc.a
}
