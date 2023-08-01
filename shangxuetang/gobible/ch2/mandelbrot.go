// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"os"
	"strconv"
)

func main(){
	http.HandleFunc("/fractal", fractalHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
	log.Fatal(err)
	}
}

func main1() {
    const (
        xmin, ymin, xmax, ymax = -2, -2, +2, +2
        width, height          = 1024, 1024
    )

    img := image.NewRGBA(image.Rect(0, 0, width, height))
    for py := 0; py < height; py++ {
        y := float64(py)/height*(ymax-ymin) + ymin
        for px := 0; px < width; px++ {
            x := float64(px)/width*(xmax-xmin) + xmin
            z := complex(x, y)
            // Image point (px, py) represents complex value z.
            img.Set(px, py, mandelbrot(z))
        }
    }
    // png.Encode(os.Stdout, img) // NOTE: ignoring errors
	f,_ := os.Create("mandelbrot.png")
	png.Encode(f, img)
	f.Close()
}

func mandelbrot(z complex128) color.Color {
    const iterations = 200
    const contrast = 15

    var v complex128
    for n := uint8(0); n < iterations; n++ {
        v = v*v + z
        if cmplx.Abs(v) > 2 {
            return color.Gray{255 - contrast*n}
        }
    }
    return color.Black
}

func fractalHandler(w http.ResponseWriter, r *http.Request) {
	x := r.URL.Query().Get("x")
	y := r.URL.Query().Get("y")
	zoom := r.URL.Query().Get("zoom")

	x0, _ := strconv.ParseFloat(x,64)
	y0, _ := strconv.ParseFloat(y,64)
	zoom0, _ := strconv.ParseFloat(zoom,64)

	img := generatedFractal(x0,y0,zoom0)
	
	png.Encode(w, img)


}

func generatedFractal(x0,y0,zoom float64) image.Image{
	xmin := -2*zoom + x0
	xmax := 2*zoom + x0
	ymin := -2*zoom + y0  
	ymax := 2*zoom + y0

	width := 1024
	height := 1024

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
		x := float64(px)/float64(width)*(xmax-xmin) + xmin
		z := complex(x, y)

		// 计算像素颜色
		img.Set(px, py, mandelbrot(z)) 
		}
	}

	return img
}