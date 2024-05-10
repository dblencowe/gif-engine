package utils

import (
	"bytes"
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"os"

	"github.com/nfnt/resize"
)

type GIFEditor struct {
	gifs []*gif.GIF
}

func (e *GIFEditor) Join() (*bytes.Buffer, error) {
	sizeX, sizeY := e.dimensions(e.gifs[0])
	newImg := gif.GIF{}
	for _, img := range e.gifs {
		for _, frame := range img.Image {
			newFrame := e.resizeImage(frame, uint(sizeX), uint(sizeY))
			newImg.Image = append(newImg.Image, newFrame)
		}

		newImg.Delay = append(newImg.Delay, img.Delay...)
	}
	var buf bytes.Buffer
	err := gif.EncodeAll(&buf, &newImg)
	return &buf, err
}

func (e *GIFEditor) LoadToBuffer(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	img, err := gif.DecodeAll(f)
	if err != nil {
		return err
	}
	e.gifs = append(e.gifs, img)
	return nil
}

func (e *GIFEditor) resizeImage(img image.PalettedImage, width, height uint) *image.Paletted {
	newImg := resize.Resize(width, height, img, resize.Lanczos3)
	b := newImg.Bounds()
	transparentPalette := append(palette.WebSafe, image.Transparent)
	pimg := image.NewPaletted(b, transparentPalette)
	draw.FloydSteinberg.Draw(pimg, b, newImg, image.Point{})
	return pimg
}

func (e *GIFEditor) dimensions(img *gif.GIF) (int, int) {
	var lowestX, lowestY, highestX, highestY int
	for _, img := range img.Image {
		if img.Rect.Min.X < lowestX {
			lowestX = img.Rect.Min.X
		}
		if img.Rect.Min.Y < lowestY {
			lowestY = img.Rect.Min.Y
		}
		if img.Rect.Max.X > highestX {
			highestX = img.Rect.Max.X
		}
		if img.Rect.Max.Y > highestY {
			highestY = img.Rect.Max.Y
		}
	}

	return highestX - lowestX, highestY - lowestY
}
