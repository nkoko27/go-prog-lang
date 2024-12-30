package main

import (
	// "fmt"
	"log"
	"net/http"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"strconv"
)

var palette = []color.Color{color.Black, color.RGBA{0x00, 0xFF, 0x00, 0xff}, color.RGBA{0x00, 0xFF, 0xFF, 0xff}, color.RGBA{0xFF, 0x00, 0x00, 0xff}}

const (
	blackIndex = 0
	greenIndex = 1
	blueIndex = 2
	redIndex = 3
)

func main() {
	http.HandleFunc("/", handler) // each req calls handler
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
// 	for k, v := range r.Header {
// 		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
// 	}
// 	fmt.Fprintf(w, "Host = %q\n", r.Host)
// 	fmt.Fprintf(w, "Remote Address = %q\n", r.RemoteAddr)
// 	if err := r.ParseForm(); err != nil {
// 		log.Print(err)
// 	}
// 	for k, v := range r.Form {
// 		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
// 	}
// }
func handler(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	cycles, err := strconv.Atoi(queryValues.Get("cycles"))
	if err != nil {
		log.Print(err)
	}
	lissajous(w, cycles)
}

func lissajous(out io.Writer, num int) {
	log.Print(num)
	const (
		res = 0.001 // angular resolution
		size = 100 // image canvas covers [-size..+size]
		nframes = 64 // number of animation frames
		delay = 8 // delay between frames in 10ms units
	)
	cycles := float64(num) // number of complete x oscillator revolutions
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0,0,2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		ind := uint8(rand.Intn(3) + 1)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), ind)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

// exercise 1.12
// modify server to read params from url
// DONE