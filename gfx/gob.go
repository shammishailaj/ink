package gfx

// TODO want to remove these dependencies from code used by sketches
import (
	"encoding/gob"
	"image"

	"github.com/buchanae/ink/color"
	"github.com/buchanae/ink/dd"
)

func init() {
	gob.Register(Shader{})
	gob.Register(Image{})
	gob.Register(color.RGBA{})
	gob.Register(dd.XY{})
	gob.Register(dd.Mesh{})
	gob.Register(dd.Rect{})
	gob.Register(dd.Quad{})
	gob.Register(dd.Line{})
	gob.Register(dd.HexCell{})
	gob.Register(dd.Triangle{})
	gob.Register(dd.Circle{})
	gob.Register(dd.Ellipse{})
	gob.Register(dd.Triangles{})
	gob.Register(dd.Quadratic{})
	gob.Register([]color.RGBA{})
	gob.Register([]dd.XY{})
	gob.Register(&image.RGBA{})
	gob.Register(&image.NRGBA{})
	gob.Register(&image.Gray{})
}
