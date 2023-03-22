// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/png"
	"io"
	"math"
	"math/rand"
	"strconv"
)

//!-main
// Packages not needed by version in book.
import (
	"log"
	"net/http"
	"time"
)

//!+main

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, cycles int) {
	const (
		cyclesConst = 5     // number of complete x oscillator revolutions
		res         = 0.001 // angular resolution
		size        = 100   // image canvas covers [-size..+size]
		nframes     = 64    // number of animation frames
		delay       = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		if cycles == 0 {
			cycles = cyclesConst
		}
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

////handler echoes the Path component of the requested URL.
//func handler(w http.ResponseWriter, r *http.Request) {
//	if err := r.ParseForm(); err != nil {
//		log.Print(err)
//	}
//	cycles := "0"
//	fmt.Fprintf(w, "Form[cycles] = %q\n", r.Form["cycles"])
//	if len(r.Form["cycles"]) > 0 {
//		cycles = r.Form["cycles"][len(r.Form["cycles"])-1]
//	}
//	if cycleInt, err := strconv.Atoi(cycles); err == nil {
//		lissajous(w, cycleInt)
//	} else {
//		fmt.Printf("Failed to convert %q to int: %v\n", cycles, err)
//	}
//	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
//}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	var cycles string
	if len(r.Form["cycles"]) > 0 {
		cycles = r.Form["cycles"][len(r.Form["cycles"])-1]
	}
	if cycleInt, err := strconv.Atoi(cycles); err == nil {
		// 将GIF动画写入缓冲区
		var buf bytes.Buffer
		lissajous(&buf, cycleInt)

		// 将GIF动画转换为PNG格式
		img, err := gif.Decode(&buf)
		if err != nil {
			log.Print(err)
			return
		}
		png.Encode(w, img) // 将PNG图像写入HTTP响应
	} else {
		fmt.Fprintf(w, "Failed to convert %q to int: %v\n", cycles, err)
	}
}

//!-main
