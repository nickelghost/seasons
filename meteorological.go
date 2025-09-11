package seasons

import "time"

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

func GetMeteorologicalForHemisphere(t time.Time, hs Hemisphere) Season {
	northern := GetMeteorological(t)

	if hs == SouthernHemisphere {
		return (northern + 2) % 4 //nolint:mnd
	}

	return northern
}
