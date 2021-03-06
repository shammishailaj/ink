package main

import (
	"github.com/buchanae/ink/app"
	. "github.com/buchanae/ink/color"
	. "github.com/buchanae/ink/dd"
	"github.com/buchanae/ink/gfx"
	"github.com/buchanae/ink/rand"
)

const (
	Size = 10
	Z    = 0.015
)

func Ink(doc *app.Doc) {
	rand.SeedNow()
	gfx.Clear(doc, White)

	palette := rand.Palette()
	grid := HexGrid{Size}
	cells := grid.Cells()

	for i := range cells {
		c := cells[i].Center
		cells[i].Center = c.Add(rand.XYRange(-Z, Z))
	}

	for _, cell := range cells {
		col := rand.Color(palette)
		gfx.Fill{cell, col}.Draw(doc)
	}

	for _, cell := range cells {
		gfx.Stroke{
			Target: cell,
			Color:  Black,
			StrokeOpt: StrokeOpt{
				Width: 0.004,
			},
		}.Draw(doc)
	}
}
