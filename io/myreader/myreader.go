package myreader

import "io"

type alphaReader struct {
	src string
	cur int
}

func NewAlphaReader(src string) *alphaReader {
	return &alphaReader{src: src}
}
func alpha(r byte) byte {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	}

	return 0
}
func (a *alphaReader) Read(p []byte) (int, error) {
	if a.cur >= len(a.src) {
		return 0, io.EOF
	}
	x := len(a.src) - a.cur
	n, bound := 0, 0
	// remain bytes > len(p)
	if x >= len(p) {
		bound = len(p)
	} else {
		bound = x
	}
	buf := make([]byte, bound)
	for n < bound {
		if char := alpha(a.src[a.cur]); char != 0 {
			buf[n] = char
		}
		n++
		a.cur++
	}
	copy(p, buf)
	return n, nil
}
