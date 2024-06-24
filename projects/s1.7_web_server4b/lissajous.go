// Lissajous generates GIF animations of random Lissajous figures.
// ch 1.4
package main

// When using the intellisense for a routine, it will
// autodetect the package import and add it here.
import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

// this is a composite literal. A compact notation
// for instantiating any of Golangs composite types
// from a sequence of element values.  Here this is a
// slice.
// var palette = []color.Color{color.Black, color.White}
// var palette = []color.Color{color.White, color.Black}
// Red Green Blue Alpha
// Red background
// var palette = []color.Color{color.RGBA{255, 0, 0, 255}, color.Black}
// Red background with alpha = 128 so its a slight red clearish
// ex 1.5
// var palette = []color.Color{color.Black, color.RGBA{0x0, 0xA0, 0x0, 0xFF}}
var palette = []color.Color{color.Black, color.RGBA{0x0, 0xA0, 0x0, 0xFF}, color.RGBA{0x0, 0x0, 0xA0, 0xFF}}

// ORIGINAL
//const (
//	whiteIndex = 0 // first color in palette
//	blackIndex = 1 // next color in palette
//)

// ex 1.5
//const (
//	blackIndex = 0 // first color in palette
//	greenIndex = 1 // next color in palette
//)

// ex 1.6
const (
	blackIndex = 0 // first color in palette
	greenIndex = 1 // next color in palette
	blueIndex  = 2 // next color in palette
)

//func main() {
//	lissajous(os.Stdout)
//}

func lissajous(out io.Writer) {

	// const is similar to var but its a constant variable.  Note
	// how he is making all these in a block rather than individual
	// const variable = value  lines.
	const (
		cycles  = 5     // number of complete x ossillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms unit
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	// this is a composite literal. A compact notation
	// for instantiating any of Golangs composite types
	// from a sequence of element values.  Here this is a
	// struct.  Meaning it has a group of values called fields.
	// The individual fields can be referenced with .name.
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	// This loop runs for 64 iterations making a single frame.
	// It creates a single frame of 201x201 pixels, two colors.
	// Initially the frame is set to the zero pallete color
	// in this case white
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		// This inner loop runs two sin functions for x and y
		// the y value varies each time for freq and phase.
		// freq is random the first time and fixed but phase starts
		// at zero and increases with each iteration.
		// Each time it draws a pixel using SetColorIndex
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// EX 1.6
			if i%2 == 0 {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
			} else {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blueIndex)
			}
		}
		phase += 0.1
		// Append frames with a specified delay of 80ms
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	// This takes the list of frames and converts to a gif file.
	// out is an io.writer object.
	gif.EncodeAll(out, &anim) // NOTE: ignorning encoding errors
}

// Exercise 1.5 change the color palette to green on black
// To make the web color #RRGGBB use color.RGBA{0xRR, 0xGG, 0xBB, 0xFF}

// Exercise 1.6 Modify the lissajous program to produce images in multiple
// colors by adding more values to palette and then displaying them by
// changing the third argument of SetColorIndex in some interesting way.
