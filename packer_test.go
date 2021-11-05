package pckr_test

import (
	"testing"

	"github.com/yohamta/pckr"
)

func TestSimplePackerAlgo(t *testing.T) {
	type result struct {
		isPressed bool
	}
	R := make([]*pckr.Rectangle, 0)

	// random 50 rectangles with different size between 16x16 and 256x256
	for _, v := range [][]int{
		{300, 600},
		{700, 400},
		{700, 600},
		{300, 400},
	} {
		r := pckr.Rect(v[0], v[1])
		R = append(R, &r)
	}
	var tests = []struct {
		name string
		a    []*pckr.Rectangle
		b    int
		want bool
	}{
		{"pack rects w/o overlaps", R, 1000, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := _TestSimpleAlgo(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("TestPaking(%s): got %v; want %v", tt.name, got, tt.want)
			}
		})
	}
}

func _TestSimpleAlgo(R []*pckr.Rectangle, size int) bool {
	algo := pckr.SimpleAlgo{}
	algo.Pack(R, size, size)
	arr := make([][]bool, size)
	for i := range arr {
		arr[i] = make([]bool, size)
	}
	for _, r := range R {
		for i := r.X; i < r.X+r.W; i++ {
			for j := r.Y; j < r.Y+r.H; j++ {
				if arr[i][j] {
					return false
				}
				arr[i][j] = true
			}
		}
	}
	return true
}
