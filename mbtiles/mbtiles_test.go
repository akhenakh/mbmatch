package mbtiles

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecodeTileData(t *testing.T) {
	f, err := os.Open("testdata/5771.pbf")
	require.NoError(t, err)

	b, err := ioutil.ReadAll(f)
	require.NoError(t, err)

	tile, err := DecodeTileData(b)
	require.NoError(t, err)
	require.NotNil(t, tile)
	t.Log(tile)
}
