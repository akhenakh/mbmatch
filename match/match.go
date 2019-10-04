package match

import (
	"log"

	"github.com/akhenakh/mbmatch/mbtiles"
)

// assume 14 is detailed enough and exists in the mbtiles
const zoomLevel = 14

type TileReader interface {
	ReadTile(uint8, uint64, uint64) ([]byte, error)
}

type Engine struct {
	TileReader
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

func (e *Engine) MapMatch(lat, lng, heading float64, hint string) (*MatchedPosition, error) {
	x, y := mbtiles.CoordinatesToXY(lat, lng, zoomLevel)

	d, err := e.ReadTile(zoomLevel, x, y)
	if err != nil {
		return nil, err
	}

	t, err := mbtiles.TileFromData(d)
	if err != nil {
		return nil, err
	}

	for _, l := range t.Layers {
		if l.Name != "transportation" {
			continue
		}

		for _, f := range l.Features {
			if f.Type != mbtiles.Tile_LINESTRING {
				continue
			}

			switch on class
			motorway
trunk
primary
secondary
tertiary

			log.Println(l.Keys, l.Values, f.Tags, f.Id)
		}
	}

	return nil, nil
}
