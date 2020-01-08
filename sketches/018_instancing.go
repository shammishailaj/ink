package main

import (
	. "github.com/buchanae/ink/color"
	. "github.com/buchanae/ink/dd"
	. "github.com/buchanae/ink/gfx"
	"github.com/buchanae/ink/rand"
)

func Ink(doc *Doc) {
	Clear(doc, White)

	const N = 10000

	pos := make([]XY, N)
	rot := make([]float32, N)
	colors := make([]RGBA, N)
	palette := rand.Palette()

	for i := 0; i < N; i++ {
		pos[i] = rand.XYRange(0.1, 0.9)
		rot[i] = rand.Angle()
		colors[i] = rand.Color(palette)
	}

	doc.AddShader(&Shader{
		Vert:      DefaultVert,
		Frag:      DefaultFrag,
		Instances: N,
		Mesh:      RectWH(0.05, 0.05),
		Attrs: Attrs{
			"a_pos":   pos,
			"a_rot":   rot,
			"a_color": colors,
		},
		Divisors: map[string]int{
			"a_pos":   1,
			"a_rot":   1,
			"a_color": 1,
		},
	})
}
