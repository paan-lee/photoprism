package geo

import (
	"math"
)

// Position represents a geo coordinate.
type Position struct {
	Lat float64
	Lng float64
}

// DegToRad converts a value from degrees to radians.
func DegToRad(d float64) float64 {
	return d * math.Pi / 180
}

// Dist returns the shortest path between two positions in km.
func Dist(p, q Position) (km float64) {
	if p.Lat == q.Lat && p.Lng == q.Lng {
		return 0.0
	}

	lat1 := DegToRad(p.Lat)
	lng1 := DegToRad(p.Lng)
	lat2 := DegToRad(q.Lat)
	lng2 := DegToRad(q.Lng)

	diffLat := lat2 - lat1
	diffLng := lng2 - lng1

	a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*
		math.Pow(math.Sin(diffLng/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return c * EarthRadiusKm
}
