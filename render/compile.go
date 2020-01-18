package render

import "errors"

// TODO would be better if the cache was per GL context
var shaderCache map[shaderOpt]program

type shaderOpt struct {
	vert, frag, geom string
}

func init() {
	shaderCache = map[shaderOpt]program{}
}

func compile(s shaderOpt) (program, error) {

	if p, ok := shaderCache[s]; ok {
		return p, nil
	}

	var prog program
	var err error

	if s.vert == "" {
		return prog, errors.New("empty vert shader")
	}
	if s.frag == "" {
		return prog, errors.New("empty frag shader")
	}

	/* TODO no longer supported due to gl_FragColor
					being required by webgl
	out := s.out
	if out == "" {
		out = "f_color"
	}
	*/

	prog, err = glBuildProgram(s.vert, s.frag, s.geom)
	if err != nil {
		return prog, err
	}

	shaderCache[s] = prog

	return prog, nil
}
