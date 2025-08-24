// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 1.12: Modify the Lissajous server to read parameter values from
// the URL. Use strconv.Atoi.

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"strconv"
)

//!-main
// Packages not needed by version in book.
import (
	"log"
	"net/http"
	"net/url"
	"time"
)

//!+main

var palette = []color.Color{
	color.Black,
	color.RGBA{0x00, 0xff, 0x00, 0xff},
}

const (
	blackIndex = 0 // first color in palette
	greenIndex = 1 // next color in palette
)

// !+URLSearchParams
type params struct {
	cycles  int     // number of complete x oscillator revolutions
	res     float64 // angular resolution
	size    int     // image canvas covers [-size..+size]
	nframes int     // number of animation frames
	delay   int     // delay between frames in 10ms units
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

//!-URLSearchParams

func main() {
	//!-main
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.

	// WARN: rand.Seed() deprecated as of Go 1.20
	// rand.Seed(time.Now().UTC().UnixNano())
	random := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			q := r.URL.Query()
			p := params{
				queryInt(q, "cycles", 5),
				queryFloat64(q, "res", 0.001),
				queryInt(q, "size", 100),
				queryInt(q, "nframes", 64),
				queryInt(q, "delay", 8),
			}
			lissajous(w, p, random)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	//!+main
	lissajous(os.Stdout, params{5, 0.001, 100, 64, 8}, random)
}

func lissajous(out io.Writer, p params, random *rand.Rand) {
	freq := random.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: p.nframes}
	phase := 0.0 // phase difference
	for range p.nframes {
		rect := image.Rect(0, 0, 2*p.size+1, 2*p.size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(p.cycles)*2*math.Pi; t += p.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(p.size+int(x*float64(p.size)+0.5), p.size+int(y*float64(p.size)+0.5),
				greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, p.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//!-main
