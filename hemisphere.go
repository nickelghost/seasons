package seasons

// Hemisphere represents either the Northern or Southern Hemisphere.
type Hemisphere int

const (
	NorthernHemisphere Hemisphere = iota // Countries north of the equator, e.g. UK, USA, China.
	SouthernHemisphere                   // Countries south of the equator, e.g. Australia, Brazil, South Africa.
)
