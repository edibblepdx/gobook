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
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		factor                 = 2 // supersampling scaling factor (don't change)
	)

	img := image.NewRGBA(image.Rect(0, 0, width*factor, height*factor))
	for py := range height * factor {
		y := float64(py)/(height*factor)*(ymax-ymin) + ymin
		for px := range width * factor {
			x := float64(px)/(width*factor)*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}

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

	png.Encode(os.Stdout, downsampled_img) // NOTE: ignoring errors
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
