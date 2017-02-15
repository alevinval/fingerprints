package main

import (
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"time"

	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/math"
	"github.com/google/gxui/samples/flags"
	"github.com/nfnt/resize"
	_ "github.com/nfnt/resize"
)

func loadImage(name string) *image.Gray {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	m, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	m = resize.Resize(300, 300, m, resize.Bilinear)

	bounds := m.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	gray := image.NewGray(bounds)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			oldColor := m.At(x, y)
			grayColor := color.GrayModel.Convert(oldColor)
			gray.Set(x, y, grayColor)
		}
	}

	return gray
}

func appMain(driver gxui.Driver) {
	original := loadImage("corpus/nist2.jpg")
	img := NewMatrixFromGray(original)
	processImage(driver, img)
}

var posX, posY = 0, 0

func showImage(driver gxui.Driver, title string, in *image.Gray) {
	widthMax := in.Bounds().Max

	theme := flags.CreateTheme(driver)
	window := theme.CreateWindow(widthMax.X, widthMax.Y, title)
	window.SetPosition(math.NewPoint(posX, posY))
	posX += widthMax.X
	if posX%(3*widthMax.X) == 0 {
		posY += widthMax.X
		posX = 0
	}
	window.SetScale(flags.DefaultScaleFactor)
	window.SetBackgroundBrush(gxui.WhiteBrush)

	img := theme.CreateImage()
	window.AddChild(img)

	var timer *time.Timer
	pause := time.Millisecond * 1000
	timer = time.AfterFunc(pause, func() {
		driver.Call(func() {
			gray := image.NewRGBA(in.Bounds())
			draw.Draw(gray, in.Bounds(), in, image.ZP, draw.Src)
			texture := driver.CreateTexture(gray, 1)
			img.SetTexture(texture)
			window.Redraw()
			timer.Reset(pause)
		})
	})
	window.OnClose(driver.Terminate)
}

func processImage(driver gxui.Driver, in *Matrix) {
	bounds := in.Bounds()
	normalized := NewMatrix(bounds)
	gx, gy := NewMatrix(bounds), NewMatrix(bounds)
	consistency, normConsistency := NewMatrix(bounds), NewMatrix(bounds)
	directional, normDirectional := NewMatrix(bounds), NewMatrix(bounds)
	filteredD, normFilteredD := NewMatrix(bounds), NewMatrix(bounds)
	segmented, _ := NewMatrix(bounds), NewMatrix(bounds)

	showImage(driver, "Original", in.ToGray())
	Normalize(in, normalized)
	showImage(driver, "Normalized", normalized.ToGray())
	Convolute(SobelDx, normalized, gx)
	Convolute(SobelDy, normalized, gy)
	Convolute(NewSqrtKernel(gx, gy), in, consistency)
	Normalize(consistency, normConsistency)
	showImage(driver, "Normalized Consistency", normConsistency.ToGray())

	Convolute(NewDirectionalKernel(gx, gy), directional, directional)
	Normalize(directional, normDirectional)
	showImage(driver, "Directional", normDirectional.ToGray())

	Convolute(NewFilteredDirectional(gx, gy), filteredD, filteredD)
	Normalize(filteredD, normFilteredD)
	showImage(driver, "Filtered Directional", normFilteredD.ToGray())

	Convolute(NewFilteredDirectionalSegmented(filteredD), normalized, segmented)
	//Normalize(segmented, normSegmented)
	showImage(driver, "Filtered Directional Segmented", segmented.ToGray())
}

func main() {
	gl.StartDriver(appMain)
}
