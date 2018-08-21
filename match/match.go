package match

import "github.com/akhenakh/mbmatch/mbtiles"

type Engine struct {
	*mbtiles.DB
	Debug bool
}

type MatchedPosition struct {
	Matched bool

	// Map matched position
	Lat, Lng, Heading float64

	// Original position
	RawLat, RawLng, RawHeading float64

	// Hint to pass to the next query
	Hint string
}

func (e *Engine) MapMatch(lat, lng, heading float64, hint string) *MatchedPosition {
	return nil
}
