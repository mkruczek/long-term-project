package handlers

import (
	"github.com/gin-gonic/gin"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"market/market/application/trade/service"
	"os"
	"time"
)

/*
!! UNDER CONSTRUCTION !!
*/

var (
	backgroundColor = color.RGBA{R: 50, G: 50, B: 50, A: 255}
)

// chart will be created from two axis: time - X and price - Y
// chart will be a background for bars
type chart struct {
	bars []bar
	Xs   []time.Time // range of time from earliest -2 to latest +2, will be converted to int
	Ys   []int       // range of prices from lowest -2 to highest +2
}

func newChart(bars []bar) chart {
	return chart{
		bars: bars,
		Xs:   []time.Time{bars[0].Time.Add(-2 * time.Minute), bars[len(bars)-1].Time.Add(2 * time.Minute)},
		Ys:   []int{bars[0].Low - 2, bars[len(bars)-1].High + 2},
	}
}

/*
bar is a single bar in chart
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
type bar struct {
	Open  int
	Close int
	High  int
	Low   int
	Time  time.Time
}

// values for bar in points, assume that i will get one minute bar
func fakeData() []bar {
	return []bar{
		{Open: 100100, Close: 100400, High: 100500, Low: 100050, Time: time.Now()},
		{Open: 100400, Close: 100100, High: 100600, Low: 100000, Time: time.Now().Add(time.Minute)},
		{Open: 100100, Close: 100200, High: 100400, Low: 100000, Time: time.Now().Add(time.Minute * 2)},
		{Open: 100200, Close: 100400, High: 100400, Low: 100200, Time: time.Now().Add(time.Minute * 3)},
	}
}

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
			img.Set(x, y, backgroundColor)
		}
	}
	//make transparent all pixels below open and to 3 pixel right
	for x := 0; x < 3; x++ {
		for y := o.Y + 3; y < height; y++ {
			img.Set(x, y, backgroundColor)
		}
	}

	//pixel for close
	c := image.Point{X: width - 1, Y: b.Close - b.Low}
	//make transparent all pixels above close and to 3 pixel left
	for x := width - 3; x < width; x++ {
		for y := 0; y < c.Y; y++ {
			img.Set(x, y, backgroundColor)
		}
	}

	//make transparent all pixels below close and to 3 pixel left
	for x := width - 3; x < width; x++ {
		for y := c.Y + 3; y < height; y++ {
			img.Set(x, y, backgroundColor)
		}
	}

	return img
}

func CreateImage(trd service.Trades) gin.HandlerFunc {
	return func(c *gin.Context) {
		//ctx := c.Request.Context()

		ch := newChart(fakeData())

		//calculate width base on time
		imgWidth := 8*len(ch.bars) + 2*8 // 8 pixels for one bar and 8 pixels for each side, first bar will be 8 pixels from left side

		//create image
		img := image.NewRGBA(image.Rect(0, 0, imgWidth, ch.Ys[1]-ch.Ys[0]))
		// set background color
		for x := 0; x < 8*len(ch.bars); x++ {
			for y := 0; y < ch.Ys[1]-ch.Ys[0]; y++ {
				img.Set(x, y, backgroundColor)
			}
		}

		//draw bars
		for i, b := range ch.bars {
			drawAt := image.Point{X: 8 * i, Y: 0}
			drawRect := image.Rectangle{Min: drawAt, Max: drawAt.Add(image.Point{X: 8, Y: ch.Ys[1] - ch.Ys[0]})}
			draw.Draw(img, drawRect, createSingleBarView(b), image.Point{}, draw.Src)
		}

		////create image
		//img := image.NewRGBA(image.Rect(0, 0, 1000, 1000))
		//// set background color
		//for x := 0; x < 8*len(fd); x++ {
		//	for y := 0; y < 100; y++ {
		//		img.Set(x, y, color.RGBA{R: 0, G: 0, B: 255, A: 255}) // white color {R:255, G:255, B:255, A:255}
		//	}
		//}
		//
		////draw bars
		//for i, b := range fd {
		//	drawAt := image.Point{X: 8 * i, Y: 0}
		//	drawRect := image.Rectangle{Min: drawAt, Max: drawAt.Add(image.Point{X: 8, Y: 1000})}
		//	draw.Draw(img, drawRect, createSingleBarView(b), image.Point{}, draw.Src)
		//}

		// Encode as PNG.
		f, _ := os.Create("image.png")
		png.Encode(f, img)

		//send image to client
		c.File("image.png")

		//remove image
		os.Remove("image.png")

	}
}
