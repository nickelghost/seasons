# seasons

Library that allows looking up the season that a date is in.

## Usage

You can find a meteorological season, for the Northern hemisphere as a default:

```go
s := seasons.GetMeteorological(time.Now())
```

or specify a hemisphere if it's known and relevant:

```go
s := seasons.GetMeteorologicalForHemisphere(time.Now(), seasons.SouthernHemisphere)
```
