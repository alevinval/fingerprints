package main

import (
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"os"
	_ "time"

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

	m = resize.Resize(400, 400, m, resize.Bilinear)

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
	original := loadImage("corpus/nist1.jpg")
	img := NewMatrixFromGray(original)
	processImage(driver, img)
}

var posX, posY = 0, 0

func showImage(driver gxui.Driver, title string, in *Matrix) {
	bounds := in.Bounds()

	theme := flags.CreateTheme(driver)
	window := theme.CreateWindow(bounds.Dx(), bounds.Dy(), title)
	window.SetPosition(math.NewPoint(posX, posY))
	posX += bounds.Dx()
	if posX%(4*bounds.Dx()) == 0 {
		posY += bounds.Dy() + 70
		posX = 0
	}
	window.SetScale(flags.DefaultScaleFactor)
	window.SetBackgroundBrush(gxui.WhiteBrush)

	img := theme.CreateImage()
	window.AddChild(img)

	gray := image.NewRGBA(in.Bounds())
	draw.Draw(gray, in.Bounds(), in.ToGray(), image.ZP, draw.Src)
	texture := driver.CreateTexture(gray, 1)
	img.SetTexture(texture)
	window.OnClose(driver.Terminate)
}

func processImage(driver gxui.Driver, in *Matrix) {
	bounds := in.Bounds()
	normalized := NewMatrix(bounds)

	showImage(driver, "Original", in)
	Normalize(in, normalized)
	showImage(driver, "Normalized", normalized)

	gx, gy := NewMatrix(bounds), NewMatrix(bounds)
	c1 := ParallelConvolution(SobelDx, normalized, gx)
	c2 := ParallelConvolution(SobelDy, normalized, gy)
	c1.Wait()
	c2.Wait()

	//Consistency matrix
	consistency, normConsistency := NewMatrix(bounds), NewMatrix(bounds)
	c1 = ParallelConvolution(NewSqrtKernel(gx, gy), in, consistency)
	c1.Wait()
	Normalize(consistency, normConsistency)
	showImage(driver, "Normalized Consistency", normConsistency)

	// Compute directional
	directional, normDirectional := NewMatrix(bounds), NewMatrix(bounds)
	c1 = ParallelConvolution(NewDirectionalKernel(gx, gy), directional, directional)
	c1.Wait()
	Normalize(directional, normDirectional)
	showImage(driver, "Directional", normDirectional)

	// Compute filtered directional
	filteredD, normFilteredD := NewMatrix(bounds), NewMatrix(bounds)
	Convolute(NewFilteredDirectional(gx, gy, 4), filteredD, filteredD)
	Normalize(filteredD, normFilteredD)
	showImage(driver, "Filtered Directional", normFilteredD)

	// Compute segmented image
	segmented, normSegmented := NewMatrix(bounds), NewMatrix(bounds)
	Convolute(NewStdDevKernel(filteredD, 8), normalized, segmented)
	Normalize(segmented, normSegmented)
	showImage(driver, "Filtered Directional Std Dev.", normSegmented)

	// Compute binarized segmented image
	binarizedSegmented := NewMatrix(bounds)
	Binarize(normSegmented, binarizedSegmented)
	showImage(driver, "Binarized Segmented", binarizedSegmented)

	binarizedNorm := NewMatrix(bounds)
	Binarize(normalized, binarizedNorm)
	showImage(driver, "Binarized Normalized", binarizedNorm)

}

func main() {
	gl.StartDriver(appMain)
}
