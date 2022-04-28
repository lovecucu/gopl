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
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	cycles := 5
	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		if _, ok := r.Form["cycles"]; ok {
			cycles, _ = strconv.Atoi(r.Form["cycles"][0])
		}
		lissajous(w, cycles)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, cycles int) {
	const (
		res     = 0.001 // 角度分辨率
		size    = 100   // 图像画布包含【-size..+size】
		nframes = 64    // 动画中的帧数
		delay   = 8     // 以10ns为单位的帧间延迟
	)
	freq := rand.Float64() * 3.0 // y振荡器的相对频率
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
