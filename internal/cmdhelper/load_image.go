package cmdhelper

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"github.com/alevinval/fingerprints/internal/matrix"
	"github.com/nfnt/resize"
)

func resizeImage(img image.Image) image.Image {
	maxDimension := 300
	dx := img.Bounds().Dx()
	dy := img.Bounds().Dy()
	if dx < maxDimension && dy < maxDimension {
		return img
	}

	xp := 0
	yp := 0
	if dx > dy {
		xp = maxDimension
		yp = int(float64(dy) / (float64(dx) / float64(maxDimension)))
	} else if dy > dx {
		yp = maxDimension
		xp = int(float64(dx) / (float64(dy) / float64(maxDimension)))
	} else {
		xp, yp = maxDimension, maxDimension
	}
	log.Printf("resizing image from (%d,%d) to (%d,%d)", dx, dy, xp, yp)
	return resize.Resize(uint(xp), uint(yp), img, resize.Bilinear)
}

func LoadImage(name string) *matrix.M {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	resizedImg := resizeImage(img)

	bounds := resizedImg.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	gray := image.NewGray(bounds)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			oldColor := resizedImg.At(x, y)
			grayColor := color.GrayModel.Convert(oldColor)
			gray.Set(x, y, grayColor)
		}
	}

	return matrix.NewFromGray(gray)
}
