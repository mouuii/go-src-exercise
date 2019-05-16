package myreader

import "io"

type alphaReader struct {
	src string
	cur int
}

// NewAlphaReader : Create  myAlphaReader
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

// also can read from any reader
type alphaChainReader struct {
	reader io.Reader
}

func (a *alphaChainReader) Read(p []byte) (int, error) {
	n, err := a.reader.Read(p)
	if err != nil {
		return n, err
	}
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		if char := alpha(p[i]); char != 0 {
			buf[i] = char
		}
	}
	copy(p, buf)
	return n, nil
}
func NewAlphaChainReader(reader io.Reader) *alphaChainReader {
	return &alphaChainReader{reader: reader}
}
