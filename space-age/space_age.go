package space

// Planet is just a string representing the planet name
type Planet string

const earthYearInSeconds = 31557600

var planetPeriods = map[Planet]float64{
	"Earth":   earthYearInSeconds,
	"Mercury": earthYearInSeconds * 0.2408467,
	"Venus":   earthYearInSeconds * 0.61519726,
	"Mars":    earthYearInSeconds * 1.8808158,
	"Jupiter": earthYearInSeconds * 11.862615,
	"Saturn":  earthYearInSeconds * 29.447498,
	"Uranus":  earthYearInSeconds * 84.016846,
	"Neptune": earthYearInSeconds * 164.79132,
}

// Age will calculate your age on a Planet
func Age(seconds float64, planet Planet) float64 {
	period, ok := planetPeriods[planet]
	if !ok {
		period = 1
	}

	return seconds / period
}
