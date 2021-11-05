# pckr

Texture Packing library for [Ebiten](https://ebiten.org/). Packs multiple textures into a single image on-the-fly.

[GoDoc](https://pkg.go.dev/github.com/yohamta/pckr)

## Simple Example
```go
func (g *Game) setup() {

	// create a new packer
	packer = pckr.NewPacker(1024, 1024)

	// add images to the packer
	packer.Add("priest", ebiten.NewImageFromImage(bytes2Image(&images.CHARACTER_HERO_PRIEST)), 0, 0, 600, 300)
	packer.Add("archor", ebiten.NewImageFromImage(bytes2Image(&images.CHARACTER_HERO_ARCHOR)), 0, 0, 600, 300)
	packer.Add("warrior", ebiten.NewImageFromImage(bytes2Image(&images.CHARACTER_HERO_WARRIOR)), 0, 0, 600, 300)

	// execute texture packing
	packer.Pack()

	packedImage := packer.Image()
	packedLocation := packer.Location("archor")

	// ...
}
```

[source code](https://github.com/yohamta/pckr/blob/master/examples/packing/main.go)
