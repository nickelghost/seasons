package seasons_test

import (
	"testing"
	"time"

	"github.com/nickelghost/seasons"
)

var meteorologicalTestData = []struct {
	month      time.Month
	hemisphere seasons.Hemisphere
	season     seasons.Season
}{
	// northern
	{time.January, seasons.NorthernHemisphere, seasons.Winter},
	{time.February, seasons.NorthernHemisphere, seasons.Winter},
	{time.March, seasons.NorthernHemisphere, seasons.Spring},
	{time.April, seasons.NorthernHemisphere, seasons.Spring},
	{time.May, seasons.NorthernHemisphere, seasons.Spring},
	{time.June, seasons.NorthernHemisphere, seasons.Summer},
	{time.July, seasons.NorthernHemisphere, seasons.Summer},
	{time.August, seasons.NorthernHemisphere, seasons.Summer},
	{time.September, seasons.NorthernHemisphere, seasons.Autumn},
	{time.October, seasons.NorthernHemisphere, seasons.Fall},
	{time.November, seasons.NorthernHemisphere, seasons.Autumn},
	{time.December, seasons.NorthernHemisphere, seasons.Winter},
	// southern
	{time.January, seasons.SouthernHemisphere, seasons.Summer},
	{time.February, seasons.SouthernHemisphere, seasons.Summer},
	{time.March, seasons.SouthernHemisphere, seasons.Autumn},
	{time.April, seasons.SouthernHemisphere, seasons.Fall},
	{time.May, seasons.SouthernHemisphere, seasons.Autumn},
	{time.June, seasons.SouthernHemisphere, seasons.Winter},
	{time.July, seasons.SouthernHemisphere, seasons.Winter},
	{time.August, seasons.SouthernHemisphere, seasons.Winter},
	{time.September, seasons.SouthernHemisphere, seasons.Spring},
	{time.October, seasons.SouthernHemisphere, seasons.Spring},
	{time.November, seasons.SouthernHemisphere, seasons.Spring},
	{time.December, seasons.SouthernHemisphere, seasons.Summer},
}

func TestGetMeteorological(t *testing.T) {
	t.Parallel()

	for _, d := range meteorologicalTestData {
		if d.hemisphere == seasons.NorthernHemisphere {
			date := time.Date(1950, d.month, 23, 13, 8, 5, 3, time.UTC)
			result := seasons.GetMeteorological(date)

			if result != d.season {
				t.Errorf("%d produced season %d instead of %d", d.month, result, d.season)
			}
		}
	}
}

func TestGetMeteorologicalForHemisphere(t *testing.T) {
	t.Parallel()

	for _, d := range meteorologicalTestData {
		date := time.Date(1950, d.month, 23, 13, 8, 5, 3, time.UTC)
		result := seasons.GetMeteorologicalForHemisphere(date, d.hemisphere)

		if result != d.season {
			t.Errorf("%d produced season %d instead of %d", d.month, result, d.season)
		}
	}
}
