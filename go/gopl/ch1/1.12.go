package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

func main() {
	http.HandleFunc("/image", handler2)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler2(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	cycles, _ := strconv.Atoi(r.Form["cycles"][0])
	size, _ := strconv.Atoi(r.Form["size"][0])
	nframes, _ := strconv.Atoi(r.Form["nframes"][0])
	delay, _ := strconv.Atoi(r.Form["delay"][0])
	res, _ := strconv.ParseFloat(r.Form["res"][0], 64)
	lissajous2(w, cycles, size, nframes, delay, res)
}

func lissajous2(out io.Writer, cycles, size ,nframes, delay int, res float64) {
	green := &green{30, 130, 76, 1}

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount:nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.Set(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), green)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

type green struct {
	r uint32
	g uint32
	b uint32
	a uint32
}

func (green green) RGBA() (r,g,b,a uint32) {
	return green.r, green.g, green.b, green.a
}



