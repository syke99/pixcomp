package pixcomp

// Color can be extended for conversions
// to more color-spaces
type Color interface {
	// RGBA converts the Color's underlying
	// values into the RGBA color-space
	RGBA() (uint8, uint8, uint8, uint8)
	// YCgCoA converts the Color's underlying
	// values into the YCgCoA color-space
	YCgCoA() (y uint8, cg, co int16, alpha uint8)
}

// NewRGBA creates a new Color based
// on RGBA values
func NewRGBA(r, g, b, a uint8) Color {
	return &rgba{
		r: r,
		g: g,
		b: b,
		a: a,
	}
}

// NewYCgCoA creates a new Color based
// on YCgCoA values
func NewYCgCoA(y uint8, cg int16, co int16, a uint8) Color {
	return &ycgco{
		y:  y,
		cg: cg,
		co: co,
		a:  a,
	}
}
