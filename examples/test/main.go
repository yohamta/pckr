package main

import (
	"image"
	"image/color"
	"log"
	"math"
	"math/rand"
	"strconv"
	"time"

	colorful "github.com/lucasb-eyer/go-colorful"
	"github.com/miyahoyo/pckr"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type textureAtlas struct {
	rect *pckr.Rectangle
	id   int
}

var (
	initialized = false
)

const (
	screenWidth   = 1024
	screenHeight  = 1024
	textureWidth  = 1024
	textureHeight = 1024
)

type Game struct {
	image *ebiten.Image
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.image == nil {
		return
	}
	screen.Clear()
	screen.DrawImage(g.image, &ebiten.DrawImageOptions{})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) SetGridAtlas(T []*textureAtlas) {
	img := ebiten.NewImage(textureWidth, textureHeight)
	for _, t := range T {
		r := t.rect
		c := colorful.HappyColor()
		FillRect(img, image.Rect(r.X, r.Y, r.X+r.W, r.Y+r.H), c)
		DrawRect(img, image.Rect(r.X, r.Y, r.X+r.W, r.Y+r.H), color.RGBA{0xff, 0, 0, 0xff}, 2)
		ebitenutil.DebugPrintAt(img, strconv.Itoa(t.id), r.X, r.Y)
	}
	g.image = img
}

type rect struct {
	W, H, X, Y int
}

func (g *Game) MakeGridAtlas() ([]*textureAtlas, time.Duration, float64) {
	rand.Seed(time.Now().UnixNano())

	id := 0
	R := make([]*pckr.Rectangle, 0)
	T := make([]*textureAtlas, 0)

	// random rectangle (16x16 ~ 256x256)
	for i := 0; i < 50; i++ {
		h := int(math.Pow(2., float64(rand.Intn(4)+4)))
		w := int(math.Pow(2., float64(rand.Intn(4)+4)))
		r := pckr.Rect(w, h)
		T = append(T, &textureAtlas{&r, id})
		R = append(R, &r)
		id++
	}

	// sprite sheet of 8 cells (16x128 ~ 64*512)
	for i := 0; i < 100; i++ {
		h := int(math.Pow(2., float64(rand.Intn(4)+2)))
		w := h * 8
		r := pckr.Rect(w, h)
		T = append(T, &textureAtlas{&r, id})
		R = append(R, &r)
		id++
	}

	// large images
	for i := 0; i < 10; i++ {
		h := rand.Intn(6)*32 + 64
		w := rand.Intn(6)*32 + 64
		r := pckr.Rect(w, h)
		T = append(T, &textureAtlas{&r, id})
		R = append(R, &r)
		id++
	}

	start := time.Now()
	pckr := pckr.SimpleAlgo{}
	pckr.Pack(R, textureWidth, textureHeight)
	elapsed := time.Since(start)

	maxY := 0
	areaSum := 0
	for _, t := range T {
		r := t.rect
		areaSum += r.W * r.H
		if r.Y+r.H > maxY {
			maxY = r.Y + r.H
		}
	}
	eff := float64(areaSum) / float64((maxY * textureWidth))

	return T, elapsed, eff
}

func NewGame() *Game {
	g := &Game{}

	T, _, _ := g.MakeGridAtlas()
	g.SetGridAtlas(T)

	var elapsed time.Duration = 0.
	eff := 0.
	for i := 0; i < 100; i++ {
		_, _elapsed, _eff := g.MakeGridAtlas()
		elapsed += _elapsed
		eff += _eff
	}
	log.Printf("Time: %s", elapsed/100.)
	log.Printf("Pack efficiency: %f", eff/100.)

	return g
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
