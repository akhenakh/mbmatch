package mbtiles

import (
	"math"
)

// CoordinatesToXY converts lat lng with z as zoom to x y tile
func CoordinatesToXY(lat, lng float64, z uint) (x, y uint64) {
	latr := lat * math.Pi / 180
	n := math.Pow(2, float64(z))
	x = uint64((lng + 180.0) / 360.0 * n)
	y = uint64((1.0 - math.Log(math.Tan(latr)+(1/math.Cos(latr)))/math.Pi) / 2.0 * n)
	return x, y
}
