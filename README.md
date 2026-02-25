# seasons

Go library for looking up the season a date falls in.

## Installation

```sh
go get github.com/nickelghost/seasons
```

## Usage

### Meteorological seasons

Meteorological seasons divide the year into four 3-month periods aligned to the calendar (March–May, June–August, September–November, December–February).

Get the season for the Northern Hemisphere (the default):

```go
s := seasons.GetMeteorological(time.Now())
```

Get the season for a specific hemisphere:

```go
s := seasons.GetMeteorologicalForHemisphere(time.Now(), seasons.NorthernHemisphere)
s := seasons.GetMeteorologicalForHemisphere(time.Now(), seasons.SouthernHemisphere)
```

Both functions return a `seasons.Season` value. The available constants are:

| Constant         | Note               |
| ---------------- | ------------------ |
| `seasons.Spring` |                    |
| `seasons.Summer` |                    |
| `seasons.Autumn` | canonical name     |
| `seasons.Winter` |                    |
| `seasons.Fall`   | alias for `Autumn` |

Hemisphere constants:

| Constant                     |
| ---------------------------- |
| `seasons.NorthernHemisphere` |
| `seasons.SouthernHemisphere` |
