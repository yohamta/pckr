package pckr

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type packerInput struct {
	id       string
	img      *ebiten.Image
	location *image.Rectangle
}

// Algo represents an algorithm for Rectangle Packing
type Algo interface {
	Pack(R []*Rectangle, maxWidth, maxHeight int)
}

// Packer packs multiple images into a single image
// for performance reasons
type Packer struct {
	inputs       []*packerInput
	packedImage  *ebiten.Image
	locationDict map[string]*image.Rectangle
	algo         Algo
	w, h         int
}

// NewPacker creates a new packer
func NewPacker(w, h int) *Packer {
	return &Packer{
		algo: SimpleAlgo{},
		w:    w,
		h:    h,
	}
}

// SetAlgo sets algo for rectangle packing
func (p *Packer) SetAlgo(algo Algo) {
	p.algo = algo
}

// Add adds img for packing
func (p *Packer) Add(id string, img *ebiten.Image, x0, y0, x1, y1 int) {
	location := image.Rect(x0, y0, x1, y1)
	size := location.Size()
	if size.X == 0 || size.Y == 0 {
		panic("p.Add() Error: Invalid rectangle specified")
	}
	p.inputs = append(p.inputs,
		&packerInput{
			id:       id,
			img:      img,
			location: &location,
		})
}

type imageRect struct {
	img      *ebiten.Image
	location *Rectangle
}

// Pack execute packing
func (p *Packer) Pack() {
	inputs := p.inputs
	R := make([]*Rectangle, 0, len(inputs))
	cache := map[*ebiten.Image]struct {
		img      *ebiten.Image
		location *Rectangle
	}{}
	imageRects := []imageRect{}

	for _, in := range inputs {
		_, ok := cache[in.img]
		if ok {
			continue
		} else {
			r := Rect(in.img.Size())

			cache[in.img] = struct {
				img      *ebiten.Image
				location *Rectangle
			}{in.img, &r}
			R = append(R, &r)
			imageRects = append(imageRects, imageRect{in.img, &r})
		}
	}

	algo := p.algo
	algo.Pack(R, p.w, p.h)

	p.packedImage = ebiten.NewImage(p.w, p.h)

	for _, imageRect := range imageRects {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(imageRect.location.X), float64(imageRect.location.Y))
		p.packedImage.DrawImage(imageRect.img, op)
	}

	rects := []*Rectangle{}
	for _, in := range inputs {
		rects = append(rects, cache[in.img].location)
	}

	dict := map[string]*image.Rectangle{}
	for i, input := range inputs {
		r := rects[i]
		loc := image.Rect(
			r.X+input.location.Min.X,
			r.Y+input.location.Min.Y,
			r.X+input.location.Max.X,
			r.Y+input.location.Max.Y,
		)
		dict[input.id] = &loc
	}
	p.locationDict = dict
}

// Image returns the packedImage
func (p *Packer) Image() *ebiten.Image {
	if p.packedImage == nil {
		panic("Packed image is nil!")
	}
	return p.packedImage
}

// Location returns the location of the image packed into
func (p *Packer) Location(id string) *image.Rectangle {
	t, ok := p.locationDict[id]
	if !ok {
		log.Fatal("Location was not found for:", id)
	}
	return t
}

// Rectangle represents a rectangle area
type Rectangle struct {
	W, H, X, Y int
}

// Rect returns a Rectangle object
func Rect(w, h int) Rectangle {
	return Rectangle{W: w, H: h, X: 0, Y: 0}
}
