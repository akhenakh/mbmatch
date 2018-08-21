package match

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

type fakeTileReader struct {
	b []byte
}

func (f *fakeTileReader) ReadTile(uint8, uint64, uint64) ([]byte, error) {
	return f.b, nil
}

func TestEngine_MapMatch(t *testing.T) {
	f, err := os.Open("../mbtiles/testdata/5771.pbf")
	require.NoError(t, err)

	b, err := ioutil.ReadAll(f)
	require.NoError(t, err)

	tr := &fakeTileReader{b: b}

	type args struct {
		lat     float64
		lng     float64
		heading float64
		hint    string
	}
	tests := []struct {
		name    string
		args    args
		want    *MatchedPosition
		wantErr bool
	}{
		{"Chemin Royal", args{lat: 46.874876, lng: -71.097252, heading: 40}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Engine{
				TileReader: tr,
			}
			got, err := e.MapMatch(tt.args.lat, tt.args.lng, tt.args.heading, tt.args.hint)
			if (err != nil) != tt.wantErr {
				t.Errorf("Engine.MapMatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Engine.MapMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
