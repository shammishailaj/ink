package main

import (
	. "github.com/buchanae/ink/color"
	. "github.com/buchanae/ink/dd"
	. "github.com/buchanae/ink/gfx"
	"github.com/buchanae/ink/rand"
)

func Ink(z *Doc) {

	colors := []RGBA{
		Red, Blue, Green, Yellow,
	}

	grid := NewGrid(20, 20)
	for _, r := range grid.Rects() {
		r = r.Shrink(0.002)
		s := z.Shader(r)
		s.Set("a_color", rand.Color(colors))
	}
}