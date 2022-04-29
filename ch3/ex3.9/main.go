package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.RequestURI() == "/favicon.ico" {
			return
		}
		log.Println(r.URL.RequestURI())
		generate(w, r)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func generate(w io.Writer, r *http.Request) {
	x1, y1 := 2.0, 2.0
	const (
		width, height = 1024, 1024
	)
	if err := r.ParseForm(); err != nil {
		fmt.Println(err)
	}
	if v, err := getParam(r, "x"); err == nil {
		x1 = v
	}
	if v, err := getParam(r, "y"); err == nil {
		y1 = v
	}
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(2*y1) - y1
		for px := 0; px < width; px++ {
			x := float64(px)/width*(2*x1) - x1
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(w, img)
}

func getParam(r *http.Request, s string) (float64, error) {
	var x float64
	if _, ok := r.Form["x"]; ok {
		tmp, _ := strconv.Atoi(r.Form["x"][0])
		x = float64(tmp)
		if tmp < 0 {
			x *= -1
		}
		return x, nil
	}
	return x, fmt.Errorf("param x not exists")
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
