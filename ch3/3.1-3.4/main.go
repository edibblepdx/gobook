// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 3.1: If the function f returns a non-finite float64 value, the SVG
// file will contain invalid <polygon> elements. Modify the program to skip
// invalid polygons.

// Exercise 3.2: Experiments with visualizations of other functions from the math
// package.

// Exercise 3.3: Color each polygon based on its height. So that
// Peaks are red (#ff0000) and valleys are blue (#0000ff).

// Exercise 3.4: Construct a web server.

// Usage: go run . -s=h web

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"flag"
	"fmt"
	"math"

	cv "gobook/ch2/2.2/colorconv"
)

// net
import (
	"io"
	"log"
	"net/http"
	"os"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

var surfaceFlag = flag.String("s", "f", "surface function (f, g, or h)")
var surfaceFunc func(x, y float64) float64

func main() {
	flag.Parse()
	switch *surfaceFlag {
	case "f":
		surfaceFunc = f
	case "g":
		surfaceFunc = g
	case "h":
		surfaceFunc = h
	default:
		surfaceFunc = f
	}

	posArgs := flag.Args()
	if len(posArgs) > 0 && posArgs[0] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "image/svg+xml")
			svg(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	svg(os.Stdout)
}

func svg(out io.Writer) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: black; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := range cells {
		for j := range cells {
			ax, ay, a_ok, az := corner(i+1, j)
			bx, by, b_ok, bz := corner(i, j)
			cx, cy, c_ok, cz := corner(i, j+1)
			dx, dy, d_ok, dz := corner(i+1, j+1)

			// Exercise 3.1: Check if the result is a finite number
			if !(a_ok && b_ok && c_ok && d_ok) {
				continue
			}

			// Exercise 3.3: Color each polygon based on height
			avgZ := (az + bz + cz + dz) / 4.0
			r, g, b := color(avgZ)

			fmt.Fprintf(out, "<polygon fill='#%02x%02x%02x' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				r, g, b, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j int) (float64, float64, bool, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5) // [-15, 15]
	y := xyrange * (float64(j)/cells - 0.5) // [-15, 15]

	// Compute surface height z.
	z := surfaceFunc(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	// Exercise 3.1: Check if the result is a finite number
	ok := true
	if math.IsNaN(sx) || math.IsNaN(sy) ||
		math.IsInf(sx, 0) || math.IsInf(sy, 0) {
		ok = false
	}

	return sx, sy, ok, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

// Exercise 3.2: Experiment with other functions from the math library
func g(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Atan(r) * math.Sin(r) * (0.5 / r)
}

func h(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Cos(r) * math.Sin(r) * 0.8
}

// Exercise 3.3: Color each polygon based on height
func color(t float64) (uint8, uint8, uint8) {
	t = clamp(t, -1, 1)
	t = (t + 1) / 2 // remap

	// linear lerp
	lrgbColor := cv.Lrgb(cv.Vec4{X: t, Y: 0, Z: (1 - t), W: 1})
	srgbColor := cv.LrgbToSrgb(lrgbColor)

	return uint8(srgbColor.X * 255), uint8(srgbColor.Y * 255), uint8(srgbColor.Z * 255)
}

func clamp(t, min, max float64) float64 {
	return math.Min(max, math.Max(min, t))
}

//!-
