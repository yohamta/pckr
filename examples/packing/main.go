package main

import (
	"bytes"
	"image"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/pckr"
	"github.com/yohamta/pckr/examples/assets/images"
)

const (
	screenWidth  = 1024
	screenHeight = 1024
)

var (
	packer *pckr.Packer
	result *ebiten.Image
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	screen.DrawImage(result, &ebiten.DrawImageOptions{})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func NewGame() *Game {
	g := &Game{}
	g.setup()

	return g
}

func (g *Game) setup() {
	packer = pckr.NewPacker(1024, 1024)

	packer.Add("priest", ebiten.NewImageFromImage(bytes2Image(&images.CHARACTER_HERO_PRIEST)), 0, 0, 600, 300)
	packer.Add("archor", ebiten.NewImageFromImage(bytes2Image(&images.CHARACTER_HERO_ARCHOR)), 0, 0, 600, 300)
	packer.Add("warrior", ebiten.NewImageFromImage(bytes2Image(&images.CHARACTER_HERO_WARRIOR)), 0, 0, 600, 300)
	packer.Pack()

	result = ebiten.NewImage(1024, 1024)
	result.DrawImage(packer.Image(), &ebiten.DrawImageOptions{})

	DrawRect(result, *packer.Location("priest"), color.RGBA{0xff, 0, 0, 0xff}, 2)
	DrawRect(result, *packer.Location("archor"), color.RGBA{0, 0xff, 0, 0xff}, 2)
	DrawRect(result, *packer.Location("warrior"), color.RGBA{0, 0xff, 0xff, 0xff}, 2)
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}

func bytes2Image(rawImage *[]byte) image.Image {
	img, format, error := image.Decode(bytes.NewReader(*rawImage))
	if error != nil {
		log.Fatal("Bytes2Image Failed: ", format, error)
	}
	return img
}
