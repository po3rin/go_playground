package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/Iwark/text2img"
)

func main() {
	width := 200
	height := 100

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	cyan := color.RGBA{100, 200, 200, 0xff}

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch {
			case x < width/2 && y < height/2: // upper left quadrant
				img.Set(x, y, cyan)
			case x >= width/2 && y >= height/2: // lower right quadrant
				img.Set(x, y, color.White)
			default:
				// Use zero value.
			}
		}
	}

	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)

	ttfpath := "mplus-1c-thin.ttf"
	imagepath := "cool-background.png"
	d, _ := text2img.NewDrawer(
		text2img.Params{
			BackgroundImagePath: imagepath,
			FontPath:            ttfpath,
			FontSize:            128,
		},
	)

	img, _ = d.Draw("Go言語でOGP画像を作る")

	file, _ := os.Create("image.jpg")
	defer file.Close()

	_ = jpeg.Encode(file, img, &jpeg.Options{Quality: 100})
}
