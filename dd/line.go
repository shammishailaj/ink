package dd

func NewLine(a, b XY) Line {
	return Line{a, b}
}

func XYsToLines(xys ...XY) []Line {
	if len(xys) < 2 {
		return nil
	}

	var lines []Line
	for i := 0; i < len(xys)-1; i++ {
		end := i + 1
		lines = append(lines, Line{xys[i], xys[end]})
	}
	return lines
}

type Line struct {
	A, B XY
}

func (l Line) Start() XY {
	return l.A
}

func (l Line) End() XY {
	return l.B
}

func (l Line) Length() float32 {
	return sqrt(l.SquaredLength())
}

func (l Line) SquaredLength() float32 {
	x := l.B.X - l.A.X
	y := l.B.Y - l.A.Y
	return x*x + y*y
}

func (l Line) Vector() XY {
	return l.B.Sub(l.A).Normalize()
}

func (l Line) Normal() XY {
	dx := l.B.X - l.A.X
	dy := l.B.Y - l.A.Y
	return XY{-dy, dx}.Normalize()
}

func (l Line) Interpolate(p float32) XY {
	return XY{
		X: l.A.X + ((l.B.X - l.A.X) * p),
		Y: l.A.Y + ((l.B.Y - l.A.Y) * p),
	}
}

func (l Line) InnerAngle(m Line) float32 {
	c := l.B.Sub(l.A)
	d := m.B.Sub(m.A)
	numer := c.X*d.X + c.Y*d.Y
	denom := l.Length() * m.Length()
	return acos(numer / denom)
}

func (l Line) RelativeAngle(m Line) float32 {
	c := l.B.Sub(l.A)
	d := m.B.Sub(m.A)
	return atan2(d.Y, d.X) - atan2(c.Y, c.X)
}

func (l Line) Stroke(opt StrokeOpt) Mesh {
	return Stroke(Path{l}, opt)
}
