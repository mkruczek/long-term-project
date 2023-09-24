package handlers

import (
	"github.com/gin-gonic/gin"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"market/market/application/service"
	"os"
)

type bar struct {
	Open  int
	Close int
	High  int
	Low   int
	//todo add time
}

// values for bar in points
func fakeData() []bar {
	return []bar{
		{Open: 100100, Close: 100400, High: 100500, Low: 100050},
		{Open: 100400, Close: 100100, High: 100600, Low: 100000},

		{Open: 107188, Close: 107188, High: 107208, Low: 107078},
		{Open: 107188, Close: 107188, High: 107208, Low: 107078},
	}
}

/*
Single bar view
example view is a pixel
---**---
---**---
---*****
---**---
---**---
---**---
*****---
---**---
---**---
*/
func createSingleBarView(b bar) image.Image {
	width := 8
	height := b.High - b.Low

	upLeft := image.Point{}
	lowRight := image.Point{X: width, Y: height}

	img := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

	mainColor := color.RGBA{R: 255, A: 255} // red for down bar, where open > close
	if b.Open < b.Close {
		mainColor = color.RGBA{G: 255, A: 255} // green for up bar, where open < close
	}

	// Set main color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, mainColor)
		}
	}

	//pixel for open
	o := image.Point{X: 0, Y: b.Open - b.Low}
	//make transparent all pixels above open and to 3 pixel right
	for x := 0; x < 3; x++ {
		for y := 0; y < o.Y; y++ {
			img.Set(x, y, color.Transparent)
		}
	}
	//make transparent all pixels below open and to 3 pixel right
	for x := 0; x < 3; x++ {
		for y := o.Y + 3; y < height; y++ {
			img.Set(x, y, color.Transparent)
		}
	}

	//pixel for close
	c := image.Point{X: width - 1, Y: b.Close - b.Low}
	//make transparent all pixels above close and to 3 pixel left
	for x := width - 3; x < width; x++ {
		for y := 0; y < c.Y; y++ {
			img.Set(x, y, color.Transparent)
		}
	}

	//make transparent all pixels below close and to 3 pixel left
	for x := width - 3; x < width; x++ {
		for y := c.Y + 3; y < height; y++ {
			img.Set(x, y, color.Transparent)
		}
	}

	return img
}

func CreateImage(trd service.Trades) gin.HandlerFunc {
	return func(c *gin.Context) {
		//ctx := c.Request.Context()

		fd := fakeData()

		//create image
		img := image.NewRGBA(image.Rect(0, 0, 8*len(fd), 1000))
		// set background color
		for x := 0; x < 8*len(fd); x++ {
			for y := 0; y < 100; y++ {
				img.Set(x, y, color.RGBA{R: 0, G: 0, B: 255, A: 255}) // white color {R:255, G:255, B:255, A:255}
			}
		}

		//draw bars
		for i, b := range fd {
			drawAt := image.Point{X: 8 * i, Y: 0}
			drawRect := image.Rectangle{Min: drawAt, Max: drawAt.Add(image.Point{X: 8, Y: 1000})}
			draw.Draw(img, drawRect, createSingleBarView(b), image.Point{}, draw.Src)
		}

		// Encode as PNG.
		f, _ := os.Create("image.png")
		png.Encode(f, img)

		//send image to client
		c.File("image.png")

		//remove image
		os.Remove("image.png")

	}
}
