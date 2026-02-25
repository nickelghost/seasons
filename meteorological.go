package seasons

import "time"

// GetMeteorological returns the meteorological season for the given time,
// assuming the Northern Hemisphere. Seasons are defined as consecutive
// 3-month periods: Spring (Mar–May), Summer (Jun–Aug), Autumn (Sep–Nov),
// Winter (Dec–Feb).
func GetMeteorological(t time.Time) Season {
	m := t.Month()

	switch {
	case m >= time.March && m <= time.May:
		return Spring
	case m >= time.June && m <= time.August:
		return Summer
	case m >= time.September && m <= time.November:
		return Fall
	default:
		return Winter
	}
}

// GetMeteorologicalForHemisphere returns the meteorological season for the
// given time and hemisphere. For the Southern Hemisphere the season is the
// opposite of the Northern Hemisphere season (e.g. Summer becomes Winter).
func GetMeteorologicalForHemisphere(t time.Time, hs Hemisphere) Season {
	northern := GetMeteorological(t)

	if hs == SouthernHemisphere {
		// Adding 2 (mod 4) flips to the opposite season. This relies on the
		// Season constants being declared in order with opposite seasons
		// exactly 2 apart: Spring(0)<->Autumn(2), Summer(1)<->Winter(3).
		return (northern + 2) % 4 //nolint:mnd
	}

	return northern
}
