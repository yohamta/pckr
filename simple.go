package pckr

import (
	"log"
	"sort"
)

// SimpleAlgo is a simple but fast algorithm
type SimpleAlgo struct{}

// Pack computes the locations for the rectangles
func (pckr SimpleAlgo) Pack(R []*Rectangle, maxWidth, maxHeight int) {
	sort.Slice(R, func(i, j int) bool { return R[i].H > R[j].H })
	x, y, currRowMaxH := 0, 0, 0
	for _, r := range R {
		if (x + r.W) > maxWidth {
			y += currRowMaxH
			x = 0
			currRowMaxH = 0
		}
		if (y + r.H) > maxHeight {
			log.Fatal("Error: texture height exceeds max height")
		}
		r.X = x
		r.Y = y
		x += r.W
		if r.H > currRowMaxH {
			currRowMaxH = r.H
		}
	}
}
