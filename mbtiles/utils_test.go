package mbtiles

import "testing"

func TestCoordinatesToXY(t *testing.T) {
	type args struct {
		lat float64
		lng float64
		z   uint
	}
	tests := []struct {
		name  string
		args  args
		wantX uint64
		wantY uint64
	}{
		{"Quebec", args{lat: 46.797358, lng: -71.228874, z: 14}, 4950, 5776},
		{"Montreal", args{lat: 45.498848, lng: -73.597851, z: 14}, 4842, 5861},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := CoordinatesToXY(tt.args.lat, tt.args.lng, tt.args.z)
			if gotX != tt.wantX {
				t.Errorf("CoordinatesToXY() gotX = %v, want %v", gotX, tt.wantX)
			}
			if gotY != tt.wantY {
				t.Errorf("CoordinatesToXY() gotY = %v, want %v", gotY, tt.wantY)
			}
		})
	}
}
