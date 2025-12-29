package main

import (
	"fmt"
	"math/rand"

	"github.com/Nadim147c/go-chroma"
	"github.com/Nadim147c/go-chroma/num"
)

func main() {
	const count = 10.0 // generated 10 colors
	const steps = 100 / (count - 1)
	hue := rand.Float64() * 360 // random hue

	palette := make([]chroma.ARGB, count)
	for i := range palette {
		tc := num.Clamp(5, 95, float64(i)*steps)
		palette[i] = chroma.Hct{Hue: hue, Chroma: tc, Tone: tc}.ToARGB()
		fmt.Println(palette[i].AnsiBg("       "))
	}
}
