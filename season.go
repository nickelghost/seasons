// Package seasons provides utilities for determining the season that a given
// date falls in, with support for both hemispheres.
package seasons

// Season represents one of the four meteorological seasons.
type Season int

// The four meteorological seasons. The iota order is significant: opposite
// seasons must be exactly 2 apart (mod 4) so that the hemisphere flip in
// GetMeteorologicalForHemisphere works correctly. Do not reorder these constants.
const (
	Spring Season = iota // March, April, May in the Northern Hemisphere
	Summer               // June, July, August in the Northern Hemisphere
	Autumn               // September, October, November in the Northern Hemisphere
	Winter               // December, January, February in the Northern Hemisphere
)

// Fall is an alias for Autumn provided for American English convenience.
// The two constants are identical in value; prefer Autumn as the canonical name,
// especially in switch statements and any future String() implementation.
const Fall = Autumn
