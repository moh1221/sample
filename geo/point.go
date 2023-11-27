package geo

import "math"

type Point struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// NewPoint Create new point
func NewPoint(lat, lng float64) *Point {
	return &Point{
		Lat: lat,
		Lng: lng,
	}
}

// Distance calculate between two points.
func Distance(p1, p2 *Point, unit ...string) float64 {
	radlat1 := float64(math.Pi * p1.Lat / 180)
	radlat2 := float64(math.Pi * p2.Lat / 180)

	theta := float64(p1.Lng - p2.Lng)
	radtheta := float64(math.Pi * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)
	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515

	if len(unit) > 0 {
		if unit[0] == "K" {
			dist = dist * 1.609344
		} else if unit[0] == "N" {
			dist = dist * 0.8684
		}
	}

	return dist
}
