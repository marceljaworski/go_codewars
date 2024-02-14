package main

import (
	"io"
	"os"
	"strings"
	"unicode"
)

type rot13Reader struct {
	src io.Reader
}

func rot13(b byte) byte {
	var char byte
	if unicode.IsLetter(rune(b)) {
		char = b + 13
		if unicode.IsLower(rune(b)) && char > 122 || unicode.IsUpper(rune(b)) && char > 90 {
			char -= 26
		}
	} else {
		return b
	}
	return char
}
func (r13 *rot13Reader) gugu(b []byte) (int, error) {
	n, err := r13.src.Read(b)
	if err == io.EOF {
		// There is no more data to read
		return 0, err
	}
	if err != nil {
		return 0, err
	}
	for i, l := 0, len(b); i < n && i < l; i++ {
		b[i] = rot13(b[i])
	}
	return n, nil
}

func (r *rot13Reader) Read(b []byte) (int, error) {
	return r.gugu(b)
}

type test struct{}

func (t test) String() string {
	return "lala"
}

func main() {
	var s io.Reader
	s = strings.NewReader("Lbh penpxrq gur pbqr!")
	// s, err := os.Open("test.txt")
	// if err != nil {
	// 	panic(err)
	// }
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
