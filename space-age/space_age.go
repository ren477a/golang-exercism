package space

// Planet ...
type Planet string

// EarthSeconds ...
const EarthSeconds float64 = 31557600

/**
Earth: orbital period 365.25 Earth days, or 31557600 seconds
Mercury: orbital period 0.2408467 Earth years
Venus: orbital period 0.61519726 Earth years
Mars: orbital period 1.8808158 Earth years
Jupiter: orbital period 11.862615 Earth years
Saturn: orbital period 29.447498 Earth years
Uranus: orbital period 84.016846 Earth years
Neptune: orbital period 164.79132 Earth years
**/

// Age ...
func Age(s float64, p Planet) float64 {
	m := make(map[Planet]float64)
	m["Earth"] = 1
	m["Mercury"] = 0.2408467
	m["Venus"] = 0.61519726
	m["Mars"] = 1.8808158
	m["Jupiter"] = 11.862615
	m["Saturn"] = 29.447498
	m["Uranus"] = 84.016846
	m["Neptune"] = 164.79132
	earthAge := s / EarthSeconds
	return earthAge / m[p]
}
