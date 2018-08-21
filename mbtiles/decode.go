package mbtiles

type CommandID uint8

const (
	MoveTo    CommandID = 1
	LineTo              = 2
	ClosePath           = 7
)

type Command struct {
	ID     CommandID
	Values []uint32
}

func DecodeGeometry(g []uint32) []Command {
	var res []Command
	for i := 0; i < len(g); i++ {
		id := CommandID(g[i] & 0x7)
		c := Command{
			ID: id,
		}

		count := g[i] >> 3
		if id == MoveTo || id == LineTo {
			count *= 2
			c.Values = make([]uint32, count)
			for j := 0; j < int(count); j++ {
				value := ((g[i+j+1] >> 1) ^ (-(g[i+j+1] & 1)))
				c.Values[j] = value
			}
		}
		i = i + int(count)
		res = append(res, c)
		continue
	}

	return res
}
