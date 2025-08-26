// Color Conversion cli
// This program is hardly tested.

// Exercise 2.2: Create a command line unit conversion program

package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	cv "gobook/ch2/2.2/colorconv"
)

var (
	inputColorSpace  = flag.String("i", "", "Input Color Space")
	outputColorSpace = flag.String("o", "", "Output Color Space")

	inputColor *cv.Vec4 = nil
)

type parser int

func (p parser) parseError(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	flag.Usage()
	os.Exit(1)
}

func (p parser) tryFloat(s string) float64 {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		p.parseError("bad format")
	}
	return v
}

func (p parser) parse() {
	flag.Parse()

	// flags
	if *inputColorSpace == "" {
		p.parseError("Error: -i is required")
	}
	if *outputColorSpace == "" {
		p.parseError("Error: -o is required")
	}

	// positional arguments
	if posArgs := flag.Args(); len(posArgs) == 1 {
		color := strings.Split(posArgs[0], ",")
		if len(color) != 3 {
			p.parseError("Error: input color requires 3 channels as 'x,y,z'")
		}

		x := p.tryFloat(color[0])
		y := p.tryFloat(color[1])
		z := p.tryFloat(color[2])

		inputColor = &cv.Vec4{X: x, Y: y, Z: z, W: 1.0}

	} else {
		fmt.Fprintln(
			os.Stderr,
			"Usage: ccli -i <colorspace> -o <colorspace> <0-1>,<0-1>,<0-1> ",
		)
		p.parseError("Error: input color is required")
	}
}

type Vec4 cv.Vec4
type converter func(Vec4) Vec4

var conversions = map[string]map[string]converter{
	"srgb": {
		"lrgb":      func(v Vec4) Vec4 { return Vec4(cv.SrgbToLrgb(cv.Srgb(v))) },
		"xyz":       func(v Vec4) Vec4 { return Vec4(cv.SrgbToXyz(cv.Srgb(v))) },
		"luminance": func(v Vec4) Vec4 { return Vec4(cv.SrgbToLuminance(cv.Srgb(v))) },
	},
	"lrgb": {
		"srgb":      func(v Vec4) Vec4 { return Vec4(cv.LrgbToSrgb(cv.Lrgb(v))) },
		"xyz":       func(v Vec4) Vec4 { return Vec4(cv.LrgbToXyz(cv.Lrgb(v))) },
		"luminance": func(v Vec4) Vec4 { return Vec4(cv.LrgbToLuminance(cv.Lrgb(v))) },
	},
	"xyz": {
		"srgb": func(v Vec4) Vec4 { return Vec4(cv.XyzToSrgb(cv.Xyz(v))) },
		"lrgb": func(v Vec4) Vec4 { return Vec4(cv.XyzToLrgb(cv.Xyz(v))) },
	},
}

func main() {
	var p parser
	p.parse()

	fmt.Println(conversions[*inputColorSpace][*outputColorSpace](Vec4(*inputColor)))
}
