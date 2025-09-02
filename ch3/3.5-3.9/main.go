// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 3.5: Implement a full-color Mandelbrot set using the function
// image.NewRGBA and the type color.RGBA or color.YCbCr

// Exercise 3.6: Implement Supersampling.

// Exercise 3.7: Newton's method fractal.

// Exercise 3.8: Implement the same fractal using complex64, complex128,
// big.Float, and big.Rat.

// Exercise 3.9: Web server. Allow parameters for x, y, and zoom.

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	//"math/big"
	"math/cmplx"
	"os"
)

// net
import (
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type params struct {
	width   int     // width of image
	height  int     // height of image
	zoom    float64 // fractal zoom
	fractal string  // fractal function
}

func queryInt(query url.Values, s string, def int) int {
	v, err := strconv.Atoi(query.Get(s))
	if err != nil {
		v = def
	}
	return v
}

func queryFloat64(query url.Values, s string, def float64) float64 {
	v, err := strconv.ParseFloat(query.Get(s), 64)
	if err != nil {
		v = def
	}
	return v
}

func queryString(query url.Values, s string, def string) string {
	v := query.Get(s)
	if v == "" {
		v = def
	}
	return v
}

type fracf = func(complex128) color.Color

var fractals = map[string]fracf{
	"mandelbrot": func(z complex128) color.Color { return mandelbrot(z) },
	"acos":       func(z complex128) color.Color { return acos(z) },
	"sqrt":       func(z complex128) color.Color { return sqrt(z) },
	"newton":     func(z complex128) color.Color { return newton(z) },
}

func main() {
	// Exercise 3.9: Web server.
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			q := r.URL.Query()
			p := params{
				queryInt(q, "width", 1024),
				queryInt(q, "height", 1024),
				queryFloat64(q, "zoom", 1.0),
				queryString(q, "fractal", "mandelbrot"),
			}

			w.Header().Set("Content-Type", "image/png")
			fractal(w, p)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	fractal(os.Stdout, params{1024, 1024, 1, "mandelbrot"})
}

func fractal(out io.Writer, p params) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		//width, height          = 1024, 1024
		factor = 2 // scaling factor (don't change)
	)

	// Exercise 3.9: Web server.
	zoom := 1.0 / p.zoom
	width, height := p.width, p.height

	img := image.NewRGBA(image.Rect(0, 0, width*factor, height*factor))
	for py := range height * factor {
		y := float64(py)/(float64(height)*factor)*(ymax-ymin)*zoom + ymin*zoom
		for px := range width * factor {
			x := float64(px)/(float64(width)*factor)*(xmax-xmin)*zoom + xmin*zoom
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, fractals[p.fractal](z))
		}
	}

	// Exercise 3.6: Supersampling
	downsampled_img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := range height {
		for px := range width {
			psx := px * factor
			psy := py * factor

			r1, g1, b1, _ := img.At(psx, psy).RGBA()
			r2, g2, b2, _ := img.At(psx+1, psy).RGBA()
			r3, g3, b3, _ := img.At(psx, psy+1).RGBA()
			r4, g4, b4, _ := img.At(psx+1, psy+1).RGBA()

			// divide by 4 times the pre-multiplied alpha
			// each color is in the range [0, 0xFFFF]
			// 0xFFFF / 255 = 257
			r := (r1 + r2 + r3 + r4) / (4 * 257)
			g := (g1 + g2 + g3 + g4) / (4 * 257)
			b := (b1 + b2 + b3 + b4) / (4 * 257)

			downsampled_img.Set(px, py, color.RGBA{uint8(r), uint8(g), uint8(b), 255})
		}
	}

	png.Encode(out, downsampled_img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {

			// Exercise 3.5: This just interpolates between
			// red and blue difference, with luminance at 128.
			blue := 255 - contrast*n
			red := contrast * n

			return color.YCbCr{128, blue, red}
		}
	}
	return color.Black
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//
//	= z - (z^4 - 1) / (4 * z^3)
//	= z - (z - 1/z^3) / 4

// Exercise 3.7: Newton's fractal.
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
