package main

import (
	"fmt"
	"unicode/utf8"

	"github.com/Nadim147c/go-chroma"
)

func createGradiant(from, to chroma.ARGB, count int) []chroma.ARGB {
	if count < 2 {
		panic("atlest 2 color has be created from gradient")
	}

	fromOkLab := from.ToOkLab()
	toOkLab := to.ToOkLab()

	step := 1.0 / float64(count-1)

	colors := make([]chroma.ARGB, count)
	for i := range count {
		ratio := step * float64(i)
		l := fromOkLab.L + (toOkLab.L-fromOkLab.L)*ratio
		a := fromOkLab.A + (toOkLab.A-fromOkLab.A)*ratio
		b := fromOkLab.B + (toOkLab.B-fromOkLab.B)*ratio
		colors[i] = chroma.NewOkLab(l, a, b).ToARGB()
	}
	return colors
}

func main() {
	from := chroma.ARGBFromHexMust("#ff0000")
	to := chroma.ARGBFromHexMust("#00ff00")

	const text = "Print me in a gradient!"

	grad := createGradiant(from, to, utf8.RuneCountInString(text))
	for i, r := range text {
		fmt.Print(grad[i].AnsiFg(string(r)))
	}
	fmt.Println()
}
