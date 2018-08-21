package mbtiles

import (
	"reflect"
	"testing"
)

func TestDecodeGeometryT(t *testing.T) {
	linestring := []uint32{9, 50, 34}
	res := DecodeGeometry(linestring)
	t.Log(res)

	linestring = []uint32{9, 4, 4, 18, 0, 16, 16, 0}
	res = DecodeGeometry(linestring)
	t.Log(res)

	// Buggy
	linestring = []uint32{17, 10, 14, 3, 9}
	res = DecodeGeometry(linestring)
	t.Log(res)

}

func TestDecodeGeometry(t *testing.T) {
	type args struct {
		g []uint32
	}
	tests := []struct {
		name string
		args args
		want []Command
	}{
		{"One Point", args{[]uint32{9, 50, 34}},
			[]Command{Command{MoveTo, []uint32{25, 17}}}},

		{"LineString", args{[]uint32{9, 4, 4, 18, 0, 16, 16, 0}},
			[]Command{Command{MoveTo, []uint32{2, 2}}, Command{LineTo, []uint32{0, 8, 8, 0}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeGeometry(tt.args.g); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeGeometry() = %v, want %v", got, tt.want)
			}
		})
	}
}
